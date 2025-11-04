// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package dedibox provides methods and message types of the dedibox v1 API.
package dedibox

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

type AttachFailoverIPToMacAddressRequestMacType string

const (
	AttachFailoverIPToMacAddressRequestMacTypeMacTypeUnknown = AttachFailoverIPToMacAddressRequestMacType("mac_type_unknown")
	AttachFailoverIPToMacAddressRequestMacTypeVmware         = AttachFailoverIPToMacAddressRequestMacType("vmware")
	AttachFailoverIPToMacAddressRequestMacTypeKvm            = AttachFailoverIPToMacAddressRequestMacType("kvm")
	AttachFailoverIPToMacAddressRequestMacTypeXen            = AttachFailoverIPToMacAddressRequestMacType("xen")
)

func (enum AttachFailoverIPToMacAddressRequestMacType) String() string {
	if enum == "" {
		// return default value if empty
		return string(AttachFailoverIPToMacAddressRequestMacTypeMacTypeUnknown)
	}
	return string(enum)
}

func (enum AttachFailoverIPToMacAddressRequestMacType) Values() []AttachFailoverIPToMacAddressRequestMacType {
	return []AttachFailoverIPToMacAddressRequestMacType{
		"mac_type_unknown",
		"vmware",
		"kvm",
		"xen",
	}
}

func (enum AttachFailoverIPToMacAddressRequestMacType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AttachFailoverIPToMacAddressRequestMacType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AttachFailoverIPToMacAddressRequestMacType(AttachFailoverIPToMacAddressRequestMacType(tmp).String())
	return nil
}

type BMCAccessStatus string

const (
	BMCAccessStatusUnknown  = BMCAccessStatus("unknown")
	BMCAccessStatusCreating = BMCAccessStatus("creating")
	BMCAccessStatusCreated  = BMCAccessStatus("created")
	BMCAccessStatusDeleting = BMCAccessStatus("deleting")
)

func (enum BMCAccessStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(BMCAccessStatusUnknown)
	}
	return string(enum)
}

func (enum BMCAccessStatus) Values() []BMCAccessStatus {
	return []BMCAccessStatus{
		"unknown",
		"creating",
		"created",
		"deleting",
	}
}

func (enum BMCAccessStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *BMCAccessStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = BMCAccessStatus(BMCAccessStatus(tmp).String())
	return nil
}

type BackupStatus string

const (
	BackupStatusUnknownBackupStatus = BackupStatus("unknown_backup_status")
	BackupStatusUninitialized       = BackupStatus("uninitialized")
	BackupStatusInactive            = BackupStatus("inactive")
	BackupStatusReady               = BackupStatus("ready")
)

func (enum BackupStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(BackupStatusUnknownBackupStatus)
	}
	return string(enum)
}

func (enum BackupStatus) Values() []BackupStatus {
	return []BackupStatus{
		"unknown_backup_status",
		"uninitialized",
		"inactive",
		"ready",
	}
}

func (enum BackupStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *BackupStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = BackupStatus(BackupStatus(tmp).String())
	return nil
}

type FailoverBlockVersion string

const (
	FailoverBlockVersionUnknownVersion = FailoverBlockVersion("unknown_version")
	FailoverBlockVersionIPv4           = FailoverBlockVersion("ipv4")
	FailoverBlockVersionIPv6           = FailoverBlockVersion("ipv6")
)

func (enum FailoverBlockVersion) String() string {
	if enum == "" {
		// return default value if empty
		return string(FailoverBlockVersionUnknownVersion)
	}
	return string(enum)
}

func (enum FailoverBlockVersion) Values() []FailoverBlockVersion {
	return []FailoverBlockVersion{
		"unknown_version",
		"ipv4",
		"ipv6",
	}
}

func (enum FailoverBlockVersion) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *FailoverBlockVersion) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = FailoverBlockVersion(FailoverBlockVersion(tmp).String())
	return nil
}

type FailoverIPInterfaceType string

const (
	FailoverIPInterfaceTypeUnknown = FailoverIPInterfaceType("unknown")
	FailoverIPInterfaceTypeNormal  = FailoverIPInterfaceType("normal")
	FailoverIPInterfaceTypeIpmi    = FailoverIPInterfaceType("ipmi")
	FailoverIPInterfaceTypeVirtual = FailoverIPInterfaceType("virtual")
)

func (enum FailoverIPInterfaceType) String() string {
	if enum == "" {
		// return default value if empty
		return string(FailoverIPInterfaceTypeUnknown)
	}
	return string(enum)
}

func (enum FailoverIPInterfaceType) Values() []FailoverIPInterfaceType {
	return []FailoverIPInterfaceType{
		"unknown",
		"normal",
		"ipmi",
		"virtual",
	}
}

func (enum FailoverIPInterfaceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *FailoverIPInterfaceType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = FailoverIPInterfaceType(FailoverIPInterfaceType(tmp).String())
	return nil
}

type FailoverIPStatus string

const (
	FailoverIPStatusUnknownStatus = FailoverIPStatus("unknown_status")
	FailoverIPStatusReady         = FailoverIPStatus("ready")
	FailoverIPStatusBusy          = FailoverIPStatus("busy")
	FailoverIPStatusLocked        = FailoverIPStatus("locked")
)

func (enum FailoverIPStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(FailoverIPStatusUnknownStatus)
	}
	return string(enum)
}

func (enum FailoverIPStatus) Values() []FailoverIPStatus {
	return []FailoverIPStatus{
		"unknown_status",
		"ready",
		"busy",
		"locked",
	}
}

func (enum FailoverIPStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *FailoverIPStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = FailoverIPStatus(FailoverIPStatus(tmp).String())
	return nil
}

type FailoverIPVersion string

const (
	FailoverIPVersionUnknownVersion = FailoverIPVersion("unknown_version")
	FailoverIPVersionIPv4           = FailoverIPVersion("ipv4")
	FailoverIPVersionIPv6           = FailoverIPVersion("ipv6")
)

func (enum FailoverIPVersion) String() string {
	if enum == "" {
		// return default value if empty
		return string(FailoverIPVersionUnknownVersion)
	}
	return string(enum)
}

func (enum FailoverIPVersion) Values() []FailoverIPVersion {
	return []FailoverIPVersion{
		"unknown_version",
		"ipv4",
		"ipv6",
	}
}

func (enum FailoverIPVersion) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *FailoverIPVersion) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = FailoverIPVersion(FailoverIPVersion(tmp).String())
	return nil
}

type GetRpnStatusResponseStatus string

const (
	GetRpnStatusResponseStatusUnknownStatus = GetRpnStatusResponseStatus("unknown_status")
	GetRpnStatusResponseStatusBusy          = GetRpnStatusResponseStatus("busy")
	GetRpnStatusResponseStatusOperational   = GetRpnStatusResponseStatus("operational")
)

func (enum GetRpnStatusResponseStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(GetRpnStatusResponseStatusUnknownStatus)
	}
	return string(enum)
}

func (enum GetRpnStatusResponseStatus) Values() []GetRpnStatusResponseStatus {
	return []GetRpnStatusResponseStatus{
		"unknown_status",
		"busy",
		"operational",
	}
}

func (enum GetRpnStatusResponseStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *GetRpnStatusResponseStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = GetRpnStatusResponseStatus(GetRpnStatusResponseStatus(tmp).String())
	return nil
}

type IPSemantic string

const (
	IPSemanticUnknown   = IPSemantic("unknown")
	IPSemanticProxad    = IPSemantic("proxad")
	IPSemanticExt       = IPSemantic("ext")
	IPSemanticPublic    = IPSemantic("public")
	IPSemanticPrivate   = IPSemantic("private")
	IPSemanticIpmi      = IPSemantic("ipmi")
	IPSemanticAdm       = IPSemantic("adm")
	IPSemanticRedirect  = IPSemantic("redirect")
	IPSemanticMigration = IPSemantic("migration")
)

func (enum IPSemantic) String() string {
	if enum == "" {
		// return default value if empty
		return string(IPSemanticUnknown)
	}
	return string(enum)
}

func (enum IPSemantic) Values() []IPSemantic {
	return []IPSemantic{
		"unknown",
		"proxad",
		"ext",
		"public",
		"private",
		"ipmi",
		"adm",
		"redirect",
		"migration",
	}
}

func (enum IPSemantic) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *IPSemantic) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = IPSemantic(IPSemantic(tmp).String())
	return nil
}

type IPStatus string

const (
	IPStatusUnknownStatus = IPStatus("unknown_status")
	IPStatusReady         = IPStatus("ready")
	IPStatusBusy          = IPStatus("busy")
	IPStatusLocked        = IPStatus("locked")
)

func (enum IPStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(IPStatusUnknownStatus)
	}
	return string(enum)
}

func (enum IPStatus) Values() []IPStatus {
	return []IPStatus{
		"unknown_status",
		"ready",
		"busy",
		"locked",
	}
}

func (enum IPStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *IPStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = IPStatus(IPStatus(tmp).String())
	return nil
}

type IPVersion string

const (
	IPVersionIPv4 = IPVersion("ipv4")
	IPVersionIPv6 = IPVersion("ipv6")
)

func (enum IPVersion) String() string {
	if enum == "" {
		// return default value if empty
		return string(IPVersionIPv4)
	}
	return string(enum)
}

func (enum IPVersion) Values() []IPVersion {
	return []IPVersion{
		"ipv4",
		"ipv6",
	}
}

func (enum IPVersion) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *IPVersion) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = IPVersion(IPVersion(tmp).String())
	return nil
}

type IPv6BlockDelegationStatus string

const (
	IPv6BlockDelegationStatusUnknownStatus = IPv6BlockDelegationStatus("unknown_status")
	IPv6BlockDelegationStatusUpdating      = IPv6BlockDelegationStatus("updating")
	IPv6BlockDelegationStatusDone          = IPv6BlockDelegationStatus("done")
)

func (enum IPv6BlockDelegationStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(IPv6BlockDelegationStatusUnknownStatus)
	}
	return string(enum)
}

func (enum IPv6BlockDelegationStatus) Values() []IPv6BlockDelegationStatus {
	return []IPv6BlockDelegationStatus{
		"unknown_status",
		"updating",
		"done",
	}
}

func (enum IPv6BlockDelegationStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *IPv6BlockDelegationStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = IPv6BlockDelegationStatus(IPv6BlockDelegationStatus(tmp).String())
	return nil
}

type InvoicePaymentMethod string

const (
	InvoicePaymentMethodUnknownPaymentMethod = InvoicePaymentMethod("unknown_payment_method")
	InvoicePaymentMethodCreditCard           = InvoicePaymentMethod("credit_card")
	InvoicePaymentMethodAmex                 = InvoicePaymentMethod("amex")
	InvoicePaymentMethodPaypal               = InvoicePaymentMethod("paypal")
	InvoicePaymentMethodTransfer             = InvoicePaymentMethod("transfer")
	InvoicePaymentMethodDirectDebit          = InvoicePaymentMethod("direct_debit")
)

func (enum InvoicePaymentMethod) String() string {
	if enum == "" {
		// return default value if empty
		return string(InvoicePaymentMethodUnknownPaymentMethod)
	}
	return string(enum)
}

func (enum InvoicePaymentMethod) Values() []InvoicePaymentMethod {
	return []InvoicePaymentMethod{
		"unknown_payment_method",
		"credit_card",
		"amex",
		"paypal",
		"transfer",
		"direct_debit",
	}
}

func (enum InvoicePaymentMethod) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *InvoicePaymentMethod) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = InvoicePaymentMethod(InvoicePaymentMethod(tmp).String())
	return nil
}

type InvoiceStatus string

const (
	InvoiceStatusUnknownInvoiceStatus = InvoiceStatus("unknown_invoice_status")
	InvoiceStatusUnpaid               = InvoiceStatus("unpaid")
	InvoiceStatusPaid                 = InvoiceStatus("paid")
	InvoiceStatusErrored              = InvoiceStatus("errored")
)

func (enum InvoiceStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(InvoiceStatusUnknownInvoiceStatus)
	}
	return string(enum)
}

func (enum InvoiceStatus) Values() []InvoiceStatus {
	return []InvoiceStatus{
		"unknown_invoice_status",
		"unpaid",
		"paid",
		"errored",
	}
}

func (enum InvoiceStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *InvoiceStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = InvoiceStatus(InvoiceStatus(tmp).String())
	return nil
}

type ListFailoverIPsRequestOrderBy string

const (
	ListFailoverIPsRequestOrderByIPAsc  = ListFailoverIPsRequestOrderBy("ip_asc")
	ListFailoverIPsRequestOrderByIPDesc = ListFailoverIPsRequestOrderBy("ip_desc")
)

func (enum ListFailoverIPsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListFailoverIPsRequestOrderByIPAsc)
	}
	return string(enum)
}

func (enum ListFailoverIPsRequestOrderBy) Values() []ListFailoverIPsRequestOrderBy {
	return []ListFailoverIPsRequestOrderBy{
		"ip_asc",
		"ip_desc",
	}
}

func (enum ListFailoverIPsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListFailoverIPsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListFailoverIPsRequestOrderBy(ListFailoverIPsRequestOrderBy(tmp).String())
	return nil
}

type ListInvoicesRequestOrderBy string

const (
	ListInvoicesRequestOrderByCreatedAtAsc  = ListInvoicesRequestOrderBy("created_at_asc")
	ListInvoicesRequestOrderByCreatedAtDesc = ListInvoicesRequestOrderBy("created_at_desc")
)

func (enum ListInvoicesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListInvoicesRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListInvoicesRequestOrderBy) Values() []ListInvoicesRequestOrderBy {
	return []ListInvoicesRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListInvoicesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListInvoicesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListInvoicesRequestOrderBy(ListInvoicesRequestOrderBy(tmp).String())
	return nil
}

type ListOSRequestOrderBy string

const (
	ListOSRequestOrderByCreatedAtAsc   = ListOSRequestOrderBy("created_at_asc")
	ListOSRequestOrderByCreatedAtDesc  = ListOSRequestOrderBy("created_at_desc")
	ListOSRequestOrderByReleasedAtAsc  = ListOSRequestOrderBy("released_at_asc")
	ListOSRequestOrderByReleasedAtDesc = ListOSRequestOrderBy("released_at_desc")
)

func (enum ListOSRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListOSRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListOSRequestOrderBy) Values() []ListOSRequestOrderBy {
	return []ListOSRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"released_at_asc",
		"released_at_desc",
	}
}

func (enum ListOSRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListOSRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListOSRequestOrderBy(ListOSRequestOrderBy(tmp).String())
	return nil
}

type ListOffersRequestOrderBy string

const (
	ListOffersRequestOrderByCreatedAtAsc  = ListOffersRequestOrderBy("created_at_asc")
	ListOffersRequestOrderByCreatedAtDesc = ListOffersRequestOrderBy("created_at_desc")
	ListOffersRequestOrderByNameAsc       = ListOffersRequestOrderBy("name_asc")
	ListOffersRequestOrderByNameDesc      = ListOffersRequestOrderBy("name_desc")
	ListOffersRequestOrderByPriceAsc      = ListOffersRequestOrderBy("price_asc")
	ListOffersRequestOrderByPriceDesc     = ListOffersRequestOrderBy("price_desc")
)

func (enum ListOffersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListOffersRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListOffersRequestOrderBy) Values() []ListOffersRequestOrderBy {
	return []ListOffersRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"name_asc",
		"name_desc",
		"price_asc",
		"price_desc",
	}
}

func (enum ListOffersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListOffersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListOffersRequestOrderBy(ListOffersRequestOrderBy(tmp).String())
	return nil
}

type ListRefundsRequestOrderBy string

const (
	ListRefundsRequestOrderByCreatedAtAsc  = ListRefundsRequestOrderBy("created_at_asc")
	ListRefundsRequestOrderByCreatedAtDesc = ListRefundsRequestOrderBy("created_at_desc")
)

func (enum ListRefundsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListRefundsRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListRefundsRequestOrderBy) Values() []ListRefundsRequestOrderBy {
	return []ListRefundsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListRefundsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListRefundsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListRefundsRequestOrderBy(ListRefundsRequestOrderBy(tmp).String())
	return nil
}

type ListRpnCapableSanServersRequestOrderBy string

const (
	ListRpnCapableSanServersRequestOrderByCreatedAtAsc  = ListRpnCapableSanServersRequestOrderBy("created_at_asc")
	ListRpnCapableSanServersRequestOrderByCreatedAtDesc = ListRpnCapableSanServersRequestOrderBy("created_at_desc")
)

func (enum ListRpnCapableSanServersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListRpnCapableSanServersRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListRpnCapableSanServersRequestOrderBy) Values() []ListRpnCapableSanServersRequestOrderBy {
	return []ListRpnCapableSanServersRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListRpnCapableSanServersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListRpnCapableSanServersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListRpnCapableSanServersRequestOrderBy(ListRpnCapableSanServersRequestOrderBy(tmp).String())
	return nil
}

type ListRpnCapableServersRequestOrderBy string

const (
	ListRpnCapableServersRequestOrderByCreatedAtAsc  = ListRpnCapableServersRequestOrderBy("created_at_asc")
	ListRpnCapableServersRequestOrderByCreatedAtDesc = ListRpnCapableServersRequestOrderBy("created_at_desc")
)

func (enum ListRpnCapableServersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListRpnCapableServersRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListRpnCapableServersRequestOrderBy) Values() []ListRpnCapableServersRequestOrderBy {
	return []ListRpnCapableServersRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListRpnCapableServersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListRpnCapableServersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListRpnCapableServersRequestOrderBy(ListRpnCapableServersRequestOrderBy(tmp).String())
	return nil
}

type ListRpnGroupMembersRequestOrderBy string

const (
	ListRpnGroupMembersRequestOrderByCreatedAtAsc  = ListRpnGroupMembersRequestOrderBy("created_at_asc")
	ListRpnGroupMembersRequestOrderByCreatedAtDesc = ListRpnGroupMembersRequestOrderBy("created_at_desc")
)

func (enum ListRpnGroupMembersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListRpnGroupMembersRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListRpnGroupMembersRequestOrderBy) Values() []ListRpnGroupMembersRequestOrderBy {
	return []ListRpnGroupMembersRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListRpnGroupMembersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListRpnGroupMembersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListRpnGroupMembersRequestOrderBy(ListRpnGroupMembersRequestOrderBy(tmp).String())
	return nil
}

type ListRpnGroupsRequestOrderBy string

const (
	ListRpnGroupsRequestOrderByCreatedAtAsc  = ListRpnGroupsRequestOrderBy("created_at_asc")
	ListRpnGroupsRequestOrderByCreatedAtDesc = ListRpnGroupsRequestOrderBy("created_at_desc")
)

func (enum ListRpnGroupsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListRpnGroupsRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListRpnGroupsRequestOrderBy) Values() []ListRpnGroupsRequestOrderBy {
	return []ListRpnGroupsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListRpnGroupsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListRpnGroupsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListRpnGroupsRequestOrderBy(ListRpnGroupsRequestOrderBy(tmp).String())
	return nil
}

type ListRpnInvitesRequestOrderBy string

const (
	ListRpnInvitesRequestOrderByCreatedAtAsc  = ListRpnInvitesRequestOrderBy("created_at_asc")
	ListRpnInvitesRequestOrderByCreatedAtDesc = ListRpnInvitesRequestOrderBy("created_at_desc")
)

func (enum ListRpnInvitesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListRpnInvitesRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListRpnInvitesRequestOrderBy) Values() []ListRpnInvitesRequestOrderBy {
	return []ListRpnInvitesRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListRpnInvitesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListRpnInvitesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListRpnInvitesRequestOrderBy(ListRpnInvitesRequestOrderBy(tmp).String())
	return nil
}

type ListRpnSansRequestOrderBy string

const (
	ListRpnSansRequestOrderByCreatedAtAsc  = ListRpnSansRequestOrderBy("created_at_asc")
	ListRpnSansRequestOrderByCreatedAtDesc = ListRpnSansRequestOrderBy("created_at_desc")
)

func (enum ListRpnSansRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListRpnSansRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListRpnSansRequestOrderBy) Values() []ListRpnSansRequestOrderBy {
	return []ListRpnSansRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListRpnSansRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListRpnSansRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListRpnSansRequestOrderBy(ListRpnSansRequestOrderBy(tmp).String())
	return nil
}

type ListRpnServerCapabilitiesRequestOrderBy string

const (
	ListRpnServerCapabilitiesRequestOrderByCreatedAtAsc  = ListRpnServerCapabilitiesRequestOrderBy("created_at_asc")
	ListRpnServerCapabilitiesRequestOrderByCreatedAtDesc = ListRpnServerCapabilitiesRequestOrderBy("created_at_desc")
)

func (enum ListRpnServerCapabilitiesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListRpnServerCapabilitiesRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListRpnServerCapabilitiesRequestOrderBy) Values() []ListRpnServerCapabilitiesRequestOrderBy {
	return []ListRpnServerCapabilitiesRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListRpnServerCapabilitiesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListRpnServerCapabilitiesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListRpnServerCapabilitiesRequestOrderBy(ListRpnServerCapabilitiesRequestOrderBy(tmp).String())
	return nil
}

type ListRpnV2CapableResourcesRequestOrderBy string

const (
	ListRpnV2CapableResourcesRequestOrderByCreatedAtAsc  = ListRpnV2CapableResourcesRequestOrderBy("created_at_asc")
	ListRpnV2CapableResourcesRequestOrderByCreatedAtDesc = ListRpnV2CapableResourcesRequestOrderBy("created_at_desc")
)

func (enum ListRpnV2CapableResourcesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListRpnV2CapableResourcesRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListRpnV2CapableResourcesRequestOrderBy) Values() []ListRpnV2CapableResourcesRequestOrderBy {
	return []ListRpnV2CapableResourcesRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListRpnV2CapableResourcesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListRpnV2CapableResourcesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListRpnV2CapableResourcesRequestOrderBy(ListRpnV2CapableResourcesRequestOrderBy(tmp).String())
	return nil
}

type ListRpnV2GroupLogsRequestOrderBy string

const (
	ListRpnV2GroupLogsRequestOrderByCreatedAtAsc  = ListRpnV2GroupLogsRequestOrderBy("created_at_asc")
	ListRpnV2GroupLogsRequestOrderByCreatedAtDesc = ListRpnV2GroupLogsRequestOrderBy("created_at_desc")
)

func (enum ListRpnV2GroupLogsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListRpnV2GroupLogsRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListRpnV2GroupLogsRequestOrderBy) Values() []ListRpnV2GroupLogsRequestOrderBy {
	return []ListRpnV2GroupLogsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListRpnV2GroupLogsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListRpnV2GroupLogsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListRpnV2GroupLogsRequestOrderBy(ListRpnV2GroupLogsRequestOrderBy(tmp).String())
	return nil
}

type ListRpnV2GroupsRequestOrderBy string

const (
	ListRpnV2GroupsRequestOrderByCreatedAtAsc  = ListRpnV2GroupsRequestOrderBy("created_at_asc")
	ListRpnV2GroupsRequestOrderByCreatedAtDesc = ListRpnV2GroupsRequestOrderBy("created_at_desc")
)

func (enum ListRpnV2GroupsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListRpnV2GroupsRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListRpnV2GroupsRequestOrderBy) Values() []ListRpnV2GroupsRequestOrderBy {
	return []ListRpnV2GroupsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListRpnV2GroupsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListRpnV2GroupsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListRpnV2GroupsRequestOrderBy(ListRpnV2GroupsRequestOrderBy(tmp).String())
	return nil
}

type ListRpnV2MembersRequestOrderBy string

const (
	ListRpnV2MembersRequestOrderByCreatedAtAsc  = ListRpnV2MembersRequestOrderBy("created_at_asc")
	ListRpnV2MembersRequestOrderByCreatedAtDesc = ListRpnV2MembersRequestOrderBy("created_at_desc")
)

func (enum ListRpnV2MembersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListRpnV2MembersRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListRpnV2MembersRequestOrderBy) Values() []ListRpnV2MembersRequestOrderBy {
	return []ListRpnV2MembersRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListRpnV2MembersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListRpnV2MembersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListRpnV2MembersRequestOrderBy(ListRpnV2MembersRequestOrderBy(tmp).String())
	return nil
}

type ListRpnV2MembersRequestType string

const (
	ListRpnV2MembersRequestTypeUnknownType = ListRpnV2MembersRequestType("unknown_type")
	ListRpnV2MembersRequestTypeRpnv1Group  = ListRpnV2MembersRequestType("rpnv1_group")
	ListRpnV2MembersRequestTypeServer      = ListRpnV2MembersRequestType("server")
)

func (enum ListRpnV2MembersRequestType) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListRpnV2MembersRequestTypeUnknownType)
	}
	return string(enum)
}

func (enum ListRpnV2MembersRequestType) Values() []ListRpnV2MembersRequestType {
	return []ListRpnV2MembersRequestType{
		"unknown_type",
		"rpnv1_group",
		"server",
	}
}

func (enum ListRpnV2MembersRequestType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListRpnV2MembersRequestType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListRpnV2MembersRequestType(ListRpnV2MembersRequestType(tmp).String())
	return nil
}

type ListServerDisksRequestOrderBy string

const (
	ListServerDisksRequestOrderByCreatedAtAsc  = ListServerDisksRequestOrderBy("created_at_asc")
	ListServerDisksRequestOrderByCreatedAtDesc = ListServerDisksRequestOrderBy("created_at_desc")
)

func (enum ListServerDisksRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListServerDisksRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListServerDisksRequestOrderBy) Values() []ListServerDisksRequestOrderBy {
	return []ListServerDisksRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListServerDisksRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListServerDisksRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListServerDisksRequestOrderBy(ListServerDisksRequestOrderBy(tmp).String())
	return nil
}

type ListServerEventsRequestOrderBy string

const (
	ListServerEventsRequestOrderByCreatedAtAsc  = ListServerEventsRequestOrderBy("created_at_asc")
	ListServerEventsRequestOrderByCreatedAtDesc = ListServerEventsRequestOrderBy("created_at_desc")
)

func (enum ListServerEventsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListServerEventsRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListServerEventsRequestOrderBy) Values() []ListServerEventsRequestOrderBy {
	return []ListServerEventsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListServerEventsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListServerEventsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListServerEventsRequestOrderBy(ListServerEventsRequestOrderBy(tmp).String())
	return nil
}

type ListServersRequestOrderBy string

const (
	ListServersRequestOrderByCreatedAtAsc  = ListServersRequestOrderBy("created_at_asc")
	ListServersRequestOrderByCreatedAtDesc = ListServersRequestOrderBy("created_at_desc")
)

func (enum ListServersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListServersRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListServersRequestOrderBy) Values() []ListServersRequestOrderBy {
	return []ListServersRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListServersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListServersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListServersRequestOrderBy(ListServersRequestOrderBy(tmp).String())
	return nil
}

type ListServicesRequestOrderBy string

const (
	ListServicesRequestOrderByCreatedAtAsc  = ListServicesRequestOrderBy("created_at_asc")
	ListServicesRequestOrderByCreatedAtDesc = ListServicesRequestOrderBy("created_at_desc")
)

func (enum ListServicesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListServicesRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListServicesRequestOrderBy) Values() []ListServicesRequestOrderBy {
	return []ListServicesRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListServicesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListServicesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListServicesRequestOrderBy(ListServicesRequestOrderBy(tmp).String())
	return nil
}

type LogAction string

const (
	LogActionUnknownLogAction        = LogAction("unknown_log_action")
	LogActionGroupCreated            = LogAction("group_created")
	LogActionGroupDeleted            = LogAction("group_deleted")
	LogActionMembersAdded            = LogAction("members_added")
	LogActionMembersDeleted          = LogAction("members_deleted")
	LogActionDescriptionUpdated      = LogAction("description_updated")
	LogActionRpnv1MembersAdded       = LogAction("rpnv1_members_added")
	LogActionRpnv1MembersDeleted     = LogAction("rpnv1_members_deleted")
	LogActionVlanUpdated             = LogAction("vlan_updated")
	LogActionVlanUpdatedOnAllServers = LogAction("vlan_updated_on_all_servers")
)

func (enum LogAction) String() string {
	if enum == "" {
		// return default value if empty
		return string(LogActionUnknownLogAction)
	}
	return string(enum)
}

func (enum LogAction) Values() []LogAction {
	return []LogAction{
		"unknown_log_action",
		"group_created",
		"group_deleted",
		"members_added",
		"members_deleted",
		"description_updated",
		"rpnv1_members_added",
		"rpnv1_members_deleted",
		"vlan_updated",
		"vlan_updated_on_all_servers",
	}
}

func (enum LogAction) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *LogAction) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = LogAction(LogAction(tmp).String())
	return nil
}

type LogStatus string

const (
	LogStatusUnknownLogStatus = LogStatus("unknown_log_status")
	LogStatusSuccess          = LogStatus("success")
	LogStatusInProgress       = LogStatus("in_progress")
	LogStatusError            = LogStatus("error")
)

func (enum LogStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(LogStatusUnknownLogStatus)
	}
	return string(enum)
}

func (enum LogStatus) Values() []LogStatus {
	return []LogStatus{
		"unknown_log_status",
		"success",
		"in_progress",
		"error",
	}
}

func (enum LogStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *LogStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = LogStatus(LogStatus(tmp).String())
	return nil
}

type MemoryType string

const (
	MemoryTypeDdr2 = MemoryType("ddr2")
	MemoryTypeDdr3 = MemoryType("ddr3")
	MemoryTypeDdr4 = MemoryType("ddr4")
	MemoryTypeDdr5 = MemoryType("ddr5")
)

func (enum MemoryType) String() string {
	if enum == "" {
		// return default value if empty
		return string(MemoryTypeDdr2)
	}
	return string(enum)
}

func (enum MemoryType) Values() []MemoryType {
	return []MemoryType{
		"ddr2",
		"ddr3",
		"ddr4",
		"ddr5",
	}
}

func (enum MemoryType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *MemoryType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = MemoryType(MemoryType(tmp).String())
	return nil
}

type NetworkInterfaceInterfaceType string

const (
	NetworkInterfaceInterfaceTypeUnknown = NetworkInterfaceInterfaceType("unknown")
	NetworkInterfaceInterfaceTypeNormal  = NetworkInterfaceInterfaceType("normal")
	NetworkInterfaceInterfaceTypeIpmi    = NetworkInterfaceInterfaceType("ipmi")
	NetworkInterfaceInterfaceTypeVirtual = NetworkInterfaceInterfaceType("virtual")
)

func (enum NetworkInterfaceInterfaceType) String() string {
	if enum == "" {
		// return default value if empty
		return string(NetworkInterfaceInterfaceTypeUnknown)
	}
	return string(enum)
}

func (enum NetworkInterfaceInterfaceType) Values() []NetworkInterfaceInterfaceType {
	return []NetworkInterfaceInterfaceType{
		"unknown",
		"normal",
		"ipmi",
		"virtual",
	}
}

func (enum NetworkInterfaceInterfaceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NetworkInterfaceInterfaceType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NetworkInterfaceInterfaceType(NetworkInterfaceInterfaceType(tmp).String())
	return nil
}

type OSArch string

const (
	OSArchUnknownArch = OSArch("unknown_arch")
	OSArchAmd64       = OSArch("amd64")
	OSArchX86         = OSArch("x86")
	OSArchArm         = OSArch("arm")
	OSArchArm64       = OSArch("arm64")
)

func (enum OSArch) String() string {
	if enum == "" {
		// return default value if empty
		return string(OSArchUnknownArch)
	}
	return string(enum)
}

func (enum OSArch) Values() []OSArch {
	return []OSArch{
		"unknown_arch",
		"amd64",
		"x86",
		"arm",
		"arm64",
	}
}

func (enum OSArch) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OSArch) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OSArch(OSArch(tmp).String())
	return nil
}

type OSType string

const (
	OSTypeUnknownType = OSType("unknown_type")
	OSTypeServer      = OSType("server")
	OSTypeVirtu       = OSType("virtu")
	OSTypePanel       = OSType("panel")
	OSTypeDesktop     = OSType("desktop")
	OSTypeCustom      = OSType("custom")
	OSTypeRescue      = OSType("rescue")
)

func (enum OSType) String() string {
	if enum == "" {
		// return default value if empty
		return string(OSTypeUnknownType)
	}
	return string(enum)
}

func (enum OSType) Values() []OSType {
	return []OSType{
		"unknown_type",
		"server",
		"virtu",
		"panel",
		"desktop",
		"custom",
		"rescue",
	}
}

func (enum OSType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OSType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OSType(OSType(tmp).String())
	return nil
}

type OfferAntiDosInfoType string

const (
	OfferAntiDosInfoTypeMinimal    = OfferAntiDosInfoType("minimal")
	OfferAntiDosInfoTypePreventive = OfferAntiDosInfoType("preventive")
	OfferAntiDosInfoTypeCurative   = OfferAntiDosInfoType("curative")
)

func (enum OfferAntiDosInfoType) String() string {
	if enum == "" {
		// return default value if empty
		return string(OfferAntiDosInfoTypeMinimal)
	}
	return string(enum)
}

func (enum OfferAntiDosInfoType) Values() []OfferAntiDosInfoType {
	return []OfferAntiDosInfoType{
		"minimal",
		"preventive",
		"curative",
	}
}

func (enum OfferAntiDosInfoType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OfferAntiDosInfoType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OfferAntiDosInfoType(OfferAntiDosInfoType(tmp).String())
	return nil
}

type OfferCatalog string

const (
	OfferCatalogAll      = OfferCatalog("all")
	OfferCatalogDefault  = OfferCatalog("default")
	OfferCatalogBeta     = OfferCatalog("beta")
	OfferCatalogReseller = OfferCatalog("reseller")
	OfferCatalogPremium  = OfferCatalog("premium")
	OfferCatalogVolume   = OfferCatalog("volume")
	OfferCatalogAdmin    = OfferCatalog("admin")
	OfferCatalogInactive = OfferCatalog("inactive")
)

func (enum OfferCatalog) String() string {
	if enum == "" {
		// return default value if empty
		return string(OfferCatalogAll)
	}
	return string(enum)
}

func (enum OfferCatalog) Values() []OfferCatalog {
	return []OfferCatalog{
		"all",
		"default",
		"beta",
		"reseller",
		"premium",
		"volume",
		"admin",
		"inactive",
	}
}

func (enum OfferCatalog) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OfferCatalog) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OfferCatalog(OfferCatalog(tmp).String())
	return nil
}

type OfferPaymentFrequency string

const (
	OfferPaymentFrequencyMonthly = OfferPaymentFrequency("monthly")
	OfferPaymentFrequencyOneshot = OfferPaymentFrequency("oneshot")
)

func (enum OfferPaymentFrequency) String() string {
	if enum == "" {
		// return default value if empty
		return string(OfferPaymentFrequencyMonthly)
	}
	return string(enum)
}

func (enum OfferPaymentFrequency) Values() []OfferPaymentFrequency {
	return []OfferPaymentFrequency{
		"monthly",
		"oneshot",
	}
}

func (enum OfferPaymentFrequency) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OfferPaymentFrequency) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OfferPaymentFrequency(OfferPaymentFrequency(tmp).String())
	return nil
}

type OfferSANInfoType string

const (
	OfferSANInfoTypeHdd = OfferSANInfoType("hdd")
	OfferSANInfoTypeSSD = OfferSANInfoType("ssd")
)

func (enum OfferSANInfoType) String() string {
	if enum == "" {
		// return default value if empty
		return string(OfferSANInfoTypeHdd)
	}
	return string(enum)
}

func (enum OfferSANInfoType) Values() []OfferSANInfoType {
	return []OfferSANInfoType{
		"hdd",
		"ssd",
	}
}

func (enum OfferSANInfoType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OfferSANInfoType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OfferSANInfoType(OfferSANInfoType(tmp).String())
	return nil
}

type OfferServerInfoStock string

const (
	OfferServerInfoStockEmpty     = OfferServerInfoStock("empty")
	OfferServerInfoStockLow       = OfferServerInfoStock("low")
	OfferServerInfoStockAvailable = OfferServerInfoStock("available")
)

func (enum OfferServerInfoStock) String() string {
	if enum == "" {
		// return default value if empty
		return string(OfferServerInfoStockEmpty)
	}
	return string(enum)
}

func (enum OfferServerInfoStock) Values() []OfferServerInfoStock {
	return []OfferServerInfoStock{
		"empty",
		"low",
		"available",
	}
}

func (enum OfferServerInfoStock) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OfferServerInfoStock) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OfferServerInfoStock(OfferServerInfoStock(tmp).String())
	return nil
}

type PartitionFileSystem string

const (
	PartitionFileSystemUnknown = PartitionFileSystem("unknown")
	PartitionFileSystemEfi     = PartitionFileSystem("efi")
	PartitionFileSystemSwap    = PartitionFileSystem("swap")
	PartitionFileSystemExt4    = PartitionFileSystem("ext4")
	PartitionFileSystemExt3    = PartitionFileSystem("ext3")
	PartitionFileSystemExt2    = PartitionFileSystem("ext2")
	PartitionFileSystemXfs     = PartitionFileSystem("xfs")
	PartitionFileSystemNtfs    = PartitionFileSystem("ntfs")
	PartitionFileSystemFat32   = PartitionFileSystem("fat32")
	PartitionFileSystemUfs     = PartitionFileSystem("ufs")
)

func (enum PartitionFileSystem) String() string {
	if enum == "" {
		// return default value if empty
		return string(PartitionFileSystemUnknown)
	}
	return string(enum)
}

func (enum PartitionFileSystem) Values() []PartitionFileSystem {
	return []PartitionFileSystem{
		"unknown",
		"efi",
		"swap",
		"ext4",
		"ext3",
		"ext2",
		"xfs",
		"ntfs",
		"fat32",
		"ufs",
	}
}

func (enum PartitionFileSystem) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PartitionFileSystem) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PartitionFileSystem(PartitionFileSystem(tmp).String())
	return nil
}

type PartitionType string

const (
	PartitionTypePrimary  = PartitionType("primary")
	PartitionTypeExtended = PartitionType("extended")
	PartitionTypeLogical  = PartitionType("logical")
)

func (enum PartitionType) String() string {
	if enum == "" {
		// return default value if empty
		return string(PartitionTypePrimary)
	}
	return string(enum)
}

func (enum PartitionType) Values() []PartitionType {
	return []PartitionType{
		"primary",
		"extended",
		"logical",
	}
}

func (enum PartitionType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PartitionType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PartitionType(PartitionType(tmp).String())
	return nil
}

type RaidArrayRaidLevel string

const (
	RaidArrayRaidLevelNoRaid = RaidArrayRaidLevel("no_raid")
	RaidArrayRaidLevelRaid0  = RaidArrayRaidLevel("raid0")
	RaidArrayRaidLevelRaid1  = RaidArrayRaidLevel("raid1")
	RaidArrayRaidLevelRaid5  = RaidArrayRaidLevel("raid5")
	RaidArrayRaidLevelRaid6  = RaidArrayRaidLevel("raid6")
	RaidArrayRaidLevelRaid10 = RaidArrayRaidLevel("raid10")
)

func (enum RaidArrayRaidLevel) String() string {
	if enum == "" {
		// return default value if empty
		return string(RaidArrayRaidLevelNoRaid)
	}
	return string(enum)
}

func (enum RaidArrayRaidLevel) Values() []RaidArrayRaidLevel {
	return []RaidArrayRaidLevel{
		"no_raid",
		"raid0",
		"raid1",
		"raid5",
		"raid6",
		"raid10",
	}
}

func (enum RaidArrayRaidLevel) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RaidArrayRaidLevel) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RaidArrayRaidLevel(RaidArrayRaidLevel(tmp).String())
	return nil
}

type RefundMethod string

const (
	RefundMethodUnknownRefundMethod = RefundMethod("unknown_refund_method")
	RefundMethodCreditCard          = RefundMethod("credit_card")
	RefundMethodAmex                = RefundMethod("amex")
	RefundMethodPaypal              = RefundMethod("paypal")
	RefundMethodTransfer            = RefundMethod("transfer")
)

func (enum RefundMethod) String() string {
	if enum == "" {
		// return default value if empty
		return string(RefundMethodUnknownRefundMethod)
	}
	return string(enum)
}

func (enum RefundMethod) Values() []RefundMethod {
	return []RefundMethod{
		"unknown_refund_method",
		"credit_card",
		"amex",
		"paypal",
		"transfer",
	}
}

func (enum RefundMethod) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RefundMethod) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RefundMethod(RefundMethod(tmp).String())
	return nil
}

type RefundStatus string

const (
	RefundStatusUnknownRefundStatus = RefundStatus("unknown_refund_status")
	RefundStatusUnpaid              = RefundStatus("unpaid")
	RefundStatusPaid                = RefundStatus("paid")
	RefundStatusErrored             = RefundStatus("errored")
)

func (enum RefundStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(RefundStatusUnknownRefundStatus)
	}
	return string(enum)
}

func (enum RefundStatus) Values() []RefundStatus {
	return []RefundStatus{
		"unknown_refund_status",
		"unpaid",
		"paid",
		"errored",
	}
}

func (enum RefundStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RefundStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RefundStatus(RefundStatus(tmp).String())
	return nil
}

type RescueProtocol string

const (
	RescueProtocolVnc = RescueProtocol("vnc")
	RescueProtocolSSH = RescueProtocol("ssh")
)

func (enum RescueProtocol) String() string {
	if enum == "" {
		// return default value if empty
		return string(RescueProtocolVnc)
	}
	return string(enum)
}

func (enum RescueProtocol) Values() []RescueProtocol {
	return []RescueProtocol{
		"vnc",
		"ssh",
	}
}

func (enum RescueProtocol) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RescueProtocol) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RescueProtocol(RescueProtocol(tmp).String())
	return nil
}

type RpnGroupMemberStatus string

const (
	RpnGroupMemberStatusUnknownRpnMemberStatus = RpnGroupMemberStatus("unknown_rpn_member_status")
	RpnGroupMemberStatusPendingInvitation      = RpnGroupMemberStatus("pending_invitation")
	RpnGroupMemberStatusActive                 = RpnGroupMemberStatus("active")
	RpnGroupMemberStatusCreating               = RpnGroupMemberStatus("creating")
	RpnGroupMemberStatusDeleting               = RpnGroupMemberStatus("deleting")
	RpnGroupMemberStatusDeleted                = RpnGroupMemberStatus("deleted")
)

func (enum RpnGroupMemberStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(RpnGroupMemberStatusUnknownRpnMemberStatus)
	}
	return string(enum)
}

func (enum RpnGroupMemberStatus) Values() []RpnGroupMemberStatus {
	return []RpnGroupMemberStatus{
		"unknown_rpn_member_status",
		"pending_invitation",
		"active",
		"creating",
		"deleting",
		"deleted",
	}
}

func (enum RpnGroupMemberStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RpnGroupMemberStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RpnGroupMemberStatus(RpnGroupMemberStatus(tmp).String())
	return nil
}

type RpnGroupType string

const (
	RpnGroupTypeUnknown = RpnGroupType("unknown")
	RpnGroupTypeLocal   = RpnGroupType("local")
	RpnGroupTypeShared  = RpnGroupType("shared")
)

func (enum RpnGroupType) String() string {
	if enum == "" {
		// return default value if empty
		return string(RpnGroupTypeUnknown)
	}
	return string(enum)
}

func (enum RpnGroupType) Values() []RpnGroupType {
	return []RpnGroupType{
		"unknown",
		"local",
		"shared",
	}
}

func (enum RpnGroupType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RpnGroupType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RpnGroupType(RpnGroupType(tmp).String())
	return nil
}

type RpnSanIPType string

const (
	RpnSanIPTypeUnknown     = RpnSanIPType("unknown")
	RpnSanIPTypeServerIP    = RpnSanIPType("server_ip")
	RpnSanIPTypeRpnv2Subnet = RpnSanIPType("rpnv2_subnet")
)

func (enum RpnSanIPType) String() string {
	if enum == "" {
		// return default value if empty
		return string(RpnSanIPTypeUnknown)
	}
	return string(enum)
}

func (enum RpnSanIPType) Values() []RpnSanIPType {
	return []RpnSanIPType{
		"unknown",
		"server_ip",
		"rpnv2_subnet",
	}
}

func (enum RpnSanIPType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RpnSanIPType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RpnSanIPType(RpnSanIPType(tmp).String())
	return nil
}

type RpnSanStatus string

const (
	RpnSanStatusUnknownStatus = RpnSanStatus("unknown_status")
	RpnSanStatusCreating      = RpnSanStatus("creating")
	RpnSanStatusActive        = RpnSanStatus("active")
	RpnSanStatusDeleting      = RpnSanStatus("deleting")
	RpnSanStatusMaintenance   = RpnSanStatus("maintenance")
)

func (enum RpnSanStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(RpnSanStatusUnknownStatus)
	}
	return string(enum)
}

func (enum RpnSanStatus) Values() []RpnSanStatus {
	return []RpnSanStatus{
		"unknown_status",
		"creating",
		"active",
		"deleting",
		"maintenance",
	}
}

func (enum RpnSanStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RpnSanStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RpnSanStatus(RpnSanStatus(tmp).String())
	return nil
}

type RpnV2GroupStatus string

const (
	RpnV2GroupStatusUnknownGroupStatus = RpnV2GroupStatus("unknown_group_status")
	RpnV2GroupStatusCreating           = RpnV2GroupStatus("creating")
	RpnV2GroupStatusActive             = RpnV2GroupStatus("active")
	RpnV2GroupStatusUpdating           = RpnV2GroupStatus("updating")
	RpnV2GroupStatusDeleting           = RpnV2GroupStatus("deleting")
)

func (enum RpnV2GroupStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(RpnV2GroupStatusUnknownGroupStatus)
	}
	return string(enum)
}

func (enum RpnV2GroupStatus) Values() []RpnV2GroupStatus {
	return []RpnV2GroupStatus{
		"unknown_group_status",
		"creating",
		"active",
		"updating",
		"deleting",
	}
}

func (enum RpnV2GroupStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RpnV2GroupStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RpnV2GroupStatus(RpnV2GroupStatus(tmp).String())
	return nil
}

type RpnV2GroupType string

const (
	RpnV2GroupTypeUnknownType = RpnV2GroupType("unknown_type")
	RpnV2GroupTypeStandard    = RpnV2GroupType("standard")
	RpnV2GroupTypeQinq        = RpnV2GroupType("qinq")
)

func (enum RpnV2GroupType) String() string {
	if enum == "" {
		// return default value if empty
		return string(RpnV2GroupTypeUnknownType)
	}
	return string(enum)
}

func (enum RpnV2GroupType) Values() []RpnV2GroupType {
	return []RpnV2GroupType{
		"unknown_type",
		"standard",
		"qinq",
	}
}

func (enum RpnV2GroupType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RpnV2GroupType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RpnV2GroupType(RpnV2GroupType(tmp).String())
	return nil
}

type RpnV2MemberStatus string

const (
	RpnV2MemberStatusUnknownMemberStatus = RpnV2MemberStatus("unknown_member_status")
	RpnV2MemberStatusCreating            = RpnV2MemberStatus("creating")
	RpnV2MemberStatusActive              = RpnV2MemberStatus("active")
	RpnV2MemberStatusUpdating            = RpnV2MemberStatus("updating")
	RpnV2MemberStatusDeleting            = RpnV2MemberStatus("deleting")
)

func (enum RpnV2MemberStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(RpnV2MemberStatusUnknownMemberStatus)
	}
	return string(enum)
}

func (enum RpnV2MemberStatus) Values() []RpnV2MemberStatus {
	return []RpnV2MemberStatus{
		"unknown_member_status",
		"creating",
		"active",
		"updating",
		"deleting",
	}
}

func (enum RpnV2MemberStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RpnV2MemberStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RpnV2MemberStatus(RpnV2MemberStatus(tmp).String())
	return nil
}

type ServerDiskType string

const (
	ServerDiskTypeSata = ServerDiskType("sata")
	ServerDiskTypeSSD  = ServerDiskType("ssd")
	ServerDiskTypeSas  = ServerDiskType("sas")
	ServerDiskTypeSshd = ServerDiskType("sshd")
	ServerDiskTypeUsb  = ServerDiskType("usb")
	ServerDiskTypeNvme = ServerDiskType("nvme")
)

func (enum ServerDiskType) String() string {
	if enum == "" {
		// return default value if empty
		return string(ServerDiskTypeSata)
	}
	return string(enum)
}

func (enum ServerDiskType) Values() []ServerDiskType {
	return []ServerDiskType{
		"sata",
		"ssd",
		"sas",
		"sshd",
		"usb",
		"nvme",
	}
}

func (enum ServerDiskType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerDiskType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerDiskType(ServerDiskType(tmp).String())
	return nil
}

type ServerInstallStatus string

const (
	ServerInstallStatusUnknown               = ServerInstallStatus("unknown")
	ServerInstallStatusBooting               = ServerInstallStatus("booting")
	ServerInstallStatusSettingUpRaid         = ServerInstallStatus("setting_up_raid")
	ServerInstallStatusPartitioning          = ServerInstallStatus("partitioning")
	ServerInstallStatusFormatting            = ServerInstallStatus("formatting")
	ServerInstallStatusInstalling            = ServerInstallStatus("installing")
	ServerInstallStatusConfiguring           = ServerInstallStatus("configuring")
	ServerInstallStatusConfiguringBootloader = ServerInstallStatus("configuring_bootloader")
	ServerInstallStatusRebooting             = ServerInstallStatus("rebooting")
	ServerInstallStatusInstalled             = ServerInstallStatus("installed")
)

func (enum ServerInstallStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(ServerInstallStatusUnknown)
	}
	return string(enum)
}

func (enum ServerInstallStatus) Values() []ServerInstallStatus {
	return []ServerInstallStatus{
		"unknown",
		"booting",
		"setting_up_raid",
		"partitioning",
		"formatting",
		"installing",
		"configuring",
		"configuring_bootloader",
		"rebooting",
		"installed",
	}
}

func (enum ServerInstallStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerInstallStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerInstallStatus(ServerInstallStatus(tmp).String())
	return nil
}

type ServerStatus string

const (
	ServerStatusUnknown    = ServerStatus("unknown")
	ServerStatusDelivering = ServerStatus("delivering")
	ServerStatusInstalling = ServerStatus("installing")
	ServerStatusReady      = ServerStatus("ready")
	ServerStatusStopped    = ServerStatus("stopped")
	ServerStatusError      = ServerStatus("error")
	ServerStatusLocked     = ServerStatus("locked")
	ServerStatusRescue     = ServerStatus("rescue")
	ServerStatusBusy       = ServerStatus("busy")
)

func (enum ServerStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(ServerStatusUnknown)
	}
	return string(enum)
}

func (enum ServerStatus) Values() []ServerStatus {
	return []ServerStatus{
		"unknown",
		"delivering",
		"installing",
		"ready",
		"stopped",
		"error",
		"locked",
		"rescue",
		"busy",
	}
}

func (enum ServerStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerStatus(ServerStatus(tmp).String())
	return nil
}

type ServiceLevelLevel string

const (
	ServiceLevelLevelUnknown  = ServiceLevelLevel("unknown")
	ServiceLevelLevelBasic    = ServiceLevelLevel("basic")
	ServiceLevelLevelBusiness = ServiceLevelLevel("business")
)

func (enum ServiceLevelLevel) String() string {
	if enum == "" {
		// return default value if empty
		return string(ServiceLevelLevelUnknown)
	}
	return string(enum)
}

func (enum ServiceLevelLevel) Values() []ServiceLevelLevel {
	return []ServiceLevelLevel{
		"unknown",
		"basic",
		"business",
	}
}

func (enum ServiceLevelLevel) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServiceLevelLevel) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServiceLevelLevel(ServiceLevelLevel(tmp).String())
	return nil
}

type ServiceProvisioningStatus string

const (
	ServiceProvisioningStatusUnknown    = ServiceProvisioningStatus("unknown")
	ServiceProvisioningStatusDelivering = ServiceProvisioningStatus("delivering")
	ServiceProvisioningStatusReady      = ServiceProvisioningStatus("ready")
	ServiceProvisioningStatusError      = ServiceProvisioningStatus("error")
	ServiceProvisioningStatusExpiring   = ServiceProvisioningStatus("expiring")
	ServiceProvisioningStatusExpired    = ServiceProvisioningStatus("expired")
)

func (enum ServiceProvisioningStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(ServiceProvisioningStatusUnknown)
	}
	return string(enum)
}

func (enum ServiceProvisioningStatus) Values() []ServiceProvisioningStatus {
	return []ServiceProvisioningStatus{
		"unknown",
		"delivering",
		"ready",
		"error",
		"expiring",
		"expired",
	}
}

func (enum ServiceProvisioningStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServiceProvisioningStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServiceProvisioningStatus(ServiceProvisioningStatus(tmp).String())
	return nil
}

type ServiceType string

const (
	ServiceTypeUnknownType = ServiceType("unknown_type")
	ServiceTypeService     = ServiceType("service")
	ServiceTypeOrder       = ServiceType("order")
)

func (enum ServiceType) String() string {
	if enum == "" {
		// return default value if empty
		return string(ServiceTypeUnknownType)
	}
	return string(enum)
}

func (enum ServiceType) Values() []ServiceType {
	return []ServiceType{
		"unknown_type",
		"service",
		"order",
	}
}

func (enum ServiceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServiceType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServiceType(ServiceType(tmp).String())
	return nil
}

// OfferAntiDosInfo: offer anti dos info.
type OfferAntiDosInfo struct {
	// Type: default value: minimal
	Type OfferAntiDosInfoType `json:"type"`
}

// OfferBackupInfo: offer backup info.
type OfferBackupInfo struct {
	Size scw.Size `json:"size"`
}

// OfferBandwidthInfo: offer bandwidth info.
type OfferBandwidthInfo struct {
	Speed uint32 `json:"speed"`
}

// OfferLicenseInfo: offer license info.
type OfferLicenseInfo struct {
	BoundToIP bool `json:"bound_to_ip"`
}

// OfferRPNInfo: offer rpn info.
type OfferRPNInfo struct {
	Speed uint32 `json:"speed"`
}

// OfferSANInfo: offer san info.
type OfferSANInfo struct {
	// Size: sAN size (in bytes).
	Size uint64 `json:"size"`

	// Ha: high availability offer.
	Ha bool `json:"ha"`

	// DeviceType: type of SAN device (hdd / ssd).
	// Default value: hdd
	DeviceType OfferSANInfoType `json:"device_type"`
}

// OfferStorageInfo: offer storage info.
type OfferStorageInfo struct {
	MaxQuota uint32 `json:"max_quota"`

	Size scw.Size `json:"size"`
}

// IP: ip.
type IP struct {
	// IPID: ID of the IP.
	IPID uint64 `json:"ip_id"`

	// Address: address of the IP.
	Address net.IP `json:"address"`

	// Reverse: reverse IP value.
	Reverse string `json:"reverse"`

	// Version: version of IP (v4 or v6).
	// Default value: ipv4
	Version IPVersion `json:"version"`

	// Cidr: classless InterDomain Routing notation of the IP.
	Cidr uint32 `json:"cidr"`

	// Netmask: network mask of IP.
	Netmask net.IP `json:"netmask"`

	// Semantic: semantic of IP.
	// Default value: unknown
	Semantic IPSemantic `json:"semantic"`

	// Gateway: gateway of IP.
	Gateway net.IP `json:"gateway"`

	// Status: status of the IP.
	// Default value: unknown_status
	Status IPStatus `json:"status"`
}

// Offer: offer.
type Offer struct {
	// ID: ID of the offer.
	ID uint64 `json:"id"`

	// Name: name of the offer.
	Name string `json:"name"`

	// Catalog: catalog of the offer.
	// Default value: all
	Catalog OfferCatalog `json:"catalog"`

	// PaymentFrequency: payment frequency of the offer.
	// Default value: monthly
	PaymentFrequency OfferPaymentFrequency `json:"payment_frequency"`

	// Pricing: price of the offer.
	Pricing *scw.Money `json:"pricing"`

	// ServerInfo: server info if it is a server offer.
	// Precisely one of ServerInfo, ServiceLevelInfo, RpnInfo, SanInfo, AntidosInfo, BackupInfo, UsbStorageInfo, StorageInfo, LicenseInfo, FailoverIPInfo, FailoverBlockInfo, BandwidthInfo must be set.
	ServerInfo *OfferServerInfo `json:"server_info,omitempty"`

	// ServiceLevelInfo: service level info if it is a service level offer.
	// Precisely one of ServerInfo, ServiceLevelInfo, RpnInfo, SanInfo, AntidosInfo, BackupInfo, UsbStorageInfo, StorageInfo, LicenseInfo, FailoverIPInfo, FailoverBlockInfo, BandwidthInfo must be set.
	ServiceLevelInfo *OfferServiceLevelInfo `json:"service_level_info,omitempty"`

	// RpnInfo: rPN info if it is a RPN offer.
	// Precisely one of ServerInfo, ServiceLevelInfo, RpnInfo, SanInfo, AntidosInfo, BackupInfo, UsbStorageInfo, StorageInfo, LicenseInfo, FailoverIPInfo, FailoverBlockInfo, BandwidthInfo must be set.
	RpnInfo *OfferRPNInfo `json:"rpn_info,omitempty"`

	// SanInfo: sAN info if it is a SAN offer.
	// Precisely one of ServerInfo, ServiceLevelInfo, RpnInfo, SanInfo, AntidosInfo, BackupInfo, UsbStorageInfo, StorageInfo, LicenseInfo, FailoverIPInfo, FailoverBlockInfo, BandwidthInfo must be set.
	SanInfo *OfferSANInfo `json:"san_info,omitempty"`

	// AntidosInfo: antiDOS info if it is a antiDOS offer.
	// Precisely one of ServerInfo, ServiceLevelInfo, RpnInfo, SanInfo, AntidosInfo, BackupInfo, UsbStorageInfo, StorageInfo, LicenseInfo, FailoverIPInfo, FailoverBlockInfo, BandwidthInfo must be set.
	AntidosInfo *OfferAntiDosInfo `json:"antidos_info,omitempty"`

	// BackupInfo: backup info if it is a backup offer.
	// Precisely one of ServerInfo, ServiceLevelInfo, RpnInfo, SanInfo, AntidosInfo, BackupInfo, UsbStorageInfo, StorageInfo, LicenseInfo, FailoverIPInfo, FailoverBlockInfo, BandwidthInfo must be set.
	BackupInfo *OfferBackupInfo `json:"backup_info,omitempty"`

	// UsbStorageInfo: uSB storage info if it is a USB storage offer.
	// Precisely one of ServerInfo, ServiceLevelInfo, RpnInfo, SanInfo, AntidosInfo, BackupInfo, UsbStorageInfo, StorageInfo, LicenseInfo, FailoverIPInfo, FailoverBlockInfo, BandwidthInfo must be set.
	UsbStorageInfo *OfferStorageInfo `json:"usb_storage_info,omitempty"`

	// StorageInfo: storage info if it is a storage offer.
	// Precisely one of ServerInfo, ServiceLevelInfo, RpnInfo, SanInfo, AntidosInfo, BackupInfo, UsbStorageInfo, StorageInfo, LicenseInfo, FailoverIPInfo, FailoverBlockInfo, BandwidthInfo must be set.
	StorageInfo *OfferStorageInfo `json:"storage_info,omitempty"`

	// LicenseInfo: license info if it is a license offer.
	// Precisely one of ServerInfo, ServiceLevelInfo, RpnInfo, SanInfo, AntidosInfo, BackupInfo, UsbStorageInfo, StorageInfo, LicenseInfo, FailoverIPInfo, FailoverBlockInfo, BandwidthInfo must be set.
	LicenseInfo *OfferLicenseInfo `json:"license_info,omitempty"`

	// FailoverIPInfo: failover IP info if it is a failover IP offer.
	// Precisely one of ServerInfo, ServiceLevelInfo, RpnInfo, SanInfo, AntidosInfo, BackupInfo, UsbStorageInfo, StorageInfo, LicenseInfo, FailoverIPInfo, FailoverBlockInfo, BandwidthInfo must be set.
	FailoverIPInfo *OfferFailoverIPInfo `json:"failover_ip_info,omitempty"`

	// FailoverBlockInfo: failover block info if it is a failover block offer.
	// Precisely one of ServerInfo, ServiceLevelInfo, RpnInfo, SanInfo, AntidosInfo, BackupInfo, UsbStorageInfo, StorageInfo, LicenseInfo, FailoverIPInfo, FailoverBlockInfo, BandwidthInfo must be set.
	FailoverBlockInfo *OfferFailoverBlockInfo `json:"failover_block_info,omitempty"`

	// BandwidthInfo: bandwidth info if it is a bandwidth offer.
	// Precisely one of ServerInfo, ServiceLevelInfo, RpnInfo, SanInfo, AntidosInfo, BackupInfo, UsbStorageInfo, StorageInfo, LicenseInfo, FailoverIPInfo, FailoverBlockInfo, BandwidthInfo must be set.
	BandwidthInfo *OfferBandwidthInfo `json:"bandwidth_info,omitempty"`
}

// NetworkInterface: network interface.
type NetworkInterface struct {
	// CardID: card ID of the network interface.
	CardID uint64 `json:"card_id"`

	// DeviceID: device ID of the network interface.
	DeviceID uint64 `json:"device_id"`

	// Mac: mAC address of the network interface.
	Mac string `json:"mac"`

	// Type: network interface type.
	// Default value: unknown
	Type NetworkInterfaceInterfaceType `json:"type"`

	// IPs: iPs of the network interface.
	IPs []*IP `json:"ips"`
}

// OS: os.
type OS struct {
	// ID: ID of the OS.
	ID uint64 `json:"id"`

	// Name: name of the OS.
	Name string `json:"name"`

	// Type: type of the OS.
	// Default value: unknown_type
	Type OSType `json:"type"`

	// Version: version of the OS.
	Version string `json:"version"`

	// Arch: architecture of the OS.
	// Default value: unknown_arch
	Arch OSArch `json:"arch"`

	// AllowCustomPartitioning: true if the OS allow custom partitioning.
	AllowCustomPartitioning bool `json:"allow_custom_partitioning"`

	// AllowSSHKeys: true if the OS allow SSH Keys.
	AllowSSHKeys bool `json:"allow_ssh_keys"`

	// RequiresUser: true if the OS requires user.
	RequiresUser bool `json:"requires_user"`

	// RequiresAdminPassword: true if the OS requires admin password.
	RequiresAdminPassword bool `json:"requires_admin_password"`

	// RequiresPanelPassword: true if the OS requires panel password.
	RequiresPanelPassword bool `json:"requires_panel_password"`

	// AllowedFilesystems: true if the OS allow file systems.
	AllowedFilesystems []PartitionFileSystem `json:"allowed_filesystems"`

	// RequiresLicense: true if the OS requires license.
	RequiresLicense bool `json:"requires_license"`

	// LicenseOffers: license offers available with the OS.
	LicenseOffers []*Offer `json:"license_offers"`

	// MaxPartitions: maximum number of partitions which can be created.
	MaxPartitions *uint32 `json:"max_partitions"`

	// DisplayName: display name of the OS.
	DisplayName string `json:"display_name"`

	// PasswordRegex: regex used to validate the installation passwords.
	PasswordRegex string `json:"password_regex"`

	// PanelPasswordRegex: regex used to validate the panel installation password.
	PanelPasswordRegex *string `json:"panel_password_regex"`

	// RequiresValidHostname: if both requires_valid_hostname & hostname_regex are set, it means that at least one of the criteria must be valid.
	RequiresValidHostname *bool `json:"requires_valid_hostname"`

	// HostnameRegex: if both requires_valid_hostname & hostname_regex are set, it means that at least one of the criteria must be valid.
	HostnameRegex *string `json:"hostname_regex"`

	// HostnameMaxLength: hostname max length.
	HostnameMaxLength uint32 `json:"hostname_max_length"`

	// ReleasedAt: oS release date.
	ReleasedAt *time.Time `json:"released_at"`
}

// ServerLocation: server location.
type ServerLocation struct {
	Rack string `json:"rack"`

	Room string `json:"room"`

	DatacenterName string `json:"datacenter_name"`
}

// ServerOption: server option.
type ServerOption struct {
	Offer *Offer `json:"offer"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	ExpiredAt *time.Time `json:"expired_at"`

	Options []*ServerOption `json:"options"`
}

// ServiceLevel: service level.
type ServiceLevel struct {
	// OfferID: offer ID of service level.
	OfferID uint32 `json:"offer_id"`

	// Level: level type of service level.
	// Default value: unknown
	Level ServiceLevelLevel `json:"level"`
}

// RpnSan: rpn san.
type RpnSan struct {
	// ID: rPN SAN  ID.
	ID uint64 `json:"id"`

	// DatacenterName: datacenter location.
	DatacenterName string `json:"datacenter_name"`

	// OrganizationID: organization ID.
	OrganizationID string `json:"organization_id"`

	// ProjectID: project ID.
	ProjectID string `json:"project_id"`

	// ServerHostname: rPN SAN server hostname.
	ServerHostname string `json:"server_hostname"`

	// IqnSuffix: iQN suffix.
	IqnSuffix string `json:"iqn_suffix"`

	// CreatedAt: date of creation of the RPN SAN.
	CreatedAt *time.Time `json:"created_at"`

	// OfferID: offer ID.
	OfferID uint64 `json:"offer_id"`

	// OfferName: offer description.
	OfferName string `json:"offer_name"`

	// Status: status.
	// Default value: unknown_status
	Status RpnSanStatus `json:"status"`

	// StorageSize: rPN SAN storage size.
	StorageSize scw.Size `json:"storage_size"`

	Iqn string `json:"iqn"`

	Offer *Offer `json:"offer"`

	// Rpnv1Compatible: true if the SAN is compatible with the RPNv1 technology.
	Rpnv1Compatible bool `json:"rpnv1_compatible"`

	// Rpnv1Implicit: true if the offer supports the RPNv1 implicitly, false if it must to be added to a group to support RPNv1.
	Rpnv1Implicit bool `json:"rpnv1_implicit"`

	// DeliveredAt: rPN SAN delivery date.
	DeliveredAt *time.Time `json:"delivered_at"`

	// TerminatedAt: rPN SAN termination date.
	TerminatedAt *time.Time `json:"terminated_at"`

	// ExpiresAt: rPN SAN expiration date.
	ExpiresAt *time.Time `json:"expires_at"`
}

// RpnGroup: rpn group.
type RpnGroup struct {
	// ID: rpn group member ID.
	ID uint64 `json:"id"`

	// Name: rpn group name.
	Name string `json:"name"`

	// Type: rpn group type (local or shared).
	// Default value: unknown
	Type RpnGroupType `json:"type"`

	// Active: whether the group is active or not.
	Active bool `json:"active"`

	// CreatedAt: rpn group creation date.
	CreatedAt *time.Time `json:"created_at"`

	// Owner: rPN group owner.
	Owner string `json:"owner"`

	// MembersCount: total number of members.
	MembersCount uint32 `json:"members_count"`

	// OrganizationID: rpn group organization ID.
	OrganizationID string `json:"organization_id"`

	// ProjectID: rpn group project ID.
	ProjectID string `json:"project_id"`
}

// RpnV2GroupSubnet: rpn v2 group subnet.
type RpnV2GroupSubnet struct {
	Address net.IP `json:"address"`

	Cidr uint32 `json:"cidr"`
}

// Server: server.
type Server struct {
	// ID: ID of the server.
	ID uint64 `json:"id"`

	// OrganizationID: organization ID the server is attached to.
	OrganizationID string `json:"organization_id"`

	// ProjectID: project ID the server is attached to.
	ProjectID string `json:"project_id"`

	// Hostname: hostname of the server.
	Hostname string `json:"hostname"`

	// RebootedAt: date of last reboot of the server.
	RebootedAt *time.Time `json:"rebooted_at"`

	// CreatedAt: date of creation of the server.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date of last modification of the server.
	UpdatedAt *time.Time `json:"updated_at"`

	// ExpiredAt: date of release of the server.
	ExpiredAt *time.Time `json:"expired_at"`

	// Offer: offer of the server.
	Offer *Offer `json:"offer"`

	// Status: status of the server.
	// Default value: unknown
	Status ServerStatus `json:"status"`

	// Location: location of the server.
	Location *ServerLocation `json:"location"`

	// AbuseContact: abuse contact of the server.
	AbuseContact string `json:"abuse_contact"`

	// Os: oS installed on the server.
	Os *OS `json:"os"`

	// Interfaces: network interfaces of the server.
	Interfaces []*NetworkInterface `json:"interfaces"`

	// Zone: the zone in which is the server.
	Zone scw.Zone `json:"zone"`

	// Options: options subscribe on the server.
	Options []*ServerOption `json:"options"`

	// Level: service level of the server.
	Level *ServiceLevel `json:"level"`

	// HasBmc: boolean if the server has a BMC.
	HasBmc bool `json:"has_bmc"`

	// RescueOs: rescue OS of the server.
	RescueOs *OS `json:"rescue_os"`

	// Tags: array of customs tags attached to the server.
	Tags []string `json:"tags"`

	// IsOutsourced: whether the server is outsourced or not.
	IsOutsourced bool `json:"is_outsourced"`

	// IPv6Slaac: whether or not you can enable/disable the IPv6.
	IPv6Slaac bool `json:"ipv6_slaac"`

	// Qinq: whether the server is compatible with QinQ.
	Qinq bool `json:"qinq"`

	// IsRpnv2Member: whether or not the server is already part of an rpnv2 group.
	IsRpnv2Member bool `json:"is_rpnv2_member"`

	// IsHds: whether or not the server is HDS.
	IsHds bool `json:"is_hds"`
}

// FailoverBlock: failover block.
type FailoverBlock struct {
	// ID: ID of the failover block.
	ID uint64 `json:"id"`

	// Address: IP of the failover block.
	Address net.IP `json:"address"`

	// Nameservers: name servers.
	Nameservers []string `json:"nameservers"`

	// IPVersion: IP version of the failover block.
	// Default value: unknown_version
	IPVersion FailoverBlockVersion `json:"ip_version"`

	// Cidr: classless InterDomain Routing notation of the failover block.
	Cidr uint32 `json:"cidr"`

	// Netmask: netmask of the failover block.
	Netmask net.IP `json:"netmask"`

	// GatewayIP: gateway IP of the failover block.
	GatewayIP net.IP `json:"gateway_ip"`
}

// RpnSanIPRpnV2Group: rpn san ip rpn v2 group.
type RpnSanIPRpnV2Group struct {
	ID uint64 `json:"id"`

	Name string `json:"name"`
}

// RpnSanIPServer: rpn san ip server.
type RpnSanIPServer struct {
	ID uint64 `json:"id"`

	Hostname string `json:"hostname"`

	DatacenterName string `json:"datacenter_name"`
}

// RpnSanServer: rpn san server.
type RpnSanServer struct {
	// ID: the RPN SAN server ID.
	ID uint64 `json:"id"`

	// DatacenterName: the RPN SAN server datacenter name.
	DatacenterName string `json:"datacenter_name"`

	// Hostname: the RPN SAN server hostname.
	Hostname string `json:"hostname"`

	// Sans: rPN SANs linked to the RPN SAN server.
	Sans []*RpnSan `json:"sans"`

	// Zone: the RPN SAN server zone.
	Zone scw.Zone `json:"zone"`
}

// RpnV2Group: rpn v2 group.
type RpnV2Group struct {
	// ID: rPN V2 group ID.
	ID uint64 `json:"id"`

	// Name: rPN V2 group name.
	Name string `json:"name"`

	// CompatibleRpnv1: whether or not the RPN V1 compatibility was enabled.
	CompatibleRpnv1 bool `json:"compatible_rpnv1"`

	// OrganizationID: organization ID of the RPN V2 group.
	OrganizationID string `json:"organization_id"`

	// ProjectID: project ID of the RPN V2 group.
	ProjectID string `json:"project_id"`

	// Type: rPN V2 group type (qing / standard).
	// Default value: unknown_type
	Type RpnV2GroupType `json:"type"`

	// Status: rPN V2 group status.
	// Default value: unknown_group_status
	Status RpnV2GroupStatus `json:"status"`

	// Owner: rPN V2 group owner.
	Owner string `json:"owner"`

	// MembersCount: total number of members.
	MembersCount uint32 `json:"members_count"`

	// Subnet: rPN V2 subnet.
	Subnet *RpnV2GroupSubnet `json:"subnet"`

	// Gateway: rPN V2 gateway.
	Gateway net.IP `json:"gateway"`

	// Rpnv1Group: the RPNv1 group (if the compatibility was enabled).
	Rpnv1Group *RpnGroup `json:"rpnv1_group"`
}

// RpnV2Member: rpn v2 member.
type RpnV2Member struct {
	// ID: rPN V2 member ID.
	ID uint64 `json:"id"`

	// Status: rPN V2 member status.
	// Default value: unknown_member_status
	Status RpnV2MemberStatus `json:"status"`

	// Vlan: rPN V2 member VLAN.
	Vlan string `json:"vlan"`

	// Server: server behind the member (may be empty).
	// Precisely one of Server, Rpnv1Group must be set.
	Server *Server `json:"server,omitempty"`

	// Rpnv1Group: rPN V1 group member.
	// Precisely one of Server, Rpnv1Group must be set.
	Rpnv1Group *RpnGroup `json:"rpnv1_group,omitempty"`

	// Speed: rPN speed.
	Speed *uint32 `json:"speed"`
}

// ServerDisk: server disk.
type ServerDisk struct {
	ID uint64 `json:"id"`

	Connector string `json:"connector"`

	// Type: default value: sata
	Type ServerDiskType `json:"type"`

	Capacity scw.Size `json:"capacity"`

	IsAddon bool `json:"is_addon"`
}

// Service: service.
type Service struct {
	// ID: ID of the service.
	ID uint64 `json:"id"`

	// ResourceID: resource ID of the service.
	ResourceID *uint64 `json:"resource_id"`

	// ProvisioningStatus: provisioning status of the service.
	// Default value: unknown
	ProvisioningStatus ServiceProvisioningStatus `json:"provisioning_status"`

	// Offer: offer of the service.
	Offer *Offer `json:"offer"`

	// CreatedAt: creation date of the service.
	CreatedAt *time.Time `json:"created_at"`

	// DeliveredAt: delivery date of the service.
	DeliveredAt *time.Time `json:"delivered_at"`

	// TerminatedAt: terminatation date of the service.
	TerminatedAt *time.Time `json:"terminated_at"`

	// ExpiresAt: expiration date of the service.
	ExpiresAt *time.Time `json:"expires_at"`

	// Type: service type, either order or service.
	// Default value: unknown_type
	Type ServiceType `json:"type"`
}

// GetIPv6BlockQuotasResponseQuota: get i pv6 block quotas response quota.
type GetIPv6BlockQuotasResponseQuota struct {
	Quota uint64 `json:"quota"`

	Cidr uint32 `json:"cidr"`
}

// InstallPartition: install partition.
type InstallPartition struct {
	// FileSystem: file system of the installation partition.
	// Default value: unknown
	FileSystem PartitionFileSystem `json:"file_system"`

	// MountPoint: mount point of the installation partition.
	MountPoint *string `json:"mount_point"`

	// RaidLevel: rAID level of the installation partition.
	// Default value: no_raid
	RaidLevel RaidArrayRaidLevel `json:"raid_level"`

	// Capacity: capacity of the installation partition.
	Capacity scw.Size `json:"capacity"`

	// Connectors: connectors of the installation partition.
	Connectors []string `json:"connectors"`
}

// FailoverIP: failover ip.
type FailoverIP struct {
	// ID: ID of the failover IP.
	ID uint64 `json:"id"`

	// Address: IP of the failover IP.
	Address net.IP `json:"address"`

	// Reverse: reverse IP value.
	Reverse string `json:"reverse"`

	// IPVersion: IP version of the failover IP.
	// Default value: unknown_version
	IPVersion FailoverIPVersion `json:"ip_version"`

	// Cidr: classless InterDomain Routing notation of the failover IP.
	Cidr uint32 `json:"cidr"`

	// Netmask: netmask of the failover IP.
	Netmask net.IP `json:"netmask"`

	// GatewayIP: gateway IP of the failover IP.
	GatewayIP net.IP `json:"gateway_ip"`

	// Mac: mAC address of the IP failover.
	Mac *string `json:"mac"`

	// ServerID: server ID linked to the IP failover.
	ServerID *uint64 `json:"server_id"`

	// Status: status of the IP failover.
	// Default value: unknown_status
	Status FailoverIPStatus `json:"status"`

	// Block: block of the IP failover.
	Block *FailoverBlock `json:"block"`

	// Type: the interface type.
	// Default value: unknown
	Type FailoverIPInterfaceType `json:"type"`

	// ServerZone: the server zone (if assigned).
	ServerZone *string `json:"server_zone"`
}

// RpnSanIP: rpn san ip.
type RpnSanIP struct {
	// Server: basic server information behind the IP.
	// Precisely one of Server, Rpnv2Group must be set.
	Server *RpnSanIPServer `json:"server,omitempty"`

	// Rpnv2Group: basic RPNv2 group information behind the IP.
	// Precisely one of Server, Rpnv2Group must be set.
	Rpnv2Group *RpnSanIPRpnV2Group `json:"rpnv2_group,omitempty"`

	// IP: an IP object.
	IP *IP `json:"ip"`

	// Type: IP type (server | rpnv2_subnet).
	// Default value: unknown
	Type RpnSanIPType `json:"type"`
}

// ListIPv6BlockSubnetsAvailableResponseSubnet: list i pv6 block subnets available response subnet.
type ListIPv6BlockSubnetsAvailableResponseSubnet struct {
	Address net.IP `json:"address"`

	Cidr uint32 `json:"cidr"`
}

// IPv6Block: i pv6 block.
type IPv6Block struct {
	// ID: ID of the IPv6.
	ID uint64 `json:"id"`

	// Address: address of the IPv6.
	Address net.IP `json:"address"`

	// Duid: dUID of the IPv6.
	Duid string `json:"duid"`

	// Nameservers: DNS linked to the IPv6.
	Nameservers []string `json:"nameservers"`

	// Cidr: classless InterDomain Routing notation of the IPv6.
	Cidr uint32 `json:"cidr"`

	// Subnets: all IPv6 subnets.
	Subnets []*IPv6Block `json:"subnets"`

	// DelegationStatus: the nameservers delegation status.
	// Default value: unknown_status
	DelegationStatus IPv6BlockDelegationStatus `json:"delegation_status"`
}

// InvoiceSummary: invoice summary.
type InvoiceSummary struct {
	ID uint64 `json:"id"`

	TotalWithTaxes *scw.Money `json:"total_with_taxes"`

	TotalWithoutTaxes *scw.Money `json:"total_without_taxes"`

	CreatedAt *time.Time `json:"created_at"`

	PaidAt *time.Time `json:"paid_at"`

	// Status: default value: unknown_invoice_status
	Status InvoiceStatus `json:"status"`

	// PaymentMethod: default value: unknown_payment_method
	PaymentMethod InvoicePaymentMethod `json:"payment_method"`

	TransactionID uint64 `json:"transaction_id"`
}

// RefundSummary: refund summary.
type RefundSummary struct {
	ID uint64 `json:"id"`

	TotalWithTaxes *scw.Money `json:"total_with_taxes"`

	TotalWithoutTaxes *scw.Money `json:"total_without_taxes"`

	CreatedAt *time.Time `json:"created_at"`

	RefundedAt *time.Time `json:"refunded_at"`

	// Status: default value: unknown_refund_status
	Status RefundStatus `json:"status"`

	// Method: default value: unknown_refund_method
	Method RefundMethod `json:"method"`
}

// RpnGroupMember: rpn group member.
type RpnGroupMember struct {
	// ID: rpn group member ID.
	ID uint64 `json:"id"`

	// Status: rPN group member status.
	// Default value: unknown_rpn_member_status
	Status RpnGroupMemberStatus `json:"status"`

	// SanServer: authorized RPN SAN server.
	SanServer *RpnSanServer `json:"san_server"`

	// Server: authorized rpn v1 capable server.
	Server *Server `json:"server"`

	// GroupID: rPN group ID.
	GroupID uint64 `json:"group_id"`

	// GroupName: rPN group name.
	GroupName string `json:"group_name"`

	// GroupOwner: rPN group owner.
	GroupOwner string `json:"group_owner"`

	// Speed: rPN speed.
	Speed *uint32 `json:"speed"`

	// Owner: rPN member owner.
	Owner string `json:"owner"`
}

// RpnSanSummary: rpn san summary.
type RpnSanSummary struct {
	// ID: rPN SAN  ID.
	ID uint64 `json:"id"`

	// DatacenterName: datacenter location.
	DatacenterName string `json:"datacenter_name"`

	// OrganizationID: organization ID.
	OrganizationID string `json:"organization_id"`

	// ProjectID: project ID.
	ProjectID string `json:"project_id"`

	// ServerHostname: rPN SAN server hostname.
	ServerHostname string `json:"server_hostname"`

	// IqnSuffix: iQN suffix.
	IqnSuffix string `json:"iqn_suffix"`

	// CreatedAt: date of creation of the RPN SAN.
	CreatedAt *time.Time `json:"created_at"`

	// OfferID: offer ID.
	OfferID uint64 `json:"offer_id"`

	// OfferName: offer description.
	OfferName string `json:"offer_name"`

	// Status: status.
	// Default value: unknown_status
	Status RpnSanStatus `json:"status"`

	// StorageSize: rPN SAN storage size.
	StorageSize scw.Size `json:"storage_size"`

	// Rpnv1Compatible: true if the SAN is compatible with the RPNv1 technology.
	Rpnv1Compatible bool `json:"rpnv1_compatible"`

	// Rpnv1Implicit: true if the offer supports the RPNv1 implicitly, false if it must to be added to a group to support RPNv1.
	Rpnv1Implicit bool `json:"rpnv1_implicit"`

	// DeliveredAt: rPN SAN delivery date.
	DeliveredAt *time.Time `json:"delivered_at"`

	// TerminatedAt: rPN SAN termination date.
	TerminatedAt *time.Time `json:"terminated_at"`

	// ExpiresAt: rPN SAN expiration date.
	ExpiresAt *time.Time `json:"expires_at"`
}

// RpnServerCapability: rpn server capability.
type RpnServerCapability struct {
	// ID: server ID.
	ID uint64 `json:"id"`

	// Hostname: server hostname.
	Hostname string `json:"hostname"`

	// DatacenterName: server datacenter name.
	DatacenterName string `json:"datacenter_name"`

	// Zone: server zone.
	Zone scw.Zone `json:"zone"`

	// IPAddress: private IP address (if rpn compatiblle).
	IPAddress *net.IP `json:"ip_address"`

	// RpnVersion: supported rpn version.
	RpnVersion *uint32 `json:"rpn_version"`

	// CompatibleQinq: true if server is compatible with QinQ protocol (rpn v2).
	CompatibleQinq bool `json:"compatible_qinq"`

	// CanJoinQinqGroup: true if server can join a QinQ group.
	CanJoinQinqGroup bool `json:"can_join_qinq_group"`

	// Rpnv1GroupCount: times server is linked in a rpnv1 group.
	Rpnv1GroupCount uint32 `json:"rpnv1_group_count"`

	// Rpnv2GroupCount: times server is linked in a rpnv2 group.
	Rpnv2GroupCount uint32 `json:"rpnv2_group_count"`

	// CanJoinRpnv2Group: true if server can join an rpnv2 group.
	CanJoinRpnv2Group bool `json:"can_join_rpnv2_group"`
}

// Log: log.
type Log struct {
	// ID: rPN V2 log ID.
	ID uint64 `json:"id"`

	// Group: rPN V2 group.
	Group *RpnV2Group `json:"group"`

	// Member: rPN V2 member (if applicable).
	Member *RpnV2Member `json:"member"`

	// Action: which action was performed.
	// Default value: unknown_log_action
	Action LogAction `json:"action"`

	// Status: action status.
	// Default value: unknown_log_status
	Status LogStatus `json:"status"`

	// CreatedAt: creation date.
	CreatedAt *time.Time `json:"created_at"`

	// FinishedAt: completion date.
	FinishedAt *time.Time `json:"finished_at"`
}

// ServerEvent: server event.
type ServerEvent struct {
	// EventID: ID of the event.
	EventID uint64 `json:"event_id"`

	// Description: description of the event.
	Description string `json:"description"`

	// Date: date of the event.
	Date *time.Time `json:"date"`
}

// ServerSummary: server summary.
type ServerSummary struct {
	// ID: ID of the server.
	ID uint64 `json:"id"`

	// DatacenterName: datacenter of the server.
	DatacenterName string `json:"datacenter_name"`

	// OrganizationID: organization ID the server is attached to.
	OrganizationID string `json:"organization_id"`

	// ProjectID: project ID the server is attached to.
	ProjectID string `json:"project_id"`

	// Hostname: hostname of the server.
	Hostname string `json:"hostname"`

	// CreatedAt: date of creation of the server.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date of last modification of the server.
	UpdatedAt *time.Time `json:"updated_at"`

	// ExpiredAt: date of release of the server.
	ExpiredAt *time.Time `json:"expired_at"`

	// OfferID: offer ID of the server.
	OfferID uint64 `json:"offer_id"`

	// OfferName: offer name of the server.
	OfferName string `json:"offer_name"`

	// Status: status of the server.
	// Default value: unknown
	Status ServerStatus `json:"status"`

	// OsID: oS ID installed on server.
	OsID *uint64 `json:"os_id"`

	// Interfaces: network interfaces of the server.
	Interfaces []*NetworkInterface `json:"interfaces"`

	// Zone: the zone in which is the server.
	Zone scw.Zone `json:"zone"`

	// Level: service level of the server.
	Level *ServiceLevel `json:"level"`

	// IsOutsourced: whether the server is outsourced or not.
	IsOutsourced bool `json:"is_outsourced"`

	// Qinq: whether the server is compatible with QinQ.
	Qinq bool `json:"qinq"`

	// RpnVersion: supported RPN version.
	RpnVersion *uint32 `json:"rpn_version"`

	// IsHds: whether or not the server is HDS.
	IsHds bool `json:"is_hds"`
}

// CPU: cpu.
type CPU struct {
	// Name: name of CPU.
	Name string `json:"name"`

	// CoreCount: number of cores of the CPU.
	CoreCount uint32 `json:"core_count"`

	// ThreadCount: number of threads of the CPU.
	ThreadCount uint32 `json:"thread_count"`

	// Frequency: frequency of the CPU.
	Frequency uint32 `json:"frequency"`
}

// Disk: disk.
type Disk struct {
	// Capacity: capacity of the disk.
	Capacity scw.Size `json:"capacity"`

	// Type: type of the disk.
	// Default value: sata
	Type ServerDiskType `json:"type"`
}

// Memory: memory.
type Memory struct {
	// Capacity: capacity of the memory.
	Capacity scw.Size `json:"capacity"`

	// Type: type of the memory.
	// Default value: ddr2
	Type MemoryType `json:"type"`

	// Frequency: frequency of the memory.
	Frequency uint32 `json:"frequency"`

	// IsEcc: true if the memory is an error-correcting code memory.
	IsEcc bool `json:"is_ecc"`
}

// PersistentMemory: persistent memory.
type PersistentMemory struct {
	// Capacity: capacity of the persistent memory.
	Capacity scw.Size `json:"capacity"`

	// Frequency: frequency of the persistent memory.
	Frequency uint32 `json:"frequency"`

	// Model: model of the persistent memory.
	Model string `json:"model"`
}

// RaidController: raid controller.
type RaidController struct {
	// Model: model of the RAID controller.
	Model string `json:"model"`

	// RaidLevel: rAID level of the RAID controller.
	RaidLevel []string `json:"raid_level"`
}

// RaidArray: raid array.
type RaidArray struct {
	// RaidLevel: the RAID level.
	// Default value: no_raid
	RaidLevel RaidArrayRaidLevel `json:"raid_level"`

	// Disks: disks on the RAID controller.
	Disks []*ServerDisk `json:"disks"`
}

// Partition: partition.
type Partition struct {
	// Type: type of the partition.
	// Default value: primary
	Type PartitionType `json:"type"`

	// FileSystem: file system of the partition.
	// Default value: unknown
	FileSystem PartitionFileSystem `json:"file_system"`

	// MountPoint: mount point of the partition.
	MountPoint *string `json:"mount_point"`

	// RaidLevel: raid level of the partition.
	// Default value: no_raid
	RaidLevel RaidArrayRaidLevel `json:"raid_level"`

	// Capacity: capacity of the partition.
	Capacity scw.Size `json:"capacity"`

	// Connectors: connectors of the partition.
	Connectors []string `json:"connectors"`
}

// UpdatableRaidArray: updatable raid array.
type UpdatableRaidArray struct {
	// RaidLevel: the RAID level.
	// Default value: no_raid
	RaidLevel RaidArrayRaidLevel `json:"raid_level"`

	// DiskIDs: the list of Disk ID of the updatable RAID.
	DiskIDs []uint64 `json:"disk_ids"`
}

// AttachFailoverIPToMacAddressRequest: attach failover ip to mac address request.
type AttachFailoverIPToMacAddressRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// IPID: ID of the failover IP.
	IPID uint64 `json:"-"`

	// Type: a mac type.
	// Default value: mac_type_unknown
	Type AttachFailoverIPToMacAddressRequestMacType `json:"type"`

	// Mac: a valid mac address (existing or not).
	Mac *string `json:"mac,omitempty"`
}

// AttachFailoverIPsRequest: attach failover i ps request.
type AttachFailoverIPsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server.
	ServerID uint64 `json:"server_id"`

	// FipsIDs: list of ID of failovers IP to attach.
	FipsIDs []uint64 `json:"fips_ids"`
}

// BMCAccess: bmc access.
type BMCAccess struct {
	// URL: URL to access to the server console.
	URL string `json:"url"`

	// Login: the login to use for the BMC (Baseboard Management Controller) access authentication.
	Login string `json:"login"`

	// Password: the password to use for the BMC (Baseboard Management Controller) access authentication.
	Password string `json:"password"`

	// ExpiresAt: the date after which the BMC (Baseboard Management Controller) access will be closed.
	ExpiresAt *time.Time `json:"expires_at"`

	// Status: status of the connection.
	// Default value: unknown
	Status BMCAccessStatus `json:"status"`
}

// Backup: backup.
type Backup struct {
	// ID: ID of the backup.
	ID uint64 `json:"id"`

	// Login: login of the backup.
	Login string `json:"login"`

	// Server: server of the backup.
	Server string `json:"server"`

	// Status: status of the backup.
	// Default value: unknown_backup_status
	Status BackupStatus `json:"status"`

	// ACLEnabled: ACL enable boolean of the backup.
	ACLEnabled bool `json:"acl_enabled"`

	// Autologin: autologin boolean of the backup.
	Autologin bool `json:"autologin"`

	// QuotaSpace: total quota space of the backup.
	QuotaSpace uint64 `json:"quota_space"`

	// QuotaSpaceUsed: quota space used of the backup.
	QuotaSpaceUsed uint64 `json:"quota_space_used"`

	// QuotaFiles: total quota files of the backup.
	QuotaFiles uint64 `json:"quota_files"`

	// QuotaFilesUsed: quota files used of the backup.
	QuotaFilesUsed uint64 `json:"quota_files_used"`
}

// BillingAPICanOrderRequest: billing api can order request.
type BillingAPICanOrderRequest struct {
	ProjectID string `json:"-"`
}

// BillingAPIDownloadInvoiceRequest: billing api download invoice request.
type BillingAPIDownloadInvoiceRequest struct {
	InvoiceID uint64 `json:"-"`
}

// BillingAPIDownloadRefundRequest: billing api download refund request.
type BillingAPIDownloadRefundRequest struct {
	RefundID uint64 `json:"-"`
}

// BillingAPIGetInvoiceRequest: billing api get invoice request.
type BillingAPIGetInvoiceRequest struct {
	InvoiceID uint64 `json:"-"`
}

// BillingAPIGetRefundRequest: billing api get refund request.
type BillingAPIGetRefundRequest struct {
	RefundID uint64 `json:"-"`
}

// BillingAPIListInvoicesRequest: billing api list invoices request.
type BillingAPIListInvoicesRequest struct {
	Page *uint32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListInvoicesRequestOrderBy `json:"-"`

	ProjectID *string `json:"-"`
}

// BillingAPIListRefundsRequest: billing api list refunds request.
type BillingAPIListRefundsRequest struct {
	Page *uint32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListRefundsRequestOrderBy `json:"-"`

	ProjectID *string `json:"-"`
}

// CanOrderResponse: can order response.
type CanOrderResponse struct {
	CanOrder bool `json:"can_order"`

	Message *string `json:"message"`

	QuotaOk bool `json:"quota_ok"`

	PhoneConfirmed bool `json:"phone_confirmed"`

	EmailConfirmed bool `json:"email_confirmed"`

	UserConfirmed bool `json:"user_confirmed"`

	PaymentMode bool `json:"payment_mode"`

	BillingOk bool `json:"billing_ok"`
}

// CancelServerInstallRequest: cancel server install request.
type CancelServerInstallRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: server ID of the server to cancel install.
	ServerID uint64 `json:"-"`
}

// CreateFailoverIPsRequest: create failover i ps request.
type CreateFailoverIPsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OfferID: failover IP offer ID.
	OfferID uint64 `json:"offer_id"`

	// ProjectID: project ID.
	ProjectID string `json:"project_id"`

	// Quantity: quantity.
	Quantity uint32 `json:"quantity"`
}

// CreateFailoverIPsResponse: create failover i ps response.
type CreateFailoverIPsResponse struct {
	TotalCount uint32 `json:"total_count"`

	Services []*Service `json:"services"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *CreateFailoverIPsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *CreateFailoverIPsResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*CreateFailoverIPsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Services = append(r.Services, results.Services...)
	r.TotalCount += uint32(len(results.Services))
	return uint32(len(results.Services)), nil
}

// CreateServerRequest: create server request.
type CreateServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OfferID: offer ID of the new server.
	OfferID uint64 `json:"offer_id"`

	// ServerOptionIDs: server option IDs of the new server.
	ServerOptionIDs []uint64 `json:"server_option_ids"`

	// ProjectID: project ID of the new server.
	ProjectID string `json:"project_id"`

	// DatacenterName: datacenter name of the new server.
	DatacenterName *string `json:"datacenter_name,omitempty"`
}

// DeleteFailoverIPRequest: delete failover ip request.
type DeleteFailoverIPRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// IPID: ID of the failover IP to delete.
	IPID uint64 `json:"-"`
}

// DeleteServerRequest: delete server request.
type DeleteServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: server ID to delete.
	ServerID uint64 `json:"-"`
}

// DeleteServiceRequest: delete service request.
type DeleteServiceRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServiceID: ID of the service.
	ServiceID uint64 `json:"-"`
}

// DetachFailoverIPFromMacAddressRequest: detach failover ip from mac address request.
type DetachFailoverIPFromMacAddressRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// IPID: ID of the failover IP.
	IPID uint64 `json:"-"`
}

// DetachFailoverIPsRequest: detach failover i ps request.
type DetachFailoverIPsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// FipsIDs: list of IDs of failovers IP to detach.
	FipsIDs []uint64 `json:"fips_ids"`
}

// GetBMCAccessRequest: get bmc access request.
type GetBMCAccessRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server to get BMC access.
	ServerID uint64 `json:"-"`
}

// GetFailoverIPRequest: get failover ip request.
type GetFailoverIPRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// IPID: ID of the failover IP.
	IPID uint64 `json:"-"`
}

// GetIPv6BlockQuotasResponse: get i pv6 block quotas response.
type GetIPv6BlockQuotasResponse struct {
	// Quotas: quota for each CIDR of IPv6 block.
	Quotas []*GetIPv6BlockQuotasResponseQuota `json:"quotas"`

	// TotalCount: total count of quotas.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *GetIPv6BlockQuotasResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *GetIPv6BlockQuotasResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*GetIPv6BlockQuotasResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Quotas = append(r.Quotas, results.Quotas...)
	r.TotalCount += uint32(len(results.Quotas))
	return uint32(len(results.Quotas)), nil
}

// GetOSRequest: get os request.
type GetOSRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OsID: ID of the OS.
	OsID uint64 `json:"-"`

	// ServerID: ID of the server.
	ServerID uint64 `json:"-"`

	// ProjectID: project ID.
	ProjectID *string `json:"-"`
}

// GetOfferRequest: get offer request.
type GetOfferRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OfferID: ID of offer.
	OfferID uint64 `json:"-"`

	// ProjectID: project ID.
	ProjectID *string `json:"project_id,omitempty"`
}

// GetOrderedServiceRequest: get ordered service request.
type GetOrderedServiceRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	OrderedServiceID uint64 `json:"-"`
}

// GetRaidRequest: get raid request.
type GetRaidRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server.
	ServerID uint64 `json:"-"`
}

// GetRemainingQuotaRequest: get remaining quota request.
type GetRemainingQuotaRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ProjectID: project ID.
	ProjectID *string `json:"project_id,omitempty"`
}

// GetRemainingQuotaResponse: get remaining quota response.
type GetRemainingQuotaResponse struct {
	// FailoverIPQuota: current failover IP quota.
	FailoverIPQuota uint32 `json:"failover_ip_quota"`

	// FailoverIPRemainingQuota: remaining failover IP quota.
	FailoverIPRemainingQuota uint32 `json:"failover_ip_remaining_quota"`

	// FailoverBlockQuota: current failover block quota.
	FailoverBlockQuota uint32 `json:"failover_block_quota"`

	// FailoverBlockRemainingQuota: remaining failover block quota.
	FailoverBlockRemainingQuota uint32 `json:"failover_block_remaining_quota"`
}

// GetRescueRequest: get rescue request.
type GetRescueRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server to get rescue.
	ServerID uint64 `json:"-"`
}

// GetRpnStatusResponse: get rpn status response.
type GetRpnStatusResponse struct {
	// Status: if status = 'operational', you can perform rpn actions in write.
	// Default value: unknown_status
	Status GetRpnStatusResponseStatus `json:"status"`

	// OperationsLeft: number of operations left to perform before being operational.
	OperationsLeft *uint32 `json:"operations_left"`
}

// GetServerBackupRequest: get server backup request.
type GetServerBackupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: server ID of the backup.
	ServerID uint64 `json:"-"`
}

// GetServerDefaultPartitioningRequest: get server default partitioning request.
type GetServerDefaultPartitioningRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server.
	ServerID uint64 `json:"-"`

	// OsID: oS ID of the default partitioning.
	OsID uint64 `json:"-"`
}

// GetServerInstallRequest: get server install request.
type GetServerInstallRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: server ID of the server to install.
	ServerID uint64 `json:"-"`
}

// GetServerRequest: get server request.
type GetServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server.
	ServerID uint64 `json:"-"`
}

// GetServiceRequest: get service request.
type GetServiceRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServiceID: ID of the service.
	ServiceID uint64 `json:"-"`
}

// IPv6BlockAPICreateIPv6BlockRequest: i pv6 block api create i pv6 block request.
type IPv6BlockAPICreateIPv6BlockRequest struct {
	// ProjectID: ID of the project.
	ProjectID *string `json:"project_id,omitempty"`
}

// IPv6BlockAPICreateIPv6BlockSubnetRequest: i pv6 block api create i pv6 block subnet request.
type IPv6BlockAPICreateIPv6BlockSubnetRequest struct {
	// BlockID: ID of the IPv6 block.
	BlockID uint64 `json:"-"`

	// Address: address of the IPv6.
	Address net.IP `json:"address"`

	// Cidr: classless InterDomain Routing notation of the IPv6.
	Cidr uint32 `json:"cidr"`
}

// IPv6BlockAPIDeleteIPv6BlockRequest: i pv6 block api delete i pv6 block request.
type IPv6BlockAPIDeleteIPv6BlockRequest struct {
	// BlockID: ID of the IPv6 block to delete.
	BlockID uint64 `json:"-"`
}

// IPv6BlockAPIGetIPv6BlockQuotasRequest: i pv6 block api get i pv6 block quotas request.
type IPv6BlockAPIGetIPv6BlockQuotasRequest struct {
	// ProjectID: ID of the project.
	ProjectID *string `json:"-"`
}

// IPv6BlockAPIGetIPv6BlockRequest: i pv6 block api get i pv6 block request.
type IPv6BlockAPIGetIPv6BlockRequest struct {
	// ProjectID: ID of the project.
	ProjectID *string `json:"-"`
}

// IPv6BlockAPIListIPv6BlockSubnetsAvailableRequest: i pv6 block api list i pv6 block subnets available request.
type IPv6BlockAPIListIPv6BlockSubnetsAvailableRequest struct {
	// BlockID: ID of the IPv6 block.
	BlockID uint64 `json:"-"`
}

// IPv6BlockAPIListIPv6BlocksRequest: i pv6 block api list i pv6 blocks request.
type IPv6BlockAPIListIPv6BlocksRequest struct {
	ProjectID *string `json:"-"`
}

// IPv6BlockAPIUpdateIPv6BlockRequest: i pv6 block api update i pv6 block request.
type IPv6BlockAPIUpdateIPv6BlockRequest struct {
	// BlockID: ID of the IPv6 block.
	BlockID uint64 `json:"-"`

	// Nameservers: DNS to link to the IPv6.
	Nameservers *[]string `json:"nameservers,omitempty"`
}

// InstallServerRequest: install server request.
type InstallServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: server ID to install.
	ServerID uint64 `json:"-"`

	// OsID: oS ID to install on the server.
	OsID uint64 `json:"os_id"`

	// Hostname: hostname of the server.
	Hostname string `json:"hostname"`

	// UserLogin: user to install on the server.
	UserLogin *string `json:"user_login,omitempty"`

	// UserPassword: user password to install on the server.
	UserPassword *string `json:"user_password,omitempty"`

	// PanelPassword: panel password to install on the server.
	PanelPassword *string `json:"panel_password,omitempty"`

	// RootPassword: root password to install on the server.
	RootPassword *string `json:"root_password,omitempty"`

	// Partitions: partitions to install on the server.
	Partitions []*InstallPartition `json:"partitions"`

	// SSHKeyIDs: SSH key IDs authorized on the server.
	SSHKeyIDs []string `json:"ssh_key_ids"`

	// LicenseOfferID: offer ID of license to install on server.
	LicenseOfferID *uint64 `json:"license_offer_id,omitempty"`

	// IPID: IP to link at the license to install on server.
	IPID *uint64 `json:"ip_id,omitempty"`
}

// Invoice: invoice.
type Invoice struct {
	ID uint64 `json:"id"`

	TotalWithTaxes *scw.Money `json:"total_with_taxes"`

	TotalWithoutTaxes *scw.Money `json:"total_without_taxes"`

	CreatedAt *time.Time `json:"created_at"`

	PaidAt *time.Time `json:"paid_at"`

	// Status: default value: unknown_invoice_status
	Status InvoiceStatus `json:"status"`

	// PaymentMethod: default value: unknown_payment_method
	PaymentMethod InvoicePaymentMethod `json:"payment_method"`

	Content string `json:"content"`

	TransactionID uint64 `json:"transaction_id"`
}

// ListFailoverIPsRequest: list failover i ps request.
type ListFailoverIPsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// Page: page number.
	Page *uint32 `json:"-"`

	// PageSize: number of failovers IP per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order of the failovers IP.
	// Default value: ip_asc
	OrderBy ListFailoverIPsRequestOrderBy `json:"-"`

	// ProjectID: filter failovers IP by project ID.
	ProjectID *string `json:"-"`

	// Search: filter failovers IP which matching with this field.
	Search *string `json:"-"`

	// OnlyAvailable: true: return all failovers IP not attached on server
	// false: return all failovers IP attached on server.
	OnlyAvailable *bool `json:"-"`
}

// ListFailoverIPsResponse: list failover i ps response.
type ListFailoverIPsResponse struct {
	// TotalCount: total count of matching failovers IP.
	TotalCount uint32 `json:"total_count"`

	// FailoverIPs: list of failover IPs that match filters.
	FailoverIPs []*FailoverIP `json:"failover_ips"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListFailoverIPsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListFailoverIPsResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListFailoverIPsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.FailoverIPs = append(r.FailoverIPs, results.FailoverIPs...)
	r.TotalCount += uint32(len(results.FailoverIPs))
	return uint32(len(results.FailoverIPs)), nil
}

// ListIPsResponse: list i ps response.
type ListIPsResponse struct {
	// TotalCount: total count of authorized IPs.
	TotalCount uint32 `json:"total_count"`

	// IPs: list of authorized IPs.
	IPs []*RpnSanIP `json:"ips"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListIPsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListIPsResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListIPsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.IPs = append(r.IPs, results.IPs...)
	r.TotalCount += uint32(len(results.IPs))
	return uint32(len(results.IPs)), nil
}

// ListIPv6BlockSubnetsAvailableResponse: list i pv6 block subnets available response.
type ListIPv6BlockSubnetsAvailableResponse struct {
	// SubnetAvailables: all available address and CIDR available in subnet.
	SubnetAvailables []*ListIPv6BlockSubnetsAvailableResponseSubnet `json:"subnet_availables"`

	// TotalCount: total count of available subnets.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListIPv6BlockSubnetsAvailableResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListIPv6BlockSubnetsAvailableResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListIPv6BlockSubnetsAvailableResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.SubnetAvailables = append(r.SubnetAvailables, results.SubnetAvailables...)
	r.TotalCount += uint32(len(results.SubnetAvailables))
	return uint32(len(results.SubnetAvailables)), nil
}

// ListIPv6BlocksResponse: list i pv6 blocks response.
type ListIPv6BlocksResponse struct {
	TotalCount uint32 `json:"total_count"`

	IPv6Blocks []*IPv6Block `json:"ipv6_blocks"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListIPv6BlocksResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListIPv6BlocksResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListIPv6BlocksResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.IPv6Blocks = append(r.IPv6Blocks, results.IPv6Blocks...)
	r.TotalCount += uint32(len(results.IPv6Blocks))
	return uint32(len(results.IPv6Blocks)), nil
}

// ListInvoicesResponse: list invoices response.
type ListInvoicesResponse struct {
	TotalCount uint32 `json:"total_count"`

	Invoices []*InvoiceSummary `json:"invoices"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListInvoicesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListInvoicesResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListInvoicesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Invoices = append(r.Invoices, results.Invoices...)
	r.TotalCount += uint32(len(results.Invoices))
	return uint32(len(results.Invoices)), nil
}

// ListOSRequest: list os request.
type ListOSRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// Page: page number.
	Page *uint32 `json:"-"`

	// PageSize: number of OS per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order of the OS.
	// Default value: created_at_asc
	OrderBy ListOSRequestOrderBy `json:"-"`

	// Type: type of the OS.
	// Default value: unknown_type
	Type OSType `json:"-"`

	// ServerID: filter OS by compatible server ID.
	ServerID uint64 `json:"-"`

	// ProjectID: project ID.
	ProjectID *string `json:"-"`
}

// ListOSResponse: list os response.
type ListOSResponse struct {
	// TotalCount: total count of matching OS.
	TotalCount uint32 `json:"total_count"`

	// Os: oS that match filters.
	Os []*OS `json:"os"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListOSResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListOSResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListOSResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Os = append(r.Os, results.Os...)
	r.TotalCount += uint32(len(results.Os))
	return uint32(len(results.Os)), nil
}

// ListOffersRequest: list offers request.
type ListOffersRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// Page: page number.
	Page *uint32 `json:"-"`

	// PageSize: number of offer per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order of the offers.
	// Default value: created_at_asc
	OrderBy ListOffersRequestOrderBy `json:"-"`

	// CommercialRange: filter on commercial range.
	CommercialRange *string `json:"-"`

	// Catalog: filter on catalog.
	// Default value: all
	Catalog OfferCatalog `json:"-"`

	// ProjectID: project ID.
	ProjectID *string `json:"-"`

	// IsFailoverIP: get the current failover IP offer.
	IsFailoverIP *bool `json:"-"`

	// IsFailoverBlock: get the current failover IP block offer.
	IsFailoverBlock *bool `json:"-"`

	// SoldIn: filter offers depending on their datacenter.
	SoldIn []string `json:"-"`

	// AvailableOnly: set this filter to true to only return available offers.
	AvailableOnly *bool `json:"-"`

	// IsRpnSan: get the RPN SAN offers.
	IsRpnSan *bool `json:"-"`
}

// ListOffersResponse: list offers response.
type ListOffersResponse struct {
	// TotalCount: total count of matching offers.
	TotalCount uint32 `json:"total_count"`

	// Offers: offers that match filters.
	Offers []*Offer `json:"offers"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListOffersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListOffersResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListOffersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Offers = append(r.Offers, results.Offers...)
	r.TotalCount += uint32(len(results.Offers))
	return uint32(len(results.Offers)), nil
}

// ListRefundsResponse: list refunds response.
type ListRefundsResponse struct {
	TotalCount uint32 `json:"total_count"`

	Refunds []*RefundSummary `json:"refunds"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListRefundsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListRefundsResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListRefundsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Refunds = append(r.Refunds, results.Refunds...)
	r.TotalCount += uint32(len(results.Refunds))
	return uint32(len(results.Refunds)), nil
}

// ListRpnCapableSanServersResponse: list rpn capable san servers response.
type ListRpnCapableSanServersResponse struct {
	// TotalCount: total count of rpn capable san servers.
	TotalCount uint32 `json:"total_count"`

	// SanServers: list of san servers.
	SanServers []*RpnSanServer `json:"san_servers"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListRpnCapableSanServersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListRpnCapableSanServersResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListRpnCapableSanServersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.SanServers = append(r.SanServers, results.SanServers...)
	r.TotalCount += uint32(len(results.SanServers))
	return uint32(len(results.SanServers)), nil
}

// ListRpnCapableServersResponse: list rpn capable servers response.
type ListRpnCapableServersResponse struct {
	// TotalCount: total count of rpn capable servers.
	TotalCount uint32 `json:"total_count"`

	// Servers: list of servers.
	Servers []*Server `json:"servers"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListRpnCapableServersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListRpnCapableServersResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListRpnCapableServersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Servers = append(r.Servers, results.Servers...)
	r.TotalCount += uint32(len(results.Servers))
	return uint32(len(results.Servers)), nil
}

// ListRpnGroupMembersResponse: list rpn group members response.
type ListRpnGroupMembersResponse struct {
	// TotalCount: total count of rpn v1 group members.
	TotalCount uint32 `json:"total_count"`

	// Members: list of rpn v1 group members.
	Members []*RpnGroupMember `json:"members"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListRpnGroupMembersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListRpnGroupMembersResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListRpnGroupMembersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Members = append(r.Members, results.Members...)
	r.TotalCount += uint32(len(results.Members))
	return uint32(len(results.Members)), nil
}

// ListRpnGroupsResponse: list rpn groups response.
type ListRpnGroupsResponse struct {
	// TotalCount: total count of rpn groups.
	TotalCount uint32 `json:"total_count"`

	// RpnGroups: list of rpn v1 groups.
	RpnGroups []*RpnGroup `json:"rpn_groups"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListRpnGroupsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListRpnGroupsResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListRpnGroupsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.RpnGroups = append(r.RpnGroups, results.RpnGroups...)
	r.TotalCount += uint32(len(results.RpnGroups))
	return uint32(len(results.RpnGroups)), nil
}

// ListRpnInvitesResponse: list rpn invites response.
type ListRpnInvitesResponse struct {
	// TotalCount: total count of invites.
	TotalCount uint32 `json:"total_count"`

	// Members: list of invites.
	Members []*RpnGroupMember `json:"members"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListRpnInvitesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListRpnInvitesResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListRpnInvitesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Members = append(r.Members, results.Members...)
	r.TotalCount += uint32(len(results.Members))
	return uint32(len(results.Members)), nil
}

// ListRpnSansResponse: list rpn sans response.
type ListRpnSansResponse struct {
	// TotalCount: total count of matching RPN SANs.
	TotalCount uint32 `json:"total_count"`

	// RpnSans: list of RPN SANs that match filters.
	RpnSans []*RpnSanSummary `json:"rpn_sans"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListRpnSansResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListRpnSansResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListRpnSansResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.RpnSans = append(r.RpnSans, results.RpnSans...)
	r.TotalCount += uint32(len(results.RpnSans))
	return uint32(len(results.RpnSans)), nil
}

// ListRpnServerCapabilitiesResponse: list rpn server capabilities response.
type ListRpnServerCapabilitiesResponse struct {
	// TotalCount: total count of servers.
	TotalCount uint32 `json:"total_count"`

	// Servers: list of servers and their RPN capabilities.
	Servers []*RpnServerCapability `json:"servers"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListRpnServerCapabilitiesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListRpnServerCapabilitiesResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListRpnServerCapabilitiesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Servers = append(r.Servers, results.Servers...)
	r.TotalCount += uint32(len(results.Servers))
	return uint32(len(results.Servers)), nil
}

// ListRpnV2CapableResourcesResponse: list rpn v2 capable resources response.
type ListRpnV2CapableResourcesResponse struct {
	// TotalCount: total count of matching rpn v2 capable resources.
	TotalCount uint32 `json:"total_count"`

	// Servers: list of rpn v2 capable resources that match filters.
	Servers []*Server `json:"servers"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListRpnV2CapableResourcesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListRpnV2CapableResourcesResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListRpnV2CapableResourcesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Servers = append(r.Servers, results.Servers...)
	r.TotalCount += uint32(len(results.Servers))
	return uint32(len(results.Servers)), nil
}

// ListRpnV2GroupLogsResponse: list rpn v2 group logs response.
type ListRpnV2GroupLogsResponse struct {
	// TotalCount: total count of matching rpn v2 logs.
	TotalCount uint32 `json:"total_count"`

	// Logs: list of rpn v2 logs that match filters.
	Logs []*Log `json:"logs"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListRpnV2GroupLogsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListRpnV2GroupLogsResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListRpnV2GroupLogsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Logs = append(r.Logs, results.Logs...)
	r.TotalCount += uint32(len(results.Logs))
	return uint32(len(results.Logs)), nil
}

// ListRpnV2GroupsResponse: list rpn v2 groups response.
type ListRpnV2GroupsResponse struct {
	// TotalCount: total count of matching rpn v2 groups.
	TotalCount uint32 `json:"total_count"`

	// RpnGroups: list of rpn v2 groups that match filters.
	RpnGroups []*RpnV2Group `json:"rpn_groups"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListRpnV2GroupsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListRpnV2GroupsResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListRpnV2GroupsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.RpnGroups = append(r.RpnGroups, results.RpnGroups...)
	r.TotalCount += uint32(len(results.RpnGroups))
	return uint32(len(results.RpnGroups)), nil
}

// ListRpnV2MembersResponse: list rpn v2 members response.
type ListRpnV2MembersResponse struct {
	// TotalCount: total count of matching rpn v2 group members.
	TotalCount uint32 `json:"total_count"`

	// Members: list of rpn v2 group members that match filters.
	Members []*RpnV2Member `json:"members"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListRpnV2MembersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListRpnV2MembersResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListRpnV2MembersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Members = append(r.Members, results.Members...)
	r.TotalCount += uint32(len(results.Members))
	return uint32(len(results.Members)), nil
}

// ListServerDisksRequest: list server disks request.
type ListServerDisksRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: server ID of the server disks.
	ServerID uint64 `json:"-"`

	// Page: page number.
	Page *uint32 `json:"-"`

	// PageSize: number of server disk per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order of the server disks.
	// Default value: created_at_asc
	OrderBy ListServerDisksRequestOrderBy `json:"-"`
}

// ListServerDisksResponse: list server disks response.
type ListServerDisksResponse struct {
	// TotalCount: total count of matching server disks.
	TotalCount uint32 `json:"total_count"`

	// Disks: server disks that match filters.
	Disks []*ServerDisk `json:"disks"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListServerDisksResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListServerDisksResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListServerDisksResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Disks = append(r.Disks, results.Disks...)
	r.TotalCount += uint32(len(results.Disks))
	return uint32(len(results.Disks)), nil
}

// ListServerEventsRequest: list server events request.
type ListServerEventsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: server ID of the server events.
	ServerID uint64 `json:"-"`

	// Page: page number.
	Page *uint32 `json:"-"`

	// PageSize: number of server event per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order of the server events.
	// Default value: created_at_asc
	OrderBy ListServerEventsRequestOrderBy `json:"-"`
}

// ListServerEventsResponse: list server events response.
type ListServerEventsResponse struct {
	// TotalCount: total count of matching server events.
	TotalCount uint32 `json:"total_count"`

	// Events: server events that match filters.
	Events []*ServerEvent `json:"events"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListServerEventsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListServerEventsResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListServerEventsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Events = append(r.Events, results.Events...)
	r.TotalCount += uint32(len(results.Events))
	return uint32(len(results.Events)), nil
}

// ListServersRequest: list servers request.
type ListServersRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// Page: page number.
	Page *uint32 `json:"-"`

	// PageSize: number of server per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order of the servers.
	// Default value: created_at_asc
	OrderBy ListServersRequestOrderBy `json:"-"`

	// ProjectID: filter servers by project ID.
	ProjectID *string `json:"-"`

	// Search: filter servers by hostname.
	Search *string `json:"-"`
}

// ListServersResponse: list servers response.
type ListServersResponse struct {
	// TotalCount: total count of matching servers.
	TotalCount uint32 `json:"total_count"`

	// Servers: servers that match filters.
	Servers []*ServerSummary `json:"servers"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListServersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListServersResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListServersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Servers = append(r.Servers, results.Servers...)
	r.TotalCount += uint32(len(results.Servers))
	return uint32(len(results.Servers)), nil
}

// ListServicesRequest: list services request.
type ListServicesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// Page: page number.
	Page *uint32 `json:"-"`

	// PageSize: number of service per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order of the services.
	// Default value: created_at_asc
	OrderBy ListServicesRequestOrderBy `json:"-"`

	// ProjectID: project ID.
	ProjectID *string `json:"-"`
}

// ListServicesResponse: list services response.
type ListServicesResponse struct {
	// TotalCount: total count of matching services.
	TotalCount uint32 `json:"total_count"`

	// Services: services that match filters.
	Services []*Service `json:"services"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListServicesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListServicesResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListServicesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Services = append(r.Services, results.Services...)
	r.TotalCount += uint32(len(results.Services))
	return uint32(len(results.Services)), nil
}

// ListSubscribableServerOptionsRequest: list subscribable server options request.
type ListSubscribableServerOptionsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: server ID of the subscribable server options.
	ServerID uint64 `json:"-"`

	// Page: page number.
	Page *uint32 `json:"-"`

	// PageSize: number of subscribable server option per page.
	PageSize *uint32 `json:"-"`
}

// ListSubscribableServerOptionsResponse: list subscribable server options response.
type ListSubscribableServerOptionsResponse struct {
	// TotalCount: total count of matching subscribable server options.
	TotalCount uint32 `json:"total_count"`

	// ServerOptions: server options that match filters.
	ServerOptions []*Offer `json:"server_options"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSubscribableServerOptionsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSubscribableServerOptionsResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListSubscribableServerOptionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.ServerOptions = append(r.ServerOptions, results.ServerOptions...)
	r.TotalCount += uint32(len(results.ServerOptions))
	return uint32(len(results.ServerOptions)), nil
}

// OfferFailoverBlockInfo: offer failover block info.
type OfferFailoverBlockInfo struct {
	OnetimeFees *Offer `json:"onetime_fees"`
}

// OfferFailoverIPInfo: offer failover ip info.
type OfferFailoverIPInfo struct {
	OnetimeFees *Offer `json:"onetime_fees"`
}

// OfferServerInfo: offer server info.
type OfferServerInfo struct {
	Bandwidth uint64 `json:"bandwidth"`

	// Stock: default value: empty
	Stock OfferServerInfoStock `json:"stock"`

	CommercialRange string `json:"commercial_range"`

	Disks []*Disk `json:"disks"`

	CPUs []*CPU `json:"cpus"`

	Memories []*Memory `json:"memories"`

	PersistentMemories []*PersistentMemory `json:"persistent_memories"`

	RaidControllers []*RaidController `json:"raid_controllers"`

	AvailableOptions []*Offer `json:"available_options"`

	RpnVersion *uint32 `json:"rpn_version"`

	Connectivity uint64 `json:"connectivity"`

	OnetimeFees *Offer `json:"onetime_fees"`

	StockByDatacenter map[string]OfferServerInfoStock `json:"stock_by_datacenter"`
}

// OfferServiceLevelInfo: offer service level info.
type OfferServiceLevelInfo struct {
	SupportTicket bool `json:"support_ticket"`

	SupportPhone bool `json:"support_phone"`

	SalesSupport bool `json:"sales_support"`

	Git string `json:"git"`

	SLA float32 `json:"sla"`

	PrioritySupport bool `json:"priority_support"`

	HighRpnBandwidth bool `json:"high_rpn_bandwidth"`

	Customization bool `json:"customization"`

	Antidos bool `json:"antidos"`

	ExtraFailoverQuota uint32 `json:"extra_failover_quota"`

	AvailableOptions []*Offer `json:"available_options"`
}

// Raid: raid.
type Raid struct {
	// RaidArrays: details about the RAID controller.
	RaidArrays []*RaidArray `json:"raid_arrays"`
}

// RebootServerRequest: reboot server request.
type RebootServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: server ID to reboot.
	ServerID uint64 `json:"-"`
}

// Refund: refund.
type Refund struct {
	ID uint64 `json:"id"`

	TotalWithTaxes *scw.Money `json:"total_with_taxes"`

	TotalWithoutTaxes *scw.Money `json:"total_without_taxes"`

	CreatedAt *time.Time `json:"created_at"`

	RefundedAt *time.Time `json:"refunded_at"`

	// Status: default value: unknown_refund_status
	Status RefundStatus `json:"status"`

	// Method: default value: unknown_refund_method
	Method RefundMethod `json:"method"`

	Content string `json:"content"`
}

// Rescue: rescue.
type Rescue struct {
	// OsID: oS ID of the rescue.
	OsID uint64 `json:"os_id"`

	// Login: login of the rescue.
	Login string `json:"login"`

	// Password: password of the rescue.
	Password string `json:"password"`

	// Protocol: protocol of the rescue.
	// Default value: vnc
	Protocol RescueProtocol `json:"protocol"`
}

// RpnAPIGetRpnStatusRequest: rpn api get rpn status request.
type RpnAPIGetRpnStatusRequest struct {
	// ProjectID: a project ID.
	ProjectID *string `json:"-"`

	// Rpnv1GroupID: an RPN v1 group ID.
	Rpnv1GroupID *uint64 `json:"-"`

	// Rpnv2GroupID: an RPN v2 group ID.
	Rpnv2GroupID *uint64 `json:"-"`
}

// RpnAPIListRpnServerCapabilitiesRequest: rpn api list rpn server capabilities request.
type RpnAPIListRpnServerCapabilitiesRequest struct {
	// Page: page number.
	Page *uint32 `json:"-"`

	// PageSize: number of servers per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order of the servers.
	// Default value: created_at_asc
	OrderBy ListRpnServerCapabilitiesRequestOrderBy `json:"-"`

	// ProjectID: filter servers by project ID.
	ProjectID *string `json:"-"`
}

// RpnSanAPIAddIPRequest: rpn san api add ip request.
type RpnSanAPIAddIPRequest struct {
	// RpnSanID: rPN SAN ID.
	RpnSanID uint64 `json:"-"`

	// IPIDs: an array of IP ID.
	IPIDs []uint64 `json:"ip_ids"`
}

// RpnSanAPICreateRpnSanRequest: rpn san api create rpn san request.
type RpnSanAPICreateRpnSanRequest struct {
	// OfferID: offer ID.
	OfferID uint64 `json:"offer_id"`

	// ProjectID: your project ID.
	ProjectID string `json:"project_id"`
}

// RpnSanAPIDeleteRpnSanRequest: rpn san api delete rpn san request.
type RpnSanAPIDeleteRpnSanRequest struct {
	// RpnSanID: rPN SAN ID.
	RpnSanID uint64 `json:"-"`
}

// RpnSanAPIGetRpnSanRequest: rpn san api get rpn san request.
type RpnSanAPIGetRpnSanRequest struct {
	// RpnSanID: rPN SAN ID.
	RpnSanID uint64 `json:"-"`
}

// RpnSanAPIListAvailableIPsRequest: rpn san api list available i ps request.
type RpnSanAPIListAvailableIPsRequest struct {
	// RpnSanID: rPN SAN ID.
	RpnSanID uint64 `json:"-"`

	// Type: filter by IP type (server | rpnv2_subnet).
	// Default value: unknown
	Type RpnSanIPType `json:"-"`
}

// RpnSanAPIListIPsRequest: rpn san api list i ps request.
type RpnSanAPIListIPsRequest struct {
	// RpnSanID: rPN SAN ID.
	RpnSanID uint64 `json:"-"`

	// Type: filter by IP type (server | rpnv2_subnet).
	// Default value: unknown
	Type RpnSanIPType `json:"-"`
}

// RpnSanAPIListRpnSansRequest: rpn san api list rpn sans request.
type RpnSanAPIListRpnSansRequest struct {
	// Page: page number.
	Page *uint32 `json:"-"`

	// PageSize: number of RPN SANs per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order of the RPN SANs.
	// Default value: created_at_asc
	OrderBy ListRpnSansRequestOrderBy `json:"-"`

	// ProjectID: filter RPN SANs by project ID.
	ProjectID *string `json:"-"`
}

// RpnSanAPIRemoveIPRequest: rpn san api remove ip request.
type RpnSanAPIRemoveIPRequest struct {
	// RpnSanID: rPN SAN ID.
	RpnSanID uint64 `json:"-"`

	// IPIDs: an array of IP ID.
	IPIDs []uint64 `json:"ip_ids"`
}

// RpnV1ApiAcceptRpnInviteRequest: rpn v1 api accept rpn invite request.
type RpnV1ApiAcceptRpnInviteRequest struct {
	// MemberID: the member ID.
	MemberID uint64 `json:"-"`
}

// RpnV1ApiAddRpnGroupMembersRequest: rpn v1 api add rpn group members request.
type RpnV1ApiAddRpnGroupMembersRequest struct {
	// GroupID: the rpn v1 group ID.
	GroupID uint64 `json:"-"`

	// ServerIDs: a collection of rpn v1 capable server IDs.
	ServerIDs []uint64 `json:"server_ids"`

	// SanServerIDs: a collection of rpn v1 capable RPN SAN server IDs.
	SanServerIDs []uint64 `json:"san_server_ids"`
}

// RpnV1ApiCreateRpnGroupRequest: rpn v1 api create rpn group request.
type RpnV1ApiCreateRpnGroupRequest struct {
	// Name: rpn v1 group name.
	Name string `json:"name"`

	// ServerIDs: a collection of rpn v1 capable servers.
	ServerIDs []uint64 `json:"server_ids"`

	// SanServerIDs: a collection of rpn v1 capable rpn sans servers.
	SanServerIDs []uint64 `json:"san_server_ids"`

	// ProjectID: a project ID.
	ProjectID string `json:"project_id"`
}

// RpnV1ApiDeleteRpnGroupMembersRequest: rpn v1 api delete rpn group members request.
type RpnV1ApiDeleteRpnGroupMembersRequest struct {
	// GroupID: the rpn v1 group ID.
	GroupID uint64 `json:"-"`

	// MemberIDs: a collection of rpn v1 group members IDs.
	MemberIDs []uint64 `json:"member_ids"`
}

// RpnV1ApiDeleteRpnGroupRequest: rpn v1 api delete rpn group request.
type RpnV1ApiDeleteRpnGroupRequest struct {
	// GroupID: rpn v1 group ID.
	GroupID uint64 `json:"-"`
}

// RpnV1ApiGetRpnGroupRequest: rpn v1 api get rpn group request.
type RpnV1ApiGetRpnGroupRequest struct {
	// GroupID: rpn v1 group ID.
	GroupID uint64 `json:"-"`
}

// RpnV1ApiLeaveRpnGroupRequest: rpn v1 api leave rpn group request.
type RpnV1ApiLeaveRpnGroupRequest struct {
	// GroupID: the RPN V1 group ID.
	GroupID uint64 `json:"-"`

	// MemberIDs: a collection of rpn v1 group members IDs.
	MemberIDs []uint64 `json:"member_ids"`
}

// RpnV1ApiListRpnCapableSanServersRequest: rpn v1 api list rpn capable san servers request.
type RpnV1ApiListRpnCapableSanServersRequest struct {
	// Page: page number.
	Page *uint32 `json:"-"`

	// PageSize: number of rpn capable resources per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order of the rpn capable resources.
	// Default value: created_at_asc
	OrderBy ListRpnCapableSanServersRequestOrderBy `json:"-"`

	// ProjectID: filter rpn capable resources by project ID.
	ProjectID *string `json:"-"`
}

// RpnV1ApiListRpnCapableServersRequest: rpn v1 api list rpn capable servers request.
type RpnV1ApiListRpnCapableServersRequest struct {
	// Page: page number.
	Page *uint32 `json:"-"`

	// PageSize: number of rpn capable resources per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order of the rpn capable resources.
	// Default value: created_at_asc
	OrderBy ListRpnCapableServersRequestOrderBy `json:"-"`

	// ProjectID: filter rpn capable resources by project ID.
	ProjectID *string `json:"-"`
}

// RpnV1ApiListRpnGroupMembersRequest: rpn v1 api list rpn group members request.
type RpnV1ApiListRpnGroupMembersRequest struct {
	// GroupID: filter rpn v1 group members by group ID.
	GroupID uint64 `json:"-"`

	// Page: page number.
	Page *uint32 `json:"-"`

	// PageSize: number of rpn v1 group members per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order of the rpn v1 group members.
	// Default value: created_at_asc
	OrderBy ListRpnGroupMembersRequestOrderBy `json:"-"`

	// ProjectID: a project ID.
	ProjectID *string `json:"-"`
}

// RpnV1ApiListRpnGroupsRequest: rpn v1 api list rpn groups request.
type RpnV1ApiListRpnGroupsRequest struct {
	// Page: page number.
	Page *uint32 `json:"-"`

	// PageSize: number of rpn v1 groups per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order of the rpn v1 groups.
	// Default value: created_at_asc
	OrderBy ListRpnGroupsRequestOrderBy `json:"-"`

	// ProjectID: filter rpn v1 groups by project ID.
	ProjectID *string `json:"-"`
}

// RpnV1ApiListRpnInvitesRequest: rpn v1 api list rpn invites request.
type RpnV1ApiListRpnInvitesRequest struct {
	// Page: page number.
	Page *uint32 `json:"-"`

	// PageSize: number of rpn capable resources per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order of the rpn capable resources.
	// Default value: created_at_asc
	OrderBy ListRpnInvitesRequestOrderBy `json:"-"`

	// ProjectID: filter rpn capable resources by project ID.
	ProjectID string `json:"-"`
}

// RpnV1ApiRefuseRpnInviteRequest: rpn v1 api refuse rpn invite request.
type RpnV1ApiRefuseRpnInviteRequest struct {
	// MemberID: the member ID.
	MemberID uint64 `json:"-"`
}

// RpnV1ApiRpnGroupInviteRequest: rpn v1 api rpn group invite request.
type RpnV1ApiRpnGroupInviteRequest struct {
	// GroupID: the RPN V1 group ID.
	GroupID uint64 `json:"-"`

	// ServerIDs: a collection of external server IDs.
	ServerIDs []uint64 `json:"server_ids"`

	// ProjectID: a project ID.
	ProjectID string `json:"project_id"`
}

// RpnV1ApiUpdateRpnGroupNameRequest: rpn v1 api update rpn group name request.
type RpnV1ApiUpdateRpnGroupNameRequest struct {
	// GroupID: rpn v1 group ID.
	GroupID uint64 `json:"-"`

	// Name: new rpn v1 group name.
	Name *string `json:"name,omitempty"`
}

// RpnV2ApiAddRpnV2MembersRequest: rpn v2 api add rpn v2 members request.
type RpnV2ApiAddRpnV2MembersRequest struct {
	// GroupID: rPN V2 group ID.
	GroupID uint64 `json:"-"`

	// Servers: a collection of server IDs.
	Servers []uint64 `json:"servers"`
}

// RpnV2ApiCreateRpnV2GroupRequest: rpn v2 api create rpn v2 group request.
type RpnV2ApiCreateRpnV2GroupRequest struct {
	// ProjectID: project ID of the RPN V2 group.
	ProjectID string `json:"project_id"`

	// Type: rPN V2 group type (qing / standard).
	// Default value: unknown_type
	Type RpnV2GroupType `json:"type"`

	// Name: rPN V2 group name.
	Name string `json:"name"`

	// Servers: a collection of server IDs.
	Servers []uint64 `json:"servers"`
}

// RpnV2ApiDeleteRpnV2GroupRequest: rpn v2 api delete rpn v2 group request.
type RpnV2ApiDeleteRpnV2GroupRequest struct {
	// GroupID: rPN V2 group ID.
	GroupID uint64 `json:"-"`
}

// RpnV2ApiDeleteRpnV2MembersRequest: rpn v2 api delete rpn v2 members request.
type RpnV2ApiDeleteRpnV2MembersRequest struct {
	// GroupID: rPN V2 group ID.
	GroupID uint64 `json:"-"`

	// MemberIDs: a collection of member IDs.
	MemberIDs []uint64 `json:"member_ids"`
}

// RpnV2ApiDisableRpnV2GroupCompatibilityRequest: rpn v2 api disable rpn v2 group compatibility request.
type RpnV2ApiDisableRpnV2GroupCompatibilityRequest struct {
	// GroupID: rPN V2 group ID.
	GroupID uint64 `json:"-"`
}

// RpnV2ApiEnableRpnV2GroupCompatibilityRequest: rpn v2 api enable rpn v2 group compatibility request.
type RpnV2ApiEnableRpnV2GroupCompatibilityRequest struct {
	// GroupID: rPN V2 group ID.
	GroupID uint64 `json:"-"`

	// Rpnv1GroupID: rPN V1 group ID.
	Rpnv1GroupID uint64 `json:"rpnv1_group_id"`
}

// RpnV2ApiGetRpnV2GroupRequest: rpn v2 api get rpn v2 group request.
type RpnV2ApiGetRpnV2GroupRequest struct {
	// GroupID: rPN V2 group ID.
	GroupID uint64 `json:"-"`
}

// RpnV2ApiListRpnV2CapableResourcesRequest: rpn v2 api list rpn v2 capable resources request.
type RpnV2ApiListRpnV2CapableResourcesRequest struct {
	// Page: page number.
	Page *uint32 `json:"-"`

	// PageSize: number of rpn v2 capable resources per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order of the rpn v2 capable resources.
	// Default value: created_at_asc
	OrderBy ListRpnV2CapableResourcesRequestOrderBy `json:"-"`

	// ProjectID: filter rpn v2 capable resources by project ID.
	ProjectID *string `json:"-"`
}

// RpnV2ApiListRpnV2GroupLogsRequest: rpn v2 api list rpn v2 group logs request.
type RpnV2ApiListRpnV2GroupLogsRequest struct {
	// GroupID: rPN V2 group ID.
	GroupID uint64 `json:"-"`

	// Page: page number.
	Page *uint32 `json:"-"`

	// PageSize: number of rpn v2 group logs per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order of the rpn v2 group logs.
	// Default value: created_at_asc
	OrderBy ListRpnV2GroupLogsRequestOrderBy `json:"-"`
}

// RpnV2ApiListRpnV2GroupsRequest: rpn v2 api list rpn v2 groups request.
type RpnV2ApiListRpnV2GroupsRequest struct {
	// Page: page number.
	Page *uint32 `json:"-"`

	// PageSize: number of rpn v2 groups per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order of the rpn v2 groups.
	// Default value: created_at_asc
	OrderBy ListRpnV2GroupsRequestOrderBy `json:"-"`

	// ProjectID: filter rpn v2 groups by project ID.
	ProjectID *string `json:"-"`
}

// RpnV2ApiListRpnV2MembersRequest: rpn v2 api list rpn v2 members request.
type RpnV2ApiListRpnV2MembersRequest struct {
	// GroupID: rPN V2 group ID.
	GroupID uint64 `json:"-"`

	// Page: page number.
	Page *uint32 `json:"-"`

	// PageSize: number of rpn v2 group members per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order of the rpn v2 group members.
	// Default value: created_at_asc
	OrderBy ListRpnV2MembersRequestOrderBy `json:"-"`

	// Type: filter members by type.
	// Default value: unknown_type
	Type ListRpnV2MembersRequestType `json:"-"`
}

// RpnV2ApiUpdateRpnV2GroupNameRequest: rpn v2 api update rpn v2 group name request.
type RpnV2ApiUpdateRpnV2GroupNameRequest struct {
	// GroupID: rPN V2 group ID.
	GroupID uint64 `json:"-"`

	// Name: rPN V2 group name.
	Name *string `json:"name,omitempty"`
}

// RpnV2ApiUpdateRpnV2VlanForMembersRequest: rpn v2 api update rpn v2 vlan for members request.
type RpnV2ApiUpdateRpnV2VlanForMembersRequest struct {
	// GroupID: rPN V2 group ID.
	GroupID uint64 `json:"-"`

	// MemberIDs: rPN V2 member IDs.
	MemberIDs []uint64 `json:"member_ids"`

	// Vlan: min: 0.
	// Max: 3967.
	Vlan *uint32 `json:"vlan,omitempty"`
}

// ServerDefaultPartitioning: server default partitioning.
type ServerDefaultPartitioning struct {
	// Partitions: default partitions.
	Partitions []*Partition `json:"partitions"`
}

// ServerInstall: server install.
type ServerInstall struct {
	OsID uint64 `json:"os_id"`

	Hostname string `json:"hostname"`

	UserLogin *string `json:"user_login"`

	Partitions []*Partition `json:"partitions"`

	SSHKeyIDs []string `json:"ssh_key_ids"`

	// Status: default value: unknown
	Status ServerInstallStatus `json:"status"`

	PanelURL *string `json:"panel_url"`
}

// StartBMCAccessRequest: start bmc access request.
type StartBMCAccessRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server to start the BMC access.
	ServerID uint64 `json:"-"`

	// IP: the IP authorized to connect to the given server.
	IP net.IP `json:"ip"`
}

// StartRescueRequest: start rescue request.
type StartRescueRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server to start rescue.
	ServerID uint64 `json:"-"`

	// OsID: oS ID to use to start rescue.
	OsID uint64 `json:"os_id"`
}

// StartServerRequest: start server request.
type StartServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: server ID to start.
	ServerID uint64 `json:"-"`
}

// StopBMCAccessRequest: stop bmc access request.
type StopBMCAccessRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server to stop BMC access.
	ServerID uint64 `json:"-"`
}

// StopRescueRequest: stop rescue request.
type StopRescueRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server to stop rescue.
	ServerID uint64 `json:"-"`
}

// StopServerRequest: stop server request.
type StopServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: server ID to stop.
	ServerID uint64 `json:"-"`
}

// SubscribeServerOptionRequest: subscribe server option request.
type SubscribeServerOptionRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: server ID to subscribe server option.
	ServerID uint64 `json:"-"`

	// OptionID: option ID to subscribe.
	OptionID uint64 `json:"option_id"`
}

// SubscribeStorageOptionsRequest: subscribe storage options request.
type SubscribeStorageOptionsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: server ID of the storage options to subscribe.
	ServerID uint64 `json:"-"`

	// OptionsIDs: option IDs of the storage options to subscribe.
	OptionsIDs []uint64 `json:"options_ids"`
}

// SubscribeStorageOptionsResponse: subscribe storage options response.
type SubscribeStorageOptionsResponse struct {
	// Services: services subscribe storage options.
	Services []*Service `json:"services"`
}

// UpdateRaidRequest: update raid request.
type UpdateRaidRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server.
	ServerID uint64 `json:"-"`

	// RaidArrays: rAIDs to update.
	RaidArrays []*UpdatableRaidArray `json:"raid_arrays"`
}

// UpdateReverseRequest: update reverse request.
type UpdateReverseRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// IPID: ID of the IP.
	IPID uint64 `json:"-"`

	// Reverse: reverse to apply on the IP.
	Reverse string `json:"reverse"`
}

// UpdateServerBackupRequest: update server backup request.
type UpdateServerBackupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: server ID to update backup.
	ServerID uint64 `json:"-"`

	// Password: password of the server backup.
	Password *string `json:"password,omitempty"`

	// Autologin: autologin of the server backup.
	Autologin *bool `json:"autologin,omitempty"`

	// ACLEnabled: boolean to enable or disable ACL.
	ACLEnabled *bool `json:"acl_enabled,omitempty"`
}

// UpdateServerRequest: update server request.
type UpdateServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: server ID to update.
	ServerID uint64 `json:"-"`

	// Hostname: hostname of the server to update.
	Hostname *string `json:"hostname,omitempty"`

	// EnableIPv6: flag to enable or not the IPv6 of server.
	EnableIPv6 *bool `json:"enable_ipv6,omitempty"`
}

// UpdateServerTagsRequest: update server tags request.
type UpdateServerTagsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: server ID to update the tags.
	ServerID uint64 `json:"-"`

	// Tags: tags of server to update.
	Tags []string `json:"tags"`
}

// Dedibox Phoenix API.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

func (s *API) Zones() []scw.Zone {
	return []scw.Zone{scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1}
}

// ListServers: List baremetal servers for project.
func (s *API) ListServers(req *ListServersRequest, opts ...scw.RequestOption) (*ListServersResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "search", req.Search)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers",
		Query:  query,
	}

	var resp ListServersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetServer: Get the server associated with the given ID.
func (s *API) GetServer(req *GetServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetServerBackup:
func (s *API) GetServerBackup(req *GetServerBackupRequest, opts ...scw.RequestOption) (*Backup, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/backups",
	}

	var resp Backup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateServerBackup:
func (s *API) UpdateServerBackup(req *UpdateServerBackupRequest, opts ...scw.RequestOption) (*Backup, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/backups",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Backup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSubscribableServerOptions: List subscribable options associated to the given server ID.
func (s *API) ListSubscribableServerOptions(req *ListSubscribableServerOptionsRequest, opts ...scw.RequestOption) (*ListSubscribableServerOptionsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/subscribable-server-options",
		Query:  query,
	}

	var resp ListSubscribableServerOptionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SubscribeServerOption: Subscribe option for the given server ID.
func (s *API) SubscribeServerOption(req *SubscribeServerOptionRequest, opts ...scw.RequestOption) (*Service, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/subscribe-server-option",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Service

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateServer: Create a new baremetal server. The order return you a service ID to follow the provisionning status you could call GetService.
func (s *API) CreateServer(req *CreateServerRequest, opts ...scw.RequestOption) (*Service, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Service

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SubscribeStorageOptions: Subscribe storage option for the given server ID.
func (s *API) SubscribeStorageOptions(req *SubscribeStorageOptionsRequest, opts ...scw.RequestOption) (*SubscribeStorageOptionsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/subscribe-storage-options",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SubscribeStorageOptionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateServer: Update the server associated with the given ID.
func (s *API) UpdateServer(req *UpdateServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateServerTags:
func (s *API) UpdateServerTags(req *UpdateServerTagsRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/tags",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RebootServer: Reboot the server associated with the given ID, use boot param to reboot in rescue.
func (s *API) RebootServer(req *RebootServerRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/reboot",
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

// StartServer: Start the server associated with the given ID.
func (s *API) StartServer(req *StartServerRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/start",
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

// StopServer: Stop the server associated with the given ID.
func (s *API) StopServer(req *StopServerRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/stop",
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

// DeleteServer: Delete the server associated with the given ID.
func (s *API) DeleteServer(req *DeleteServerRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListServerEvents: List events associated to the given server ID.
func (s *API) ListServerEvents(req *ListServerEventsRequest, opts ...scw.RequestOption) (*ListServerEventsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/events",
		Query:  query,
	}

	var resp ListServerEventsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListServerDisks: List disks associated to the given server ID.
func (s *API) ListServerDisks(req *ListServerDisksRequest, opts ...scw.RequestOption) (*ListServerDisksResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/disks",
		Query:  query,
	}

	var resp ListServerDisksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOrderedService:
func (s *API) GetOrderedService(req *GetOrderedServiceRequest, opts ...scw.RequestOption) (*Service, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.OrderedServiceID) == "" {
		return nil, errors.New("field OrderedServiceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/ordered-services/" + fmt.Sprint(req.OrderedServiceID) + "",
	}

	var resp Service

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetService: Get the service associated with the given ID.
func (s *API) GetService(req *GetServiceRequest, opts ...scw.RequestOption) (*Service, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServiceID) == "" {
		return nil, errors.New("field ServiceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/services/" + fmt.Sprint(req.ServiceID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Service

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteService: Delete the service associated with the given ID.
func (s *API) DeleteService(req *DeleteServiceRequest, opts ...scw.RequestOption) (*Service, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServiceID) == "" {
		return nil, errors.New("field ServiceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/services/" + fmt.Sprint(req.ServiceID) + "",
	}

	var resp Service

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListServices: List services.
func (s *API) ListServices(req *ListServicesRequest, opts ...scw.RequestOption) (*ListServicesResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/services",
		Query:  query,
	}

	var resp ListServicesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// InstallServer: Install an OS on the server associated with the given ID.
func (s *API) InstallServer(req *InstallServerRequest, opts ...scw.RequestOption) (*ServerInstall, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/install",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ServerInstall

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetServerInstall: Get the server installation status associated with the given server ID.
func (s *API) GetServerInstall(req *GetServerInstallRequest, opts ...scw.RequestOption) (*ServerInstall, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/install",
	}

	var resp ServerInstall

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CancelServerInstall: Cancels the current server installation associated with the given server ID.
func (s *API) CancelServerInstall(req *CancelServerInstallRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/cancel-install",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetServerDefaultPartitioning: Get the server default partitioning schema associated with the given server ID and OS ID.
func (s *API) GetServerDefaultPartitioning(req *GetServerDefaultPartitioningRequest, opts ...scw.RequestOption) (*ServerDefaultPartitioning, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	if fmt.Sprint(req.OsID) == "" {
		return nil, errors.New("field OsID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/partitioning/" + fmt.Sprint(req.OsID) + "",
	}

	var resp ServerDefaultPartitioning

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StartBMCAccess: Start BMC (Baseboard Management Controller) access associated with the given ID.
// The BMC (Baseboard Management Controller) access is available one hour after the installation of the server.
func (s *API) StartBMCAccess(req *StartBMCAccessRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/bmc-access",
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

// GetBMCAccess: Get the BMC (Baseboard Management Controller) access associated with the given ID.
func (s *API) GetBMCAccess(req *GetBMCAccessRequest, opts ...scw.RequestOption) (*BMCAccess, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/bmc-access",
	}

	var resp BMCAccess

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StopBMCAccess: Stop BMC (Baseboard Management Controller) access associated with the given ID.
func (s *API) StopBMCAccess(req *StopBMCAccessRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/bmc-access",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListOffers: List all available server offers.
func (s *API) ListOffers(req *ListOffersRequest, opts ...scw.RequestOption) (*ListOffersResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "commercial_range", req.CommercialRange)
	parameter.AddToQuery(query, "catalog", req.Catalog)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "is_failover_ip", req.IsFailoverIP)
	parameter.AddToQuery(query, "is_failover_block", req.IsFailoverBlock)
	if len(req.SoldIn) != 0 {
		parameter.AddToQuery(query, "sold_in", strings.Join(req.SoldIn, ","))
	}
	parameter.AddToQuery(query, "available_only", req.AvailableOnly)
	parameter.AddToQuery(query, "is_rpn_san", req.IsRpnSan)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/offers",
		Query:  query,
	}

	var resp ListOffersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOffer: Return specific offer for the given ID.
func (s *API) GetOffer(req *GetOfferRequest, opts ...scw.RequestOption) (*Offer, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.OfferID) == "" {
		return nil, errors.New("field OfferID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/offers/" + fmt.Sprint(req.OfferID) + "",
		Query:  query,
	}

	var resp Offer

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListOS: List all available OS that can be install on a baremetal server.
func (s *API) ListOS(req *ListOSRequest, opts ...scw.RequestOption) (*ListOSResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "type", req.Type)
	parameter.AddToQuery(query, "server_id", req.ServerID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/os",
		Query:  query,
	}

	var resp ListOSResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOS: Return specific OS for the given ID.
func (s *API) GetOS(req *GetOSRequest, opts ...scw.RequestOption) (*OS, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "server_id", req.ServerID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.OsID) == "" {
		return nil, errors.New("field OsID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/os/" + fmt.Sprint(req.OsID) + "",
		Query:  query,
	}

	var resp OS

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateReverse: Update reverse of ip associated with the given ID.
func (s *API) UpdateReverse(req *UpdateReverseRequest, opts ...scw.RequestOption) (*IP, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return nil, errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/reverses/" + fmt.Sprint(req.IPID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateFailoverIPs: Order X failover IPs.
func (s *API) CreateFailoverIPs(req *CreateFailoverIPsRequest, opts ...scw.RequestOption) (*CreateFailoverIPsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/failover-ips",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateFailoverIPsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachFailoverIPs: Attach failovers on the server associated with the given ID.
func (s *API) AttachFailoverIPs(req *AttachFailoverIPsRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/failover-ips/attach",
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

// DetachFailoverIPs: Detach failovers on the server associated with the given ID.
func (s *API) DetachFailoverIPs(req *DetachFailoverIPsRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/failover-ips/detach",
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

// AttachFailoverIPToMacAddress: Attach a failover IP to a MAC address.
func (s *API) AttachFailoverIPToMacAddress(req *AttachFailoverIPToMacAddressRequest, opts ...scw.RequestOption) (*IP, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return nil, errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/failover-ips/" + fmt.Sprint(req.IPID) + "/attach-to-mac-address",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachFailoverIPFromMacAddress: Detach a failover IP from a MAC address.
func (s *API) DetachFailoverIPFromMacAddress(req *DetachFailoverIPFromMacAddressRequest, opts ...scw.RequestOption) (*IP, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return nil, errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/failover-ips/" + fmt.Sprint(req.IPID) + "/detach-from-mac-address",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteFailoverIP: Delete the failover associated with the given ID.
func (s *API) DeleteFailoverIP(req *DeleteFailoverIPRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/failover-ips/" + fmt.Sprint(req.IPID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListFailoverIPs: List failovers servers for project.
func (s *API) ListFailoverIPs(req *ListFailoverIPsRequest, opts ...scw.RequestOption) (*ListFailoverIPsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "search", req.Search)
	parameter.AddToQuery(query, "only_available", req.OnlyAvailable)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/failover-ips",
		Query:  query,
	}

	var resp ListFailoverIPsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFailoverIP: Get the server associated with the given ID.
func (s *API) GetFailoverIP(req *GetFailoverIPRequest, opts ...scw.RequestOption) (*FailoverIP, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return nil, errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/failover-ips/" + fmt.Sprint(req.IPID) + "",
	}

	var resp FailoverIP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRemainingQuota: Get remaining quota.
func (s *API) GetRemainingQuota(req *GetRemainingQuotaRequest, opts ...scw.RequestOption) (*GetRemainingQuotaResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/remaining-quota",
		Query:  query,
	}

	var resp GetRemainingQuotaResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRaid: Return raid for the given server ID.
func (s *API) GetRaid(req *GetRaidRequest, opts ...scw.RequestOption) (*Raid, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/raid",
	}

	var resp Raid

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateRaid: Update RAID associated with the given server ID.
func (s *API) UpdateRaid(req *UpdateRaidRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/update-raid",
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

// StartRescue: Start in rescue the server associated with the given ID.
func (s *API) StartRescue(req *StartRescueRequest, opts ...scw.RequestOption) (*Rescue, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/rescue",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Rescue

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRescue: Return rescue information for the given server ID.
func (s *API) GetRescue(req *GetRescueRequest, opts ...scw.RequestOption) (*Rescue, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/rescue",
	}

	var resp Rescue

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StopRescue: Stop rescue on the server associated with the given ID.
func (s *API) StopRescue(req *StopRescueRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/rescue",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// Dedibox Phoenix Billing API.
type BillingAPI struct {
	client *scw.Client
}

// NewBillingAPI returns a BillingAPI object from a Scaleway client.
func NewBillingAPI(client *scw.Client) *BillingAPI {
	return &BillingAPI{
		client: client,
	}
}

// ListInvoices:
func (s *BillingAPI) ListInvoices(req *BillingAPIListInvoicesRequest, opts ...scw.RequestOption) (*ListInvoicesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/invoices",
		Query:  query,
	}

	var resp ListInvoicesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetInvoice:
func (s *BillingAPI) GetInvoice(req *BillingAPIGetInvoiceRequest, opts ...scw.RequestOption) (*Invoice, error) {
	var err error

	if fmt.Sprint(req.InvoiceID) == "" {
		return nil, errors.New("field InvoiceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/invoices/" + fmt.Sprint(req.InvoiceID) + "",
	}

	var resp Invoice

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DownloadInvoice:
func (s *BillingAPI) DownloadInvoice(req *BillingAPIDownloadInvoiceRequest, opts ...scw.RequestOption) (*scw.File, error) {
	var err error

	if fmt.Sprint(req.InvoiceID) == "" {
		return nil, errors.New("field InvoiceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/invoices/" + fmt.Sprint(req.InvoiceID) + "/download",
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListRefunds:
func (s *BillingAPI) ListRefunds(req *BillingAPIListRefundsRequest, opts ...scw.RequestOption) (*ListRefundsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/refunds",
		Query:  query,
	}

	var resp ListRefundsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRefund:
func (s *BillingAPI) GetRefund(req *BillingAPIGetRefundRequest, opts ...scw.RequestOption) (*Refund, error) {
	var err error

	if fmt.Sprint(req.RefundID) == "" {
		return nil, errors.New("field RefundID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/refunds/" + fmt.Sprint(req.RefundID) + "",
	}

	var resp Refund

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DownloadRefund:
func (s *BillingAPI) DownloadRefund(req *BillingAPIDownloadRefundRequest, opts ...scw.RequestOption) (*scw.File, error) {
	var err error

	if fmt.Sprint(req.RefundID) == "" {
		return nil, errors.New("field RefundID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/refunds/" + fmt.Sprint(req.RefundID) + "/download",
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CanOrder:
func (s *BillingAPI) CanOrder(req *BillingAPICanOrderRequest, opts ...scw.RequestOption) (*CanOrderResponse, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/can-order",
		Query:  query,
	}

	var resp CanOrderResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Dedibox Phoenix IPv6 Block API.
type IPv6BlockAPI struct {
	client *scw.Client
}

// NewIPv6BlockAPI returns a IPv6BlockAPI object from a Scaleway client.
func NewIPv6BlockAPI(client *scw.Client) *IPv6BlockAPI {
	return &IPv6BlockAPI{
		client: client,
	}
}

// GetIPv6BlockQuotas: Get IPv6 block quota with the given project ID.
// /48 one per organization.
// /56 link to your number of server.
// /64 link to your number of failover IP.
func (s *IPv6BlockAPI) GetIPv6BlockQuotas(req *IPv6BlockAPIGetIPv6BlockQuotasRequest, opts ...scw.RequestOption) (*GetIPv6BlockQuotasResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/ipv6-block-quotas",
		Query:  query,
	}

	var resp GetIPv6BlockQuotasResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateIPv6Block: Create IPv6 block associated with the given project ID.
func (s *IPv6BlockAPI) CreateIPv6Block(req *IPv6BlockAPICreateIPv6BlockRequest, opts ...scw.RequestOption) (*IPv6Block, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/ipv6-block",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IPv6Block

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListIPv6Blocks: List IPv6 blocks associated given project ID.
func (s *IPv6BlockAPI) ListIPv6Blocks(req *IPv6BlockAPIListIPv6BlocksRequest, opts ...scw.RequestOption) (*ListIPv6BlocksResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/ipv6-blocks",
		Query:  query,
	}

	var resp ListIPv6BlocksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIPv6Block: Get the first IPv6 block associated with the given project ID.
func (s *IPv6BlockAPI) GetIPv6Block(req *IPv6BlockAPIGetIPv6BlockRequest, opts ...scw.RequestOption) (*IPv6Block, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/ipv6-block",
		Query:  query,
	}

	var resp IPv6Block

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateIPv6Block: Update DNS associated to IPv6 block.
// If DNS is used, minimum of 2 is necessary and maximum of 5 (no duplicate).
func (s *IPv6BlockAPI) UpdateIPv6Block(req *IPv6BlockAPIUpdateIPv6BlockRequest, opts ...scw.RequestOption) (*IPv6Block, error) {
	var err error

	if fmt.Sprint(req.BlockID) == "" {
		return nil, errors.New("field BlockID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/dedibox/v1/ipv6-blocks/" + fmt.Sprint(req.BlockID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IPv6Block

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteIPv6Block: Delete IPv6 block subnet with the given ID.
func (s *IPv6BlockAPI) DeleteIPv6Block(req *IPv6BlockAPIDeleteIPv6BlockRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.BlockID) == "" {
		return errors.New("field BlockID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/dedibox/v1/ipv6-blocks/" + fmt.Sprint(req.BlockID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// CreateIPv6BlockSubnet: Create IPv6 block subnet for the given IP ID.
// /48 could create subnet in /56 (quota link to your number of server).
// /56 could create subnet in /64 (quota link to your number of failover IP).
func (s *IPv6BlockAPI) CreateIPv6BlockSubnet(req *IPv6BlockAPICreateIPv6BlockSubnetRequest, opts ...scw.RequestOption) (*IPv6Block, error) {
	var err error

	if fmt.Sprint(req.BlockID) == "" {
		return nil, errors.New("field BlockID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/ipv6-blocks/" + fmt.Sprint(req.BlockID) + "/subnets",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IPv6Block

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListIPv6BlockSubnetsAvailable: List all available IPv6 block subnets for given IP ID.
func (s *IPv6BlockAPI) ListIPv6BlockSubnetsAvailable(req *IPv6BlockAPIListIPv6BlockSubnetsAvailableRequest, opts ...scw.RequestOption) (*ListIPv6BlockSubnetsAvailableResponse, error) {
	var err error

	if fmt.Sprint(req.BlockID) == "" {
		return nil, errors.New("field BlockID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/ipv6-blocks/" + fmt.Sprint(req.BlockID) + "/subnets",
	}

	var resp ListIPv6BlockSubnetsAvailableResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Dedibox Phoenix RPN API.
type RpnAPI struct {
	client *scw.Client
}

// NewRpnAPI returns a RpnAPI object from a Scaleway client.
func NewRpnAPI(client *scw.Client) *RpnAPI {
	return &RpnAPI{
		client: client,
	}
}

// ListRpnServerCapabilities:
func (s *RpnAPI) ListRpnServerCapabilities(req *RpnAPIListRpnServerCapabilitiesRequest, opts ...scw.RequestOption) (*ListRpnServerCapabilitiesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/rpn/server-capabilities",
		Query:  query,
	}

	var resp ListRpnServerCapabilitiesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRpnStatus:
func (s *RpnAPI) GetRpnStatus(req *RpnAPIGetRpnStatusRequest, opts ...scw.RequestOption) (*GetRpnStatusResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "rpnv1_group_id", req.Rpnv1GroupID)
	parameter.AddToQuery(query, "rpnv2_group_id", req.Rpnv2GroupID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/rpn/status",
		Query:  query,
	}

	var resp GetRpnStatusResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Dedibox Phoenix RPN SAN API.
type RpnSanAPI struct {
	client *scw.Client
}

// NewRpnSanAPI returns a RpnSanAPI object from a Scaleway client.
func NewRpnSanAPI(client *scw.Client) *RpnSanAPI {
	return &RpnSanAPI{
		client: client,
	}
}

// ListRpnSans:
func (s *RpnSanAPI) ListRpnSans(req *RpnSanAPIListRpnSansRequest, opts ...scw.RequestOption) (*ListRpnSansResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/rpn-sans",
		Query:  query,
	}

	var resp ListRpnSansResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRpnSan:
func (s *RpnSanAPI) GetRpnSan(req *RpnSanAPIGetRpnSanRequest, opts ...scw.RequestOption) (*RpnSan, error) {
	var err error

	if fmt.Sprint(req.RpnSanID) == "" {
		return nil, errors.New("field RpnSanID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/rpn-sans/" + fmt.Sprint(req.RpnSanID) + "",
	}

	var resp RpnSan

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteRpnSan:
func (s *RpnSanAPI) DeleteRpnSan(req *RpnSanAPIDeleteRpnSanRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.RpnSanID) == "" {
		return errors.New("field RpnSanID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/dedibox/v1/rpn-sans/" + fmt.Sprint(req.RpnSanID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// CreateRpnSan:
func (s *RpnSanAPI) CreateRpnSan(req *RpnSanAPICreateRpnSanRequest, opts ...scw.RequestOption) (*Service, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/rpn-sans",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Service

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListIPs:
func (s *RpnSanAPI) ListIPs(req *RpnSanAPIListIPsRequest, opts ...scw.RequestOption) (*ListIPsResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "type", req.Type)

	if fmt.Sprint(req.RpnSanID) == "" {
		return nil, errors.New("field RpnSanID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/rpn-sans/" + fmt.Sprint(req.RpnSanID) + "/ips",
		Query:  query,
	}

	var resp ListIPsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddIP:
func (s *RpnSanAPI) AddIP(req *RpnSanAPIAddIPRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.RpnSanID) == "" {
		return errors.New("field RpnSanID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/rpn-sans/" + fmt.Sprint(req.RpnSanID) + "/ips",
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

// RemoveIP:
func (s *RpnSanAPI) RemoveIP(req *RpnSanAPIRemoveIPRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.RpnSanID) == "" {
		return errors.New("field RpnSanID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/dedibox/v1/rpn-sans/" + fmt.Sprint(req.RpnSanID) + "/ips",
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

// ListAvailableIPs:
func (s *RpnSanAPI) ListAvailableIPs(req *RpnSanAPIListAvailableIPsRequest, opts ...scw.RequestOption) (*ListIPsResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "type", req.Type)

	if fmt.Sprint(req.RpnSanID) == "" {
		return nil, errors.New("field RpnSanID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/rpn-sans/" + fmt.Sprint(req.RpnSanID) + "/available-ips",
		Query:  query,
	}

	var resp ListIPsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Dedibox Phoenix RPN v1 API.
type RpnV1API struct {
	client *scw.Client
}

// NewRpnV1API returns a RpnV1API object from a Scaleway client.
func NewRpnV1API(client *scw.Client) *RpnV1API {
	return &RpnV1API{
		client: client,
	}
}

// ListRpnGroups:
func (s *RpnV1API) ListRpnGroups(req *RpnV1ApiListRpnGroupsRequest, opts ...scw.RequestOption) (*ListRpnGroupsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/rpnv1/groups",
		Query:  query,
	}

	var resp ListRpnGroupsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRpnGroup:
func (s *RpnV1API) GetRpnGroup(req *RpnV1ApiGetRpnGroupRequest, opts ...scw.RequestOption) (*RpnGroup, error) {
	var err error

	if fmt.Sprint(req.GroupID) == "" {
		return nil, errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/rpnv1/groups/" + fmt.Sprint(req.GroupID) + "",
	}

	var resp RpnGroup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateRpnGroup:
func (s *RpnV1API) CreateRpnGroup(req *RpnV1ApiCreateRpnGroupRequest, opts ...scw.RequestOption) (*RpnGroup, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/rpnv1/groups",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp RpnGroup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteRpnGroup:
func (s *RpnV1API) DeleteRpnGroup(req *RpnV1ApiDeleteRpnGroupRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.GroupID) == "" {
		return errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/dedibox/v1/rpnv1/groups/" + fmt.Sprint(req.GroupID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// UpdateRpnGroupName:
func (s *RpnV1API) UpdateRpnGroupName(req *RpnV1ApiUpdateRpnGroupNameRequest, opts ...scw.RequestOption) (*RpnGroup, error) {
	var err error

	if fmt.Sprint(req.GroupID) == "" {
		return nil, errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/dedibox/v1/rpnv1/groups/" + fmt.Sprint(req.GroupID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp RpnGroup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListRpnGroupMembers:
func (s *RpnV1API) ListRpnGroupMembers(req *RpnV1ApiListRpnGroupMembersRequest, opts ...scw.RequestOption) (*ListRpnGroupMembersResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.GroupID) == "" {
		return nil, errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/rpnv1/groups/" + fmt.Sprint(req.GroupID) + "/members",
		Query:  query,
	}

	var resp ListRpnGroupMembersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RpnGroupInvite:
func (s *RpnV1API) RpnGroupInvite(req *RpnV1ApiRpnGroupInviteRequest, opts ...scw.RequestOption) error {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.GroupID) == "" {
		return errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/rpnv1/groups/" + fmt.Sprint(req.GroupID) + "/invite",
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

// LeaveRpnGroup:
func (s *RpnV1API) LeaveRpnGroup(req *RpnV1ApiLeaveRpnGroupRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.GroupID) == "" {
		return errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/rpnv1/groups/" + fmt.Sprint(req.GroupID) + "/leave",
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

// AddRpnGroupMembers:
func (s *RpnV1API) AddRpnGroupMembers(req *RpnV1ApiAddRpnGroupMembersRequest, opts ...scw.RequestOption) (*RpnGroup, error) {
	var err error

	if fmt.Sprint(req.GroupID) == "" {
		return nil, errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/rpnv1/groups/" + fmt.Sprint(req.GroupID) + "/members",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp RpnGroup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteRpnGroupMembers:
func (s *RpnV1API) DeleteRpnGroupMembers(req *RpnV1ApiDeleteRpnGroupMembersRequest, opts ...scw.RequestOption) (*RpnGroup, error) {
	var err error

	if fmt.Sprint(req.GroupID) == "" {
		return nil, errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/dedibox/v1/rpnv1/groups/" + fmt.Sprint(req.GroupID) + "/members",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp RpnGroup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListRpnCapableServers:
func (s *RpnV1API) ListRpnCapableServers(req *RpnV1ApiListRpnCapableServersRequest, opts ...scw.RequestOption) (*ListRpnCapableServersResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/rpnv1/capable-servers",
		Query:  query,
	}

	var resp ListRpnCapableServersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListRpnCapableSanServers:
func (s *RpnV1API) ListRpnCapableSanServers(req *RpnV1ApiListRpnCapableSanServersRequest, opts ...scw.RequestOption) (*ListRpnCapableSanServersResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/rpnv1/capable-san-servers",
		Query:  query,
	}

	var resp ListRpnCapableSanServersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListRpnInvites:
func (s *RpnV1API) ListRpnInvites(req *RpnV1ApiListRpnInvitesRequest, opts ...scw.RequestOption) (*ListRpnInvitesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/rpnv1/invites",
		Query:  query,
	}

	var resp ListRpnInvitesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AcceptRpnInvite:
func (s *RpnV1API) AcceptRpnInvite(req *RpnV1ApiAcceptRpnInviteRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.MemberID) == "" {
		return errors.New("field MemberID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/rpnv1/invites/" + fmt.Sprint(req.MemberID) + "/accept",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// RefuseRpnInvite:
func (s *RpnV1API) RefuseRpnInvite(req *RpnV1ApiRefuseRpnInviteRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.MemberID) == "" {
		return errors.New("field MemberID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/rpnv1/invites/" + fmt.Sprint(req.MemberID) + "/refuse",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// Dedibox Phoenix RPN v2 API.
type RpnV2API struct {
	client *scw.Client
}

// NewRpnV2API returns a RpnV2API object from a Scaleway client.
func NewRpnV2API(client *scw.Client) *RpnV2API {
	return &RpnV2API{
		client: client,
	}
}

// ListRpnV2Groups:
func (s *RpnV2API) ListRpnV2Groups(req *RpnV2ApiListRpnV2GroupsRequest, opts ...scw.RequestOption) (*ListRpnV2GroupsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/rpnv2/groups",
		Query:  query,
	}

	var resp ListRpnV2GroupsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListRpnV2Members:
func (s *RpnV2API) ListRpnV2Members(req *RpnV2ApiListRpnV2MembersRequest, opts ...scw.RequestOption) (*ListRpnV2MembersResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "type", req.Type)

	if fmt.Sprint(req.GroupID) == "" {
		return nil, errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/rpnv2/groups/" + fmt.Sprint(req.GroupID) + "/members",
		Query:  query,
	}

	var resp ListRpnV2MembersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRpnV2Group:
func (s *RpnV2API) GetRpnV2Group(req *RpnV2ApiGetRpnV2GroupRequest, opts ...scw.RequestOption) (*RpnV2Group, error) {
	var err error

	if fmt.Sprint(req.GroupID) == "" {
		return nil, errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/rpnv2/groups/" + fmt.Sprint(req.GroupID) + "",
	}

	var resp RpnV2Group

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateRpnV2Group:
func (s *RpnV2API) CreateRpnV2Group(req *RpnV2ApiCreateRpnV2GroupRequest, opts ...scw.RequestOption) (*RpnV2Group, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/rpnv2/groups",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp RpnV2Group

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteRpnV2Group:
func (s *RpnV2API) DeleteRpnV2Group(req *RpnV2ApiDeleteRpnV2GroupRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.GroupID) == "" {
		return errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/dedibox/v1/rpnv2/groups/" + fmt.Sprint(req.GroupID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// UpdateRpnV2GroupName:
func (s *RpnV2API) UpdateRpnV2GroupName(req *RpnV2ApiUpdateRpnV2GroupNameRequest, opts ...scw.RequestOption) (*RpnV2Group, error) {
	var err error

	if fmt.Sprint(req.GroupID) == "" {
		return nil, errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/dedibox/v1/rpnv2/groups/" + fmt.Sprint(req.GroupID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp RpnV2Group

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddRpnV2Members:
func (s *RpnV2API) AddRpnV2Members(req *RpnV2ApiAddRpnV2MembersRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.GroupID) == "" {
		return errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/rpnv2/groups/" + fmt.Sprint(req.GroupID) + "/members",
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

// DeleteRpnV2Members:
func (s *RpnV2API) DeleteRpnV2Members(req *RpnV2ApiDeleteRpnV2MembersRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.GroupID) == "" {
		return errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/dedibox/v1/rpnv2/groups/" + fmt.Sprint(req.GroupID) + "/members",
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

// ListRpnV2CapableResources:
func (s *RpnV2API) ListRpnV2CapableResources(req *RpnV2ApiListRpnV2CapableResourcesRequest, opts ...scw.RequestOption) (*ListRpnV2CapableResourcesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/rpnv2/groups/capable",
		Query:  query,
	}

	var resp ListRpnV2CapableResourcesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListRpnV2GroupLogs:
func (s *RpnV2API) ListRpnV2GroupLogs(req *RpnV2ApiListRpnV2GroupLogsRequest, opts ...scw.RequestOption) (*ListRpnV2GroupLogsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.GroupID) == "" {
		return nil, errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/rpnv2/groups/" + fmt.Sprint(req.GroupID) + "/logs",
		Query:  query,
	}

	var resp ListRpnV2GroupLogsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateRpnV2VlanForMembers:
func (s *RpnV2API) UpdateRpnV2VlanForMembers(req *RpnV2ApiUpdateRpnV2VlanForMembersRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.GroupID) == "" {
		return errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/dedibox/v1/rpnv2/groups/" + fmt.Sprint(req.GroupID) + "/vlan",
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

// EnableRpnV2GroupCompatibility:
func (s *RpnV2API) EnableRpnV2GroupCompatibility(req *RpnV2ApiEnableRpnV2GroupCompatibilityRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.GroupID) == "" {
		return errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/rpnv2/groups/" + fmt.Sprint(req.GroupID) + "/enable-compatibility",
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

// DisableRpnV2GroupCompatibility:
func (s *RpnV2API) DisableRpnV2GroupCompatibility(req *RpnV2ApiDisableRpnV2GroupCompatibilityRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.GroupID) == "" {
		return errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/rpnv2/groups/" + fmt.Sprint(req.GroupID) + "/disable-compatibility",
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

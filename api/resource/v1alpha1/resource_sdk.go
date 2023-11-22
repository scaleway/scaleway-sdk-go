// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package resource provides methods and message types of the resource v1alpha1 API.
package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/marshaler"
	"github.com/scaleway/scaleway-sdk-go/internal/parameter"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
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

type AccountAbuseInfoScwStatus string

const (
	AccountAbuseInfoScwStatusUnknownScwStatus  = AccountAbuseInfoScwStatus("unknown_scw_status")
	AccountAbuseInfoScwStatusScwStatusClosed   = AccountAbuseInfoScwStatus("scw_status_closed")
	AccountAbuseInfoScwStatusCustomerSent      = AccountAbuseInfoScwStatus("customer_sent")
	AccountAbuseInfoScwStatusFetched           = AccountAbuseInfoScwStatus("fetched")
	AccountAbuseInfoScwStatusIPInfoNotFound    = AccountAbuseInfoScwStatus("ip_info_not_found")
	AccountAbuseInfoScwStatusNoOwner           = AccountAbuseInfoScwStatus("no_owner")
	AccountAbuseInfoScwStatusScwStatusPending  = AccountAbuseInfoScwStatus("scw_status_pending")
	AccountAbuseInfoScwStatusPosted            = AccountAbuseInfoScwStatus("posted")
	AccountAbuseInfoScwStatusPostedNoCloseWarn = AccountAbuseInfoScwStatus("posted_no_close_warn")
	AccountAbuseInfoScwStatusToPush            = AccountAbuseInfoScwStatus("to_push")
	AccountAbuseInfoScwStatusToPushNoCloseWarn = AccountAbuseInfoScwStatus("to_push_no_close_warn")
	AccountAbuseInfoScwStatusRejected          = AccountAbuseInfoScwStatus("rejected")
)

func (enum AccountAbuseInfoScwStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_scw_status"
	}
	return string(enum)
}

func (enum AccountAbuseInfoScwStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AccountAbuseInfoScwStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AccountAbuseInfoScwStatus(AccountAbuseInfoScwStatus(tmp).String())
	return nil
}

type AccountAbuseInfoStatus string

const (
	AccountAbuseInfoStatusUnknownStatus = AccountAbuseInfoStatus("unknown_status")
	AccountAbuseInfoStatusPending       = AccountAbuseInfoStatus("pending")
	AccountAbuseInfoStatusResolved      = AccountAbuseInfoStatus("resolved")
	AccountAbuseInfoStatusNew           = AccountAbuseInfoStatus("new")
	AccountAbuseInfoStatusClosed        = AccountAbuseInfoStatus("closed")
	AccountAbuseInfoStatusCancelled     = AccountAbuseInfoStatus("cancelled")
	AccountAbuseInfoStatusValidated     = AccountAbuseInfoStatus("validated")
	AccountAbuseInfoStatusConfirmed     = AccountAbuseInfoStatus("confirmed")
	AccountAbuseInfoStatusReopened      = AccountAbuseInfoStatus("reopened")
)

func (enum AccountAbuseInfoStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum AccountAbuseInfoStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AccountAbuseInfoStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AccountAbuseInfoStatus(AccountAbuseInfoStatus(tmp).String())
	return nil
}

type AccountAbuseInfoType string

const (
	AccountAbuseInfoTypeUnknownType  = AccountAbuseInfoType("unknown_type")
	AccountAbuseInfoTypeBotnet       = AccountAbuseInfoType("botnet")
	AccountAbuseInfoTypeBruteForce   = AccountAbuseInfoType("brute_force")
	AccountAbuseInfoTypeCopyright    = AccountAbuseInfoType("copyright")
	AccountAbuseInfoTypeDos          = AccountAbuseInfoType("dos")
	AccountAbuseInfoTypeIrc          = AccountAbuseInfoType("irc")
	AccountAbuseInfoTypeOpenRelay    = AccountAbuseInfoType("open_relay")
	AccountAbuseInfoTypePhishing     = AccountAbuseInfoType("phishing")
	AccountAbuseInfoTypeSecurityHole = AccountAbuseInfoType("security_hole")
	AccountAbuseInfoTypeSpam         = AccountAbuseInfoType("spam")
	AccountAbuseInfoTypeVirus        = AccountAbuseInfoType("virus")
)

func (enum AccountAbuseInfoType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum AccountAbuseInfoType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AccountAbuseInfoType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AccountAbuseInfoType(AccountAbuseInfoType(tmp).String())
	return nil
}

type AccountOrganizationInfoCustomerClass string

const (
	AccountOrganizationInfoCustomerClassCustomerClassUnknown = AccountOrganizationInfoCustomerClass("customer_class_unknown")
	AccountOrganizationInfoCustomerClassIndividual           = AccountOrganizationInfoCustomerClass("individual")
	AccountOrganizationInfoCustomerClassCorporate            = AccountOrganizationInfoCustomerClass("corporate")
	AccountOrganizationInfoCustomerClassInternal             = AccountOrganizationInfoCustomerClass("internal")
	AccountOrganizationInfoCustomerClassIntragroup           = AccountOrganizationInfoCustomerClass("intragroup")
)

func (enum AccountOrganizationInfoCustomerClass) String() string {
	if enum == "" {
		// return default value if empty
		return "customer_class_unknown"
	}
	return string(enum)
}

func (enum AccountOrganizationInfoCustomerClass) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AccountOrganizationInfoCustomerClass) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AccountOrganizationInfoCustomerClass(AccountOrganizationInfoCustomerClass(tmp).String())
	return nil
}

type AccountTokenInfoCategory string

const (
	AccountTokenInfoCategoryCategoryUnknown = AccountTokenInfoCategory("category_unknown")
	AccountTokenInfoCategorySession         = AccountTokenInfoCategory("session")
	AccountTokenInfoCategoryUserCreated     = AccountTokenInfoCategory("user_created")
	AccountTokenInfoCategoryAdminCreated    = AccountTokenInfoCategory("admin_created")
	AccountTokenInfoCategoryResellerCreated = AccountTokenInfoCategory("reseller_created")
)

func (enum AccountTokenInfoCategory) String() string {
	if enum == "" {
		// return default value if empty
		return "category_unknown"
	}
	return string(enum)
}

func (enum AccountTokenInfoCategory) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AccountTokenInfoCategory) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AccountTokenInfoCategory(AccountTokenInfoCategory(tmp).String())
	return nil
}

type AccountUserLogInfoAction string

const (
	AccountUserLogInfoActionUnknown            = AccountUserLogInfoAction("unknown")
	AccountUserLogInfoActionTokenCreated       = AccountUserLogInfoAction("token_created")
	AccountUserLogInfoActionTokenDeleted       = AccountUserLogInfoAction("token_deleted")
	AccountUserLogInfoActionPasswordChanged    = AccountUserLogInfoAction("password_changed")
	AccountUserLogInfoActionInformationChanged = AccountUserLogInfoAction("information_changed")
	AccountUserLogInfoActionUserCreated        = AccountUserLogInfoAction("user_created")
	AccountUserLogInfoActionUserDeleted        = AccountUserLogInfoAction("user_deleted")
	AccountUserLogInfoActionTwofaEnabled       = AccountUserLogInfoAction("twofa_enabled")
	AccountUserLogInfoActionTwofaDisabled      = AccountUserLogInfoAction("twofa_disabled")
)

func (enum AccountUserLogInfoAction) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum AccountUserLogInfoAction) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AccountUserLogInfoAction) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AccountUserLogInfoAction(AccountUserLogInfoAction(tmp).String())
	return nil
}

type BillingDiscountInfoDiscountMode string

const (
	BillingDiscountInfoDiscountModeUnknown                = BillingDiscountInfoDiscountMode("unknown")
	BillingDiscountInfoDiscountModeDiscountModeRate       = BillingDiscountInfoDiscountMode("discount_mode_rate")
	BillingDiscountInfoDiscountModeDiscountModeValue      = BillingDiscountInfoDiscountMode("discount_mode_value")
	BillingDiscountInfoDiscountModeDiscountModeSplittable = BillingDiscountInfoDiscountMode("discount_mode_splittable")
)

func (enum BillingDiscountInfoDiscountMode) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum BillingDiscountInfoDiscountMode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *BillingDiscountInfoDiscountMode) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = BillingDiscountInfoDiscountMode(BillingDiscountInfoDiscountMode(tmp).String())
	return nil
}

type BillingInvoiceInfoState string

const (
	BillingInvoiceInfoStateUnknown    = BillingInvoiceInfoState("unknown")
	BillingInvoiceInfoStateDraft      = BillingInvoiceInfoState("draft")
	BillingInvoiceInfoStateOutdated   = BillingInvoiceInfoState("outdated")
	BillingInvoiceInfoStateStopped    = BillingInvoiceInfoState("stopped")
	BillingInvoiceInfoStateIncomplete = BillingInvoiceInfoState("incomplete")
	BillingInvoiceInfoStateIssued     = BillingInvoiceInfoState("issued")
	BillingInvoiceInfoStatePaid       = BillingInvoiceInfoState("paid")
	BillingInvoiceInfoStateIssuing    = BillingInvoiceInfoState("issuing")
	BillingInvoiceInfoStateCancelled  = BillingInvoiceInfoState("cancelled")
)

func (enum BillingInvoiceInfoState) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum BillingInvoiceInfoState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *BillingInvoiceInfoState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = BillingInvoiceInfoState(BillingInvoiceInfoState(tmp).String())
	return nil
}

type FipIPInfoStatus string

const (
	FipIPInfoStatusUnknownStatus = FipIPInfoStatus("unknown_status")
	FipIPInfoStatusReady         = FipIPInfoStatus("ready")
	FipIPInfoStatusUpdating      = FipIPInfoStatus("updating")
	FipIPInfoStatusAttached      = FipIPInfoStatus("attached")
	FipIPInfoStatusError         = FipIPInfoStatus("error")
	FipIPInfoStatusDetaching     = FipIPInfoStatus("detaching")
	FipIPInfoStatusLocked        = FipIPInfoStatus("locked")
)

func (enum FipIPInfoStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum FipIPInfoStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *FipIPInfoStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = FipIPInfoStatus(FipIPInfoStatus(tmp).String())
	return nil
}

type InstanceImageInfoArch string

const (
	InstanceImageInfoArchUnknownArch = InstanceImageInfoArch("unknown_arch")
	InstanceImageInfoArchX86_64      = InstanceImageInfoArch("x86_64")
	InstanceImageInfoArchArm         = InstanceImageInfoArch("arm")
	InstanceImageInfoArchArm64       = InstanceImageInfoArch("arm64")
)

func (enum InstanceImageInfoArch) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_arch"
	}
	return string(enum)
}

func (enum InstanceImageInfoArch) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *InstanceImageInfoArch) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = InstanceImageInfoArch(InstanceImageInfoArch(tmp).String())
	return nil
}

type InstanceVolumeInfoType string

const (
	InstanceVolumeInfoTypeTypeUnknown = InstanceVolumeInfoType("type_unknown")
	InstanceVolumeInfoTypeLHdd        = InstanceVolumeInfoType("l_hdd")
	InstanceVolumeInfoTypeLSSD        = InstanceVolumeInfoType("l_ssd")
	InstanceVolumeInfoTypeBSSD        = InstanceVolumeInfoType("b_ssd")
	InstanceVolumeInfoTypeForeign     = InstanceVolumeInfoType("foreign")
	InstanceVolumeInfoTypeScratch     = InstanceVolumeInfoType("scratch")
)

func (enum InstanceVolumeInfoType) String() string {
	if enum == "" {
		// return default value if empty
		return "type_unknown"
	}
	return string(enum)
}

func (enum InstanceVolumeInfoType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *InstanceVolumeInfoType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = InstanceVolumeInfoType(InstanceVolumeInfoType(tmp).String())
	return nil
}

type LBServerInfoStatus string

const (
	LBServerInfoStatusToCreate    = LBServerInfoStatus("to_create")
	LBServerInfoStatusCreating    = LBServerInfoStatus("creating")
	LBServerInfoStatusRunning     = LBServerInfoStatus("running")
	LBServerInfoStatusToDelete    = LBServerInfoStatus("to_delete")
	LBServerInfoStatusToConfigure = LBServerInfoStatus("to_configure")
	LBServerInfoStatusDeleting    = LBServerInfoStatus("deleting")
	LBServerInfoStatusDeleled     = LBServerInfoStatus("deleled")
	LBServerInfoStatusToMigrate   = LBServerInfoStatus("to_migrate")
	LBServerInfoStatusMigrating   = LBServerInfoStatus("migrating")
	LBServerInfoStatusLocked      = LBServerInfoStatus("locked")
	LBServerInfoStatusLocking     = LBServerInfoStatus("locking")
	LBServerInfoStatusUnlocking   = LBServerInfoStatus("unlocking")
	LBServerInfoStatusError       = LBServerInfoStatus("error")
)

func (enum LBServerInfoStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "to_create"
	}
	return string(enum)
}

func (enum LBServerInfoStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *LBServerInfoStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = LBServerInfoStatus(LBServerInfoStatus(tmp).String())
	return nil
}

type PaymentCardInfoCardType string

const (
	PaymentCardInfoCardTypeCardTypeUnknown         = PaymentCardInfoCardType("card_type_unknown")
	PaymentCardInfoCardTypeAmericanExpress         = PaymentCardInfoCardType("american_express")
	PaymentCardInfoCardTypeBancontact              = PaymentCardInfoCardType("bancontact")
	PaymentCardInfoCardTypeCarteBleue              = PaymentCardInfoCardType("carte_bleue")
	PaymentCardInfoCardTypeDinersClubInternational = PaymentCardInfoCardType("diners_club_international")
	PaymentCardInfoCardTypeDiscover                = PaymentCardInfoCardType("discover")
	PaymentCardInfoCardTypeElectron                = PaymentCardInfoCardType("electron")
	PaymentCardInfoCardTypeEnroute                 = PaymentCardInfoCardType("enroute")
	PaymentCardInfoCardTypeJcb                     = PaymentCardInfoCardType("jcb")
	PaymentCardInfoCardTypeMaestro                 = PaymentCardInfoCardType("maestro")
	PaymentCardInfoCardTypeMastercard              = PaymentCardInfoCardType("mastercard")
	PaymentCardInfoCardTypeVisa                    = PaymentCardInfoCardType("visa")
	PaymentCardInfoCardTypeVpay                    = PaymentCardInfoCardType("vpay")
	PaymentCardInfoCardTypeUnionpay                = PaymentCardInfoCardType("unionpay")
)

func (enum PaymentCardInfoCardType) String() string {
	if enum == "" {
		// return default value if empty
		return "card_type_unknown"
	}
	return string(enum)
}

func (enum PaymentCardInfoCardType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PaymentCardInfoCardType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PaymentCardInfoCardType(PaymentCardInfoCardType(tmp).String())
	return nil
}

type PaymentCardInfoPsp string

const (
	PaymentCardInfoPspPspUnknown   = PaymentCardInfoPsp("psp_unknown")
	PaymentCardInfoPspBe2bill      = PaymentCardInfoPsp("be2bill")
	PaymentCardInfoPspStripe       = PaymentCardInfoPsp("stripe")
	PaymentCardInfoPspSepa         = PaymentCardInfoPsp("sepa")
	PaymentCardInfoPspWireTransfer = PaymentCardInfoPsp("wire_transfer")
)

func (enum PaymentCardInfoPsp) String() string {
	if enum == "" {
		// return default value if empty
		return "psp_unknown"
	}
	return string(enum)
}

func (enum PaymentCardInfoPsp) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PaymentCardInfoPsp) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PaymentCardInfoPsp(PaymentCardInfoPsp(tmp).String())
	return nil
}

type RdbBackupInfoStatus string

const (
	RdbBackupInfoStatusUnknown   = RdbBackupInfoStatus("unknown")
	RdbBackupInfoStatusCreating  = RdbBackupInfoStatus("creating")
	RdbBackupInfoStatusReady     = RdbBackupInfoStatus("ready")
	RdbBackupInfoStatusRestoring = RdbBackupInfoStatus("restoring")
	RdbBackupInfoStatusDeleting  = RdbBackupInfoStatus("deleting")
	RdbBackupInfoStatusError     = RdbBackupInfoStatus("error")
	RdbBackupInfoStatusExporting = RdbBackupInfoStatus("exporting")
	RdbBackupInfoStatusLocked    = RdbBackupInfoStatus("locked")
)

func (enum RdbBackupInfoStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum RdbBackupInfoStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RdbBackupInfoStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RdbBackupInfoStatus(RdbBackupInfoStatus(tmp).String())
	return nil
}

type RdbInstanceInfoStatus string

const (
	RdbInstanceInfoStatusUnknown      = RdbInstanceInfoStatus("unknown")
	RdbInstanceInfoStatusReady        = RdbInstanceInfoStatus("ready")
	RdbInstanceInfoStatusProvisioning = RdbInstanceInfoStatus("provisioning")
	RdbInstanceInfoStatusConfiguring  = RdbInstanceInfoStatus("configuring")
	RdbInstanceInfoStatusDeleting     = RdbInstanceInfoStatus("deleting")
	RdbInstanceInfoStatusError        = RdbInstanceInfoStatus("error")
	RdbInstanceInfoStatusAutohealing  = RdbInstanceInfoStatus("autohealing")
	RdbInstanceInfoStatusLocked       = RdbInstanceInfoStatus("locked")
	RdbInstanceInfoStatusInitializing = RdbInstanceInfoStatus("initializing")
	RdbInstanceInfoStatusDiskFull     = RdbInstanceInfoStatus("disk_full")
	RdbInstanceInfoStatusBackuping    = RdbInstanceInfoStatus("backuping")
	RdbInstanceInfoStatusSnapshotting = RdbInstanceInfoStatus("snapshotting")
	RdbInstanceInfoStatusRestarting   = RdbInstanceInfoStatus("restarting")
)

func (enum RdbInstanceInfoStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum RdbInstanceInfoStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RdbInstanceInfoStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RdbInstanceInfoStatus(RdbInstanceInfoStatus(tmp).String())
	return nil
}

type RdbSnapshotInfoStatus string

const (
	RdbSnapshotInfoStatusUnknown   = RdbSnapshotInfoStatus("unknown")
	RdbSnapshotInfoStatusCreating  = RdbSnapshotInfoStatus("creating")
	RdbSnapshotInfoStatusReady     = RdbSnapshotInfoStatus("ready")
	RdbSnapshotInfoStatusRestoring = RdbSnapshotInfoStatus("restoring")
	RdbSnapshotInfoStatusDeleting  = RdbSnapshotInfoStatus("deleting")
	RdbSnapshotInfoStatusError     = RdbSnapshotInfoStatus("error")
	RdbSnapshotInfoStatusLocked    = RdbSnapshotInfoStatus("locked")
)

func (enum RdbSnapshotInfoStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum RdbSnapshotInfoStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RdbSnapshotInfoStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RdbSnapshotInfoStatus(RdbSnapshotInfoStatus(tmp).String())
	return nil
}

type RedisClusterInfoStatus string

const (
	RedisClusterInfoStatusUnknown      = RedisClusterInfoStatus("unknown")
	RedisClusterInfoStatusReady        = RedisClusterInfoStatus("ready")
	RedisClusterInfoStatusProvisioning = RedisClusterInfoStatus("provisioning")
	RedisClusterInfoStatusConfiguring  = RedisClusterInfoStatus("configuring")
	RedisClusterInfoStatusDestroying   = RedisClusterInfoStatus("destroying")
	RedisClusterInfoStatusError        = RedisClusterInfoStatus("error")
	RedisClusterInfoStatusAutohealing  = RedisClusterInfoStatus("autohealing")
	RedisClusterInfoStatusLocked       = RedisClusterInfoStatus("locked")
	RedisClusterInfoStatusSuspended    = RedisClusterInfoStatus("suspended")
	RedisClusterInfoStatusDestroyed    = RedisClusterInfoStatus("destroyed")
	RedisClusterInfoStatusInitializing = RedisClusterInfoStatus("initializing")
	RedisClusterInfoStatusDeleting     = RedisClusterInfoStatus("deleting")
)

func (enum RedisClusterInfoStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum RedisClusterInfoStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RedisClusterInfoStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RedisClusterInfoStatus(RedisClusterInfoStatus(tmp).String())
	return nil
}

type ResourceLocality string

const (
	ResourceLocalityUnknownLocality = ResourceLocality("unknown_locality")
	ResourceLocalityGlobal          = ResourceLocality("global")
	ResourceLocalityFrRz            = ResourceLocality("fr_rz")
	ResourceLocalityFrSrr           = ResourceLocality("fr_srr")
	ResourceLocalityFrSrr1          = ResourceLocality("fr_srr_1")
	ResourceLocalityFrPar           = ResourceLocality("fr_par")
	ResourceLocalityFrPar1          = ResourceLocality("fr_par_1")
	ResourceLocalityFrPar2          = ResourceLocality("fr_par_2")
	ResourceLocalityFrPar3          = ResourceLocality("fr_par_3")
	ResourceLocalityNlAms           = ResourceLocality("nl_ams")
	ResourceLocalityNlAms1          = ResourceLocality("nl_ams_1")
	ResourceLocalityNlAms2          = ResourceLocality("nl_ams_2")
	ResourceLocalityNlAms3          = ResourceLocality("nl_ams_3")
	ResourceLocalityPlWaw           = ResourceLocality("pl_waw")
	ResourceLocalityPlWaw1          = ResourceLocality("pl_waw_1")
	ResourceLocalityPlWaw2          = ResourceLocality("pl_waw_2")
	ResourceLocalityPlWaw3          = ResourceLocality("pl_waw_3")
	ResourceLocalityFrInt           = ResourceLocality("fr_int")
	ResourceLocalityFrInt1          = ResourceLocality("fr_int_1")
	ResourceLocalityFrLab           = ResourceLocality("fr_lab")
	ResourceLocalityFrLab1          = ResourceLocality("fr_lab_1")
)

func (enum ResourceLocality) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_locality"
	}
	return string(enum)
}

func (enum ResourceLocality) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ResourceLocality) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ResourceLocality(ResourceLocality(tmp).String())
	return nil
}

type ResourceType string

const (
	ResourceTypeUnknownType                   = ResourceType("unknown_type")
	ResourceTypeInstanceServer                = ResourceType("instance_server")
	ResourceTypeInstanceIP                    = ResourceType("instance_ip")
	ResourceTypeInstanceVolume                = ResourceType("instance_volume")
	ResourceTypeInstanceImage                 = ResourceType("instance_image")
	ResourceTypeInstanceCluster               = ResourceType("instance_cluster")
	ResourceTypeInstanceHypervisor            = ResourceType("instance_hypervisor")
	ResourceTypeInstanceSession               = ResourceType("instance_session")
	ResourceTypeAccountOrganization           = ResourceType("account_organization")
	ResourceTypeAccountUser                   = ResourceType("account_user")
	ResourceTypeAccountToken                  = ResourceType("account_token")
	ResourceTypeAccountUserLog                = ResourceType("account_user_log")
	ResourceTypeAccountAbuse                  = ResourceType("account_abuse")
	ResourceTypeAccountProject                = ResourceType("account_project")
	ResourceTypeK8sCluster                    = ResourceType("k8s_cluster")
	ResourceTypeK8sPool                       = ResourceType("k8s_pool")
	ResourceTypeK8sNode                       = ResourceType("k8s_node")
	ResourceTypeBillingDiscount               = ResourceType("billing_discount")
	ResourceTypeBillingInvoice                = ResourceType("billing_invoice")
	ResourceTypePaymentCard                   = ResourceType("payment_card")
	ResourceTypeDomainDomain                  = ResourceType("domain_domain")
	ResourceTypeDNSZone                       = ResourceType("dns_zone")
	ResourceTypeVpcPrivateNetwork             = ResourceType("vpc_private_network")
	ResourceTypeVpcVpc                        = ResourceType("vpc_vpc")
	ResourceTypeVpgGateway                    = ResourceType("vpg_gateway")
	ResourceTypeVpgInstance                   = ResourceType("vpg_instance")
	ResourceTypeVpgIP                         = ResourceType("vpg_ip")
	ResourceTypeVpgNetwork                    = ResourceType("vpg_network")
	ResourceTypeAppleSiliconMac               = ResourceType("apple_silicon_mac")
	ResourceTypeAppleSiliconServer            = ResourceType("apple_silicon_server")
	ResourceTypeRdbInstance                   = ResourceType("rdb_instance")
	ResourceTypeRdbSnapshot                   = ResourceType("rdb_snapshot")
	ResourceTypeRdbBackup                     = ResourceType("rdb_backup")
	ResourceTypeBaremetalServer               = ResourceType("baremetal_server")
	ResourceTypeBaremetalDcimServer           = ResourceType("baremetal_dcim_server")
	ResourceTypeTemDomain                     = ResourceType("tem_domain")
	ResourceTypeTemEmail                      = ResourceType("tem_email")
	ResourceTypeFipIP                         = ResourceType("fip_ip")
	ResourceTypeLBServer                      = ResourceType("lb_server")
	ResourceTypeServerlessFunctionsFunction   = ResourceType("serverless_functions_function")
	ResourceTypeServerlessFunctionsNamespace  = ResourceType("serverless_functions_namespace")
	ResourceTypeServerlessFunctionsDomain     = ResourceType("serverless_functions_domain")
	ResourceTypeServerlessFunctionsCron       = ResourceType("serverless_functions_cron")
	ResourceTypeServerlessContainersContainer = ResourceType("serverless_containers_container")
	ResourceTypeServerlessContainersNamespace = ResourceType("serverless_containers_namespace")
	ResourceTypeServerlessContainersDomain    = ResourceType("serverless_containers_domain")
	ResourceTypeServerlessContainersCron      = ResourceType("serverless_containers_cron")
	ResourceTypeWbhHosting                    = ResourceType("wbh_hosting")
	ResourceTypeRedisCluster                  = ResourceType("redis_cluster")
	ResourceTypeSmSecret                      = ResourceType("sm_secret")
	ResourceTypeKmsKey                        = ResourceType("kms_key")
	ResourceTypeIpamIP                        = ResourceType("ipam_ip")
	ResourceTypeObsCockpit                    = ResourceType("obs_cockpit")
	ResourceTypeSdbInstance                   = ResourceType("sdb_instance")
	ResourceTypeServerlessJobDefinition       = ResourceType("serverless_job_definition")
	ResourceTypeServerlessJobRun              = ResourceType("serverless_job_run")
	ResourceTypeEdgPipeline                   = ResourceType("edg_pipeline")
	ResourceTypeMnqNatsAccount                = ResourceType("mnq_nats_account")
	ResourceTypeMnqNatsCredentials            = ResourceType("mnq_nats_credentials")
	ResourceTypeMnqSqsAccount                 = ResourceType("mnq_sqs_account")
	ResourceTypeMnqSqsCredentials             = ResourceType("mnq_sqs_credentials")
	ResourceTypeMnqSnsAccount                 = ResourceType("mnq_sns_account")
	ResourceTypeMnqSnsCredentials             = ResourceType("mnq_sns_credentials")
	ResourceTypeSbsVolume                     = ResourceType("sbs_volume")
	ResourceTypeSbsSnapshot                   = ResourceType("sbs_snapshot")
)

func (enum ResourceType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum ResourceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ResourceType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ResourceType(ResourceType(tmp).String())
	return nil
}

type SecretInfoType string

const (
	SecretInfoTypeUnknownType = SecretInfoType("unknown_type")
	SecretInfoTypeOpaque      = SecretInfoType("opaque")
	SecretInfoTypeCertificate = SecretInfoType("certificate")
	SecretInfoTypeKeyValue    = SecretInfoType("key_value")
)

func (enum SecretInfoType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum SecretInfoType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SecretInfoType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SecretInfoType(SecretInfoType(tmp).String())
	return nil
}

type VpgGatewayInfoStatus string

const (
	VpgGatewayInfoStatusUnknownStatus = VpgGatewayInfoStatus("unknown_status")
	VpgGatewayInfoStatusStopped       = VpgGatewayInfoStatus("stopped")
	VpgGatewayInfoStatusAllocating    = VpgGatewayInfoStatus("allocating")
	VpgGatewayInfoStatusConfiguring   = VpgGatewayInfoStatus("configuring")
	VpgGatewayInfoStatusRunning       = VpgGatewayInfoStatus("running")
	VpgGatewayInfoStatusStopping      = VpgGatewayInfoStatus("stopping")
	VpgGatewayInfoStatusFailed        = VpgGatewayInfoStatus("failed")
	VpgGatewayInfoStatusDeleting      = VpgGatewayInfoStatus("deleting")
	VpgGatewayInfoStatusDeleted       = VpgGatewayInfoStatus("deleted")
)

func (enum VpgGatewayInfoStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum VpgGatewayInfoStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *VpgGatewayInfoStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = VpgGatewayInfoStatus(VpgGatewayInfoStatus(tmp).String())
	return nil
}

type VpgInstanceInfoStatus string

const (
	VpgInstanceInfoStatusUnknownStatus = VpgInstanceInfoStatus("unknown_status")
	VpgInstanceInfoStatusAllocating    = VpgInstanceInfoStatus("allocating")
	VpgInstanceInfoStatusFree          = VpgInstanceInfoStatus("free")
	VpgInstanceInfoStatusRunning       = VpgInstanceInfoStatus("running")
	VpgInstanceInfoStatusStopping      = VpgInstanceInfoStatus("stopping")
	VpgInstanceInfoStatusFailed        = VpgInstanceInfoStatus("failed")
)

func (enum VpgInstanceInfoStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum VpgInstanceInfoStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *VpgInstanceInfoStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = VpgInstanceInfoStatus(VpgInstanceInfoStatus(tmp).String())
	return nil
}

type VpgNetworkInfoStatus string

const (
	VpgNetworkInfoStatusUnknownStatus = VpgNetworkInfoStatus("unknown_status")
	VpgNetworkInfoStatusCreated       = VpgNetworkInfoStatus("created")
	VpgNetworkInfoStatusAttaching     = VpgNetworkInfoStatus("attaching")
	VpgNetworkInfoStatusConfiguring   = VpgNetworkInfoStatus("configuring")
	VpgNetworkInfoStatusReady         = VpgNetworkInfoStatus("ready")
	VpgNetworkInfoStatusDetaching     = VpgNetworkInfoStatus("detaching")
	VpgNetworkInfoStatusDeleted       = VpgNetworkInfoStatus("deleted")
)

func (enum VpgNetworkInfoStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum VpgNetworkInfoStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *VpgNetworkInfoStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = VpgNetworkInfoStatus(VpgNetworkInfoStatus(tmp).String())
	return nil
}

// AccountAbuseInfo: account abuse info.
type AccountAbuseInfo struct {
	OnlineID *int32 `json:"online_id"`

	Date *time.Time `json:"date"`

	// Type: default value: unknown_type
	Type AccountAbuseInfoType `json:"type"`

	// Status: default value: unknown_status
	Status AccountAbuseInfoStatus `json:"status"`

	SentDate *time.Time `json:"sent_date"`

	Sender string `json:"sender"`

	ResolvedDate *time.Time `json:"resolved_date"`

	// ScwStatus: default value: unknown_scw_status
	ScwStatus AccountAbuseInfoScwStatus `json:"scw_status"`

	WarningID *string `json:"warning_id"`

	ServerZone *string `json:"server_zone"`

	ServerID *string `json:"server_id"`

	SessionID *string `json:"session_id"`
}

// AccountOrganizationInfo: account organization info.
type AccountOrganizationInfo struct {
	// CustomerClass: default value: customer_class_unknown
	CustomerClass AccountOrganizationInfoCustomerClass `json:"customer_class"`

	CreatorID string `json:"creator_id"`

	SupportID string `json:"support_id"`
}

// AccountProjectInfo: account project info.
type AccountProjectInfo struct {
}

// AccountTokenInfo: account token info.
type AccountTokenInfo struct {
	AccessKey string `json:"access_key"`

	// Category: default value: category_unknown
	Category AccountTokenInfoCategory `json:"category"`

	CreationIP *string `json:"creation_ip"`
}

// AccountUserInfo: account user info.
type AccountUserInfo struct {
	Email string `json:"email"`

	FirstName *string `json:"first_name"`

	LastName *string `json:"last_name"`

	PhoneNumbers []string `json:"phone_numbers"`

	OldEmails []string `json:"old_emails"`
}

// AccountUserLogInfo: account user log info.
type AccountUserLogInfo struct {
	UserID string `json:"user_id"`

	FromUserID *string `json:"from_user_id"`

	// Action: default value: unknown
	Action AccountUserLogInfoAction `json:"action"`

	IP *string `json:"ip"`

	Port *int32 `json:"port"`

	Country *string `json:"country"`
}

// AppleSiliconMacInfo: apple silicon mac info.
type AppleSiliconMacInfo struct {
	IP string `json:"ip"`

	SerialNumber string `json:"serial_number"`
}

// AppleSiliconServerInfo: apple silicon server info.
type AppleSiliconServerInfo struct {
}

// BillingDiscountInfo: billing discount info.
type BillingDiscountInfo struct {
	Description string `json:"description"`

	InternalDescription string `json:"internal_description"`

	Value float64 `json:"value"`

	ValueUsed *float64 `json:"value_used"`

	// Mode: default value: unknown
	Mode BillingDiscountInfoDiscountMode `json:"mode"`

	StartDate *time.Time `json:"start_date"`

	StopDate *time.Time `json:"stop_date"`

	CouponID *string `json:"coupon_id"`
}

// BillingInvoiceInfo: billing invoice info.
type BillingInvoiceInfo struct {
	VatNumber *string `json:"vat_number"`

	Number *int32 `json:"number"`

	// State: default value: unknown
	State BillingInvoiceInfoState `json:"state"`

	IssuedDate *time.Time `json:"issued_date"`

	PayerID string `json:"payer_id"`
}

// BrmDcimServerInfo: brm dcim server info.
type BrmDcimServerInfo struct {
	IP string `json:"ip"`

	IPIpmi string `json:"ip_ipmi"`

	NetboxID int32 `json:"netbox_id"`
}

// BrmServerInfo: brm server info.
type BrmServerInfo struct {
	IP string `json:"ip"`

	NetboxID int32 `json:"netbox_id"`

	ConsoleServerID int32 `json:"console_server_id"`

	IPIpmi string `json:"ip_ipmi"`
}

// DNSZoneInfo: dns zone info.
type DNSZoneInfo struct {
	DNSZone string `json:"dns_zone"`
}

// DomainDomainInfo: domain domain info.
type DomainDomainInfo struct {
	Domain string `json:"domain"`
}

// FipIPInfo: fip ip info.
type FipIPInfo struct {
	IP net.IP `json:"ip"`

	// Status: default value: unknown_status
	Status FipIPInfoStatus `json:"status"`

	MacAddress string `json:"mac_address"`

	ReverseDNS *string `json:"reverse_dns"`

	ServerID *string `json:"server_id"`
}

// InstanceClusterInfo: instance cluster info.
type InstanceClusterInfo struct {
	PlatformID int32 `json:"platform_id"`

	ClusterID int32 `json:"cluster_id"`
}

// InstanceHypervisorInfo: instance hypervisor info.
type InstanceHypervisorInfo struct {
	PlatformID int32 `json:"platform_id"`

	ClusterID int32 `json:"cluster_id"`

	HypervisorID int32 `json:"hypervisor_id"`

	NbCPU int32 `json:"nb_cpu"`

	Memory int64 `json:"memory"`

	VMSupported []string `json:"vm_supported"`
}

// InstanceIPInfo: instance ip info.
type InstanceIPInfo struct {
	Address string `json:"address"`

	Pool string `json:"pool"`

	ServerID *string `json:"server_id"`
}

// InstanceImageInfo: instance image info.
type InstanceImageInfo struct {
	IsPublic bool `json:"is_public"`

	// Arch: default value: unknown_arch
	Arch InstanceImageInfoArch `json:"arch"`
}

// InstanceServerInfo: instance server info.
type InstanceServerInfo struct {
	PublicIP string `json:"public_ip"`

	PrivateIP string `json:"private_ip"`

	PublicIPV6 string `json:"public_ip_v6"`

	Type string `json:"type"`
}

// InstanceSessionInfo: instance session info.
type InstanceSessionInfo struct {
	PublicIPV4 string `json:"public_ip_v4"`

	PublicIPV6 string `json:"public_ip_v6"`

	ServerID string `json:"server_id"`
}

// InstanceVolumeInfo: instance volume info.
type InstanceVolumeInfo struct {
	// Type: default value: type_unknown
	Type InstanceVolumeInfoType `json:"type"`
}

// K8SClusterInfo: k8s cluster info.
type K8SClusterInfo struct {
	ClusterID string `json:"cluster_id"`
}

// K8SNodeInfo: k8s node info.
type K8SNodeInfo struct {
	NodeID string `json:"node_id"`

	PoolID string `json:"pool_id"`

	ClusterID string `json:"cluster_id"`

	PublicIPV4 *string `json:"public_ip_v4"`

	// Precisely one of InstanceID, BaremetalID must be set.
	InstanceID *string `json:"instance_id,omitempty"`

	// Precisely one of InstanceID, BaremetalID must be set.
	BaremetalID *string `json:"baremetal_id,omitempty"`
}

// K8SPoolInfo: k8s pool info.
type K8SPoolInfo struct {
	PoolID string `json:"pool_id"`

	ClusterID string `json:"cluster_id"`
}

// LBServerInfo: lb server info.
type LBServerInfo struct {
	IPID *string `json:"ip_id"`

	LBTypeID string `json:"lb_type_id"`

	// Status: default value: to_create
	Status LBServerInfoStatus `json:"status"`
}

// MnqNatsAccountInfo: mnq nats account info.
type MnqNatsAccountInfo struct {
}

// MnqNatsCredentialsInfo: mnq nats credentials info.
type MnqNatsCredentialsInfo struct {
	AccountID string `json:"account_id"`

	PublicKey string `json:"public_key"`
}

// MnqSnsAccountInfo: mnq sns account info.
type MnqSnsAccountInfo struct {
}

// MnqSnsCredentialsInfo: mnq sns credentials info.
type MnqSnsCredentialsInfo struct {
	AccountID string `json:"account_id"`

	AccessKey string `json:"access_key"`
}

// MnqSqsAccountInfo: mnq sqs account info.
type MnqSqsAccountInfo struct {
}

// MnqSqsCredentialsInfo: mnq sqs credentials info.
type MnqSqsCredentialsInfo struct {
	AccountID string `json:"account_id"`

	AccessKey string `json:"access_key"`
}

// PaymentCardInfo: payment card info.
type PaymentCardInfo struct {
	Alias *string `json:"alias"`

	Valid bool `json:"valid"`

	Selected bool `json:"selected"`

	LastNumberDigits *string `json:"last_number_digits"`

	Country *string `json:"country"`

	// CardType: default value: card_type_unknown
	CardType PaymentCardInfoCardType `json:"card_type"`

	ValidityDate *time.Time `json:"validity_date"`

	// Psp: default value: psp_unknown
	Psp PaymentCardInfoPsp `json:"psp"`
}

// RdbBackupInfo: rdb backup info.
type RdbBackupInfo struct {
	InstanceID string `json:"instance_id"`

	DatabaseName string `json:"database_name"`

	// Status: default value: unknown
	Status RdbBackupInfoStatus `json:"status"`

	Size *scw.Size `json:"size"`

	ExpiresAt *time.Time `json:"expires_at"`

	InstanceName string `json:"instance_name"`

	SameRegion bool `json:"same_region"`
}

// RdbInstanceInfo: rdb instance info.
type RdbInstanceInfo struct {
	// Status: default value: unknown
	Status RdbInstanceInfoStatus `json:"status"`

	Engine string `json:"engine"`

	IsHaCluster bool `json:"is_ha_cluster"`

	NodeType string `json:"node_type"`

	BackupSameRegion bool `json:"backup_same_region"`
}

// RdbSnapshotInfo: rdb snapshot info.
type RdbSnapshotInfo struct {
	InstanceID string `json:"instance_id"`

	// Status: default value: unknown
	Status RdbSnapshotInfoStatus `json:"status"`

	Size *scw.Size `json:"size"`

	ExpiresAt *time.Time `json:"expires_at"`

	InstanceName string `json:"instance_name"`

	NodeType string `json:"node_type"`
}

// RedisClusterInfo: redis cluster info.
type RedisClusterInfo struct {
	// Status: default value: unknown
	Status RedisClusterInfoStatus `json:"status"`

	Version string `json:"version"`

	ClusterSize uint32 `json:"cluster_size"`

	NodeType string `json:"node_type"`
}

// SbsSnapshotInfo: sbs snapshot info.
type SbsSnapshotInfo struct {
	CephName string `json:"ceph_name"`
}

// SbsVolumeInfo: sbs volume info.
type SbsVolumeInfo struct {
	CephName string `json:"ceph_name"`
}

// SecretInfo: secret info.
type SecretInfo struct {
	VersionCount uint32 `json:"version_count"`

	// Type: default value: unknown_type
	Type SecretInfoType `json:"type"`

	Path string `json:"path"`
}

// ServerlessContainerInfo: serverless container info.
type ServerlessContainerInfo struct {
	NamespaceID string `json:"namespace_id"`

	DomainName string `json:"domain_name"`

	Image string `json:"image"`
}

// ServerlessCronInfo: serverless cron info.
type ServerlessCronInfo struct {
	FunctionID string `json:"function_id"`

	FunctionName string `json:"function_name"`
}

// ServerlessDomainInfo: serverless domain info.
type ServerlessDomainInfo struct {
	Hostname string `json:"hostname"`

	FunctionID string `json:"function_id"`

	FunctionName string `json:"function_name"`

	URL string `json:"url"`
}

// ServerlessFunctionInfo: serverless function info.
type ServerlessFunctionInfo struct {
	NamespaceID string `json:"namespace_id"`

	Handler string `json:"handler"`

	DomainName string `json:"domain_name"`

	Image string `json:"image"`
}

// ServerlessNamespaceInfo: serverless namespace info.
type ServerlessNamespaceInfo struct {
	RegistryNamespaceID string `json:"registry_namespace_id"`

	RegistryEndpoint string `json:"registry_endpoint"`
}

// VpcPrivateNetworkInfo: vpc private network info.
type VpcPrivateNetworkInfo struct {
	VxlanNetworkIDentifier int32 `json:"vxlan_network_identifier"`

	VpcID string `json:"vpc_id"`
}

// VpcVpcInfo: vpc vpc info.
type VpcVpcInfo struct {
}

// VpgGatewayInfo: vpg gateway info.
type VpgGatewayInfo struct {
	PublicIPID *string `json:"public_ip_id"`

	GatewayTypeID string `json:"gateway_type_id"`

	// Status: default value: unknown_status
	Status VpgGatewayInfoStatus `json:"status"`
}

// VpgIPInfo: vpg ip info.
type VpgIPInfo struct {
	FlexibleIPID string `json:"flexible_ip_id"`

	IP net.IP `json:"ip"`

	ReverseDNS *string `json:"reverse_dns"`
}

// VpgInstanceInfo: vpg instance info.
type VpgInstanceInfo struct {
	InstanceType string `json:"instance_type"`

	// Status: default value: unknown_status
	Status VpgInstanceInfoStatus `json:"status"`

	IP *net.IP `json:"ip"`

	Location *string `json:"location"`

	ImageID string `json:"image_id"`

	ImageVersion string `json:"image_version"`

	AllocatedAt *time.Time `json:"allocated_at"`

	GatewayID *string `json:"gateway_id"`
}

// VpgNetworkInfo: vpg network info.
type VpgNetworkInfo struct {
	GatewayID string `json:"gateway_id"`

	PrivateNetworkID *string `json:"private_network_id"`

	PrivateNicID *string `json:"private_nic_id"`

	// Status: default value: unknown_status
	Status VpgNetworkInfoStatus `json:"status"`

	PrivateNicMac *string `json:"private_nic_mac"`

	DHCPID *string `json:"dhcp_id"`

	DHCPEnabled bool `json:"dhcp_enabled"`

	IP *net.IP `json:"ip"`
}

// Resource: resource.
type Resource struct {
	ID string `json:"id"`

	// Type: default value: unknown_type
	Type ResourceType `json:"type"`

	// Locality: default value: unknown_locality
	Locality ResourceLocality `json:"locality"`

	Name *string `json:"name"`

	Tags []string `json:"tags"`

	OrganizationID *string `json:"organization_id"`

	Deleted bool `json:"deleted"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	DeletedAt *time.Time `json:"deleted_at"`

	ProjectID *string `json:"project_id"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	InstanceServerInfo *InstanceServerInfo `json:"instance_server_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	InstanceIPInfo *InstanceIPInfo `json:"instance_ip_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	InstanceVolumeInfo *InstanceVolumeInfo `json:"instance_volume_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	InstanceImageInfo *InstanceImageInfo `json:"instance_image_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	InstanceClusterInfo *InstanceClusterInfo `json:"instance_cluster_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	InstanceHypervisorInfo *InstanceHypervisorInfo `json:"instance_hypervisor_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	InstanceSessionInfo *InstanceSessionInfo `json:"instance_session_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	AccountOrganizationInfo *AccountOrganizationInfo `json:"account_organization_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	AccountUserInfo *AccountUserInfo `json:"account_user_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	AccountTokenInfo *AccountTokenInfo `json:"account_token_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	AccountUserLogInfo *AccountUserLogInfo `json:"account_user_log_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	AccountAbuseInfo *AccountAbuseInfo `json:"account_abuse_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	AccountProjectInfo *AccountProjectInfo `json:"account_project_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	K8sClusterInfo *K8SClusterInfo `json:"k8s_cluster_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	K8sPoolInfo *K8SPoolInfo `json:"k8s_pool_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	K8sNodeInfo *K8SNodeInfo `json:"k8s_node_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	BillingDiscountInfo *BillingDiscountInfo `json:"billing_discount_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	BillingInvoiceInfo *BillingInvoiceInfo `json:"billing_invoice_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	PaymentCardInfo *PaymentCardInfo `json:"payment_card_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	DomainDomainInfo *DomainDomainInfo `json:"domain_domain_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	DNSZoneInfo *DNSZoneInfo `json:"dns_zone_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	VpcPrivateNetworkInfo *VpcPrivateNetworkInfo `json:"vpc_private_network_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	VpcVpcInfo *VpcVpcInfo `json:"vpc_vpc_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	VpgGatewayInfo *VpgGatewayInfo `json:"vpg_gateway_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	VpgInstanceInfo *VpgInstanceInfo `json:"vpg_instance_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	VpgIPInfo *VpgIPInfo `json:"vpg_ip_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	VpgNetworkInfo *VpgNetworkInfo `json:"vpg_network_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	AppleSiliconMacInfo *AppleSiliconMacInfo `json:"apple_silicon_mac_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	AppleSiliconServerInfo *AppleSiliconServerInfo `json:"apple_silicon_server_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	RdbInstanceInfo *RdbInstanceInfo `json:"rdb_instance_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	RdbSnapshotInfo *RdbSnapshotInfo `json:"rdb_snapshot_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	RdbBackupInfo *RdbBackupInfo `json:"rdb_backup_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	BrmServerInfo *BrmServerInfo `json:"brm_server_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	BrmDcimServerInfo *BrmDcimServerInfo `json:"brm_dcim_server_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	FipIPInfo *FipIPInfo `json:"fip_ip_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	LbaasServerInfo *LBServerInfo `json:"lbaas_server_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	ServerlessNamespaceInfo *ServerlessNamespaceInfo `json:"serverless_namespace_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	ServerlessFunctionInfo *ServerlessFunctionInfo `json:"serverless_function_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	ServerlessContainerInfo *ServerlessContainerInfo `json:"serverless_container_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	ServerlessDomainInfo *ServerlessDomainInfo `json:"serverless_domain_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	ServerlessCronInfo *ServerlessCronInfo `json:"serverless_cron_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	RedisClusterInfo *RedisClusterInfo `json:"redis_cluster_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	SecretInfo *SecretInfo `json:"secret_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	MnqNatsAccountInfo *MnqNatsAccountInfo `json:"mnq_nats_account_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	MnqNatsCredentialsInfo *MnqNatsCredentialsInfo `json:"mnq_nats_credentials_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	MnqSqsAccountInfo *MnqSqsAccountInfo `json:"mnq_sqs_account_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	MnqSqsCredentialsInfo *MnqSqsCredentialsInfo `json:"mnq_sqs_credentials_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	MnqSnsAccountInfo *MnqSnsAccountInfo `json:"mnq_sns_account_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	MnqSnsCredentialsInfo *MnqSnsCredentialsInfo `json:"mnq_sns_credentials_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	SbsVolumeInfo *SbsVolumeInfo `json:"sbs_volume_info,omitempty"`

	// Precisely one of InstanceServerInfo, InstanceIPInfo, InstanceVolumeInfo, InstanceImageInfo, InstanceClusterInfo, InstanceHypervisorInfo, InstanceSessionInfo, AccountOrganizationInfo, AccountUserInfo, AccountTokenInfo, AccountUserLogInfo, AccountAbuseInfo, AccountProjectInfo, K8sClusterInfo, K8sPoolInfo, K8sNodeInfo, BillingDiscountInfo, BillingInvoiceInfo, PaymentCardInfo, DomainDomainInfo, DNSZoneInfo, VpcPrivateNetworkInfo, VpcVpcInfo, VpgGatewayInfo, VpgInstanceInfo, VpgIPInfo, VpgNetworkInfo, AppleSiliconMacInfo, AppleSiliconServerInfo, RdbInstanceInfo, RdbSnapshotInfo, RdbBackupInfo, BrmServerInfo, BrmDcimServerInfo, FipIPInfo, LbaasServerInfo, ServerlessNamespaceInfo, ServerlessFunctionInfo, ServerlessContainerInfo, ServerlessDomainInfo, ServerlessCronInfo, RedisClusterInfo, SecretInfo, MnqNatsAccountInfo, MnqNatsCredentialsInfo, MnqSqsAccountInfo, MnqSqsCredentialsInfo, MnqSnsAccountInfo, MnqSnsCredentialsInfo, SbsVolumeInfo, SbsSnapshotInfo must be set.
	SbsSnapshotInfo *SbsSnapshotInfo `json:"sbs_snapshot_info,omitempty"`
}

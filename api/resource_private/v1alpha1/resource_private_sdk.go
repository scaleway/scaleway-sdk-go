// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package resource_private provides methods and message types of the resource_private v1alpha1 API.
package resource_private

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/errors"
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

type ResourceCountType string

const (
	ResourceCountTypeUnknown                        = ResourceCountType("unknown")
	ResourceCountTypeInstance                       = ResourceCountType("instance")
	ResourceCountTypeInstanceGpu                    = ResourceCountType("instance_gpu")
	ResourceCountTypeSnapshots                      = ResourceCountType("snapshots")
	ResourceCountTypeImages                         = ResourceCountType("images")
	ResourceCountTypeVolumes                        = ResourceCountType("volumes")
	ResourceCountTypeSecurityGroups                 = ResourceCountType("security_groups")
	ResourceCountTypePlacementGroups                = ResourceCountType("placement_groups")
	ResourceCountTypeInstanceIP                     = ResourceCountType("instance_ip")
	ResourceCountTypeRelationalDatabase             = ResourceCountType("relational_database")
	ResourceCountTypeRdbSnapshots                   = ResourceCountType("rdb_snapshots")
	ResourceCountTypeRdbBackups                     = ResourceCountType("rdb_backups")
	ResourceCountTypeRedis                          = ResourceCountType("redis")
	ResourceCountTypeDocumentDb                     = ResourceCountType("document_db")
	ResourceCountTypeObjectStorage                  = ResourceCountType("object_storage")
	ResourceCountTypeObjectStorageBucket            = ResourceCountType("object_storage_bucket")
	ResourceCountTypeLoadBalancer                   = ResourceCountType("load_balancer")
	ResourceCountTypeLoadBalancerIP                 = ResourceCountType("load_balancer_ip")
	ResourceCountTypeRegistry                       = ResourceCountType("registry")
	ResourceCountTypeKubernetes                     = ResourceCountType("kubernetes")
	ResourceCountTypeBaremetal                      = ResourceCountType("baremetal")
	ResourceCountTypeIot                            = ResourceCountType("iot")
	ResourceCountTypeServerless                     = ResourceCountType("serverless")
	ResourceCountTypeDomain                         = ResourceCountType("domain")
	ResourceCountTypeAppleSilicon                   = ResourceCountType("apple_silicon")
	ResourceCountTypeFlexibleIP                     = ResourceCountType("flexible_ip")
	ResourceCountTypeSmartLabeling                  = ResourceCountType("smart_labeling")
	ResourceCountTypeDedibox                        = ResourceCountType("dedibox")
	ResourceCountTypeContainers                     = ResourceCountType("containers")
	ResourceCountTypeFunctions                      = ResourceCountType("functions")
	ResourceCountTypeServerlessDb                   = ResourceCountType("serverless_db")
	ResourceCountTypeTransactionalEmail             = ResourceCountType("transactional_email")
	ResourceCountTypeMessaging                      = ResourceCountType("messaging")
	ResourceCountTypeWebhosting                     = ResourceCountType("webhosting")
	ResourceCountTypePublicGateway                  = ResourceCountType("public_gateway")
	ResourceCountTypeIP                             = ResourceCountType("ip")
	ResourceCountTypePrivateNetworks                = ResourceCountType("private_networks")
	ResourceCountTypeVpc                            = ResourceCountType("vpc")
	ResourceCountTypeObservabilityCockpit           = ResourceCountType("observability_cockpit")
	ResourceCountTypeObservabilityToken             = ResourceCountType("observability_token")
	ResourceCountTypeObservabilityAlertContactPoint = ResourceCountType("observability_alert_contact_point")
	ResourceCountTypeObservabilityGrafanaUser       = ResourceCountType("observability_grafana_user")
	ResourceCountTypeSecret                         = ResourceCountType("secret")
	ResourceCountTypeSecretVersion                  = ResourceCountType("secret_version")
	ResourceCountTypeKey                            = ResourceCountType("key")
	ResourceCountTypeKeyVersion                     = ResourceCountType("key_version")
	ResourceCountTypeManagedKey                     = ResourceCountType("managed_key")
	ResourceCountTypeManagedKeyVersion              = ResourceCountType("managed_key_version")
	ResourceCountTypeRegistryNs                     = ResourceCountType("registry_ns")
	ResourceCountTypeContainersNs                   = ResourceCountType("containers_ns")
	ResourceCountTypeFunctionsNs                    = ResourceCountType("functions_ns")
	ResourceCountTypeIamAPIKeys                     = ResourceCountType("iam_api_keys")
	ResourceCountTypeIamSSHKeys                     = ResourceCountType("iam_ssh_keys")
	ResourceCountTypeIamUsers                       = ResourceCountType("iam_users")
	ResourceCountTypeIamApplications                = ResourceCountType("iam_applications")
	ResourceCountTypeIamGroups                      = ResourceCountType("iam_groups")
	ResourceCountTypeIamPolicies                    = ResourceCountType("iam_policies")
	ResourceCountTypeIamPermissionSets              = ResourceCountType("iam_permission_sets")
)

func (enum ResourceCountType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum ResourceCountType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ResourceCountType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ResourceCountType(ResourceCountType(tmp).String())
	return nil
}

type ResourceCountUnit string

const (
	ResourceCountUnitNone = ResourceCountUnit("none")
	ResourceCountUnitByte = ResourceCountUnit("byte")
)

func (enum ResourceCountUnit) String() string {
	if enum == "" {
		// return default value if empty
		return "none"
	}
	return string(enum)
}

func (enum ResourceCountUnit) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ResourceCountUnit) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ResourceCountUnit(ResourceCountUnit(tmp).String())
	return nil
}

type ResourceCountWarning string

const (
	ResourceCountWarningNoWarning                = ResourceCountWarning("no_warning")
	ResourceCountWarningComputeUnassignedIP      = ResourceCountWarning("compute_unassigned_ip")
	ResourceCountWarningLoadBalancerUnassignedIP = ResourceCountWarning("load_balancer_unassigned_ip")
	ResourceCountWarningDediboxProfileNotLinked  = ResourceCountWarning("dedibox_profile_not_linked")
)

func (enum ResourceCountWarning) String() string {
	if enum == "" {
		// return default value if empty
		return "no_warning"
	}
	return string(enum)
}

func (enum ResourceCountWarning) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ResourceCountWarning) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ResourceCountWarning(ResourceCountWarning(tmp).String())
	return nil
}

// ResourceCount: resource count.
type ResourceCount struct {
	// Type: default value: unknown
	Type ResourceCountType `json:"type"`

	Values map[string]uint64 `json:"values"`

	// Unit: default value: none
	Unit ResourceCountUnit `json:"unit"`

	// Warning: default value: no_warning
	Warning ResourceCountWarning `json:"warning"`

	HasError bool `json:"has_error"`

	PreventDelete bool `json:"prevent_delete"`
}

// ConsoleAPIGetDashboardRequest: console api get dashboard request.
type ConsoleAPIGetDashboardRequest struct {
	// Precisely one of ProjectID, OrganizationID must be set.
	ProjectID *string `json:"project_id,omitempty"`

	// Precisely one of ProjectID, OrganizationID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`
}

// Dashboard: dashboard.
type Dashboard struct {
	ResourceCounts []*ResourceCount `json:"resource_counts"`
}

// Provides information for the console about all product resources.
type ConsoleAPI struct {
	client *scw.Client
}

// NewConsoleAPI returns a ConsoleAPI object from a Scaleway client.
func NewConsoleAPI(client *scw.Client) *ConsoleAPI {
	return &ConsoleAPI{
		client: client,
	}
}

// GetDashboard:
func (s *ConsoleAPI) GetDashboard(req *ConsoleAPIGetDashboardRequest, opts ...scw.RequestOption) (*Dashboard, error) {
	var err error

	defaultProjectID, exist := s.client.GetDefaultProjectID()
	if exist && req.ProjectID == nil && req.OrganizationID == nil {
		req.ProjectID = &defaultProjectID
	}

	defaultOrganizationID, exist := s.client.GetDefaultOrganizationID()
	if exist && req.ProjectID == nil && req.OrganizationID == nil {
		req.OrganizationID = &defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/resource-private/v1alpha1/dashboard",
		Query:  query,
	}

	var resp Dashboard

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package trustandsafety_private provides methods and message types of the trustandsafety_private v1 API.
package trustandsafety_private

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/internal/marshaler"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/internal/parameter"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/namegenerator"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/scw"
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

type ResourceType string

const (
	ResourceTypeUnknown                       = ResourceType("unknown")
	ResourceTypeInstanceServer                = ResourceType("instance_server")
	ResourceTypeInstanceIP                    = ResourceType("instance_ip")
	ResourceTypeInstanceSecurityGroup         = ResourceType("instance_security_group")
	ResourceTypeInstanceVolume                = ResourceType("instance_volume")
	ResourceTypeInstanceSnapshot              = ResourceType("instance_snapshot")
	ResourceTypeInstanceImage                 = ResourceType("instance_image")
	ResourceTypeLBLB                          = ResourceType("lb_lb")
	ResourceTypeLBIP                          = ResourceType("lb_ip")
	ResourceTypeKubernetesCluster             = ResourceType("kubernetes_cluster")
	ResourceTypeBaremetalServer               = ResourceType("baremetal_server")
	ResourceTypeBaremetalIPFailover           = ResourceType("baremetal_ip_failover")
	ResourceTypeObjectBucket                  = ResourceType("object_bucket")
	ResourceTypeObjectObject                  = ResourceType("object_object")
	ResourceTypeDomainDomain                  = ResourceType("domain_domain")
	ResourceTypeDomainDNS                     = ResourceType("domain_dns")
	ResourceTypeRegistryNamespace             = ResourceType("registry_namespace")
	ResourceTypeRegistryImage                 = ResourceType("registry_image")
	ResourceTypeRegistryTag                   = ResourceType("registry_tag")
	ResourceTypeVpcPrivateNetwork             = ResourceType("vpc_private_network")
	ResourceTypeRdbInstance                   = ResourceType("rdb_instance")
	ResourceTypeRdbBackup                     = ResourceType("rdb_backup")
	ResourceTypeFlexibleIPFip                 = ResourceType("flexible_ip_fip")
	ResourceTypeRdbSnapshot                   = ResourceType("rdb_snapshot")
	ResourceTypeAsServer                      = ResourceType("as_server")
	ResourceTypeVpcGwGateway                  = ResourceType("vpc_gw_gateway")
	ResourceTypeVpcGwDHCP                     = ResourceType("vpc_gw_dhcp")
	ResourceTypeVpcGwIP                       = ResourceType("vpc_gw_ip")
	ResourceTypeIotIP                         = ResourceType("iot_ip")
	ResourceTypeIotHub                        = ResourceType("iot_hub")
	ResourceTypeFncFunction                   = ResourceType("fnc_function")
	ResourceTypeFncNamespace                  = ResourceType("fnc_namespace")
	ResourceTypeCtnContainer                  = ResourceType("ctn_container")
	ResourceTypeCtnNamespace                  = ResourceType("ctn_namespace")
	ResourceTypeBaremetalPrivateNetworkMember = ResourceType("baremetal_private_network_member")
	ResourceTypeTemDomain                     = ResourceType("tem_domain")
	ResourceTypeRkvCluster                    = ResourceType("rkv_cluster")
	ResourceTypeWbhHosting                    = ResourceType("wbh_hosting")
	ResourceTypeObsCockpit                    = ResourceType("obs_cockpit")
	ResourceTypeSmSecret                      = ResourceType("sm_secret")
	ResourceTypeKmsKey                        = ResourceType("kms_key")
	ResourceTypeMnqNamespace                  = ResourceType("mnq_namespace")
	ResourceTypeIpamIP                        = ResourceType("ipam_ip")
	ResourceTypeIpfsCid                       = ResourceType("ipfs_cid")
	ResourceTypeIpfsVolume                    = ResourceType("ipfs_volume")
	ResourceTypeSdbInstance                   = ResourceType("sdb_instance")
	ResourceTypeSljJob                        = ResourceType("slj_job")
	ResourceTypeVpcVpc                        = ResourceType("vpc_vpc")
	ResourceTypeSbsVolume                     = ResourceType("sbs_volume")
	ResourceTypeSbsSnapshot                   = ResourceType("sbs_snapshot")
	ResourceTypeEdgPipeline                   = ResourceType("edg_pipeline")
	ResourceTypeIpfsNames                     = ResourceType("ipfs_names")
	ResourceTypeQaasSessions                  = ResourceType("qaas_sessions")
	ResourceTypeServerlessSqldbDatabase       = ResourceType("serverless_sqldb_database")
	ResourceTypeServerlessSqldbBackup         = ResourceType("serverless_sqldb_backup")
)

func (enum ResourceType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
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

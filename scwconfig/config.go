package scwconfig

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

// Environment variables
const (

	// Up-to-date
	scwConfigPathEnv            = "SCW_CONFIG_PATH"
	scwAccessKeyEnv             = "SCW_ACCESS_KEY"
	scwSecretKeyEnv             = "SCW_SECRET_KEY"
	scwActiveProfileEnv         = "SCW_PROFILE"
	scwApiUrlEnv                = "SCW_API_URL"
	scwInsecureEnv              = "SCW_INSECURE"
	scwDefaultOrganizationIDEnv = "SCW_DEFAULT_ORGANIZATION_ID"
	scwDefaultRegionEnv         = "SCW_DEFAULT_REGION"
	scwDefaultZoneEnv           = "SCW_DEFAULT_ZONE"

	// All deprecated (cli&terraform)
	terraformAccessKeyEnv    = "SCALEWAY_ACCESS_KEY" // used both as access key and secret key
	terraformSecretKeyEnv    = "SCALEWAY_TOKEN"
	terraformOrganizationEnv = "SCALEWAY_ORGANIZATION"
	terraformRegionEnv       = "SCALEWAY_REGION"
	cliTLSVerifyEnv          = "SCW_TLSVERIFY"
	cliOrganizationEnv       = "SCW_ORGANIZATION"
	cliRegionEnv             = "SCW_REGION"
	cliSecretKeyEnv          = "SCW_TOKEN"

	// TBD
	cliVerboseEnv         = "SCW_VERBOSE_API"
	cliDebugEnv           = "DEBUG"
	cliNoCheckVersionEnv  = "SCW_NOCHECKVERSION"
	cliTestWithRealAPIEnv = "TEST_WITH_REAL_API"
	cliSecureExecEnv      = "SCW_SECURE_EXEC"
	cliGatewayEnv         = "SCW_GATEWAY"
	cliSensitiveEnv       = "SCW_SENSITIVE"
	cliAccountApiEnv      = "SCW_ACCOUNT_API"
	cliMetadataAPIEnv     = "SCW_METADATA_API"
	cliMarketPlaceAPIEnv  = "SCW_MARKETPLACE_API"
	cliComputePar1APIEnv  = "SCW_COMPUTE_PAR1_API"
	cliComputeAms1APIEnv  = "SCW_COMPUTE_AMS1_API"
	cliCommercialTypeEnv  = "SCW_COMMERCIAL_TYPE"
	cliTargetArchEnv      = "SCW_TARGET_ARCH"
)

type Config interface {
	GetAccessKey() string
	GetSecretKey() string
	GetApiUrl() string
	GetInsecure() bool
	GetDefaultOrganizationId() string
	GetDefaultRegion() string
	GetDefaultZone() string
}

type configV2 struct {
	profile
	ActiveProfile *string             `yaml:"active_profile"`
	Profiles      map[string]*profile `yaml:"profiles"`

	withProfile string
}

type profile struct {
	AccessKey             *string `yaml:"access_key"`
	SecretKey             *string `yaml:"secret_key"`
	ApiUrl                *string `yaml:"api_url"`
	Insecure              *bool   `yaml:"insecure"`
	DefaultOrganizationId *string `yaml:"default_region"`
	DefaultRegion         *string `yaml:"default_zone"`
	DefaultZone           *string `yaml:"default_organization_id"`
}

func unmarshalConfV2(content []byte) (*configV2, error) {
	var config configV2

	err := yaml.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil

}

func (c *configV2) catchInvalidProfile() (*configV2, error) {
	activeProfile, err := c.getActiveProfile()
	if err != nil {
		return nil, err
	}

	_, exist := c.Profiles[activeProfile]
	if !exist {
		return nil, fmt.Errorf("profile %s does not exist %s", activeProfile, inConfigFile())
	}
	return c, nil
}

func (c *configV2) getActiveProfile() (string, error) {
	switch {
	case c.withProfile != "":
		return c.withProfile, nil
	case os.Getenv(scwActiveProfileEnv) != "":
		return os.Getenv(scwActiveProfileEnv), nil
	case c.ActiveProfile != nil:
		if *c.ActiveProfile == "" {
			return "", fmt.Errorf("empty active profile %s", inConfigFile())
		}
		return *c.ActiveProfile, nil
	default:
		return "", nil
	}
}
func (c *configV2) GetAccessKey() string {
	// STEP 1: getenv
	accessKey := getenv(scwAccessKeyEnv, terraformAccessKeyEnv)
	if accessKey != "" {
		return accessKey
	}

	// STEP 2: active profile
	activeProfile, _ := c.getActiveProfile()
	if activeProfile != "" && c.Profiles[activeProfile].AccessKey != nil {
		return *c.Profiles[activeProfile].AccessKey
	}

	// STEP 3: default profile
	if c.AccessKey != nil {
		return *c.AccessKey
	}
	return ""
}

func (c *configV2) GetSecretKey() string {
	// STEP 1: getenv
	secretKey := getenv(scwSecretKeyEnv, cliSecretKeyEnv, terraformSecretKeyEnv, terraformAccessKeyEnv)
	if secretKey != "" {
		return secretKey
	}

	// STEP 2: active profile
	activeProfile, _ := c.getActiveProfile()
	if activeProfile != "" && c.Profiles[activeProfile].SecretKey != nil {
		return *c.Profiles[activeProfile].SecretKey
	}

	// STEP 3: default profile
	if c.SecretKey != nil {
		return *c.SecretKey
	}
	return ""
}

func (c *configV2) GetApiUrl() string {
	// STEP 1: getenv
	apiUrl := getenv(scwApiUrlEnv)
	if apiUrl != "" {
		return apiUrl
	}

	// STEP 2: active profile
	activeProfile, _ := c.getActiveProfile()
	if activeProfile != "" && c.Profiles[activeProfile].ApiUrl != nil {
		return *c.Profiles[activeProfile].ApiUrl
	}

	// STEP 3: default profile
	if c.ApiUrl != nil {
		return *c.ApiUrl
	}
	return ""
}

func (c *configV2) GetInsecure() bool {
	// STEP 1: getenv
	insecure := getenv(scwInsecureEnv, cliTLSVerifyEnv)
	if insecure != "" {
		insecureBool, _ := strconv.ParseBool(insecure)
		return insecureBool
	}

	// STEP 2: active profile
	activeProfile, err := c.getActiveProfile()
	if err != nil {
		// TODO: log invalid bool
	}
	if activeProfile != "" && c.Profiles[activeProfile].Insecure != nil {
		return *c.Profiles[activeProfile].Insecure
	}

	// STEP 3: default profile
	if c.Insecure != nil {
		return *c.Insecure
	}
	return false
}

func (c *configV2) GetDefaultOrganizationId() string {
	// STEP 1: getenv
	defaultOrg := getenv(scwDefaultOrganizationIDEnv, cliOrganizationEnv, terraformOrganizationEnv)
	if defaultOrg != "" {
		return defaultOrg
	}

	// STEP 2: active profile
	activeProfile, _ := c.getActiveProfile()
	if activeProfile != "" && c.Profiles[activeProfile].DefaultOrganizationId != nil {
		return *c.Profiles[activeProfile].DefaultOrganizationId
	}

	// STEP 3: default profile
	if c.DefaultOrganizationId != nil {
		return *c.DefaultOrganizationId
	}
	return ""
}

func (c *configV2) GetDefaultRegion() string {
	// STEP 1: getenv
	defaultRegion := getenv(scwDefaultRegionEnv, cliRegionEnv, terraformRegionEnv)
	if defaultRegion != "" {
		return v1RegionToV2(defaultRegion)
	}

	// STEP 2: active profile
	activeProfile, _ := c.getActiveProfile()
	if activeProfile != "" && c.Profiles[activeProfile].DefaultRegion != nil {
		return v1RegionToV2(*c.Profiles[activeProfile].DefaultRegion)
	}

	// STEP 3: default profile
	if c.DefaultOrganizationId != nil {
		return v1RegionToV2(*c.DefaultRegion)
	}
	return ""
}

func (c *configV2) GetDefaultZone() string {
	// STEP 1: getenv
	defaultZone := getenv(scwDefaultZoneEnv)
	if defaultZone != "" {
		return defaultZone
	}

	// STEP 2: active profile
	activeProfile, _ := c.getActiveProfile()
	if activeProfile != "" && c.Profiles[activeProfile].DefaultZone != nil {
		return *c.Profiles[activeProfile].DefaultZone
	}

	// STEP 3: default profile
	if c.DefaultZone != nil {
		return *c.DefaultZone
	}
	return ""
}

func getenv(upToDateKey string, deprecatedKeys ...string) string {
	value := os.Getenv(upToDateKey)
	if value != "" {
		return value
	}

	for _, key := range deprecatedKeys {
		value = os.Getenv(key)
		if value != "" {
			// TODO: log deprecated
			return value
		}
	}

	return ""
}

const (
	v1RegionFrPar = "par1"
	v1RegionNlAms = "ams1"
)

// configV1 is a Scaleway CLI configuration file
type configV1 struct {
	// Organization is the identifier of the Scaleway organization
	Organization string `json:"organization"`

	// Token is the authentication token for the Scaleway organization
	Token string `json:"token"`

	// Version is the actual version of scw CLI
	Version string `json:"version"`
}

func unmarshalConfV1(content []byte) (*configV1, error) {
	var config configV1
	err := json.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}
	return &config, err
}

func (v1 *configV1) toV2() *configV2 {
	return &configV2{
		profile: profile{
			DefaultOrganizationId: &v1.Organization,
			SecretKey:             &v1.Token,
			// ignore v1 version
		},
	}
}

func v1RegionToV2(region string) string {
	switch region {
	case v1RegionFrPar:
		return "fr-par"
	case v1RegionNlAms:
		return "nl-ams"
	default:
		return region
	}
}

package scwconfig

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/scaleway/scaleway-sdk-go/utils"
	"gopkg.in/yaml.v2"
)

// Environment variables
const (
	// Up-to-date
	scwConfigPathEnv            = "SCW_CONFIG_PATH"
	scwAccessKeyEnv             = "SCW_ACCESS_KEY"
	scwSecretKeyEnv             = "SCW_SECRET_KEY"
	scwActiveProfileEnv         = "SCW_PROFILE"
	scwAPIURLEnv                = "SCW_API_URL"
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
	cliAccountAPIEnv      = "SCW_ACCOUNT_API"
	cliMetadataAPIEnv     = "SCW_METADATA_API"
	cliMarketPlaceAPIEnv  = "SCW_MARKETPLACE_API"
	cliComputePar1APIEnv  = "SCW_COMPUTE_PAR1_API"
	cliComputeAms1APIEnv  = "SCW_COMPUTE_AMS1_API"
	cliCommercialTypeEnv  = "SCW_COMMERCIAL_TYPE"
	cliTargetArchEnv      = "SCW_TARGET_ARCH"
)

// Config interface is made of getters to retrieve
// the config field by field.
type Config interface {
	GetAccessKey() (accessKey string, exist bool)
	GetSecretKey() (secretKey string, exist bool)
	GetAPIURL() (apiURL string, exist bool)
	GetInsecure() (insecure bool, exist bool)
	GetDefaultOrganizationID() (defaultOrganizationID string, exist bool)
	GetDefaultRegion() (defaultRegion utils.Region, exist bool)
	GetDefaultZone() (defaultZone utils.Zone, exist bool)
}

type configV2 struct {
	profile       `yaml:",inline"`
	ActiveProfile *string             `yaml:"active_profile,omitempty"`
	Profiles      map[string]*profile `yaml:"profiles,omitempty"`

	// withProfile is used by LoadWithProfile to handle the following priority order:
	// c.withProfile > os.Getenv("SCW_PROFILE") > c.ActiveProfile
	withProfile string
}

type profile struct {
	AccessKey             *string `yaml:"access_key,omitempty"`
	SecretKey             *string `yaml:"secret_key,omitempty"`
	APIURL                *string `yaml:"api_url,omitempty"`
	Insecure              *bool   `yaml:"insecure,omitempty"`
	DefaultOrganizationID *string `yaml:"default_organization_id,omitempty"`
	DefaultRegion         *string `yaml:"default_region,omitempty"`
	DefaultZone           *string `yaml:"default_zone,omitempty"`
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
	if activeProfile == "" {
		return c, nil
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
			return "", fmt.Errorf("active_profile key cannot be empty %s", inConfigFile())
		}
		return *c.ActiveProfile, nil
	default:
		return "", nil
	}
}

// GetAccessKey retrieve the access key from the config.
// It will check the following order:
// env, legacy env, active profile, default profile
//
// If the config is present in one of the above environment the
// value (which may be empty) is returned and the boolean is true.
// Otherwise the returned value will be empty and the boolean will
// be false.
func (c *configV2) GetAccessKey() (string, bool) {
	envValue, envExist := getenv(scwAccessKeyEnv, terraformAccessKeyEnv)
	activeProfile, _ := c.getActiveProfile()

	var accessKey string
	switch {
	case envExist:
		accessKey = envValue
	case activeProfile != "" && c.Profiles[activeProfile].AccessKey != nil:
		accessKey = *c.Profiles[activeProfile].AccessKey
	case c.AccessKey != nil:
		accessKey = *c.AccessKey
	default:
		// warning:
		return "", false
	}

	if accessKey == "" {
		// warning : empty value
	}

	return accessKey, true
}

// GetSecretKey retrieve the secret key from the config.
// It will check the following order:
// env, legacy env, active profile, default profile
//
// If the config is present in one of the above environment the
// value (which may be empty) is returned and the boolean is true.
// Otherwise the returned value will be empty and the boolean will
// be false.
func (c *configV2) GetSecretKey() (string, bool) {
	envValue, envExist := getenv(scwSecretKeyEnv, cliSecretKeyEnv, terraformSecretKeyEnv, terraformAccessKeyEnv)
	activeProfile, _ := c.getActiveProfile()

	var secretKey string
	switch {
	case envExist:
		secretKey = envValue
	case activeProfile != "" && c.Profiles[activeProfile].SecretKey != nil:
		secretKey = *c.Profiles[activeProfile].SecretKey
	case c.SecretKey != nil:
		secretKey = *c.SecretKey
	default:
		// warning:
		return "", false
	}

	if secretKey == "" {
		// warning : empty value
	}

	return secretKey, true
}

// GetAPIURL retrieve the api url from the config.
// It will check the following order:
// env, legacy env, active profile, default profile
//
// If the config is present in one of the above environment the
// value (which may be empty) is returned and the boolean is true.
// Otherwise the returned value will be empty and the boolean will
// be false.
func (c *configV2) GetAPIURL() (string, bool) {
	envValue, envExist := getenv(scwAPIURLEnv)
	activeProfile, _ := c.getActiveProfile()

	var apiURL string
	switch {
	case envExist:
		apiURL = envValue
	case activeProfile != "" && c.Profiles[activeProfile].APIURL != nil:
		apiURL = *c.Profiles[activeProfile].APIURL
	case c.APIURL != nil:
		apiURL = *c.APIURL
	default:
		return "", false
	}

	if apiURL == "" {
		// warning
	}

	return apiURL, true
}

// GetInsecure retrieve the insecure flag from the config.
// It will check the following order:
// env, legacy env, active profile, default profile
//
// If the config is present in one of the above environment the
// value (which may be empty) is returned and the boolean is true.
// Otherwise the returned value will be empty and the boolean will
// be false.
func (c *configV2) GetInsecure() (bool, bool) {
	envValue, envExist := getenv(scwInsecureEnv, cliTLSVerifyEnv)
	activeProfile, _ := c.getActiveProfile()

	var insecure bool
	var err error
	switch {
	case envExist:
		insecure, err = strconv.ParseBool(envValue)
		if err != nil {
			// todo: warning
		}
	case activeProfile != "" && c.Profiles[activeProfile].Insecure != nil:
		insecure = *c.Profiles[activeProfile].Insecure
	case c.Insecure != nil:
		insecure = *c.Insecure
	default:
		return false, false
	}

	return insecure, true
}

// GetDefaultOrganizationID retrieve the default organization ID
// from the config. It will check the following order:
// env, legacy env, active profile, default profile
//
// If the config is present in one of the above environment the
// value (which may be empty) is returned and the boolean is true.
// Otherwise the returned value will be empty and the boolean will
// be false.
func (c *configV2) GetDefaultOrganizationID() (string, bool) {
	envValue, envExist := getenv(scwDefaultOrganizationIDEnv, cliOrganizationEnv, terraformOrganizationEnv)
	activeProfile, _ := c.getActiveProfile()

	var defaultOrg string
	switch {
	case envExist:
		defaultOrg = envValue
	case activeProfile != "" && c.Profiles[activeProfile].DefaultOrganizationID != nil:
		defaultOrg = *c.Profiles[activeProfile].DefaultOrganizationID
	case c.DefaultOrganizationID != nil:
		defaultOrg = *c.DefaultOrganizationID
	default:
		return "", false
	}

	// todo: validate format
	if defaultOrg == "" {
		// todo: warning
	}

	return defaultOrg, true
}

// GetDefaultRegion retrieve the default region
// from the config. It will check the following order:
// env, legacy env, active profile, default profile
//
// If the config is present in one of the above environment the
// value (which may be empty) is returned and the boolean is true.
// Otherwise the returned value will be empty and the boolean will
// be false.
func (c *configV2) GetDefaultRegion() (utils.Region, bool) {
	envValue, envExist := getenv(scwDefaultRegionEnv, cliRegionEnv, terraformRegionEnv)
	activeProfile, _ := c.getActiveProfile()

	var defaultRegion string
	switch {
	case envExist:
		defaultRegion = v1RegionToV2(envValue)
	case activeProfile != "" && c.Profiles[activeProfile].DefaultRegion != nil:
		defaultRegion = *c.Profiles[activeProfile].DefaultRegion
	case c.DefaultRegion != nil:
		defaultRegion = *c.DefaultRegion
	default:
		return "", false
	}

	// todo: validate format
	if defaultRegion == "" {
		// todo: warning
	}

	return utils.Region(defaultRegion), true
}

// GetDefaultZone retrieve the default zone
// from the config. It will check the following order:
// env, legacy env, active profile, default profile
//
// If the config is present in one of the above environment the
// value (which may be empty) is returned and the boolean is true.
// Otherwise the returned value will be empty and the boolean will
// be false.
func (c *configV2) GetDefaultZone() (utils.Zone, bool) {
	envValue, envExist := getenv(scwDefaultZoneEnv)
	activeProfile, _ := c.getActiveProfile()

	var defaultZone string
	switch {
	case envExist:
		defaultZone = envValue
	case activeProfile != "" && c.Profiles[activeProfile].DefaultZone != nil:
		defaultZone = *c.Profiles[activeProfile].DefaultZone
	case c.DefaultZone != nil:
		defaultZone = *c.DefaultZone
	default:
		return "", false
	}

	// todo: validate format
	if defaultZone == "" {
		// todo: warning
	}

	return utils.Zone(defaultZone), true
}

func getenv(upToDateKey string, deprecatedKeys ...string) (string, bool) {
	value, exist := os.LookupEnv(upToDateKey)
	if exist {
		return value, true
	}

	for _, key := range deprecatedKeys {
		value, exist := os.LookupEnv(key)
		if exist {
			// TODO: log deprecated
			return value, true
		}
	}

	return "", false
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
			DefaultOrganizationID: &v1.Organization,
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

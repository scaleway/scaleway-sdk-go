package scw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"gopkg.in/yaml.v2"
)

type configV2 struct {
	profile       `yaml:",inline"`
	ActiveProfile *string             `yaml:"active_profile,omitempty"`
	Profiles      map[string]*profile `yaml:"profiles,omitempty"`

	// withProfile is used by LoadConfigWithProfile to handle the following priority order:
	// c.withProfile > os.Getenv("SCW_PROFILE") > c.ActiveProfile
	withProfile string
}

type profile struct {
	AccessKey        *string `yaml:"access_key,omitempty"`
	SecretKey        *string `yaml:"secret_key,omitempty"`
	APIURL           *string `yaml:"api_url,omitempty"`
	Insecure         *bool   `yaml:"insecure,omitempty"`
	DefaultProjectID *string `yaml:"default_project_id,omitempty"`
	DefaultRegion    *string `yaml:"default_region,omitempty"`
	DefaultZone      *string `yaml:"default_zone,omitempty"`
}

func (c *configV2) String() string {
	configRaw, err := yaml.Marshal(c)
	if err != nil {
		return "cannot print config:" + err.Error()
	}

	return string(configRaw)
}

func unmarshalConfV2(content []byte) (*configV2, error) {
	var config configV2

	err := yaml.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
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
			DefaultProjectID: &v1.Organization,
			SecretKey:        &v1.Token,
			// ignore v1 version
		},
	}
}

func v1RegionToV2(region string) string {
	switch region {
	case v1RegionFrPar:
		logger.Warningf("par1 is a deprecated name for region, use fr-par instead")
		return "fr-par"
	case v1RegionNlAms:
		logger.Warningf("ams1 is a deprecated name for region, use nl-ams instead")
		return "nl-ams"
	default:
		return region
	}
}

// MigrateLegacyConfig will migrate the legacy config to the V2 when none exist yet.
// Returns a boolean set to true when the migration happened.
// TODO: get accesskey from account?
func MigrateLegacyConfig() (bool, error) {
	// STEP 1: try to load config file V2
	v2Path, v2PathOk := GetConfigV2FilePath()
	if !v2PathOk || fileExist(v2Path) {
		return false, nil
	}

	// STEP 2: try to load config file V1
	v1Path, v1PathOk := GetConfigV1FilePath()
	if !v1PathOk {
		return false, nil
	}
	file, err := ioutil.ReadFile(v1Path)
	if err != nil {
		return false, nil
	}
	confV1, err := unmarshalConfV1(file)
	if err != nil {
		return false, errors.Wrap(err, "content of config file %s is invalid json", v1Path)
	}

	// STEP 3: create dir
	err = os.MkdirAll(filepath.Dir(v2Path), 0700)
	if err != nil {
		return false, errors.Wrap(err, "mkdir did not work on %s", filepath.Dir(v2Path))
	}

	// STEP 4: marshal yaml config
	newConfig := confV1.toV2()
	file, err = yaml.Marshal(newConfig)
	if err != nil {
		return false, err
	}

	// STEP 5: save config
	err = ioutil.WriteFile(v2Path, file, defaultConfigPermission)
	if err != nil {
		return false, fmt.Errorf("cannot write file %s: %s", v2Path, err)
	}

	// STEP 6: log success
	logger.Warningf("migrated existing config to %s", v2Path)
	return true, nil
}

func MustLoadConfig() *configV2 {
	c, err := LoadConfigFromPath(GetConfigPath())
	if err != nil {
		panic(err)
	}
	return c
}
func LoadConfig2() (*configV2, error) {
	return LoadConfigFromPath(GetConfigPath())
}

func LoadConfigFromPath(path string) (*configV2, error) {

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read config file: %s", err)
	}

	confV2, err := unmarshalConfV2(file)
	if err != nil {
		return nil, fmt.Errorf("content of config file %s is invalid: %s", path, err)
	}

	return confV2, nil
}

// GetProfile returns the profile corresponding to the given profile name.
func (c *configV2) GetProfile(profileName string) (*profile, error) {
	if profileName == "" {
		return nil, errors.New("active profile cannot be empty")
	}

	p, exist := c.Profiles[profileName]
	if !exist {
		return nil, errors.New("given profile %s does not exist", profileName)
	}

	return p, nil
}

// GetActiveProfile returns the active profile of the config based on the following order:
// env SCW_PROFILE > config active_profile > config root profile
func (c *configV2) GetActiveProfile() (*profile, error) {
	switch {
	case os.Getenv(scwActiveProfileEnv) != "":
		logger.Debugf("using active profile from env: %s=%s", scwActiveProfileEnv, os.Getenv(scwActiveProfileEnv))
		return c.GetProfile(os.Getenv(scwActiveProfileEnv))
	case c.ActiveProfile != nil:
		logger.Debugf("using active profile from config: active_profile=%s", scwActiveProfileEnv, *c.ActiveProfile)
		return c.GetProfile(*c.ActiveProfile)
	default:
		return &c.profile, nil
	}

}

// SaveTo will save the config to the default config path. This
// action will overwrite the previous file when it exists.
func (c *configV2) Save() error {
	return c.SaveTo(GetConfigPath())
}

// SaveTo will save the config to the given path. This action will
// overwrite the previous file when it exists.
func (c *configV2) SaveTo(path string) error {
	path = filepath.Clean(path)

	// STEP 1: marshal config
	file, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	// STEP 2: create config path dir in cases it didn't exist before
	err = os.MkdirAll(filepath.Dir(path), 0700)
	if err != nil {
		return err
	}

	// STEP 3: write new config file
	err = ioutil.WriteFile(path, file, defaultConfigPermission)
	if err != nil {
		return err
	}

	return nil

}

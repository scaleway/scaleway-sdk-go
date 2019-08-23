package scw

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"gopkg.in/yaml.v2"
)

const (
	documentationLink       = "https://github.com/scaleway/scaleway-sdk-go/blob/master/scw/README.md"
	defaultConfigPermission = 0600
)

type Config struct {
	Profile       `yaml:",inline"`
	ActiveProfile *string             `yaml:"active_profile,omitempty"`
	Profiles      map[string]*Profile `yaml:"profiles,omitempty"`
}

type Profile struct {
	AccessKey        *string `yaml:"access_key,omitempty"`
	SecretKey        *string `yaml:"secret_key,omitempty"`
	APIURL           *string `yaml:"api_url,omitempty"`
	Insecure         *bool   `yaml:"insecure,omitempty"`
	DefaultProjectID *string `yaml:"default_project_id,omitempty"`
	DefaultRegion    *string `yaml:"default_region,omitempty"`
	DefaultZone      *string `yaml:"default_zone,omitempty"`
}

func (p *Profile) String() string {

	// deep copy profile object
	copy, err := copyYamlConfig(p)
	if err != nil {
		return "cannot print config profile:" + err.Error()
	}
	c2 := copy.(*Profile)
	c2.SecretKey = hideSecretKey(c2.SecretKey)

	configRaw, _ := yaml.Marshal(c2)
	return string(configRaw)
}

func (c *Config) String() string {

	// deep copy config object
	copy, err := copyYamlConfig(c)
	if err != nil {
		return "cannot print config:" + err.Error()
	}
	c2 := copy.(*Config)
	c2.SecretKey = hideSecretKey(c2.SecretKey)
	for _, p := range c2.Profiles {
		p.SecretKey = hideSecretKey(p.SecretKey)
	}

	configRaw, _ := yaml.Marshal(c2)
	return string(configRaw)
}

func copyYamlConfig(v interface{}) (interface{}, error) {
	configRaw, err := yaml.Marshal(v)
	if err != nil {
		return nil, err
	}
	c2 := newVariableFromType(v)
	err = yaml.Unmarshal(configRaw, c2)
	if err != nil {
		return nil, err
	}
	return c2, nil
}

func hideSecretKey(key *string) *string {
	if key == nil {
		return nil
	}
	if len(*key) > 27 {
		*key = (*key)[0:9] + "xxxx-xxxx-xxxx-xxxxxxxx" + (*key)[28:]
	}

	return key
}

func unmarshalConfV2(content []byte) (*Config, error) {
	var config Config

	err := yaml.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

// MustLoadConfig is like LoadConfig but panic instead of returning an error.
func MustLoadConfig() *Config {
	c, err := LoadConfigFromPath(GetConfigPath())
	if err != nil {
		panic(err)
	}
	return c
}

// LoadConfig read the config from the default path.
func LoadConfig() (*Config, error) {
	return LoadConfigFromPath(GetConfigPath())
}

// LoadConfigFromPath read the config from the given path.
func LoadConfigFromPath(path string) (*Config, error) {

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read config file: %s", err)
	}

	_, err = unmarshalConfV1(file)
	if err == nil {
		// reject V1 config
		return nil, errors.New("found legacy config in %s: legacy config is not allowed, please switch to the new config file format: %s", path, documentationLink)
	}

	confV2, err := unmarshalConfV2(file)
	if err != nil {
		return nil, errors.Wrap(err, "content of config file %s is invalid", path)
	}

	return confV2, nil
}

// GetProfile returns the profile corresponding to the given profile name.
func (c *Config) GetProfile(profileName string) (*Profile, error) {
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
func (c *Config) GetActiveProfile() (*Profile, error) {
	switch {
	case os.Getenv(scwActiveProfileEnv) != "":
		logger.Debugf("using active profile from env: %s=%s", scwActiveProfileEnv, os.Getenv(scwActiveProfileEnv))
		return c.GetProfile(os.Getenv(scwActiveProfileEnv))
	case c.ActiveProfile != nil:
		logger.Debugf("using active profile from config: active_profile=%s", scwActiveProfileEnv, *c.ActiveProfile)
		return c.GetProfile(*c.ActiveProfile)
	default:
		return &c.Profile, nil
	}

}

// SaveTo will save the config to the default config path. This
// action will overwrite the previous file when it exists.
func (c *Config) Save() error {
	return c.SaveTo(GetConfigPath())
}

// SaveTo will save the config to the given path. This action will
// overwrite the previous file when it exists.
func (c *Config) SaveTo(path string) error {
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

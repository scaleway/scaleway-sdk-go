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

// MigrateLegacyConfig will migrate the legacy config to the V2 when none exist yet.
// Returns a boolean set to true when the migration happened.
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
	v2Path = filepath.Clean(v2Path)
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

func LoadConfig2() (*configV2, error) {
	return nil, nil
}

func LoadConfigFromPath(path string) (*configV2, error) {
	return nil, nil

}

func (c *configV2) GetProfile(profileName string) (*profile, error) {
	return nil, nil

}

func (c *configV2) GetActiveProfile() (*profile, error) {
	return nil, nil

}
func (c *configV2) Save() error {

	return nil
}
func (c *configV2) SaveTo(path string) error {
	return nil

}

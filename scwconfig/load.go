package scwconfig

import (
	"fmt"
	"io/ioutil"
	"os"
)

func LoadWithProfile(profileName string) (*configV2, error) {
	config, err := Load()
	if err != nil {
		return nil, err
	}

	config.withProfile = profileName
	return config.catchInvalidProfile()
}

func Load() (*configV2, error) {

	// STEP 1: try to load config file from SCW_CONFIG_PATH
	configPath := os.Getenv(scwConfigPathEnv)
	if configPath != "" {
		content, err := ioutil.ReadFile(configPath)
		if err != nil {
			return nil, fmt.Errorf("cannot read $SCW_CONFIG_PATH (%s): %s", configPath, err)
		}
		confV1, err := unmarshalConfV1(content)
		if err == nil {
			return confV1.toV2().catchInvalidProfile()
		}
		confV2, err := unmarshalConfV2(content)
		if err != nil {
			return nil, fmt.Errorf("content of $SCW_CONFIG_PATH (%s) is invalid yaml: %s", configPath, err)
		}

		return confV2.catchInvalidProfile()
	}

	// STEP 2: try to load new config file
	V2Path, V2PathOk := GetConfigV2FilePath()
	if V2PathOk && fileExist(V2Path) {
		file, err := ioutil.ReadFile(configPath)
		if err != nil {
			return nil, fmt.Errorf("cannot read config file %s: %s", V2Path, err)
		}

		confV2, err := unmarshalConfV2(file)
		if err != nil {
			return nil, fmt.Errorf("content of config file %s is invalid yaml: %s", configPath, err)
		}
		return confV2.catchInvalidProfile()
	}

	// STEP 3: try to load V1 config file
	V1Path, V1PathOk := GetConfigV1FilePath()
	if !V1PathOk {
		return (&configV2{}).catchInvalidProfile()
	}
	file, err := ioutil.ReadFile(V1Path)
	if err != nil {
		return (&configV2{}).catchInvalidProfile() // ignore if file doesn't exist
	}
	confV1, err := unmarshalConfV1(file)
	if err != nil {
		return nil, fmt.Errorf("content of config file %s is invalid yaml: %s", V1Path, err)
	}

	// STEP 4: migrate V1 config to V2 config file
	if V2PathOk {
		err = migrateV1toV2(confV1, V2Path)
		if err != nil {
			return nil, err
		}
	}

	return confV1.toV2().catchInvalidProfile()

}

func fileExist(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

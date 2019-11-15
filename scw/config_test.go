package scw

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/logger"
)

const emptyFile = ""

// v2 config
var (
	v2ValidAccessKey2             = "SCW234567890ABCDEFGH"
	v2ValidSecretKey2             = "6f6e6574-6f72-756c-6c74-68656d616c6c" // hint: | xxd -ps -r
	v2ValidAPIURL2                = "api-fr-par.scaleway.com"
	v2ValidInsecure2              = "true"
	v2ValidDefaultOrganizationID2 = "6d6f7264-6f72-6772-6561-74616761696e" // hint: | xxd -ps -r
	v2ValidDefaultRegion2         = string(RegionFrPar)
	v2ValidDefaultZone2           = string(ZoneFrPar2)

	v2ValidAccessKey             = "SCW1234567890ABCDEFG"
	v2ValidSecretKey             = "7363616c-6577-6573-6862-6f7579616161" // hint: | xxd -ps -r
	v2ValidAPIURL                = "api.scaleway.com"
	v2ValidInsecure              = "false"
	v2ValidDefaultOrganizationID = "6170692e-7363-616c-6577-61792e636f6d" // hint: | xxd -ps -r
	v2ValidDefaultRegion         = string(RegionNlAms)
	v2ValidDefaultZone           = string(ZoneNlAms1)
	v2ValidProfile               = "flantier"

	v2InvalidAccessKey             = "invalid"
	v2InvalidSecretKey             = "invalid"
	v2InvalidDefaultOrganizationID = "invalid"
	v2InvalidDefaultRegion         = "invalid"
	v2InvalidDefaultZone           = "invalid"

	v2SimpleValidConfig = &Config{
		Profile: Profile{
			AccessKey:             &v2ValidAccessKey,
			SecretKey:             &v2ValidSecretKey,
			DefaultOrganizationID: &v2ValidDefaultOrganizationID,
			DefaultRegion:         &v2ValidDefaultRegion,
		},
	}
	v2PartialValidConfigFile = `
access_key: ` + v2ValidAccessKey + `
secret_key: ` + v2ValidSecretKey + `
api_url: ` + v2ValidAPIURL + `
insecure: ` + v2ValidInsecure + `
default_organization_id: ` + v2ValidDefaultOrganizationID + `
default_region: ` + v2ValidDefaultRegion + `
default_zone: ` + v2ValidDefaultZone

	v2CompleteValidConfigFile = v2PartialValidConfigFile + `
profiles:
  ` + v2ValidProfile + `:
    access_key: ` + v2ValidAccessKey2 + `
    secret_key: ` + v2ValidSecretKey2 + `
    api_url: ` + v2ValidAPIURL2 + `
    insecure: ` + v2ValidInsecure2 + `
    default_organization_id: ` + v2ValidDefaultOrganizationID2 + `
    default_region: ` + v2ValidDefaultRegion2 + `
    default_zone: ` + v2ValidDefaultZone2 + `
`

	v2CompleteValidConfigWithActiveProfileFile = `
access_key: ` + v2ValidAccessKey + `
secret_key: ` + v2ValidSecretKey + `
api_url: ` + v2ValidAPIURL + `
insecure: ` + v2ValidInsecure + `
default_organization_id: ` + v2ValidDefaultOrganizationID + `
default_region: ` + v2ValidDefaultRegion + `
default_zone: ` + v2ValidDefaultZone + `
active_profile: ` + v2ValidProfile + `
profiles:
  ` + v2ValidProfile + `:
    access_key: ` + v2ValidAccessKey2 + `
    secret_key: ` + v2ValidSecretKey2 + `
    api_url: ` + v2ValidAPIURL2 + `
    insecure: ` + v2ValidInsecure2 + `
    default_organization_id: ` + v2ValidDefaultOrganizationID2 + `
    default_region: ` + v2ValidDefaultRegion2 + `
    default_zone: ` + v2ValidDefaultZone2 + `
`

	v2MixedValidConfigWithActiveProfileFile = `
access_key: ` + v2ValidAccessKey + `
secret_key: ` + v2ValidSecretKey + `
api_url: ` + v2ValidAPIURL + `
insecure: ` + v2ValidInsecure + `
default_organization_id: ` + v2ValidDefaultOrganizationID + `
default_region: ` + v2ValidDefaultRegion + `
default_zone: ` + v2ValidDefaultZone + `
active_profile: ` + v2ValidProfile + `
profiles:
  ` + v2ValidProfile + `:
    access_key: ` + v2ValidAccessKey2 + `
    secret_key: ` + v2ValidSecretKey2 + `
`

	v2SimpleValidConfigFile = `
access_key: ` + v2ValidAccessKey + `
secret_key: ` + v2ValidSecretKey + `
default_organization_id: ` + v2ValidDefaultOrganizationID + `
default_region: ` + v2ValidDefaultRegion + `
`

	v2SimpleInvalidConfigFile            = `insecure: "bool""`
	v2SimpleConfigFileWithInvalidProfile = `active_profile: flantier`

	v2FromV1ConfigFile = `secret_key: ` + v1ValidToken + `
default_organization_id: ` + v1ValidOrganizationID + `
`
)

// TestSaveConfig tests config write the correct values in the config file
func TestSaveConfig(t *testing.T) {
	tests := []struct {
		name             string
		env              map[string]string
		files            map[string]string
		config           *Config
		funcUpdateConfig func(*Config)

		expectedFiles map[string]string
	}{
		{
			name: "Custom-path config",
			env: map[string]string{
				scwConfigPathEnv: "{HOME}/valid1/test.conf",
			},
			files: map[string]string{
				"valid1/test.conf": emptyFile,
			},
			config: &Config{
				Profile: Profile{
					AccessKey:             s(v2ValidAccessKey),
					SecretKey:             s(v2ValidSecretKey),
					DefaultOrganizationID: s(v2ValidDefaultOrganizationID),
					DefaultRegion:         s(v2ValidDefaultRegion),
				},
			},
			expectedFiles: map[string]string{
				"valid1/test.conf": v2SimpleValidConfigFile,
			},
		},
		{
			name: "Default config path",
			env: map[string]string{
				"HOME": "{HOME}",
			},
			config: &Config{
				Profile: Profile{
					AccessKey:             s(v2ValidAccessKey),
					SecretKey:             s(v2ValidSecretKey),
					DefaultOrganizationID: s(v2ValidDefaultOrganizationID),
					DefaultRegion:         s(v2ValidDefaultRegion),
				},
			},
			expectedFiles: map[string]string{
				".config/scw/config.yaml": v2SimpleValidConfigFile,
			},
		},
		{
			name: "Add zone only",
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".config/scw/config.yaml": v2SimpleValidConfigFile,
			},
			config: &Config{},
			funcUpdateConfig: func(c *Config) {
				*c = *MustLoadConfig()
				c.DefaultZone = s(v2ValidDefaultZone)
			},
			expectedFiles: map[string]string{
				".config/scw/config.yaml": v2SimpleValidConfigFile + "default_zone: " + v2ValidDefaultZone + "\n",
			},
		},
		{
			name: "Add new profile",
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".config/scw/config.yaml": v2PartialValidConfigFile,
			},
			config: &Config{},
			funcUpdateConfig: func(c *Config) {
				*c = *MustLoadConfig()
				c.Profiles = map[string]*Profile{v2ValidProfile: {
					AccessKey:             s(v2ValidAccessKey2),
					SecretKey:             s(v2ValidSecretKey2),
					APIURL:                s(v2ValidAPIURL2),
					Insecure:              b(true),
					DefaultOrganizationID: s(v2ValidDefaultOrganizationID2),
					DefaultRegion:         s(v2ValidDefaultRegion2),
					DefaultZone:           s(v2ValidDefaultZone2),
				}}
			},
			expectedFiles: map[string]string{
				".config/scw/config.yaml": v2CompleteValidConfigFile,
			},
		},
	}
	// create home dir
	dir := initEnv(t)

	// delete home dir and reset env variables
	defer resetEnv(t, os.Environ(), dir)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// set up env and config file(s)
			setEnv(t, test.env, test.files, dir)

			// remove config file(s)
			defer cleanEnv(t, test.files, dir)
			defer cleanEnv(t, test.expectedFiles, dir)
			if test.funcUpdateConfig != nil {
				test.funcUpdateConfig(test.config)
			}
			testhelpers.AssertNoError(t, test.config.Save())

			// test expected files
			for fileName, expectedContent := range test.expectedFiles {
				content, err := ioutil.ReadFile(filepath.Join(dir, fileName))
				testhelpers.AssertNoError(t, err)
				testhelpers.Equals(t, expectedContent, "\n"+string(content))
			}

		})
	}
}

// TestLoadConfig tests config getters return correct values
func TestLoadProfileAndActiveProfile(t *testing.T) {

	tests := []struct {
		name       string
		env        map[string]string
		files      map[string]string
		isMigrated bool

		expectedError                 string
		expectedAccessKey             *string
		expectedSecretKey             *string
		expectedAPIURL                *string
		expectedInsecure              *bool
		expectedDefaultOrganizationID *string
		expectedDefaultRegion         *string
		expectedDefaultZone           *string
	}{
		// no env variables
		{
			name:          "No config without home dir",
			expectedError: "scaleway-sdk-go: cannot read config file: read .: is a directory",
		},
		{
			name:          "No config",
			expectedError: "scaleway-sdk-go: cannot read config file: open {HOME}/.config/scw/config.yaml: no such file or directory",
			env: map[string]string{
				"HOME": "{HOME}",
			},
		},
		{
			name: "Custom-path config is empty", // custom config path
			env: map[string]string{
				scwConfigPathEnv: "{HOME}/valid1/test.conf",
			},
			files: map[string]string{
				"valid1/test.conf": emptyFile,
			},
		},
		{
			name: "Custom-path config with valid V2",
			env: map[string]string{
				scwConfigPathEnv: "{HOME}/valid3/test.conf",
			},
			files: map[string]string{
				"valid3/test.conf": v2SimpleValidConfigFile,
			},
			expectedAccessKey:             s(v2ValidAccessKey),
			expectedSecretKey:             s(v2ValidSecretKey),
			expectedDefaultOrganizationID: s(v2ValidDefaultOrganizationID),
			expectedDefaultRegion:         s(v2ValidDefaultRegion),
		},
		{
			name: "Simple config with valid V2", // default config path
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".config/scw/config.yaml": v2SimpleValidConfigFile,
			},
			expectedAccessKey:             s(v2ValidAccessKey),
			expectedSecretKey:             s(v2ValidSecretKey),
			expectedDefaultOrganizationID: s(v2ValidDefaultOrganizationID),
			expectedDefaultRegion:         s(v2ValidDefaultRegion),
		},
		{
			name: "Simple config with valid V1",
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".scwrc": v1ValidConfigFile,
			},
			isMigrated:                    true,
			expectedSecretKey:             s(v1ValidToken),
			expectedDefaultOrganizationID: s(v1ValidOrganizationID),
		},
		{
			name: "Simple config with valid V2 and valid V1",
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".config/scw/config.yaml": v2SimpleValidConfigFile,
				".scwrc":                  v1ValidConfigFile,
			},
			expectedAccessKey:             s(v2ValidAccessKey),
			expectedSecretKey:             s(v2ValidSecretKey),
			expectedDefaultOrganizationID: s(v2ValidDefaultOrganizationID),
			expectedDefaultRegion:         s(v2ValidDefaultRegion),
		},
		{
			name: "Complete config",
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".config/scw/config.yaml": v2CompleteValidConfigFile,
			},
			expectedAccessKey:             s(v2ValidAccessKey),
			expectedSecretKey:             s(v2ValidSecretKey),
			expectedAPIURL:                s(v2ValidAPIURL),
			expectedInsecure:              b(false),
			expectedDefaultOrganizationID: s(v2ValidDefaultOrganizationID),
			expectedDefaultRegion:         s(v2ValidDefaultRegion),
			expectedDefaultZone:           s(v2ValidDefaultZone),
		},
		{
			name: "Complete config with active profile",
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".config/scw/config.yaml": v2CompleteValidConfigWithActiveProfileFile,
			},
			expectedAccessKey:             s(v2ValidAccessKey2),
			expectedSecretKey:             s(v2ValidSecretKey2),
			expectedAPIURL:                s(v2ValidAPIURL2),
			expectedInsecure:              b(true),
			expectedDefaultOrganizationID: s(v2ValidDefaultOrganizationID2),
			expectedDefaultRegion:         s(v2ValidDefaultRegion2),
			expectedDefaultZone:           s(v2ValidDefaultZone2),
		},
		{
			name: "Mixed config with active profile",
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".config/scw/config.yaml": v2MixedValidConfigWithActiveProfileFile,
			},
			expectedAccessKey:             s(v2ValidAccessKey2),
			expectedSecretKey:             s(v2ValidSecretKey2),
			expectedAPIURL:                s(v2ValidAPIURL),
			expectedInsecure:              b(false),
			expectedDefaultOrganizationID: s(v2ValidDefaultOrganizationID),
			expectedDefaultRegion:         s(v2ValidDefaultRegion),
			expectedDefaultZone:           s(v2ValidDefaultZone),
		},
		{
			name: "Complete config with active profile env variable",
			env: map[string]string{
				"HOME":              "{HOME}",
				scwActiveProfileEnv: v2ValidProfile,
			},
			files: map[string]string{
				".config/scw/config.yaml": v2CompleteValidConfigFile,
			},
			expectedAccessKey:             s(v2ValidAccessKey2),
			expectedSecretKey:             s(v2ValidSecretKey2),
			expectedAPIURL:                s(v2ValidAPIURL2),
			expectedInsecure:              b(true),
			expectedDefaultOrganizationID: s(v2ValidDefaultOrganizationID2),
			expectedDefaultRegion:         s(v2ValidDefaultRegion2),
			expectedDefaultZone:           s(v2ValidDefaultZone2),
		},

		// Valid V1 is not allowed
		{
			name: "Err: custom-path config with valid V1",
			env: map[string]string{
				scwConfigPathEnv: "{HOME}/valid2/test.conf",
			},
			files: map[string]string{
				"valid2/test.conf": v1ValidConfigFile,
			},
			expectedError: "scaleway-sdk-go: found legacy config in {HOME}/valid2/test.conf: legacy config is not allowed, please switch to the new config file format: " + documentationLink,
		},
		{
			name: "Err: default config with invalid V1",
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".scwrc": v1InvalidConfigFile,
			},
			expectedError: "scaleway-sdk-go: content of config file {HOME}/.scwrc is invalid json: invalid character ':' after top-level value",
		},
	}

	// create home dir
	dir := initEnv(t)

	// delete home dir and reset env variables
	defer resetEnv(t, os.Environ(), dir)
	logger.EnableDebugMode()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// set up env and config file(s)
			setEnv(t, test.env, test.files, dir)
			test.expectedError = strings.Replace(test.expectedError, "{HOME}", dir, -1)

			// remove config file(s)
			defer cleanEnv(t, test.files, dir)

			// load config
			isMigrated, err := MigrateLegacyConfig()

			if test.expectedError != "" && err != nil {
				testhelpers.Equals(t, test.expectedError, err.Error())
				return
			} else {
				testhelpers.AssertNoError(t, err)
			}
			testhelpers.Equals(t, test.isMigrated, isMigrated)
			config, err := LoadConfig()
			if test.expectedError == "" {
				testhelpers.AssertNoError(t, err)
				p, err := config.GetActiveProfile()
				testhelpers.AssertNoError(t, err)

				// assert getters
				testhelpers.Equals(t, test.expectedAccessKey, p.AccessKey)
				testhelpers.Equals(t, test.expectedSecretKey, p.SecretKey)
				testhelpers.Equals(t, test.expectedAPIURL, p.APIURL)
				testhelpers.Equals(t, test.expectedDefaultOrganizationID, p.DefaultOrganizationID)
				testhelpers.Equals(t, test.expectedDefaultRegion, p.DefaultRegion)
				testhelpers.Equals(t, test.expectedDefaultZone, p.DefaultZone)
				testhelpers.Equals(t, test.expectedInsecure, p.Insecure)
			} else {
				testhelpers.Equals(t, test.expectedError, err.Error())
			}

		})
	}
}

func TestConfigString(t *testing.T) {
	var c = &Config{
		Profile: Profile{
			AccessKey: s(v2ValidAccessKey),
			SecretKey: s(v2ValidSecretKey),
		},
		ActiveProfile: s(v2ValidProfile),
		Profiles: map[string]*Profile{
			v2ValidProfile: {
				AccessKey: s(v2ValidAccessKey2),
				SecretKey: s(v2ValidSecretKey2),
			},
		},
	}

	testhelpers.Equals(t, `access_key: SCW1234567890ABCDEFG
secret_key: 7363616c-xxxx-xxxx-xxxx-xxxxxxxxxxxx
active_profile: flantier
profiles:
  flantier:
    access_key: SCW234567890ABCDEFGH
    secret_key: 6f6e6574-xxxx-xxxx-xxxx-xxxxxxxxxxxx
`, c.String())
	testhelpers.Equals(t, v2ValidSecretKey, *c.SecretKey)

	p, err := c.GetActiveProfile()
	testhelpers.AssertNoError(t, err)
	testhelpers.Equals(t, `access_key: SCW234567890ABCDEFGH
secret_key: 6f6e6574-xxxx-xxxx-xxxx-xxxxxxxxxxxx
`, p.String())
	testhelpers.Equals(t, v2ValidSecretKey2, *p.SecretKey)

}

func TestMergeProfiles(t *testing.T) {
	p1 := &Profile{
		AccessKey: StringPtr("1"),
		SecretKey: StringPtr("1"),
	}
	p2 := &Profile{
		AccessKey: StringPtr("2"),
		Insecure:  BoolPtr(true),
	}
	p3 := &Profile{
		Insecure:    BoolPtr(false),
		DefaultZone: StringPtr(string(ZoneFrPar1)),
	}

	act := MergeProfiles(p1, p2, p3)
	exp := &Profile{
		AccessKey:   StringPtr("2"),
		SecretKey:   StringPtr("1"),
		Insecure:    BoolPtr(false),
		DefaultZone: StringPtr(string(ZoneFrPar1)),
	}

	testhelpers.Equals(t, exp, act)
}

func initEnv(t *testing.T) string {
	dir, err := ioutil.TempDir("", "home")
	if err != nil {
		t.Fatal(err)
	}
	return dir
}

func cleanEnv(t *testing.T, files map[string]string, homeDir string) {
	for path := range files {
		testhelpers.AssertNoError(t, os.RemoveAll(filepath.Join(homeDir, path)))
	}
}

func setEnv(t *testing.T, env, files map[string]string, homeDir string) {
	os.Clearenv()
	for key, value := range env {
		value = strings.Replace(value, "{HOME}", homeDir, -1)
		testhelpers.AssertNoError(t, os.Setenv(key, value))
	}

	for path, content := range files {
		targetPath := filepath.Join(homeDir, path)
		testhelpers.AssertNoError(t, os.MkdirAll(filepath.Dir(targetPath), 0700))
		testhelpers.AssertNoError(t, ioutil.WriteFile(targetPath, []byte(content), defaultConfigPermission))
	}
}

// function taken from https://golang.org/src/os/env_test.go
func resetEnv(t *testing.T, origEnv []string, homeDir string) {
	testhelpers.AssertNoError(t, os.RemoveAll(homeDir))
	for _, pair := range origEnv {
		// Environment variables on Windows can begin with =
		// https://blogs.msdn.com/b/oldnewthing/archive/2010/05/06/10008132.aspx
		i := strings.Index(pair[1:], "=") + 1
		if err := os.Setenv(pair[:i], pair[i+1:]); err != nil {
			t.Errorf("Setenv(%q, %q) failed during reset: %v", pair[:i], pair[i+1:], err)
		}
	}
}

func s(value string) *string {
	return &value
}

func r(value Region) *Region {
	return &value
}

func z(value Zone) *Zone {
	return &value
}

func b(value bool) *bool {
	return &value
}

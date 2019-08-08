package scw

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

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

// TestLoadConfig tests config getters return correct values
func TestLoadConfig(t *testing.T) {

	tests := []struct {
		name  string
		env   map[string]string
		files map[string]string

		expectedAccessKey        *string
		expectedSecretKey        *string
		expectedAPIURL           *string
		expectedInsecure         *bool
		expectedDefaultProjectID *string
		expectedDefaultRegion    *Region
		expectedDefaultZone      *Zone
	}{
		// no env variables
		{
			name: "No config without home dir",
		},
		{
			name: "No config",
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
			name: "Custom-path config with valid V1",
			env: map[string]string{
				scwConfigPathEnv: "{HOME}/valid2/test.conf",
			},
			files: map[string]string{
				"valid2/test.conf": v1ValidConfigFile,
			},
			expectedSecretKey:        s(v1ValidToken),
			expectedDefaultProjectID: s(v1ValidProjectID),
		},
		{
			name: "Custom-path config with valid V2",
			env: map[string]string{
				scwConfigPathEnv: "{HOME}/valid3/test.conf",
			},
			files: map[string]string{
				"valid3/test.conf": v2SimpleValidConfigFile,
			},
			expectedAccessKey:        s(v2ValidAccessKey),
			expectedSecretKey:        s(v2ValidSecretKey),
			expectedDefaultProjectID: s(v2ValidDefaultProjectID),
			expectedDefaultRegion:    r(Region(v2ValidDefaultRegion)),
		},
		{
			name: "Simple config with valid V2", // default config path
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".config/scw/config.yaml": v2SimpleValidConfigFile,
			},
			expectedAccessKey:        s(v2ValidAccessKey),
			expectedSecretKey:        s(v2ValidSecretKey),
			expectedDefaultProjectID: s(v2ValidDefaultProjectID),
			expectedDefaultRegion:    r(Region(v2ValidDefaultRegion)),
		},
		{
			name: "Simple config with valid V1",
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".scwrc": v1ValidConfigFile,
			},
			expectedSecretKey:        s(v1ValidToken),
			expectedDefaultProjectID: s(v1ValidProjectID),
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
			expectedAccessKey:        s(v2ValidAccessKey),
			expectedSecretKey:        s(v2ValidSecretKey),
			expectedDefaultProjectID: s(v2ValidDefaultProjectID),
			expectedDefaultRegion:    r(Region(v2ValidDefaultRegion)),
		},
		{
			name: "Complete config",
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".config/scw/config.yaml": v2CompleteValidConfigFile,
			},
			expectedAccessKey:        s(v2ValidAccessKey),
			expectedSecretKey:        s(v2ValidSecretKey),
			expectedAPIURL:           s(v2ValidAPIURL),
			expectedInsecure:         b(false),
			expectedDefaultProjectID: s(v2ValidDefaultProjectID),
			expectedDefaultRegion:    r(Region(v2ValidDefaultRegion)),
			expectedDefaultZone:      z(Zone(v2ValidDefaultZone)),
		},
		{
			name: "Complete config with active profile",
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".config/scw/config.yaml": v2CompleteValidConfigWithActiveProfileFile,
			},
			expectedAccessKey:        s(v2ValidAccessKey2),
			expectedSecretKey:        s(v2ValidSecretKey2),
			expectedAPIURL:           s(v2ValidAPIURL2),
			expectedInsecure:         b(true),
			expectedDefaultProjectID: s(v2ValidDefaultProjectID2),
			expectedDefaultRegion:    r(Region(v2ValidDefaultRegion2)),
			expectedDefaultZone:      z(Zone(v2ValidDefaultZone2)),
		},

		// up-to-date env variables
		{
			name: "No config with env variables",
			env: map[string]string{
				scwAccessKeyEnv:        v2ValidAccessKey,
				scwSecretKeyEnv:        v2ValidSecretKey,
				scwAPIURLEnv:           v2ValidAPIURL,
				scwInsecureEnv:         "false",
				scwDefaultProjectIDEnv: v2ValidDefaultProjectID,
				scwDefaultRegionEnv:    v2ValidDefaultRegion,
				scwDefaultZoneEnv:      v2ValidDefaultZone,
			},
			expectedAccessKey:        s(v2ValidAccessKey),
			expectedSecretKey:        s(v2ValidSecretKey),
			expectedAPIURL:           s(v2ValidAPIURL),
			expectedInsecure:         b(false),
			expectedDefaultProjectID: s(v2ValidDefaultProjectID),
			expectedDefaultRegion:    r(Region(v2ValidDefaultRegion)),
			expectedDefaultZone:      z(Zone(v2ValidDefaultZone)),
		},
		{
			name: "Complete config file with env variables",
			env: map[string]string{
				"HOME":                 "{HOME}",
				scwAccessKeyEnv:        v2ValidAccessKey2,
				scwSecretKeyEnv:        v2ValidSecretKey2,
				scwAPIURLEnv:           v2ValidAPIURL2,
				scwInsecureEnv:         v2ValidInsecure2,
				scwDefaultProjectIDEnv: v2ValidDefaultProjectID2,
				scwDefaultRegionEnv:    v2ValidDefaultRegion2,
				scwDefaultZoneEnv:      v2ValidDefaultZone2,
			},
			files: map[string]string{
				".config/scw/config.yaml": v2CompleteValidConfigFile,
			},
			expectedAccessKey:        s(v2ValidAccessKey2),
			expectedSecretKey:        s(v2ValidSecretKey2),
			expectedAPIURL:           s(v2ValidAPIURL2),
			expectedInsecure:         b(true),
			expectedDefaultProjectID: s(v2ValidDefaultProjectID2),
			expectedDefaultRegion:    r(Region(v2ValidDefaultRegion2)),
			expectedDefaultZone:      z(Zone(v2ValidDefaultZone2)),
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
			expectedAccessKey:        s(v2ValidAccessKey2),
			expectedSecretKey:        s(v2ValidSecretKey2),
			expectedAPIURL:           s(v2ValidAPIURL2),
			expectedInsecure:         b(true),
			expectedDefaultProjectID: s(v2ValidDefaultProjectID2),
			expectedDefaultRegion:    r(Region(v2ValidDefaultRegion2)),
			expectedDefaultZone:      z(Zone(v2ValidDefaultZone2)),
		},
		{
			name: "Complete config with active profile env variable and all env variables",
			env: map[string]string{
				"HOME":                 "{HOME}",
				scwActiveProfileEnv:    v2ValidProfile,
				scwAccessKeyEnv:        v2ValidAccessKey,
				scwSecretKeyEnv:        v2ValidSecretKey,
				scwAPIURLEnv:           v2ValidAPIURL,
				scwInsecureEnv:         "false",
				scwDefaultProjectIDEnv: v2ValidDefaultProjectID,
				scwDefaultRegionEnv:    v2ValidDefaultRegion,
				scwDefaultZoneEnv:      v2ValidDefaultZone,
			},
			files: map[string]string{
				".config/scw/config.yaml": v2CompleteValidConfigFile,
			},
			expectedAccessKey:        s(v2ValidAccessKey),
			expectedSecretKey:        s(v2ValidSecretKey),
			expectedAPIURL:           s(v2ValidAPIURL),
			expectedInsecure:         b(false),
			expectedDefaultProjectID: s(v2ValidDefaultProjectID),
			expectedDefaultRegion:    r(Region(v2ValidDefaultRegion)),
			expectedDefaultZone:      z(Zone(v2ValidDefaultZone)),
		},

		// legacy env variables
		{
			name: "No config with terraform legacy env variables",
			env: map[string]string{
				terraformAccessKeyEnv:    v2ValidAccessKey,
				terraformSecretKeyEnv:    v2ValidSecretKey,
				terraformOrganizationEnv: v2ValidDefaultProjectID,
				terraformRegionEnv:       v2ValidDefaultRegion,
			},
			expectedAccessKey:        s(v2ValidAccessKey),
			expectedSecretKey:        s(v2ValidSecretKey),
			expectedDefaultProjectID: s(v2ValidDefaultProjectID),
			expectedDefaultRegion:    r(Region(v2ValidDefaultRegion)),
		},
		{
			name: "No config with CLI legacy env variables",
			env: map[string]string{
				cliSecretKeyEnv:    v2ValidSecretKey2,
				cliOrganizationEnv: v2ValidDefaultProjectID2,
				cliRegionEnv:       v2ValidDefaultRegion2,
				cliTLSVerifyEnv:    "false",
			},
			expectedSecretKey:        s(v2ValidSecretKey2),
			expectedInsecure:         b(true),
			expectedDefaultProjectID: s(v2ValidDefaultProjectID2),
			expectedDefaultRegion:    r(Region(v2ValidDefaultRegion2)),
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

			// load config
			config, err := LoadConfig()
			testhelpers.AssertNoError(t, err)

			// assert getters
			accessKey, exist := config.GetAccessKey()
			if exist {
				testhelpers.Equals(t, test.expectedAccessKey, &accessKey)
			} else {
				testhelpers.Assert(t, test.expectedAccessKey == nil, "expected accessKey must be nil")
			}
			secretKey, exist := config.GetSecretKey()
			if exist {
				testhelpers.Equals(t, test.expectedSecretKey, &secretKey)
			} else {
				testhelpers.Assert(t, test.expectedSecretKey == nil, "expected secretKey must be nil")
			}
			apiURL, exist := config.GetAPIURL()
			if exist {
				testhelpers.Equals(t, test.expectedAPIURL, &apiURL)
			} else {
				testhelpers.Assert(t, test.expectedAPIURL == nil, "expected apiURL must be nil")
			}
			defaultProjectID, exist := config.GetDefaultProjectID()
			if exist {
				testhelpers.Equals(t, test.expectedDefaultProjectID, &defaultProjectID)
			} else {
				testhelpers.Assert(t, test.expectedDefaultProjectID == nil, "expected defaultProjectID must be nil")
			}
			defaultRegion, exist := config.GetDefaultRegion()
			if exist {
				testhelpers.Equals(t, test.expectedDefaultRegion, &defaultRegion)
			} else {
				testhelpers.Assert(t, test.expectedDefaultRegion == nil, "expected defaultRegion must be nil")
			}
			defaultZone, exist := config.GetDefaultZone()
			if exist {
				testhelpers.Equals(t, test.expectedDefaultZone, &defaultZone)
			} else {
				testhelpers.Assert(t, test.expectedDefaultZone == nil, "expected defaultZone must be nil")
			}
			insecure, exist := config.GetInsecure()
			if exist {
				testhelpers.Equals(t, test.expectedInsecure, &insecure)
			} else {
				testhelpers.Assert(t, test.expectedInsecure == nil, "expected insecure must be nil")
			}
		})
	}
}

const emptyFile = ""

// v2 config
var (
	v2ValidAccessKey2        = "ACCESS_KEY2"
	v2ValidSecretKey2        = "6f6e6574-6f72-756c-6c74-68656d616c6c" // hint: | xxd -ps -r
	v2ValidAPIURL2           = "api-fr-par.scaleway.com"
	v2ValidInsecure2         = "true"
	v2ValidDefaultProjectID2 = "6d6f7264-6f72-6772-6561-74616761696e" // hint: | xxd -ps -r
	v2ValidDefaultRegion2    = string(RegionFrPar)
	v2ValidDefaultZone2      = string(ZoneFrPar2)

	v2ValidAccessKey        = "ACCESS_KEY"
	v2ValidSecretKey        = "7363616c-6577-6573-6862-6f7579616161" // hint: | xxd -ps -r
	v2ValidAPIURL           = "api.scaleway.com"
	v2ValidInsecure         = "false"
	v2ValidDefaultProjectID = "6170692e-7363-616c-6577-61792e636f6d" // hint: | xxd -ps -r
	v2ValidDefaultRegion    = string(RegionNlAms)
	v2ValidDefaultZone      = string(ZoneNlAms1)
	v2ValidProfile          = "flantier"

	v2SimpleValidConfig = &configV2{
		profile: profile{
			AccessKey:        &v2ValidAccessKey,
			SecretKey:        &v2ValidSecretKey,
			DefaultProjectID: &v2ValidDefaultProjectID,
			DefaultRegion:    &v2ValidDefaultRegion,
		},
	}
	v2PartialValidConfigFile = `
access_key: ` + v2ValidAccessKey + `
secret_key: ` + v2ValidSecretKey + `
api_url: ` + v2ValidAPIURL + `
insecure: ` + v2ValidInsecure + `
default_project_id: ` + v2ValidDefaultProjectID + `
default_region: ` + v2ValidDefaultRegion + `
default_zone: ` + v2ValidDefaultZone

	v2CompleteValidConfigFile = v2PartialValidConfigFile + `
profiles:
  ` + v2ValidProfile + `:
    access_key: ` + v2ValidAccessKey2 + `
    secret_key: ` + v2ValidSecretKey2 + `
    api_url: ` + v2ValidAPIURL2 + `
    insecure: ` + v2ValidInsecure2 + `
    default_project_id: ` + v2ValidDefaultProjectID2 + `
    default_region: ` + v2ValidDefaultRegion2 + `
    default_zone: ` + v2ValidDefaultZone2 + `
`

	v2CompleteValidConfigWithActiveProfileFile = `
access_key: ` + v2ValidAccessKey + `
secret_key: ` + v2ValidSecretKey + `
api_url: ` + v2ValidAPIURL + `
insecure: ` + v2ValidInsecure + `
default_project_id: ` + v2ValidDefaultProjectID + `
default_region: ` + v2ValidDefaultRegion + `
default_zone: ` + v2ValidDefaultZone + `
active_profile: ` + v2ValidProfile + `
profiles:
  ` + v2ValidProfile + `:
    access_key: ` + v2ValidAccessKey2 + `
    secret_key: ` + v2ValidSecretKey2 + `
    api_url: ` + v2ValidAPIURL2 + `
    insecure: ` + v2ValidInsecure2 + `
    default_project_id: ` + v2ValidDefaultProjectID2 + `
    default_region: ` + v2ValidDefaultRegion2 + `
    default_zone: ` + v2ValidDefaultZone2 + `
`

	v2SimpleValidConfigFile = `
access_key: ` + v2ValidAccessKey + `
secret_key: ` + v2ValidSecretKey + `
default_project_id: ` + v2ValidDefaultProjectID + `
default_region: ` + v2ValidDefaultRegion + `
`

	v2SimpleInvalidConfigFile            = `insecure: "bool""`
	v2SimpleConfigFileWithInvalidProfile = `active_profile: flantier`

	v2FromV1ConfigFile = `secret_key: ` + v1ValidToken + `
default_project_id: ` + v1ValidProjectID + `
`
)

// v1 config
var (
	v1ValidProjectID = "29aa5db6-1d6d-404e-890d-f896913f9ec1"
	v1ValidToken     = "a057b0c1-eb47-4bf8-a589-72c1f2029515"
	v1Version        = "1.19"

	v1ValidConfig = &configV2{
		profile: profile{
			SecretKey:        &v1ValidToken,
			DefaultProjectID: &v1ValidProjectID,
		},
	}

	v1ValidConfigFile = `{
"organization":"` + v1ValidProjectID + `",
"token":"` + v1ValidToken + `",
"version":"` + v1Version + `"
}`

	v1InvalidConfigFile = `
"organization":"` + v1ValidProjectID + `",
"token":"` + v1ValidToken + `",
"version":"` + v1Version + `"
`
)

// TestSaveConfig tests config write the correct values in the config file
func TestSaveConfig(t *testing.T) {
	tests := []struct {
		name        string
		env         map[string]string
		files       map[string]string
		newConfig   *configV2
		profileName []string

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
			newConfig: &configV2{
				profile: profile{
					AccessKey:        s(v2ValidAccessKey),
					SecretKey:        s(v2ValidSecretKey),
					DefaultProjectID: s(v2ValidDefaultProjectID),
					DefaultRegion:    s(v2ValidDefaultRegion),
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
			newConfig: &configV2{
				profile: profile{
					AccessKey:        s(v2ValidAccessKey),
					SecretKey:        s(v2ValidSecretKey),
					DefaultProjectID: s(v2ValidDefaultProjectID),
					DefaultRegion:    s(v2ValidDefaultRegion),
				},
			},

			expectedFiles: map[string]string{
				".config/scw/config.yaml": v2SimpleValidConfigFile,
			},
		},
		{
			name: "Add region only",
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".config/scw/config.yaml": v2SimpleValidConfigFile,
			},
			newConfig: &configV2{
				profile: profile{
					DefaultZone: s(v2ValidDefaultZone),
				},
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
			profileName: []string{v2ValidProfile},
			newConfig: &configV2{
				Profiles: map[string]*profile{v2ValidProfile: {
					AccessKey:        s(v2ValidAccessKey2),
					SecretKey:        s(v2ValidSecretKey2),
					APIURL:           s(v2ValidAPIURL2),
					Insecure:         b(true),
					DefaultProjectID: s(v2ValidDefaultProjectID2),
					DefaultRegion:    s(v2ValidDefaultRegion2),
					DefaultZone:      s(v2ValidDefaultZone2),
				}},
				withProfile: v2ValidProfile,
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

			testhelpers.AssertNoError(t, SaveConfig(test.newConfig, test.profileName...))

			// test expected files
			for fileName, expectedContent := range test.expectedFiles {
				content, err := ioutil.ReadFile(filepath.Join(dir, fileName))
				testhelpers.AssertNoError(t, err)
				testhelpers.Equals(t, expectedContent, "\n"+string(content))
			}

		})
	}
}

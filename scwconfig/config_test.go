package scwconfig

import (
	"os"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

func TestConfig(t *testing.T) {
	tests := []struct {
		name  string
		env   map[string]string
		files map[string]string

		expectedAccessKey             string
		expectedSecretKey             string
		expectedApiUrl                string
		expectedInsecure              bool
		expectedDefaultOrganizationId string
		expectedDefaultRegion         string
		expectedDefaultZone           string
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
			expectedSecretKey:             v1ValidToken,
			expectedDefaultOrganizationId: v1ValidOrganizationID,
		},
		{
			name: "Custom-path config with valid V2",
			env: map[string]string{
				scwConfigPathEnv: "{HOME}/valid3/test.conf",
			},
			files: map[string]string{
				"valid3/test.conf": v2SimpleValidConfigFile,
			},
			expectedAccessKey:             v2ValidAccessKey,
			expectedSecretKey:             v2ValidSecretKey,
			expectedDefaultOrganizationId: v2ValidDefaultOrganizationID,
			expectedDefaultRegion:         v2ValidDefaultRegion,
		},
		{
			name: "Simple config with valid V2", // default config path
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".config/scw/config.yaml": v2SimpleValidConfigFile,
			},
			expectedAccessKey:             v2ValidAccessKey,
			expectedSecretKey:             v2ValidSecretKey,
			expectedDefaultOrganizationId: v2ValidDefaultOrganizationID,
			expectedDefaultRegion:         v2ValidDefaultRegion,
		},
		{
			name: "Simple config with valid V1",
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".scwrc": v1ValidConfigFile,
			},
			expectedSecretKey:             v1ValidToken,
			expectedDefaultOrganizationId: v1ValidOrganizationID,
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
			expectedAccessKey:             v2ValidAccessKey,
			expectedSecretKey:             v2ValidSecretKey,
			expectedDefaultOrganizationId: v2ValidDefaultOrganizationID,
			expectedDefaultRegion:         v2ValidDefaultRegion,
		},
		{
			name: "Complete config",
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".config/scw/config.yaml": v2CompleteValidConfigFile,
			},
			expectedAccessKey:             v2ValidAccessKey,
			expectedSecretKey:             v2ValidSecretKey,
			expectedApiUrl:                v2ValidApiUrl,
			expectedInsecure:              false,
			expectedDefaultOrganizationId: v2ValidDefaultOrganizationID,
			expectedDefaultRegion:         v2ValidDefaultRegion,
			expectedDefaultZone:           v2ValidDefaultZone,
		},
		{
			name: "Complete config with active profile",
			env: map[string]string{
				"HOME": "{HOME}",
			},
			files: map[string]string{
				".config/scw/config.yaml": v2CompleteValidConfigWithActiveProfileFile,
			},
			expectedAccessKey:             v2ValidAccessKey2,
			expectedSecretKey:             v2ValidSecretKey2,
			expectedApiUrl:                v2ValidApiUrl2,
			expectedInsecure:              true,
			expectedDefaultOrganizationId: v2ValidDefaultOrganizationID2,
			expectedDefaultRegion:         v2ValidDefaultRegion2,
			expectedDefaultZone:           v2ValidDefaultZone2,
		},

		// up-to-date env variables
		{
			name: "No config with env variables",
			env: map[string]string{
				scwAccessKeyEnv:             v2ValidAccessKey,
				scwSecretKeyEnv:             v2ValidSecretKey,
				scwApiUrlEnv:                v2ValidApiUrl,
				scwInsecureEnv:              "false",
				scwDefaultOrganizationIDEnv: v2ValidDefaultOrganizationID,
				scwDefaultRegionEnv:         v2ValidDefaultRegion,
				scwDefaultZoneEnv:           v2ValidDefaultZone,
			},
			expectedAccessKey:             v2ValidAccessKey,
			expectedSecretKey:             v2ValidSecretKey,
			expectedApiUrl:                v2ValidApiUrl,
			expectedInsecure:              false,
			expectedDefaultOrganizationId: v2ValidDefaultOrganizationID,
			expectedDefaultRegion:         v2ValidDefaultRegion,
			expectedDefaultZone:           v2ValidDefaultZone,
		},
		{
			name: "Complete config file with env variables",
			env: map[string]string{
				"HOME":                      "{HOME}",
				scwAccessKeyEnv:             v2ValidAccessKey2,
				scwSecretKeyEnv:             v2ValidSecretKey2,
				scwApiUrlEnv:                v2ValidApiUrl2,
				scwInsecureEnv:              v2ValidInsecure2,
				scwDefaultOrganizationIDEnv: v2ValidDefaultOrganizationID2,
				scwDefaultRegionEnv:         v2ValidDefaultRegion2,
				scwDefaultZoneEnv:           v2ValidDefaultZone2,
			},
			files: map[string]string{
				".config/scw/config.yaml": v2CompleteValidConfigFile,
			},
			expectedAccessKey:             v2ValidAccessKey2,
			expectedSecretKey:             v2ValidSecretKey2,
			expectedApiUrl:                v2ValidApiUrl2,
			expectedInsecure:              true,
			expectedDefaultOrganizationId: v2ValidDefaultOrganizationID2,
			expectedDefaultRegion:         v2ValidDefaultRegion2,
			expectedDefaultZone:           v2ValidDefaultZone2,
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
			expectedAccessKey:             v2ValidAccessKey2,
			expectedSecretKey:             v2ValidSecretKey2,
			expectedApiUrl:                v2ValidApiUrl2,
			expectedInsecure:              true,
			expectedDefaultOrganizationId: v2ValidDefaultOrganizationID2,
			expectedDefaultRegion:         v2ValidDefaultRegion2,
			expectedDefaultZone:           v2ValidDefaultZone2,
		},

		// todo: legacy env variables
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
			config, err := Load()
			testhelpers.Ok(t, err)

			// assert getters
			testhelpers.Equals(t, test.expectedAccessKey, config.GetAccessKey())
			testhelpers.Equals(t, test.expectedSecretKey, config.GetSecretKey())
			testhelpers.Equals(t, test.expectedApiUrl, config.GetApiUrl())
			testhelpers.Equals(t, test.expectedDefaultOrganizationId, config.GetDefaultOrganizationId())
			testhelpers.Equals(t, test.expectedDefaultRegion, config.GetDefaultRegion())
			testhelpers.Equals(t, test.expectedDefaultZone, config.GetDefaultZone())
			testhelpers.Equals(t, test.expectedInsecure, config.GetInsecure())
		})
	}
}

const emptyFile = ""

// v2 config
var (
	v2ValidAccessKey2             = "ACCESS_KEY2"
	v2ValidSecretKey2             = "23cd9bbf-519f-4fc8-912e-83bbd18e4247"
	v2ValidApiUrl2                = "api-fr-par.scaleway.com"
	v2ValidInsecure2              = "true"
	v2ValidDefaultOrganizationID2 = "76ca3ea8-0a29-4e39-b4ae-f7a770868072"
	v2ValidDefaultRegion2         = "fr-par"
	v2ValidDefaultZone2           = "fr-par-2"

	v2ValidAccessKey             = "ACCESS_KEY"
	v2ValidSecretKey             = "539a6564-bf92-4dc9-a0d4-50e3ca827ecb"
	v2ValidApiUrl                = "api.scaleway.com"
	v2ValidInsecure              = "false"
	v2ValidDefaultOrganizationID = "d45a075f-18a1-4e9f-824e-43914a3ae8bd"
	v2ValidDefaultRegion         = "nl-ams"
	v2ValidDefaultZone           = "nl-ams-1"
	v2ValidProfile               = "flantier"

	v2SimpleValidConfig = &configV2{
		profile: profile{
			AccessKey:             &v2ValidAccessKey,
			SecretKey:             &v2ValidSecretKey,
			DefaultOrganizationId: &v2ValidDefaultOrganizationID,
			DefaultRegion:         &v2ValidDefaultRegion,
		},
	}

	v2CompleteValidConfigFile = `
access_key: "` + v2ValidAccessKey + `"
secret_key: "` + v2ValidSecretKey + `"
api_url: "` + v2ValidApiUrl + `"
insecure: ` + v2ValidInsecure + `
default_organization_id: "` + v2ValidDefaultOrganizationID + `"
default_region: "` + v2ValidDefaultRegion + `"
default_zone: "` + v2ValidDefaultZone + `"
profiles:
  ` + v2ValidProfile + `:
    access_key: "` + v2ValidAccessKey2 + `"
    secret_key: "` + v2ValidSecretKey2 + `"
    api_url: "` + v2ValidApiUrl2 + `"
    insecure: ` + v2ValidInsecure2 + `
    default_organization_id: "` + v2ValidDefaultOrganizationID2 + `"
    default_region: "` + v2ValidDefaultRegion2 + `"
    default_zone: "` + v2ValidDefaultZone2 + `"
`

	v2CompleteValidConfigWithActiveProfileFile = `
access_key: "` + v2ValidAccessKey + `"
secret_key: "` + v2ValidSecretKey + `"
api_url: "` + v2ValidApiUrl + `"
insecure: ` + v2ValidInsecure + `
default_organization_id: "` + v2ValidDefaultOrganizationID + `"
default_region: "` + v2ValidDefaultRegion + `"
default_zone: "` + v2ValidDefaultZone + `"
active_profile: ` + v2ValidProfile + `
profiles:
  ` + v2ValidProfile + `:
    access_key: "` + v2ValidAccessKey2 + `"
    secret_key: "` + v2ValidSecretKey2 + `"
    api_url: "` + v2ValidApiUrl2 + `"
    insecure: ` + v2ValidInsecure2 + `
    default_organization_id: "` + v2ValidDefaultOrganizationID2 + `"
    default_region: "` + v2ValidDefaultRegion2 + `"
    default_zone: "` + v2ValidDefaultZone2 + `"
`

	v2SimpleValidConfigFile = `
access_key: "` + v2ValidAccessKey + `"
secret_key: "` + v2ValidSecretKey + `"
default_organization_id: "` + v2ValidDefaultOrganizationID + `"
default_region: "` + v2ValidDefaultRegion + `"
`

	v2SimpleInvalidConfigFile            = `insecure: "bool""`
	v2SimpleConfigFileWithInvalidProfile = `active_profile: flantier`

	v2FromV1ConfigFile = `secret_key: ` + v1ValidToken + `
default_organization_id: ` + v1ValidOrganizationID + `
`
)

// v1 config
var (
	v1ValidOrganizationID = "29aa5db6-1d6d-404e-890d-f896913f9ec1"
	v1ValidToken          = "a057b0c1-eb47-4bf8-a589-72c1f2029515"
	v1Version             = "1.19"

	v1ValidConfig = &configV2{
		profile: profile{
			SecretKey:             &v1ValidToken,
			DefaultOrganizationId: &v1ValidOrganizationID,
		},
	}

	v1ValidConfigFile = `{
"organization":"` + v1ValidOrganizationID + `",
"token":"` + v1ValidToken + `",
"version":"` + v1Version + `"
}`

	v1InvalidConfigFile = `
"organization":"` + v1ValidOrganizationID + `",
"token":"` + v1ValidToken + `",
"version":"` + v1Version + `"
`
)

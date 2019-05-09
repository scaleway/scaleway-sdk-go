package scwconfig

import (
	"os"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

func s(value string) *string {
	return &value
}

func b(value bool) *bool {
	return &value
}

// TestConfig tests config getters return correct values
func TestConfig(t *testing.T) {

	tests := []struct {
		name  string
		env   map[string]string
		files map[string]string

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
			expectedSecretKey:             s(v1ValidToken),
			expectedDefaultOrganizationID: s(v1ValidOrganizationID),
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

		// up-to-date env variables
		{
			name: "No config with env variables",
			env: map[string]string{
				scwAccessKeyEnv:             v2ValidAccessKey,
				scwSecretKeyEnv:             v2ValidSecretKey,
				scwAPIURLEnv:                v2ValidAPIURL,
				scwInsecureEnv:              "false",
				scwDefaultOrganizationIDEnv: v2ValidDefaultOrganizationID,
				scwDefaultRegionEnv:         v2ValidDefaultRegion,
				scwDefaultZoneEnv:           v2ValidDefaultZone,
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
			name: "Complete config file with env variables",
			env: map[string]string{
				"HOME":                      "{HOME}",
				scwAccessKeyEnv:             v2ValidAccessKey2,
				scwSecretKeyEnv:             v2ValidSecretKey2,
				scwAPIURLEnv:                v2ValidAPIURL2,
				scwInsecureEnv:              v2ValidInsecure2,
				scwDefaultOrganizationIDEnv: v2ValidDefaultOrganizationID2,
				scwDefaultRegionEnv:         v2ValidDefaultRegion2,
				scwDefaultZoneEnv:           v2ValidDefaultZone2,
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
		{
			name: "Complete config with active profile env variable and all env variables",
			env: map[string]string{
				"HOME":                      "{HOME}",
				scwActiveProfileEnv:         v2ValidProfile,
				scwAccessKeyEnv:             v2ValidAccessKey,
				scwSecretKeyEnv:             v2ValidSecretKey,
				scwAPIURLEnv:                v2ValidAPIURL,
				scwInsecureEnv:              "false",
				scwDefaultOrganizationIDEnv: v2ValidDefaultOrganizationID,
				scwDefaultRegionEnv:         v2ValidDefaultRegion,
				scwDefaultZoneEnv:           v2ValidDefaultZone,
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

		// legacy env variables
		{
			name: "No config with terraform legacy env variables",
			env: map[string]string{
				terraformAccessKeyEnv:    v2ValidAccessKey,
				terraformSecretKeyEnv:    v2ValidSecretKey,
				terraformOrganizationEnv: v2ValidDefaultOrganizationID,
				terraformRegionEnv:       v2ValidDefaultRegion,
			},
			expectedAccessKey:             s(v2ValidAccessKey),
			expectedSecretKey:             s(v2ValidSecretKey),
			expectedDefaultOrganizationID: s(v2ValidDefaultOrganizationID),
			expectedDefaultRegion:         s(v2ValidDefaultRegion),
		},
		{
			name: "No config with CLI legacy env variables",
			env: map[string]string{
				cliSecretKeyEnv:    v2ValidSecretKey2,
				cliOrganizationEnv: v2ValidDefaultOrganizationID2,
				cliRegionEnv:       v2ValidDefaultRegion2,
				cliTLSVerifyEnv:    v2ValidInsecure2,
			},
			expectedSecretKey:             s(v2ValidSecretKey2),
			expectedInsecure:              b(true),
			expectedDefaultOrganizationID: s(v2ValidDefaultOrganizationID2),
			expectedDefaultRegion:         s(v2ValidDefaultRegion2),
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
			config, err := Load()
			testhelpers.Ok(t, err)

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
			defaultOrganizationID, exist := config.GetDefaultOrganizationID()
			if exist {
				testhelpers.Equals(t, test.expectedDefaultOrganizationID, &defaultOrganizationID)
			} else {
				testhelpers.Assert(t, test.expectedDefaultOrganizationID == nil, "expected defaultOrganizationID must be nil")
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
	v2ValidAccessKey2             = "ACCESS_KEY2"
	v2ValidSecretKey2             = "23cd9bbf-519f-4fc8-912e-83bbd18e4247"
	v2ValidAPIURL2                = "api-fr-par.scaleway.com"
	v2ValidInsecure2              = "true"
	v2ValidDefaultOrganizationID2 = "76ca3ea8-0a29-4e39-b4ae-f7a770868072"
	v2ValidDefaultRegion2         = "fr-par"
	v2ValidDefaultZone2           = "fr-par-2"

	v2ValidAccessKey             = "ACCESS_KEY"
	v2ValidSecretKey             = "539a6564-bf92-4dc9-a0d4-50e3ca827ecb"
	v2ValidAPIURL                = "api.scaleway.com"
	v2ValidInsecure              = "false"
	v2ValidDefaultOrganizationID = "d45a075f-18a1-4e9f-824e-43914a3ae8bd"
	v2ValidDefaultRegion         = "nl-ams"
	v2ValidDefaultZone           = "nl-ams-1"
	v2ValidProfile               = "flantier"

	v2SimpleValidConfig = &configV2{
		profile: profile{
			AccessKey:             &v2ValidAccessKey,
			SecretKey:             &v2ValidSecretKey,
			DefaultOrganizationID: &v2ValidDefaultOrganizationID,
			DefaultRegion:         &v2ValidDefaultRegion,
		},
	}

	v2CompleteValidConfigFile = `
access_key: "` + v2ValidAccessKey + `"
secret_key: "` + v2ValidSecretKey + `"
api_url: "` + v2ValidAPIURL + `"
insecure: ` + v2ValidInsecure + `
default_organization_id: "` + v2ValidDefaultOrganizationID + `"
default_region: "` + v2ValidDefaultRegion + `"
default_zone: "` + v2ValidDefaultZone + `"
profiles:
  ` + v2ValidProfile + `:
    access_key: "` + v2ValidAccessKey2 + `"
    secret_key: "` + v2ValidSecretKey2 + `"
    api_url: "` + v2ValidAPIURL2 + `"
    insecure: ` + v2ValidInsecure2 + `
    default_organization_id: "` + v2ValidDefaultOrganizationID2 + `"
    default_region: "` + v2ValidDefaultRegion2 + `"
    default_zone: "` + v2ValidDefaultZone2 + `"
`

	v2CompleteValidConfigWithActiveProfileFile = `
access_key: "` + v2ValidAccessKey + `"
secret_key: "` + v2ValidSecretKey + `"
api_url: "` + v2ValidAPIURL + `"
insecure: ` + v2ValidInsecure + `
default_organization_id: "` + v2ValidDefaultOrganizationID + `"
default_region: "` + v2ValidDefaultRegion + `"
default_zone: "` + v2ValidDefaultZone + `"
active_profile: ` + v2ValidProfile + `
profiles:
  ` + v2ValidProfile + `:
    access_key: "` + v2ValidAccessKey2 + `"
    secret_key: "` + v2ValidSecretKey2 + `"
    api_url: "` + v2ValidAPIURL2 + `"
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
			DefaultOrganizationID: &v1ValidOrganizationID,
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

package scw

import (
	"os"
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

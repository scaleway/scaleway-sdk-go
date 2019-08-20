package scw

import (
	"os"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/logger"
)

// TestLoadConfig tests config getters return correct values
func TestLoadEnvProfile(t *testing.T) {

	tests := []struct {
		name string
		env  map[string]string

		expectedAccessKey        *string
		expectedSecretKey        *string
		expectedAPIURL           *string
		expectedInsecure         *bool
		expectedDefaultProjectID *string
		expectedDefaultRegion    *string
		expectedDefaultZone      *string
	}{
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
			expectedDefaultRegion:    s(v2ValidDefaultRegion),
			expectedDefaultZone:      s(v2ValidDefaultZone),
		},
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
			expectedDefaultRegion:    s(v2ValidDefaultRegion),
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
			expectedDefaultRegion:    s(v2ValidDefaultRegion2),
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
			setEnv(t, test.env, nil, dir)

			// remove config file(s)
			defer cleanEnv(t, nil, dir)

			// load config
			p, err := LoadEnvProfile()
			testhelpers.AssertNoError(t, err)

			// assert getters
			testhelpers.Equals(t, test.expectedAccessKey, p.AccessKey)
			testhelpers.Equals(t, test.expectedSecretKey, p.SecretKey)
			testhelpers.Equals(t, test.expectedAPIURL, p.APIURL)
			testhelpers.Equals(t, test.expectedDefaultProjectID, p.DefaultProjectID)
			testhelpers.Equals(t, test.expectedDefaultRegion, p.DefaultRegion)
			testhelpers.Equals(t, test.expectedDefaultZone, p.DefaultZone)
			testhelpers.Equals(t, test.expectedInsecure, p.Insecure)

		})
	}
}

/* removed tests

//{
//	name: "Complete config file with env variables",
//	env: map[string]string{
//		"HOME":                 "{HOME}",
//		scwAccessKeyEnv:        v2ValidAccessKey2,
//		scwSecretKeyEnv:        v2ValidSecretKey2,
//		scwAPIURLEnv:           v2ValidAPIURL2,
//		scwInsecureEnv:         v2ValidInsecure2,
//		scwDefaultProjectIDEnv: v2ValidDefaultProjectID2,
//		scwDefaultRegionEnv:    v2ValidDefaultRegion2,
//		scwDefaultZoneEnv:      v2ValidDefaultZone2,
//	},
//	files: map[string]string{
//		".config/scw/config.yaml": v2CompleteValidConfigFile,
//	},
//	expectedAccessKey:        s(v2ValidAccessKey2),
//	expectedSecretKey:        s(v2ValidSecretKey2),
//	expectedAPIURL:           s(v2ValidAPIURL2),
//	expectedInsecure:         b(true),
//	expectedDefaultProjectID: s(v2ValidDefaultProjectID2),
//	expectedDefaultRegion:    s(v2ValidDefaultRegion2),
//	expectedDefaultZone:      s(v2ValidDefaultZone2),
//},
//{
//	name: "Complete config with active profile env variable and all env variables",
//	env: map[string]string{
//		"HOME":                 "{HOME}",
//		scwActiveProfileEnv:    v2ValidProfile,
//		scwAccessKeyEnv:        v2ValidAccessKey,
//		scwSecretKeyEnv:        v2ValidSecretKey,
//		scwAPIURLEnv:           v2ValidAPIURL,
//		scwInsecureEnv:         "false",
//		scwDefaultProjectIDEnv: v2ValidDefaultProjectID,
//		scwDefaultRegionEnv:    v2ValidDefaultRegion,
//		scwDefaultZoneEnv:      v2ValidDefaultZone,
//	},
//	files: map[string]string{
//		".config/scw/config.yaml": v2CompleteValidConfigFile,
//	},
//	expectedAccessKey:        s(v2ValidAccessKey),
//	expectedSecretKey:        s(v2ValidSecretKey),
//	expectedAPIURL:           s(v2ValidAPIURL),
//	expectedInsecure:         b(false),
//	expectedDefaultProjectID: s(v2ValidDefaultProjectID),
//	expectedDefaultRegion:    s(v2ValidDefaultRegion),
//	expectedDefaultZone:      s(v2ValidDefaultZone),
//},



*/

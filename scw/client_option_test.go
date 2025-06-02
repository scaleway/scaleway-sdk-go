package scw_test

import (
	"os"
	"strings"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/auth"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

var (
	defaultOrganizationID = "6170692e-7363-616c-6577-61792e636f6d" // hint: | xxd -ps -r
	defaultProjectID      = "6170692e-7363-616c-6577-61792e636f6d" // hint: | xxd -ps -r
	defaultRegion         = scw.RegionNlAms
	defaultZone           = scw.ZoneNlAms1
)

func TestClientOptions(t *testing.T) {
	testCases := []struct {
		name         string
		clientOption scw.ClientOption
		errStr       string
	}{
		{
			name: "Create a valid client option",
			clientOption: func(s *settings) {
				s.token = auth.NewToken(v2ValidAccessKey, v2ValidSecretKey)
				s.apiURL = v2ValidAPIURL
				s.defaultOrganizationID = &defaultOrganizationID
				s.defaultProjectID = &defaultProjectID
				s.defaultRegion = &defaultRegion
				s.defaultZone = &defaultZone
			},
		},
		{
			name: "Should throw an empty access key error",
			clientOption: func(s *settings) {
				s.token = auth.NewToken("", v2ValidSecretKey)
			},
			errStr: "scaleway-sdk-go: access key cannot be empty",
		},
		{
			name: "Should throw a bad access key error",
			clientOption: func(s *settings) {
				s.token = auth.NewToken(v2InvalidAccessKey, v2ValidSecretKey)
			},
			errStr: "scaleway-sdk-go: invalid access key format 'invalid', expected SCWXXXXXXXXXXXXXXXXX format",
		},
		{
			name: "Should throw an empty secret key error",
			clientOption: func(s *settings) {
				s.token = auth.NewToken(v2ValidAccessKey, "")
			},
			errStr: "scaleway-sdk-go: secret key cannot be empty",
		},
		{
			name: "Should throw a bad secret key error",
			clientOption: func(s *settings) {
				s.token = auth.NewToken(v2ValidAccessKey, v2InvalidSecretKey)
			},
			errStr: "scaleway-sdk-go: invalid secret key format 'invalid', expected a UUID: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		},
		{
			name: "Should throw an url error",
			clientOption: func(s *settings) {
				s.apiURL = ":test"
				s.token = auth.NewToken(v2ValidAccessKey, v2ValidSecretKey)
			},
			errStr: "scaleway-sdk-go: invalid API url ':test'",
		},
		{
			name: "Should throw an empty url error",
			clientOption: func(s *settings) {
				s.apiURL = ""
				s.token = auth.NewToken(v2ValidAccessKey, v2ValidSecretKey)
			},
			errStr: "scaleway-sdk-go: invalid API url ''",
		},
		{
			name: "Should throw an empty organization ID error",
			clientOption: func(s *settings) {
				s.token = auth.NewToken(v2ValidAccessKey, v2ValidSecretKey)
				s.defaultOrganizationID = StringPtr("")
			},
			errStr: "scaleway-sdk-go: default organization ID cannot be empty",
		},
		{
			name: "Should throw a bad organization ID error",
			clientOption: func(s *settings) {
				s.token = auth.NewToken(v2ValidAccessKey, v2ValidSecretKey)
				s.defaultOrganizationID = StringPtr(v2InvalidDefaultOrganizationID)
			},
			errStr: "scaleway-sdk-go: invalid organization ID format 'invalid', expected a UUID: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		},
		{
			name: "Should throw an empty project ID error",
			clientOption: func(s *settings) {
				s.token = auth.NewToken(v2ValidAccessKey, v2ValidSecretKey)
				s.defaultProjectID = StringPtr("")
			},
			errStr: "scaleway-sdk-go: default project ID cannot be empty",
		},
		{
			name: "Should throw a bad project ID error",
			clientOption: func(s *settings) {
				s.token = auth.NewToken(v2ValidAccessKey, v2ValidSecretKey)
				s.defaultProjectID = StringPtr(v2InvalidDefaultProjectID)
			},
			errStr: "scaleway-sdk-go: invalid project ID format 'invalid', expected a UUID: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		},
		{
			name: "Should throw a region error",
			clientOption: func(s *settings) {
				v := Region("")
				s.token = auth.NewToken(testAccessKey, testSecretKey)
				s.defaultRegion = &v
			},
			errStr: "scaleway-sdk-go: default region cannot be empty",
		},
		{
			name: "Should throw a bad region error",
			clientOption: func(s *settings) {
				v := Region(v2InvalidDefaultRegion)
				s.token = auth.NewToken(testAccessKey, testSecretKey)
				s.defaultRegion = &v
			},
			errStr: "scaleway-sdk-go: invalid default region format 'invalid', available regions are: fr-par, nl-ams, pl-waw",
		},
		{
			name: "Should throw a zone error",
			clientOption: func(s *settings) {
				v := Zone("")
				s.token = auth.NewToken(testAccessKey, testSecretKey)
				s.defaultZone = &v
			},
			errStr: "scaleway-sdk-go: default zone cannot be empty",
		},
		{
			name: "Should throw a bad zone error",
			clientOption: func(s *settings) {
				v := Zone(v2InvalidDefaultZone)
				s.token = auth.NewToken(testAccessKey, testSecretKey)
				s.defaultZone = &v
			},
			errStr: "scaleway-sdk-go: invalid default zone format 'invalid', available zones are: fr-par-1, fr-par-2, fr-par-3, nl-ams-1, nl-ams-2, nl-ams-3, pl-waw-1, pl-waw-2, pl-waw-3",
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			// New
			s := newSettings()

			// Apply
			s.apply([]scw.ClientOption{c.clientOption})

			// Validate
			err := s.validate()

			if c.errStr != "" {
				testhelpers.Assert(t, err != nil, "Should have error")
				testhelpers.Equals(t, c.errStr, err.Error())
			} else {
				testhelpers.AssertNoError(t, err)
			}
		})
	}
}

func TestCombinedClientOptions(t *testing.T) {
	tests := []struct {
		name  string
		env   map[string]string
		files map[string]string

		expectedError                 string
		expectedAccessKey             string
		expectedSecretKey             string
		expectedAPIURL                string
		expectedDefaultOrganizationID *string
		expectedDefaultProjectID      *string
		expectedDefaultRegion         *scw.Region
		expectedDefaultZone           *scw.Zone
	}{
		{
			name: "Complete config file with env variables",
			env: map[string]string{
				"HOME":                          "{HOME}",
				scw.ScwAccessKeyEnv:             v2ValidAccessKey2,
				scw.ScwSecretKeyEnv:             v2ValidSecretKey2,
				scw.ScwAPIURLEnv:                v2ValidAPIURL2,
				scw.ScwDefaultOrganizationIDEnv: v2ValidDefaultOrganizationID2,
				scw.ScwDefaultProjectIDEnv:      v2ValidDefaultProjectID2,
				scw.ScwDefaultRegionEnv:         v2ValidDefaultRegion2,
				scw.ScwDefaultZoneEnv:           v2ValidDefaultZone2,
			},
			files: map[string]string{
				".config/scw/config.yaml": v2CompleteValidConfigFile,
			},
			expectedAccessKey:             v2ValidAccessKey2,
			expectedSecretKey:             v2ValidSecretKey2,
			expectedAPIURL:                v2ValidAPIURL2,
			expectedDefaultOrganizationID: s(v2ValidDefaultOrganizationID2),
			expectedDefaultProjectID:      s(v2ValidDefaultProjectID2),
			expectedDefaultRegion:         r(scw.Region(v2ValidDefaultRegion2)),
			expectedDefaultZone:           z(scw.Zone(v2ValidDefaultZone2)),
		},
		{
			name: "Complete config with active profile env variable and all env variables",
			env: map[string]string{
				"HOME":                          "{HOME}",
				scw.ScwActiveProfileEnv:         v2ValidProfile,
				scw.ScwAccessKeyEnv:             v2ValidAccessKey,
				scw.ScwSecretKeyEnv:             v2ValidSecretKey,
				scw.ScwAPIURLEnv:                v2ValidAPIURL,
				scw.ScwDefaultOrganizationIDEnv: v2ValidDefaultOrganizationID,
				scw.ScwDefaultProjectIDEnv:      v2ValidDefaultProjectID,
				scw.ScwDefaultRegionEnv:         v2ValidDefaultRegion,
				scw.ScwDefaultZoneEnv:           v2ValidDefaultZone,
			},
			files: map[string]string{
				".config/scw/config.yaml": v2CompleteValidConfigFile,
			},
			expectedAccessKey:             v2ValidAccessKey,
			expectedSecretKey:             v2ValidSecretKey,
			expectedAPIURL:                v2ValidAPIURL,
			expectedDefaultOrganizationID: s(v2ValidDefaultOrganizationID),
			expectedDefaultProjectID:      s(v2ValidDefaultProjectID),
			expectedDefaultRegion:         r(scw.Region(v2ValidDefaultRegion)),
			expectedDefaultZone:           z(scw.Zone(v2ValidDefaultZone)),
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
			test.expectedError = strings.ReplaceAll(test.expectedError, "{HOME}", dir)

			// remove config file(s)
			defer cleanEnv(t, test.files, dir)

			config, err := scw.LoadConfig()
			testhelpers.AssertNoError(t, err)

			p, err := config.GetActiveProfile()
			testhelpers.AssertNoError(t, err)

			client, err := scw.NewClient(scw.WithProfile(p), scw.WithEnv())
			if test.expectedError == "" {
				testhelpers.AssertNoError(t, err)

				// assert getters
				testhelpers.Equals(t, test.expectedAccessKey, client.auth.(*auth.Token).AccessKey)
				testhelpers.Equals(t, test.expectedSecretKey, client.auth.(*auth.Token).SecretKey)
				testhelpers.Equals(t, test.expectedAPIURL, client.apiURL)
				testhelpers.Equals(t, test.expectedDefaultOrganizationID, client.defaultOrganizationID)
				testhelpers.Equals(t, test.expectedDefaultProjectID, client.defaultProjectID)
				testhelpers.Equals(t, test.expectedDefaultRegion, client.defaultRegion)
				testhelpers.Equals(t, test.expectedDefaultZone, client.defaultZone)
				// skip insecure tests
			} else {
				testhelpers.Equals(t, test.expectedError, err.Error())
			}
		})
	}
}

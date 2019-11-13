package scw

import (
	"os"
	"strings"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/auth"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

const (
	apiURL = "https://example.com/"
)

var (
	defaultOrganizationID = "6170692e-7363-616c-6577-61792e636f6d" // hint: | xxd -ps -r
	defaultRegion         = RegionNlAms
	defaultZone           = ZoneNlAms1
)

func TestClientOptions(t *testing.T) {

	testCases := []struct {
		name         string
		clientOption ClientOption
		errStr       string
	}{
		{
			name: "Create a valid client option",
			clientOption: func(s *settings) {
				s.token = auth.NewToken(testAccessKey, testSecretKey)
				s.apiURL = apiURL
				s.defaultOrganizationID = &defaultOrganizationID
				s.defaultRegion = &defaultRegion
				s.defaultZone = &defaultZone
			},
		},
		{
			name: "Should throw an access key error",
			clientOption: func(s *settings) {
				s.apiURL = apiURL
				s.token = auth.NewToken("", testSecretKey)
			},
			errStr: "scaleway-sdk-go: access key cannot be empty",
		},
		{
			name: "Should throw a secret key error",
			clientOption: func(s *settings) {
				s.apiURL = apiURL
				s.token = auth.NewToken(testSecretKey, "")
			},
			errStr: "scaleway-sdk-go: secret key cannot be empty",
		},
		{
			name: "Should throw an url error",
			clientOption: func(s *settings) {
				s.apiURL = ":test"
				s.token = auth.NewToken(testAccessKey, testSecretKey)
			},
			errStr: "scaleway-sdk-go: invalid url :test: parse :test: missing protocol scheme",
		},
		{
			name: "Should throw a organization id error",
			clientOption: func(s *settings) {
				v := ""
				s.token = auth.NewToken(testAccessKey, testSecretKey)
				s.defaultOrganizationID = &v
			},
			errStr: "scaleway-sdk-go: default organization id cannot be empty",
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
			name: "Should throw a zone error",
			clientOption: func(s *settings) {
				v := Zone("")
				s.token = auth.NewToken(testAccessKey, testSecretKey)
				s.defaultZone = &v
			},
			errStr: "scaleway-sdk-go: default zone cannot be empty",
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			// New
			s := newSettings()

			// Apply
			s.apply([]ClientOption{c.clientOption})

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
		expectedDefaultRegion         *Region
		expectedDefaultZone           *Zone
	}{
		{
			name: "Complete config file with env variables",
			env: map[string]string{
				"HOME":                      "{HOME}",
				scwAccessKeyEnv:             v2ValidAccessKey2,
				scwSecretKeyEnv:             v2ValidSecretKey2,
				scwAPIURLEnv:                v2ValidAPIURL2,
				scwDefaultOrganizationIDEnv: v2ValidDefaultOrganizationID2,
				scwDefaultRegionEnv:         v2ValidDefaultRegion2,
				scwDefaultZoneEnv:           v2ValidDefaultZone2,
			},
			files: map[string]string{
				".config/scw/config.yaml": v2CompleteValidConfigFile,
			},
			expectedAccessKey:             v2ValidAccessKey2,
			expectedSecretKey:             v2ValidSecretKey2,
			expectedAPIURL:                v2ValidAPIURL2,
			expectedDefaultOrganizationID: s(v2ValidDefaultOrganizationID2),
			expectedDefaultRegion:         r(Region(v2ValidDefaultRegion2)),
			expectedDefaultZone:           z(Zone(v2ValidDefaultZone2)),
		},
		{
			name: "Complete config with active profile env variable and all env variables",
			env: map[string]string{
				"HOME":                      "{HOME}",
				scwActiveProfileEnv:         v2ValidProfile,
				scwAccessKeyEnv:             v2ValidAccessKey,
				scwSecretKeyEnv:             v2ValidSecretKey,
				scwAPIURLEnv:                v2ValidAPIURL,
				scwDefaultOrganizationIDEnv: v2ValidDefaultOrganizationID,
				scwDefaultRegionEnv:         v2ValidDefaultRegion,
				scwDefaultZoneEnv:           v2ValidDefaultZone,
			},
			files: map[string]string{
				".config/scw/config.yaml": v2CompleteValidConfigFile,
			},
			expectedAccessKey:             v2ValidAccessKey,
			expectedSecretKey:             v2ValidSecretKey,
			expectedAPIURL:                v2ValidAPIURL,
			expectedDefaultOrganizationID: s(v2ValidDefaultOrganizationID),
			expectedDefaultRegion:         r(Region(v2ValidDefaultRegion)),
			expectedDefaultZone:           z(Zone(v2ValidDefaultZone)),
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
			test.expectedError = strings.Replace(test.expectedError, "{HOME}", dir, -1)

			// remove config file(s)
			defer cleanEnv(t, test.files, dir)

			config, err := LoadConfig()
			testhelpers.AssertNoError(t, err)

			p, err := config.GetActiveProfile()
			testhelpers.AssertNoError(t, err)

			client, err := NewClient(WithProfile(p), WithEnv())
			if test.expectedError == "" {
				testhelpers.AssertNoError(t, err)

				// assert getters
				testhelpers.Equals(t, test.expectedAccessKey, client.auth.(*auth.Token).AccessKey)
				testhelpers.Equals(t, test.expectedSecretKey, client.auth.(*auth.Token).SecretKey)
				testhelpers.Equals(t, test.expectedAPIURL, client.apiURL)
				testhelpers.Equals(t, test.expectedDefaultOrganizationID, client.defaultOrganizationID)
				testhelpers.Equals(t, test.expectedDefaultRegion, client.defaultRegion)
				testhelpers.Equals(t, test.expectedDefaultZone, client.defaultZone)
				// skip insecure tests
			} else {
				testhelpers.Equals(t, test.expectedError, err.Error())
			}

		})
	}
}

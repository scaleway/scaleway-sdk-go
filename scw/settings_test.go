package scw

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/auth"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/utils"
)

const (
	apiURL = "https://example.com/"
)

var (
	defaultProjectID = "6170692e-7363-616c-6577-61792e636f6d" // hint: | xxd -ps -r
	defaultRegion    = utils.RegionNlAms
	defaultZone      = utils.ZoneNlAms1
)

func TestSettings(t *testing.T) {

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
				s.defaultProjectID = &defaultProjectID
				s.defaultRegion = &defaultRegion
				s.defaultZone = &defaultZone
			},
		},
		{
			name: "Should throw an credential error",
			clientOption: func(s *settings) {
				s.apiURL = apiURL
			},
			errStr: "no credential option provided",
		},
		{
			name: "Should throw an url error",
			clientOption: func(s *settings) {
				s.apiURL = ":test"
				s.token = auth.NewToken(testAccessKey, testSecretKey)
			},
			errStr: "invalid url :test: parse :test: missing protocol scheme",
		},
		{
			name: "Should throw a project id error",
			clientOption: func(s *settings) {
				v := ""
				s.token = auth.NewToken(testAccessKey, testSecretKey)
				s.defaultProjectID = &v
			},
			errStr: "default project id cannot be empty",
		},
		{
			name: "Should throw a region error",
			clientOption: func(s *settings) {
				v := utils.Region("")
				s.token = auth.NewToken(testAccessKey, testSecretKey)
				s.defaultRegion = &v
			},
			errStr: "default region cannot be empty",
		},
		{
			name: "Should throw a zone error",
			clientOption: func(s *settings) {
				v := utils.Zone("")
				s.token = auth.NewToken(testAccessKey, testSecretKey)
				s.defaultZone = &v
			},
			errStr: "default zone cannot be empty",
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
				testhelpers.Ok(t, err)
			}
		})
	}
}

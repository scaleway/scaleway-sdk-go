package scw

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/auth"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

const (
	apiURL = "https://example.com/"
)

func TestSettings(t *testing.T) {
	// New
	s := newSettings()

	// Apply
	var funcToApply ClientOption = func(s *settings) {
		s.token = auth.NewToken(testAccessKey, testSecretKey)
		s.apiURL = apiURL
	}

	s.apply([]ClientOption{funcToApply})

	// Validate
	err := s.validate()
	testhelpers.Ok(t, err)
}

func TestSettingsInvalidToken(t *testing.T) {
	// New
	s := newSettings()

	// Apply
	var setURL ClientOption = func(s *settings) {
		s.apiURL = apiURL
	}

	s.apply([]ClientOption{setURL})

	// Validate without auth
	err := s.validate()
	testhelpers.Assert(t, err != nil, "Should have error")
	testhelpers.Equals(t, "no credential option provided", err.Error())
}
func TestSettingsInvalidURL(t *testing.T) {
	// New
	s := newSettings()

	// Apply
	var setTokenUnsetURL ClientOption = func(s *settings) {
		s.apiURL = ":test"
		s.token = auth.NewToken(testAccessKey, testSecretKey)
	}

	s.apply([]ClientOption{setTokenUnsetURL})

	err := s.validate()

	testhelpers.Assert(t, err != nil, "Should have error")
	testhelpers.Equals(t, "invalid url :test: parse :test: missing protocol scheme", err.Error())
}

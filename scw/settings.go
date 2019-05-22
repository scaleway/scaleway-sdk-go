package scw

import (
	"fmt"
	"net/url"

	"github.com/scaleway/scaleway-sdk-go/internal/auth"
	"github.com/scaleway/scaleway-sdk-go/utils"
)

type settings struct {
	apiURL                string
	token                 auth.Auth
	userAgent             string
	httpClient            httpClient
	insecure              bool
	defaultOrganizationID *string
	defaultRegion         *utils.Region
	defaultZone           *utils.Zone
}

func newSettings() *settings {
	return &settings{}
}

func (s *settings) apply(opts []ClientOption) {
	for _, opt := range opts {
		opt(s)
	}
}

func (s *settings) validate() error {
	var err error
	if s.token == nil {
		return fmt.Errorf("no credential option provided")
	}

	_, err = url.Parse(s.apiURL)
	if err != nil {
		return fmt.Errorf("invalid url %s: %s", s.apiURL, err)
	}

	// TODO: Check OrganizationID format
	if s.defaultOrganizationID != nil && *s.defaultOrganizationID == "" {
		return fmt.Errorf("default organization id cannot be empty")
	}

	// TODO: Check Region format
	if s.defaultRegion != nil && *s.defaultRegion == "" {
		return fmt.Errorf("default region cannot be empty")
	}

	// TODO: Check Zone format
	if s.defaultZone != nil && *s.defaultZone == "" {
		return fmt.Errorf("default zone cannot be empty")
	}

	return nil
}

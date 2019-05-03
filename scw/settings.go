package scw

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/scaleway/scaleway-sdk-go/internal/auth"
)

type settings struct {
	apiUrl                string
	token                 auth.Auth
	userAgent             string
	httpClient            *http.Client
	insecure              bool
	defaultOrganizationId string
	defaultRegion         Region
	defaultZone           Zone
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

	_, err = url.Parse(s.apiUrl)
	if err != nil {
		return fmt.Errorf("invalid url %s: %s", s.apiUrl, err)
	}

	return nil
}

package scw

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/scaleway/scaleway-sdk-go/internal/auth"
)

type settings struct {
	Url        string
	Token      auth.Auth
	UserAgent  string
	HttpClient *http.Client
	Insecure   bool

	DefaultOrganizationId string
	DefaultRegion         Region
	DefaultZone           Zone
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
	if s.Token == nil {
		return fmt.Errorf("no credential option provided")
	}

	_, err = url.Parse(s.Url)
	if err != nil {
		return fmt.Errorf("invalid url %s: %s", s.Url, err)
	}

	return nil
}

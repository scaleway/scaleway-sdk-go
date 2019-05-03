package scw

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/auth"
)

// Client is the Scaleway client which performs API requests.
//
// This client should be passed in the `NewApi` functions whenever an API instance is created.
// Creating a Client is done with the `NewClient` function.
type Client struct {
	httpClient            *http.Client
	auth                  auth.Auth
	apiUrl                string
	userAgent             string
	defaultOrganizationId string
	defaultRegion         Region
	defaultZone           Zone
}

// NewClient instantiates a new Client object.
//
// Zero or more ClientOption object can be passed as a parameter.
// These options will then be applied to the client.
func NewClient(opts ...ClientOption) (*Client, error) {
	s := newSettings()

	// apply options
	s.apply(append(defaultOptions(), opts...))

	// validate settings
	err := s.validate()
	if err != nil {
		return nil, err
	}

	// dial the API
	if s.HttpClient == nil {
		s.HttpClient = newHttpClient()
	}

	// insecure mode
	if s.Insecure {
		clientTransport, ok := s.HttpClient.Transport.(*http.Transport)
		if !ok {
			return nil, fmt.Errorf("cannot use insecure mode with HTTP client of type %T", s.HttpClient.Transport)
		}
		if clientTransport.TLSClientConfig == nil {
			clientTransport.TLSClientConfig = &tls.Config{}
		}
		clientTransport.TLSClientConfig.InsecureSkipVerify = true
	}

	return &Client{
		auth:                  s.Token,
		httpClient:            s.HttpClient,
		apiUrl:                s.ApiUrl,
		userAgent:             s.UserAgent,
		defaultOrganizationId: s.DefaultOrganizationId,
		defaultRegion:         s.DefaultRegion,
		defaultZone:           s.DefaultZone,
	}, nil
}

func (c *Client) GetDefaultOrganizationId() string {
	return c.defaultOrganizationId
}

func (c *Client) GetDefaultRegion() Region {
	return c.defaultRegion
}

func (c *Client) GetDefaultZone() Zone {
	return c.defaultZone
}

func defaultOptions() []ClientOption {
	return []ClientOption{
		WithApiUrl("https://api.scaleway.com"),
		WithUserAgent(userAgent),
	}
}

func newHttpClient() *http.Client {
	return &http.Client{
		Timeout: 9 * time.Second,
		Transport: &http.Transport{
			DialContext:           (&net.Dialer{Timeout: 5 * time.Second}).DialContext,
			TLSHandshakeTimeout:   5 * time.Second,
			ResponseHeaderTimeout: 8 * time.Second,
			MaxIdleConnsPerHost:   20,
		},
	}
}

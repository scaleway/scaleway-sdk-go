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
	apiURL                string
	userAgent             string
	defaultOrganizationID string
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
	if s.httpClient == nil {
		s.httpClient = newHTTPClient()
	}

	// insecure mode
	if s.insecure {
		clientTransport, ok := s.httpClient.Transport.(*http.Transport)
		if !ok {
			return nil, fmt.Errorf("cannot use insecure mode with HTTP client of type %T", s.httpClient.Transport)
		}
		if clientTransport.TLSClientConfig == nil {
			clientTransport.TLSClientConfig = &tls.Config{}
		}
		clientTransport.TLSClientConfig.InsecureSkipVerify = true
	}

	return &Client{
		auth:                  s.token,
		httpClient:            s.httpClient,
		apiURL:                s.apiURL,
		userAgent:             s.userAgent,
		defaultOrganizationID: s.defaultOrganizationID,
		defaultRegion:         s.defaultRegion,
		defaultZone:           s.defaultZone,
	}, nil
}

// GetDefaultOrganizationID return the default organization ID
// of the client. This value can be set from the client option
// WithDefaultOrganizationID().
func (c *Client) GetDefaultOrganizationID() string {
	return c.defaultOrganizationID
}

// GetDefaultRegion return the default region of the client.
// This value can be set from the client option
// WithDefaultRegion().
func (c *Client) GetDefaultRegion() Region {
	return c.defaultRegion
}

// GetDefaultZone return the default zone of the client.
// This value can be set from the client option
// WithDefaultZone().
func (c *Client) GetDefaultZone() Zone {
	return c.defaultZone
}

func defaultOptions() []ClientOption {
	return []ClientOption{
		WithAPIURL("https://api.scaleway.com"),
		WithUserAgent(userAgent),
	}
}

func newHTTPClient() *http.Client {
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

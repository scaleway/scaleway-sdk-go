package scw

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/auth"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/utils"
)

// Client is the Scaleway client which performs API requests.
//
// This client should be passed in the `NewApi` functions whenever an API instance is created.
// Creating a Client is done with the `NewClient` function.
type Client struct {
	httpClient            httpClient
	auth                  auth.Auth
	apiURL                string
	userAgent             string
	defaultOrganizationID string
	defaultRegion         utils.Region
	defaultZone           utils.Zone
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
		logger.Debugf("client: using insecure mode")
		setInsecureMode(s.httpClient)
	}

	logger.Debugf("client: using sdk version " + version)

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
// WithDefaultOrganizationID(). Be aware this value can be empty.
func (c *Client) GetDefaultOrganizationID() string {
	return c.defaultOrganizationID
}

// GetDefaultRegion return the default region of the client.
// This value can be set from the client option
// WithDefaultRegion(). Be aware this value can be empty.
func (c *Client) GetDefaultRegion() utils.Region {
	return c.defaultRegion
}

// GetDefaultZone return the default zone of the client.
// This value can be set from the client option
// WithDefaultZone(). Be aware this value can be empty.
func (c *Client) GetDefaultZone() utils.Zone {
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

func setInsecureMode(c httpClient) {
	standardHTTPClient, ok := c.(*http.Client)
	if !ok {
		logger.Warningf("client: cannot use insecure mode with HTTP client of type %T", c)
		return
	}
	transportClient, ok := standardHTTPClient.Transport.(*http.Transport)
	if !ok {
		logger.Warningf("client: cannot use insecure mode with Transport client of type %T", standardHTTPClient.Transport)
		return
	}
	if transportClient.TLSClientConfig == nil {
		transportClient.TLSClientConfig = &tls.Config{}
	}
	transportClient.TLSClientConfig.InsecureSkipVerify = true
}

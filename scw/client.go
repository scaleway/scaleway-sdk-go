package scw

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
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
	defaultOrganizationID *string
	defaultRegion         *utils.Region
	defaultZone           *utils.Zone
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
// of the client. This value can be set in the client option
// WithDefaultOrganizationID(). Be aware this value can be empty.
func (c *Client) GetDefaultOrganizationID() (string, bool) {
	if c.defaultOrganizationID != nil {
		return *c.defaultOrganizationID, true
	}
	return "", false
}

// GetDefaultRegion return the default region of the client.
// This value can be set in the client option
// WithDefaultRegion(). Be aware this value can be empty.
func (c *Client) GetDefaultRegion() (utils.Region, bool) {
	if c.defaultRegion != nil {
		return *c.defaultRegion, true
	}
	return utils.Region(""), false
}

// GetDefaultZone return the default zone of the client.
// This value can be set in the client option
// WithDefaultZone(). Be aware this value can be empty.
func (c *Client) GetDefaultZone() (utils.Zone, bool) {
	if c.defaultZone != nil {
		return *c.defaultZone, true
	}
	return utils.Zone(""), false
}

// Do performs an HTTP request based on the ScalewayRequest object.
// RequestOptions are executed on the ScalewayRequest.
func (c *Client) Do(req *ScalewayRequest, res interface{}, opts ...RequestOption) (err error) {
	if req == nil {
		return fmt.Errorf("request must be non-nil")
	}

	// apply request options
	for _, opt := range opts {
		opt(req)
	}

	// build url
	url, err := req.getURL(c.apiURL)
	if err != nil {
		return err
	}
	logger.Debugf("creating %s request on %s", req.Method, url.String())

	// build request
	request, err := http.NewRequest(req.Method, url.String(), req.Body)
	if err != nil {
		return err
	}

	request.Header = req.getAllHeaders(c.auth, c.userAgent, false)

	if req.Ctx != nil {
		request = request.WithContext(req.Ctx)
	}

	if logger.ShouldLog(logger.LogLevelDebug) {

		// Keep original headers (before anonymization)
		originalHeaders := request.Header

		// Get anonymized headers
		request.Header = req.getAllHeaders(c.auth, c.userAgent, true)

		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			logger.Warningf("cannot dump outgoing request: %s", err)
		} else {
			logger.Debugf("dumping http request:\n" + string(dump))
		}

		// Restore original headers before sending the request
		request.Header = originalHeaders
	}

	// execute request
	resp, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}

	defer func() {
		tmpErr := resp.Body.Close()
		if err == nil && tmpErr != nil {
			err = tmpErr
		}
	}()

	err = hasResponseError(resp)
	if err != nil {
		return err
	}

	if res != nil {
		err = json.NewDecoder(resp.Body).Decode(&res)
		if err != nil {
			return err
		}
	}

	return nil
}

func defaultOptions() []ClientOption {
	return []ClientOption{
		WithAPIURL("https://api.scaleway.com"),
		WithUserAgent(userAgent),
	}
}

func newHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			DialContext:           (&net.Dialer{Timeout: 5 * time.Second}).DialContext,
			TLSHandshakeTimeout:   5 * time.Second,
			ResponseHeaderTimeout: 30 * time.Second,
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

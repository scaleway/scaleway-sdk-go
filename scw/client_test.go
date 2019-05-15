package scw

import (
	"bytes"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/auth"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/utils"
)

const (
	testAPIURL                = "https://api.example.com/"
	defaultAPIURL             = "https://api.scaleway.com"
	testAccessKey             = "ACCESS_KEY"
	testSecretKey             = "7363616c-6577-6573-6862-6f7579616161" // hint: | xxd -ps -r
	testDefaultOrganizationID = "6170692e-7363-616c-6577-61792e636f6d" // hint: | xxd -ps -r
	testDefaultRegion         = utils.RegionFrPar
	testDefaultZone           = utils.ZoneFrPar1
	testInsecure              = true
)

func TestNewClientWithDefaults(t *testing.T) {

	options := []ClientOption{
		WithoutAuth(),
		WithInsecure(),
	}

	client, err := NewClient(options...)
	testhelpers.Ok(t, err)

	testhelpers.Equals(t, defaultAPIURL, client.apiURL)
	testhelpers.Equals(t, auth.NewNoAuth(), client.auth)

}

type TestConfig struct{}

func (c *TestConfig) GetAccessKey() (string, bool) {
	return testAccessKey, true
}
func (c *TestConfig) GetSecretKey() (string, bool) {
	return testSecretKey, true
}
func (c *TestConfig) GetAPIURL() (string, bool) {
	return testAPIURL, true
}
func (c *TestConfig) GetInsecure() (bool, bool) {
	return testInsecure, true
}
func (c *TestConfig) GetDefaultOrganizationID() (string, bool) {
	return testDefaultOrganizationID, true
}
func (c *TestConfig) GetDefaultRegion() (utils.Region, bool) {
	return testDefaultRegion, true
}
func (c *TestConfig) GetDefaultZone() (utils.Zone, bool) {
	return testDefaultZone, true
}

func TestNewClientWithOptions(t *testing.T) {

	t.Run("Basic", func(t *testing.T) {
		someHTTPClient := &http.Client{}

		options := []ClientOption{
			WithAPIURL(testAPIURL),
			WithAuth(testAccessKey, testSecretKey),
			WithHTTPClient(someHTTPClient),
			WithDefaultOrganizationID(testDefaultOrganizationID),
			WithDefaultRegion(testDefaultRegion),
			WithDefaultZone(testDefaultZone),
		}

		client, err := NewClient(options...)
		testhelpers.Ok(t, err)

		testhelpers.Equals(t, testAPIURL, client.apiURL)
		testhelpers.Equals(t, auth.NewToken(testAccessKey, testSecretKey), client.auth)

		testhelpers.Equals(t, someHTTPClient, client.httpClient)

		testhelpers.Equals(t, testDefaultOrganizationID, client.GetDefaultOrganizationID())
		testhelpers.Equals(t, testDefaultRegion, client.GetDefaultRegion())
		testhelpers.Equals(t, testDefaultZone, client.GetDefaultZone())
	})

	t.Run("With scwconfig", func(t *testing.T) {
		config := &TestConfig{}

		client, err := NewClient(WithConfig(config))
		testhelpers.Ok(t, err)

		testhelpers.Equals(t, auth.NewToken(testAccessKey, testSecretKey), client.auth)
		testhelpers.Equals(t, testAPIURL, client.apiURL)

		clientTransport, ok := client.httpClient.(*http.Client).Transport.(*http.Transport)
		testhelpers.Assert(t, ok, "clientTransport must be not nil")
		testhelpers.Assert(t, clientTransport.TLSClientConfig != nil, "TLSClientConfig must be not nil")
		testhelpers.Equals(t, testInsecure, clientTransport.TLSClientConfig.InsecureSkipVerify)

		testhelpers.Equals(t, testDefaultOrganizationID, client.GetDefaultOrganizationID())
		testhelpers.Equals(t, testDefaultRegion, client.GetDefaultRegion())
		testhelpers.Equals(t, testDefaultZone, client.GetDefaultZone())
	})

}

type fakeHTTPClient struct{}

func (fakeHTTPClient) Do(*http.Request) (*http.Response, error) {
	return nil, nil
}

func (fakeHTTPClient) RoundTrip(*http.Request) (*http.Response, error) {
	return nil, nil
}

// TestSetInsecureMode test if setInsecureMode panic when given custom HTTP client
func TestSetInsecureMode(t *testing.T) {
	var buf bytes.Buffer
	logger.DefaultLogger.Init(&buf, logger.LogLevelWarning)

	// custom Transport client
	clientWithFakeTransport := newHTTPClient()
	clientWithFakeTransport.Transport = fakeHTTPClient{}
	setInsecureMode(clientWithFakeTransport)

	// custom HTTP client
	setInsecureMode(fakeHTTPClient{})

	// check log messages
	lines := strings.Split(buf.String(), "\n")
	getLogMessage := func(s string) string {
		return strings.Join(strings.Split(s, " ")[3:], " ")
	}
	testhelpers.Equals(t, "cannot use insecure mode with Transport client of type scw.fakeHTTPClient", getLogMessage(lines[0]))
	testhelpers.Equals(t, "cannot use insecure mode with HTTP client of type scw.fakeHTTPClient", getLogMessage(lines[1]))

	logger.DefaultLogger.Init(os.Stderr, logger.LogLevelWarning)
}

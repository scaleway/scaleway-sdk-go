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
)

const (
	testAPIURL           = "https://api.example.com/"
	defaultAPIURL        = "https://api.scaleway.com"
	testAccessKey        = "ACCESS_KEY"
	testSecretKey        = "7363616c-6577-6573-6862-6f7579616161" // hint: | xxd -ps -r
	testDefaultProjectID = "6170692e-7363-616c-6577-61792e636f6d" // hint: | xxd -ps -r
	testDefaultRegion    = RegionFrPar
	testDefaultZone      = ZoneFrPar1
	testDefaultPageSize  = int32(5)
	testInsecure         = true
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

type mockConfig struct{}

func (c *mockConfig) GetAccessKey() (string, bool) {
	return testAccessKey, true
}
func (c *mockConfig) GetSecretKey() (string, bool) {
	return testSecretKey, true
}
func (c *mockConfig) GetAPIURL() (string, bool) {
	return testAPIURL, true
}
func (c *mockConfig) GetInsecure() (bool, bool) {
	return testInsecure, true
}
func (c *mockConfig) GetDefaultProjectID() (string, bool) {
	return testDefaultProjectID, true
}
func (c *mockConfig) GetDefaultRegion() (Region, bool) {
	return testDefaultRegion, true
}
func (c *mockConfig) GetDefaultZone() (Zone, bool) {
	return testDefaultZone, true
}

func TestNewClientWithOptions(t *testing.T) {

	t.Run("Basic", func(t *testing.T) {
		someHTTPClient := &http.Client{}

		options := []ClientOption{
			WithAPIURL(testAPIURL),
			WithAuth(testAccessKey, testSecretKey),
			WithHTTPClient(someHTTPClient),
			WithDefaultProjectID(testDefaultProjectID),
			WithDefaultRegion(testDefaultRegion),
			WithDefaultZone(testDefaultZone),
			WithDefaultPageSize(testDefaultPageSize),
		}

		client, err := NewClient(options...)
		testhelpers.Ok(t, err)

		testhelpers.Equals(t, testAPIURL, client.apiURL)
		testhelpers.Equals(t, auth.NewToken(testAccessKey, testSecretKey), client.auth)

		testhelpers.Equals(t, someHTTPClient, client.httpClient)

		defaultProjectID, exist := client.GetDefaultProjectID()
		testhelpers.Equals(t, testDefaultProjectID, defaultProjectID)
		testhelpers.Assert(t, exist, "defaultProjectID must exist")

		defaultRegion, exist := client.GetDefaultRegion()
		testhelpers.Equals(t, testDefaultRegion, defaultRegion)
		testhelpers.Assert(t, exist, "defaultRegion must exist")

		defaultZone, exist := client.GetDefaultZone()
		testhelpers.Equals(t, testDefaultZone, defaultZone)
		testhelpers.Assert(t, exist, "defaultZone must exist")

		defaultPageSize, exist := client.GetDefaultPageSize()
		testhelpers.Equals(t, testDefaultPageSize, defaultPageSize)
		testhelpers.Assert(t, exist, "defaultPageSize must exist")
	})

	t.Run("With scwconfig", func(t *testing.T) {
		config := &mockConfig{}

		client, err := NewClient(WithConfig(config))
		testhelpers.Ok(t, err)

		testhelpers.Equals(t, auth.NewToken(testAccessKey, testSecretKey), client.auth)
		testhelpers.Equals(t, testAPIURL, client.apiURL)

		clientTransport, ok := client.httpClient.(*http.Client).Transport.(*http.Transport)
		testhelpers.Assert(t, ok, "clientTransport must be not nil")
		testhelpers.Assert(t, clientTransport.TLSClientConfig != nil, "TLSClientConfig must be not nil")
		testhelpers.Equals(t, testInsecure, clientTransport.TLSClientConfig.InsecureSkipVerify)

		defaultProjectID, exist := client.GetDefaultProjectID()
		testhelpers.Equals(t, testDefaultProjectID, defaultProjectID)
		testhelpers.Assert(t, exist, "defaultProjectID must exist")

		defaultRegion, exist := client.GetDefaultRegion()
		testhelpers.Equals(t, testDefaultRegion, defaultRegion)
		testhelpers.Assert(t, exist, "defaultRegion must exist")

		defaultZone, exist := client.GetDefaultZone()
		testhelpers.Equals(t, testDefaultZone, defaultZone)
		testhelpers.Assert(t, exist, "defaultZone must exist")

		_, exist = client.GetDefaultPageSize()
		testhelpers.Assert(t, !exist, "defaultPageSize must not exist")
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
	testhelpers.Equals(t, "client: cannot use insecure mode with Transport client of type scw.fakeHTTPClient", getLogMessage(lines[0]))
	testhelpers.Equals(t, "client: cannot use insecure mode with HTTP client of type scw.fakeHTTPClient", getLogMessage(lines[1]))

	logger.DefaultLogger.Init(os.Stderr, logger.LogLevelWarning)
}

func TestNewPage(t *testing.T) {
	type fakeType struct {
		plop int
	}

	testhelpers.Equals(t, &fakeType{}, newPage(&fakeType{3}))
}

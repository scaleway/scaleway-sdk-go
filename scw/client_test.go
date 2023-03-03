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
	testAPIURL                = "https://api.example.com"
	defaultAPIURL             = "https://api.scaleway.com"
	testAccessKey             = "SCW1234567890ABCDEFG"
	testSecretKey             = "7363616c-6577-6573-6862-6f7579616161" // hint: | xxd -ps -r
	testDefaultOrganizationID = "6170692e-7363-616c-6577-61792e636f6d" // hint: | xxd -ps -r
	testDefaultProjectID      = "6170692e-7363-616c-6577-61792e636f6e"
	testDefaultRegion         = RegionFrPar
	testDefaultZone           = ZoneFrPar1
	testDefaultPageSize       = uint32(5)
	testInsecure              = true
)

func TestNewClientWithNoAuth(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		client, err := NewClient()
		testhelpers.AssertNoError(t, err)

		secretKey, exist := client.GetSecretKey()
		testhelpers.Equals(t, "", secretKey)
		testhelpers.Assert(t, !exist, "secretKey must not exist")

		accessKey, exist := client.GetAccessKey()
		testhelpers.Equals(t, "", accessKey)
		testhelpers.Assert(t, !exist, "accessKey must not exist")
	})
}

func TestNewClientMultipleClients(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		logger.EnableDebugMode()
		httpClient := &http.Client{}
		_, err := NewClient(WithHTTPClient(httpClient))
		testhelpers.AssertNoError(t, err)

		_, isLogger := httpClient.Transport.(*requestLoggerTransport)
		testhelpers.Assert(t, isLogger, "transport should be a request logger")

		_, err = NewClient(WithHTTPClient(httpClient))
		testhelpers.AssertNoError(t, err)

		transport, isLogger := httpClient.Transport.(*requestLoggerTransport)
		testhelpers.Assert(t, isLogger, "transport should be a request logger")
		_, isLogger = transport.rt.(*requestLoggerTransport)
		testhelpers.Assert(t, !isLogger, "nested transport should not be a request logger")

	})
}

func TestNewClientWithDefaults(t *testing.T) {
	options := []ClientOption{
		WithInsecure(),
	}

	client, err := NewClient(options...)
	testhelpers.AssertNoError(t, err)

	testhelpers.Equals(t, defaultAPIURL, client.apiURL)
	testhelpers.Equals(t, auth.NewNoAuth(), client.auth)
}

func TestNewClientWithOptions(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		someHTTPClient := &http.Client{}

		options := []ClientOption{
			WithAPIURL(testAPIURL),
			WithAuth(testAccessKey, testSecretKey),
			WithHTTPClient(someHTTPClient),
			WithDefaultOrganizationID(testDefaultOrganizationID),
			WithDefaultProjectID(testDefaultProjectID),
			WithDefaultRegion(testDefaultRegion),
			WithDefaultZone(testDefaultZone),
			WithDefaultPageSize(testDefaultPageSize),
		}

		client, err := NewClient(options...)
		testhelpers.AssertNoError(t, err)

		testhelpers.Equals(t, testAPIURL, client.apiURL)
		testhelpers.Equals(t, auth.NewToken(testAccessKey, testSecretKey), client.auth)

		testhelpers.Equals(t, someHTTPClient, client.httpClient)

		defaultOrganizationID, exist := client.GetDefaultOrganizationID()
		testhelpers.Equals(t, testDefaultOrganizationID, defaultOrganizationID)
		testhelpers.Assert(t, exist, "defaultOrganizationID must exist")

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

		secretKey, exist := client.GetSecretKey()
		testhelpers.Equals(t, testSecretKey, secretKey)
		testhelpers.Assert(t, exist, "secretKey must exist")

		accessKey, exist := client.GetAccessKey()
		testhelpers.Equals(t, testAccessKey, accessKey)
		testhelpers.Assert(t, exist, "accessKey must exist")
	})

	t.Run("With custom profile", func(t *testing.T) {
		profile := &Profile{
			s(testAccessKey),
			s(testSecretKey),
			s(testAPIURL),
			b(testInsecure),
			s(testDefaultOrganizationID),
			s(testDefaultProjectID),
			s(string(testDefaultRegion)),
			s(string(testDefaultZone)),
			b(true),
		}

		client, err := NewClient(WithProfile(profile))
		testhelpers.AssertNoError(t, err)

		testhelpers.Equals(t, auth.NewToken(testAccessKey, testSecretKey), client.auth)
		testhelpers.Equals(t, testAPIURL, client.apiURL)

		clientTransport, ok := client.httpClient.(*http.Client).Transport.(*http.Transport)
		if loggerTransport, isLogger := client.httpClient.(*http.Client).Transport.(*requestLoggerTransport); !ok && isLogger {
			clientTransport, ok = loggerTransport.rt.(*http.Transport)
		}
		testhelpers.Assert(t, ok, "clientTransport must be not nil")
		testhelpers.Assert(t, clientTransport.TLSClientConfig != nil, "TLSClientConfig must be not nil")
		testhelpers.Equals(t, testInsecure, clientTransport.TLSClientConfig.InsecureSkipVerify)

		defaultOrganizationID, exist := client.GetDefaultOrganizationID()
		testhelpers.Equals(t, testDefaultOrganizationID, defaultOrganizationID)
		testhelpers.Assert(t, exist, "defaultOrganizationID must exist")

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

		secretKey, exist := client.GetSecretKey()
		testhelpers.Equals(t, testSecretKey, secretKey)
		testhelpers.Assert(t, exist, "secretKey must exist")

		accessKey, exist := client.GetAccessKey()
		testhelpers.Equals(t, testAccessKey, accessKey)
		testhelpers.Assert(t, exist, "accessKey must exist")
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

func TestNewVariableFromType(t *testing.T) {
	type fakeType struct {
		plop int
	}

	testhelpers.Equals(t, &fakeType{}, newVariableFromType(&fakeType{3}))
}

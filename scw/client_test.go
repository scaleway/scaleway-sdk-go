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
	testS3Endpoint            = "https://s3.example.com"
	testS3UsePathStyle        = true
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
	t.Run("Only access key", func(t *testing.T) {
		client, err := NewClient(WithProfile(&Profile{
			AccessKey: StringPtr(testAccessKey),
		}))
		testhelpers.AssertNoError(t, err)

		secretKey, exist := client.GetSecretKey()
		testhelpers.Equals(t, "", secretKey)
		testhelpers.Assert(t, !exist, "secretKey must not exist")

		accessKey, exist := client.GetAccessKey()
		testhelpers.Assert(t, exist, "accessKey must exist")
		testhelpers.Equals(t, accessKey, testAccessKey)
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

func TestNewClientS3Endpoint(t *testing.T) {
	t.Run("Empty S3 endpoint", func(t *testing.T) {
		options := []ClientOption{
			WithInsecure(),
		}

		client, err := NewClient(options...)
		testhelpers.AssertNoError(t, err)

		testhelpers.Equals(t, "", client.s3Endpoint)
	})

	t.Run("Default S3 endpoint for fr-par", func(t *testing.T) {
		options := []ClientOption{
			WithInsecure(),
			WithDefaultRegion("fr-par"),
		}

		client, err := NewClient(options...)
		testhelpers.AssertNoError(t, err)

		testhelpers.Equals(t, "https://s3.fr-par.scw.cloud", client.s3Endpoint)
	})

	t.Run("Default S3 endpoint for nl-ams", func(t *testing.T) {
		options := []ClientOption{
			WithInsecure(),
			WithDefaultRegion("nl-ams"),
		}

		client, err := NewClient(options...)
		testhelpers.AssertNoError(t, err)

		testhelpers.Equals(t, "https://s3.nl-ams.scw.cloud", client.s3Endpoint)
	})

	t.Run("Custom S3 endpoint, with default region", func(t *testing.T) {
		options := []ClientOption{
			WithInsecure(),
			WithDefaultRegion("fr-par"),
			WithS3Endpoint("https://my-s3-endpoint.com"),
		}

		client, err := NewClient(options...)
		testhelpers.AssertNoError(t, err)

		testhelpers.Equals(t, "https://my-s3-endpoint.com", client.s3Endpoint)
	})

	t.Run("Custom S3 endpoint, without default region", func(t *testing.T) {
		options := []ClientOption{
			WithInsecure(),
			WithS3Endpoint("https://my-s3-endpoint.com"),
		}

		client, err := NewClient(options...)
		testhelpers.AssertNoError(t, err)

		testhelpers.Equals(t, "https://my-s3-endpoint.com", client.s3Endpoint)
	})

	t.Run("Custom S3 endpoint, validation error (trailing slash)", func(t *testing.T) {
		options := []ClientOption{
			WithInsecure(),
			WithS3Endpoint("https://my-s3-endpoint.com/"),
		}

		expectedErr := "invalid S3 endpoint 'https://my-s3-endpoint.com/': trailing slash is not allowed"

		_, err := NewClient(options...)
		if err == nil {
			t.Fatal("expected error, got none")
		} else if !strings.Contains(err.Error(), expectedErr) {
			t.Fatalf("expected error to contain '%s', got '%s'", expectedErr, err.Error())
		}
	})
}

func TestNewClientWithOptions(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		someHTTPClient := &http.Client{}

		options := []ClientOption{
			WithAPIURL(testAPIURL),
			WithS3Endpoint(testS3Endpoint),
			WithS3UsePathStyle(testS3UsePathStyle),
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

		s3Endpoint, exist := client.GetS3Endpoint()
		testhelpers.Equals(t, testS3Endpoint, s3Endpoint)
		testhelpers.Assert(t, exist, "s3Endpoint must exist")

		s3UsePathStyle := client.GetS3UsePathStyle()
		testhelpers.Equals(t, testS3UsePathStyle, s3UsePathStyle)

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
			s(testS3Endpoint),
			b(testS3UsePathStyle),
			b(testInsecure),
			s(testDefaultOrganizationID),
			s(testDefaultProjectID),
			s(string(testDefaultRegion)),
			s(string(testDefaultZone)),
			b(true),
			s(testUserAgent),
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

		s3Endpoint, exist := client.GetS3Endpoint()
		testhelpers.Equals(t, testS3Endpoint, s3Endpoint)
		testhelpers.Assert(t, exist, "s3Endpoint must exist")

		s3UsePathStyle := client.GetS3UsePathStyle()
		testhelpers.Equals(t, testS3UsePathStyle, s3UsePathStyle)

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

type fakeTransport struct{}

func (fakeTransport) RoundTrip(*http.Request) (*http.Response, error) {
	return nil, nil
}

type fakeAltTransport struct {
	insecure bool
}

func (f *fakeAltTransport) RoundTrip(*http.Request) (*http.Response, error) {
	return nil, nil
}

func (f *fakeAltTransport) SetInsecureTransport() {
	f.insecure = true
}

// TestSetInsecureMode test if setInsecureMode panic when given custom HTTP client
func TestSetInsecureMode(t *testing.T) {
	var buf bytes.Buffer
	logger.DefaultLogger.Init(&buf, logger.LogLevelWarning)

	// custom Transport client
	clientWithFakeTransport := newHTTPClient()
	clientWithFakeTransport.Transport = fakeTransport{}
	setInsecureMode(clientWithFakeTransport)

	// custom Alt transport client
	customTransport := &fakeAltTransport{}
	clientWithFakeTransport.Transport = customTransport
	setInsecureMode(clientWithFakeTransport)
	testhelpers.Equals(t, true, customTransport.insecure)

	// custom HTTP client
	setInsecureMode(fakeHTTPClient{})

	// check log messages
	lines := strings.Split(buf.String(), "\n")
	getLogMessage := func(s string) string {
		return strings.Join(strings.Split(s, " ")[3:], " ")
	}
	testhelpers.Equals(t, "client: cannot use insecure mode with Transport client of type scw.fakeTransport", getLogMessage(lines[0]))
	testhelpers.Equals(t, "client: cannot use insecure mode with HTTP client of type scw.fakeHTTPClient", getLogMessage(lines[1]))

	logger.DefaultLogger.Init(os.Stderr, logger.LogLevelWarning)
}

func TestNewVariableFromType(t *testing.T) {
	type fakeType struct {
		plop int
	}

	testhelpers.Equals(t, &fakeType{}, newVariableFromType(&fakeType{3}))
}

func TestExtractRequestLocality(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name             string
		path             string
		expectedLocality string
	}{
		{
			name:             "extract zone from path",
			path:             "/instance/v1/zones/fr-par-1/servers/11111111-1111-4111-8111-111111111111",
			expectedLocality: string(ZoneFrPar1),
		},
		{
			name:             "extract region from path",
			path:             "/k8s/v1/regions/fr-par/clusters/11111111-1111-4111-8111-111111111111",
			expectedLocality: string(RegionFrPar),
		},
		{
			name:             "prefer zone over region",
			path:             "/instance/v1/zones/fr-par-1/servers",
			expectedLocality: string(ZoneFrPar1),
		},
		{
			name:             "no locality in path",
			path:             "/instance/v1/servers/11111111-1111-4111-8111-111111111111",
			expectedLocality: "",
		},
		{
			name:             "unknown locality in path",
			path:             "/instance/v1/zones/xx-xxx-9/servers/11111111-1111-4111-8111-111111111111",
			expectedLocality: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			locality := extractRequestLocality(tc.path)
			testhelpers.Equals(t, tc.expectedLocality, locality)
		})
	}
}

func TestClientGetAPIMetadata(t *testing.T) {
	t.Run("APIMetadata", func(t *testing.T) {
		client, err := NewClient()
		testhelpers.AssertNoError(t, err)

		metadata, err := client.GetAPIMetadata()
		testhelpers.AssertNoError(t, err)
		testhelpers.Equals(t, "scw.eu", metadata.Domain)
		testhelpers.Equals(t, "scw", metadata.Partition)
		testhelpers.Equals(t, "external", metadata.Platform)

		// let's make sure the client cannot call home anymore
		// (in fact, if it tries to use httpClient, it will panic)
		client.httpClient = nil
		metadata, err = client.GetAPIMetadata()
		testhelpers.AssertNoError(t, err)
		testhelpers.Equals(t, "scw.eu", metadata.Domain)
		testhelpers.Equals(t, "scw", metadata.Partition)
		testhelpers.Equals(t, "external", metadata.Platform)
	})
}

package scw

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
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

		var tlsConfig *tls.Config
		switch t := client.httpClient.(*http.Client).Transport.(type) {
		case *http.Transport:
			tlsConfig = t.TLSClientConfig
		case *RateLimitTransport:
			tlsConfig = t.base().(*http.Transport).TLSClientConfig
		case *requestLoggerTransport:
			switch rt := t.rt.(type) {
			case *http.Transport:
				tlsConfig = rt.TLSClientConfig
			case *RateLimitTransport:
				tlsConfig = rt.base().(*http.Transport).TLSClientConfig
			}
		}

		testhelpers.Assert(t, tlsConfig != nil, "TLSClientConfig must be not nil")
		testhelpers.Equals(t, testInsecure, tlsConfig.InsecureSkipVerify)

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

// TestSetInsecureModeRateLimitTransport verifies that insecure mode actually
// disables certificate verification on the transport used by RoundTrip.
func TestSetInsecureModeRateLimitTransport(t *testing.T) {
	// Use a self-signed certificate
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Build the SDK HTTP client and enable insecure mode
	httpClient := newHTTPClient()
	setInsecureMode(httpClient)

	// A request must succeed despite the self-signed cert
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, server.URL, nil)
	testhelpers.AssertNoError(t, err)

	resp, err := httpClient.Do(req)
	testhelpers.AssertNoError(t, err)
	testhelpers.Equals(t, http.StatusOK, resp.StatusCode)
	_ = resp.Body.Close()
}

func TestNewVariableFromType(t *testing.T) {
	type fakeType struct {
		plop int
	}

	testhelpers.Equals(t, &fakeType{}, newVariableFromType(&fakeType{3}))
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

func TestRateLimit(t *testing.T) {
	cases := []struct {
		name           string
		clientMsg      []byte
		responseStatus string
		responseHeader http.Header
	}{
		{
			name:           "basic",
			clientMsg:      []byte(`{"code": 200, "headers": {"x-ratelimit-limit": "50, 50;w=1", "x-ratelimit-remaining": "49", "x-ratelimit-reset": "2"}}`),
			responseStatus: "200 OK",
			responseHeader: http.Header{
				"X-Ratelimit-Limit": {"50, 50;w=1"}, "X-Ratelimit-Remaining": {"49"}, "X-Ratelimit-Reset": {"2"},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			client := newHTTPClient()
			server := NewTestServer()

			req, err := http.NewRequestWithContext(
				context.Background(), http.MethodPost, server.URL, bytes.NewBuffer(c.clientMsg),
			)
			testhelpers.AssertNoError(t, err)

			req.Header.Set("Content-Type", "application/json; charset=UTF-8")

			resp, err := client.Do(req)
			testhelpers.AssertNoError(t, err)

			testhelpers.Equals(t, c.responseStatus, resp.Status)
			testhelpers.HeaderContains(t, c.responseHeader, resp.Header)
		})
	}
}

type ClientMessage struct {
	Code    int               `json:"code"`
	Headers map[string]string `json:"headers"`
}

// NewTestServer creates a dummy test server to try out the HTTP client
func NewTestServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var msg ClientMessage
		err := json.Unmarshal(body, &msg)
		if err != nil {
			panic(err)
		}

		for k, v := range msg.Headers {
			w.Header().Set(k, v)
		}

		w.WriteHeader(msg.Code)

		return
	}))
}

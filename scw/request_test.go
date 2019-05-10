package scw

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/auth"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

const (
	testBaseURL   = "http://example.com"
	testPath      = "/some/path/"
	testKey       = "some_key"
	testValue     = "some_value"
	testBody      = "some body"
	testUserAgent = "user/agent"

	testHeaderKey = "Some-Header-Key"
	testHeaderVal = "some_header_val"
	testTokenKey  = "some_secret_key"
)

func TestGetURL(t *testing.T) {

	req := ScalewayRequest{
		Path: testPath,
		Query: url.Values{
			testKey: []string{testValue},
		},
	}

	newURL, err := req.getURL(testBaseURL)
	testhelpers.Ok(t, err)

	expectedURL := fmt.Sprintf("%s%s?%s=%s", testBaseURL, testPath, testKey, testValue)

	testhelpers.Equals(t, expectedURL, newURL.String())

}

func TestGetHeadersWithoutBody(t *testing.T) {

	req := ScalewayRequest{
		Headers: http.Header{
			testHeaderKey: []string{testHeaderVal},
		},
	}
	token := auth.NewToken(testAccessKey, testTokenKey)

	expectedHeaders := http.Header{
		testHeaderKey:  []string{testHeaderVal},
		"x-auth-token": []string{testTokenKey},
		"User-Agent":   []string{testUserAgent},
	}

	allHeaders := req.getAllHeaders(token, testUserAgent)

	testhelpers.Equals(t, expectedHeaders, allHeaders)

}

func TestGetHeadersWithBody(t *testing.T) {
	req := ScalewayRequest{
		Headers: http.Header{
			testHeaderKey: []string{testHeaderVal},
		},
		Body: bytes.NewReader([]byte(testBody)),
	}
	token := auth.NewToken(testSecretKey, testTokenKey)

	expectedHeaders := http.Header{
		testHeaderKey:  []string{testHeaderVal},
		"x-auth-token": []string{testTokenKey},
		"Content-Type": []string{"application/json"},
		"User-Agent":   []string{testUserAgent},
	}

	allHeaders := req.getAllHeaders(token, testUserAgent)

	testhelpers.Equals(t, expectedHeaders, allHeaders)
}

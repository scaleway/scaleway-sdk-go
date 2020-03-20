package scw

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"

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
	testhelpers.AssertNoError(t, err)

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
		"X-Auth-Token": []string{testTokenKey},
		"User-Agent":   []string{testUserAgent},
	}

	allHeaders := req.getAllHeaders(token, testUserAgent, false)

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
		"X-Auth-Token": []string{testTokenKey},
		"Content-Type": []string{"application/json"},
		"User-Agent":   []string{testUserAgent},
	}

	allHeaders := req.getAllHeaders(token, testUserAgent, false)

	testhelpers.Equals(t, expectedHeaders, allHeaders)
}

func TestSetBody(t *testing.T) {
	body := struct {
		Region  Region         `json:"-"`
		ID      string         `json:"-"`
		Name    string         `json:"name,omitempty"`
		Slice   []string       `json:"slice,omitempty"`
		Flag    bool           `json:"flag,omitempty"`
		Timeout *time.Duration `json:"timeout,omitempty"`
	}{
		Region:  RegionNlAms,
		ID:      "plop",
		Name:    "plop",
		Slice:   []string{"plop", "plop"},
		Flag:    true,
		Timeout: TimeDurationPtr(time.Second),
	}

	req := ScalewayRequest{
		Headers: http.Header{},
	}

	testhelpers.AssertNoError(t, req.SetBody(body))

	r, isBytesReader := req.Body.(*bytes.Reader)

	testhelpers.Assert(t, isBytesReader, "req.Body should be bytes Reader")

	b := make([]byte, r.Len())
	_, err := r.Read(b)
	testhelpers.AssertNoError(t, err)

	testhelpers.Equals(t, []string{"application/json"}, req.Headers["Content-Type"])
	testhelpers.Equals(t, `{"name":"plop","slice":["plop","plop"],"flag":true,"timeout":1000000000}`, string(b))
}

func TestSetFileBody(t *testing.T) {
	body := &File{
		Content:     bytes.NewReader([]byte(testBody)),
		ContentType: "plain/text",
	}

	req := ScalewayRequest{
		Headers: http.Header{},
	}

	testhelpers.AssertNoError(t, req.SetBody(body))

	r, isBytesReader := req.Body.(*bytes.Reader)

	testhelpers.Assert(t, isBytesReader, "req.Body should be bytes Reader")

	b := make([]byte, r.Len())
	_, err := r.Read(b)
	testhelpers.AssertNoError(t, err)

	testhelpers.Equals(t, []string{"plain/text"}, req.Headers["Content-Type"])
	testhelpers.Equals(t, `some body`, string(b))
}

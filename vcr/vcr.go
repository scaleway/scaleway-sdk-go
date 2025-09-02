package vcr

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/strcase"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/cassette"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/recorder"
)

const (
	windowsDockerEngine     = "//./pipe/docker_engine"
	unixDockerEngine        = "/var/run/docker.sock"
	escapedUnixDockerEngine = "%2Fvar%2Frun%2Fdocker.sock"
)

// QueryMatcherIgnore contains the list of query value that should be ignored when matching requests with cassettes
var QueryMatcherIgnore = []string{
	"organization_id",
}

// getTestFilePath returns a valid filename path based on the go test name and suffix. (Take care of non fs friendly char)
func getTestFilePath(t *testing.T, pkgFolder string, suffix string) string {
	t.Helper()

	specialChars := regexp.MustCompile(`[\\?%*:|"<>. ]`)

	// Replace nested tests separators.
	fileName := strings.ReplaceAll(t.Name(), "/", "-")

	fileName = strcase.ToBashArg(fileName)

	// Replace special characters.
	fileName = specialChars.ReplaceAllLiteralString(fileName, "") + suffix

	// Remove prefix to simplify
	fileName = strings.TrimPrefix(fileName, "test-acc-")

	return filepath.Join(pkgFolder, "testdata", fileName)
}

func cassetteRequestFilter(i *cassette.Interaction) error {
	// Remove authorization headers
	delete(i.Request.Headers, "x-auth-token")
	delete(i.Request.Headers, "X-Auth-Token")
	delete(i.Request.Headers, "Authorization")

	// URL substitutions
	// 	Make URLs generic
	i.Request.URL = regexp.MustCompile(`(.+)organization_id=[0-9a-f-]{36}(.+)`).
		ReplaceAllString(i.Request.URL,
			"${1}organization_id=11111111-1111-1111-1111-111111111111${2}")
	i.Request.URL = regexp.MustCompile(`^https://api\.scaleway\.com/account/v1/tokens/[0-9a-f-]{36}$`).
		ReplaceAllString(
			i.Request.URL,
			"api.scaleway.com/account/v1/tokens/11111111-1111-1111-1111-111111111111")
	i.Request.URL = regexp.MustCompile(`(.+)?SCW[0-9A-Z]{17}(.+)?`).
		ReplaceAllString(
			i.Request.URL,
			"${1}SCWXXXXXXXXXXXXXXXXX${2}")
	//	Buildpacks
	i.Request.URL = regexp.MustCompile(`pack\.local%2Fbuilder%2F[0-9a-f]{20}`).
		ReplaceAllString(i.Request.URL, "pack.local%2Fbuilder%2F11111111111111111111")
	i.Request.URL = regexp.MustCompile(`pack\.local/builder/[0-9a-f]{20}`).
		ReplaceAllString(i.Request.URL, "pack.local/builder/11111111111111111111")

	// Body substitutions
	i.Request.Body = regexp.MustCompile(`pack\.local/builder/[0-9a-f]{20}`).
		ReplaceAllString(i.Request.Body, "pack.local/builder/11111111111111111111")

	return nil
}

func cassetteResponseFilter(i *cassette.Interaction) error {
	delete(i.Response.Headers, "Content-Security-Policy")
	delete(i.Response.Headers, "Strict-Transport-Security")
	delete(i.Response.Headers, "X-Content-Type-Options")
	delete(i.Response.Headers, "X-Frame-Options")

	i.Response.Body = regexp.MustCompile(`"secret_key":"[0-9a-f-]{36}"`).
		ReplaceAllString(i.Response.Body, `"secret_key":"11111111-1111-1111-1111-111111111111"`)

	i.Response.Body = regexp.MustCompile(`pack\.local/builder/[0-9a-f]{20}`).
		ReplaceAllString(i.Response.Body, "pack.local/builder/11111111111111111111")

	return nil
}

// This hook is needed for container deploy tests on the CLI
func unescapeDockerURL(i *cassette.Interaction) error {
	i.Request.URL = regexp.MustCompile(`http://`+escapedUnixDockerEngine+`(.+)?`).
		ReplaceAllString(
			i.Request.URL,
			"http://"+unixDockerEngine+"${1}")

	return nil
}

func s3Encoder(i *cassette.Interaction) error {
	if strings.HasSuffix(i.Request.Host, "scw.cloud") {
		if i.Request.Body != "" && i.Request.Headers.Get("Content-Type") == "application/octet-stream" {
			requestBody := []byte(i.Request.Body)
			if !json.Valid(requestBody) {
				err := xml.Unmarshal(requestBody, new(any))
				if err != nil {
					i.Request.Body = base64.StdEncoding.EncodeToString(requestBody)
				}
			}
		}

		if i.Response.Body != "" && i.Response.Headers.Get("Content-Type") == "binary/octet-stream" {
			responseBody := []byte(i.Response.Body)
			if !json.Valid(responseBody) {
				err := xml.Unmarshal(responseBody, new(any))
				if err != nil {
					i.Response.Body = base64.StdEncoding.EncodeToString(responseBody)
				}
			}
		}
	}

	return nil
}

// NewHTTPRecorder creates a new httpClient that records all HTTP requests in a cassette.
// This cassette is then replayed whenever tests are executed again. This means that once the
// requests are recorded in the cassette, no more real HTTP requests must be made to run the tests.
//
// It is important to add a `defer cleanup()` so the given cassette files are correctly
// closed and saved after the requests.
func NewHTTPRecorder(t *testing.T, pkgFolder string, update bool, realTransport ...http.RoundTripper) (*recorder.Recorder, error) {
	t.Helper()

	recorderOptions := []recorder.Option{
		// Ignore requests durations when replaying cassettes
		recorder.WithSkipRequestLatency(true),
		// Remove information that is either sensitive or that would disrupt matching from the request
		recorder.WithHook(cassetteRequestFilter, recorder.BeforeSaveHook),
		// Remove secrets and unnecessary information from the response
		recorder.WithHook(cassetteResponseFilter, recorder.BeforeSaveHook),
		// Starting with v3, go-vcr now calls net/url.Parse to build the interaction which results in an error for escaped
		// Docker URLs on Test_Deploy (container), so we need to unescape these paths after capture
		recorder.WithHook(unescapeDockerURL, recorder.AfterCaptureHook),
		// Serialization changed on v4, we now need to encode binary payloads as base64
		recorder.WithHook(s3Encoder, recorder.AfterCaptureHook),

		// Add custom matcher for requests and cassettes
		recorder.WithMatcher(CassetteMatcher),
	}

	// Set up recorder mode
	recorderMode := recorder.ModeReplayOnly
	if update {
		recorderMode = recorder.ModeRecordOnly
	}
	recorderOptions = append(recorderOptions, recorder.WithMode(recorderMode))

	// Set up custom transport for the CLI
	if len(realTransport) > 0 {
		recorderOptions = append(recorderOptions, recorder.WithRealTransport(realTransport[0]))
	}

	// Setup recorder and scw client
	cassetteFilePath := getTestFilePath(t, pkgFolder, ".cassette")
	r, err := recorder.New(cassetteFilePath, recorderOptions...)
	if err != nil {
		return nil, err
	}
	defer func(r *recorder.Recorder) {
		_ = r.Stop()
	}(r)

	return r, nil
}

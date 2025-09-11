package vcr

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"gopkg.in/dnaeon/go-vcr.v4/pkg/cassette"
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

func customDockerMatcher(r *http.Request, i cassette.Request) bool {
	escapedRecordedURL := regexp.MustCompile(`http://`+unixDockerEngine+`(.+)?`).
		ReplaceAllString(
			i.URL,
			"http://"+escapedUnixDockerEngine+"${1}")

	return r.URL.String() == escapedRecordedURL
}

func customS3BodyMatcher(r *http.Request, i cassette.Request) bool {
	reqBody, err := r.GetBody()
	if err != nil {
		panic(fmt.Errorf("cassette body matcher: failed to copy request body: %w", err))
	}

	requestBody, err := io.ReadAll(reqBody)
	if err != nil {
		panic(fmt.Errorf("cassette body matcher: failed to read actualRequest body: %w", err))
	}

	// Handle binary content
	if r.Header.Get("Content-Type") == "application/octet-stream" {
		requestEncodedBody := base64.StdEncoding.EncodeToString(requestBody)
		return requestEncodedBody == i.Body
	}

	// Handle XML
	err = xml.Unmarshal(requestBody, new(any))
	if err == nil {
		// We would need a specific struct for each type of xml to unmarshal so instead we just consider that it's a match
		return true
	}

	// Handle other data types (bucket policies)
	requestBucket, cassetteBucket := extractGenericBucketNames(r, i)
	if requestBucket != cassetteBucket {
		return false
	}
	bucketNameRegex := fmt.Sprintf("%s-[0-9]+", requestBucket)
	requestBodyGeneric := regexp.MustCompile(bucketNameRegex).ReplaceAllString(string(requestBody), requestBucket)
	cassetteBodyGeneric := regexp.MustCompile(bucketNameRegex).ReplaceAllString(i.Body, requestBucket)
	if requestBodyGeneric == cassetteBodyGeneric {
		return true
	}

	requestValues, err := url.ParseQuery(string(requestBody))
	if err != nil {
		panic(fmt.Errorf("cassette body matcher: failed to parse body as url values: %w", err)) // lintignore: R009
	}

	return compareFormBodies(requestValues, i.Form)
}

// cassetteBodyMatcher is a custom matcher that checks equivalence of request bodies
func cassetteBodyMatcher(request *http.Request, cassette cassette.Request) bool {
	if request.Body == nil || request.ContentLength == 0 {
		if cassette.Body == "" {
			return true // Body match if both are empty
		}

		if _, isFile := request.Body.(*os.File); isFile {
			return true // Body match if request is sending a file, maybe do more check here
		}

		return false
	}

	if strings.HasSuffix(request.URL.Host, "scw.cloud") {
		return customS3BodyMatcher(request, cassette)
	}

	r, err := request.GetBody()
	if err != nil {
		panic(fmt.Errorf("cassette body matcher: failed to copy request body: %w", err)) // lintignore: R009
	}

	requestBody, err := io.ReadAll(r)
	if err != nil {
		panic(fmt.Errorf("cassette body matcher: failed to read actualRequest body: %w", err)) // lintignore: R009
	}

	// Try to match raw bodies if they are not JSON (ex: cloud-init config)
	if string(requestBody) == cassette.Body {
		return true
	}

	requestJSON := make(map[string]any)
	cassetteJSON := make(map[string]any)

	// Match if content is xml (we would need a specific struct for each type of xml to unmarshal so instead we just consider that it's a match)
	err = xml.Unmarshal(requestBody, new(any))
	if err == nil {
		return true
	}

	if !json.Valid(requestBody) {
		requestValues, err := url.ParseQuery(string(requestBody))
		if err != nil {
			panic(fmt.Errorf("cassette body matcher: failed to parse body as url values: %w", err)) // lintignore: R009
		}

		// Remove keys that should be ignored during comparison
		for _, key := range BodyMatcherIgnore {
			requestValues.Del(key)
		}

		return compareFormBodies(requestValues, cassette.Form)
	}

	err = json.Unmarshal(requestBody, &requestJSON)
	if err != nil {
		panic(fmt.Errorf("cassette body matcher: failed to parse request body as json: %w", err)) // lintignore: R009
	}

	err = json.Unmarshal([]byte(cassette.Body), &cassetteJSON)
	if err != nil {
		// actualRequest contains JSON but cassette may not contain JSON, this doesn't match in this case
		return false
	}

	return compareJSONBodies(requestJSON, cassetteJSON, false)
}

// CassetteMatcher is a custom matcher that checks equivalence of a played request against a recorded one
// It compares method, path and query but will remove unwanted values from query
func CassetteMatcher(request *http.Request, cassette cassette.Request) bool {
	cassetteURL, _ := url.Parse(cassette.URL)
	requestURL := request.URL

	requestURLValues := requestURL.Query()
	cassetteURLValues := cassetteURL.Query()

	for _, query := range QueryMatcherIgnore {
		requestURLValues.Del(query)
		cassetteURLValues.Del(query)
	}

	requestURL.RawQuery = requestURLValues.Encode()
	cassetteURL.RawQuery = cassetteURLValues.Encode()

	// Make API key generic for URL comparison
	request.URL.Path = regexp.MustCompile(`(.+)?SCW[0-9A-Z]{17}(.+)?`).
		ReplaceAllString(
			request.URL.Path,
			"${1}SCWXXXXXXXXXXXXXXXXX${2}")

	// Specific handling of s3 URLs
	// URL format is https://test-acc-scaleway-object-bucket-lifecycle-8445817190507446251.s3.fr-par.scw.cloud/?lifecycle=
	if strings.HasSuffix(requestURL.Host, "scw.cloud") {
		if !strings.HasSuffix(cassetteURL.Host, "scw.cloud") {
			return false
		}

		requestBucket, cassetteBucket := extractGenericBucketNames(request, cassette)
		if requestBucket != cassetteBucket {
			return false
		}
	}

	// Specific handling of Docker URLs
	// URL needs to be escaped before comparison
	if request.URL.Host == unixDockerEngine || request.URL.Host == windowsDockerEngine {
		return customDockerMatcher(request, cassette)
	}

	return request.Method == cassette.Method &&
		request.URL.Path == cassetteURL.Path &&
		requestURL.RawQuery == cassetteURL.RawQuery &&
		cassetteBodyMatcher(request, cassette)
}

package scw

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

func TestHasResponseErrorWithStatus200(t *testing.T) {

	res := &http.Response{StatusCode: 200}

	newErr := hasResponseError(res)
	testhelpers.AssertNoError(t, newErr)

}

func TestHasResponseErrorWithoutBody(t *testing.T) {

	res := &http.Response{StatusCode: 400}

	newErr := hasResponseError(res)
	testhelpers.Assert(t, newErr != nil, "Should have error")

}

func TestNonStandardError(t *testing.T) {
	// Create expected error response
	testErrorReponse := &InvalidArgumentsError{
		Details: []InvalidArgumentsErrorDetail{
			{
				ArgumentName: "volumes.5.id",
				Reason:       "constraint",
				HelpMessage:  "92 is not a valid UUID.",
			},
			{
				ArgumentName: "volumes.5.name",
				Reason:       "constraint",
				HelpMessage:  "required key not provided",
			},
		},
		RawBody: []byte(`{"fields":{"volumes.5.id":["92 is not a valid UUID."],"volumes.5.name":["required key not provided"]},"message":"Validation Error","type":"invalid_request_error"}`),
	}

	// Create response body with marshalled error response
	body := `{"fields":{"volumes.5.id":["92 is not a valid UUID."],"volumes.5.name":["required key not provided"]},"message":"Validation Error","type":"invalid_request_error"}`
	res := &http.Response{Status: "400 Bad Request", StatusCode: 400, Body: ioutil.NopCloser(strings.NewReader(body))}

	// Test hasResponseError convert the response into a InvalidArgumentsError error
	newErr := hasResponseError(res)
	testhelpers.Assert(t, newErr != nil, "Should have error")
	testhelpers.Equals(t, testErrorReponse, newErr)
}

func TestHasResponseErrorWithValidError(t *testing.T) {
	var (
		errorMessage    = "some message"
		errorType       = "some type"
		errorFields     = map[string][]string{"some_field": {"some_value"}}
		errorStatusCode = 400
		errorStatus     = "400 Bad Request"
	)

	// Create expected error response
	testErrorReponse := &ResponseError{
		Message:    errorMessage,
		Type:       errorType,
		Fields:     errorFields,
		StatusCode: errorStatusCode,
		Status:     errorStatus,
		RawBody:    []byte(`{"message":"some message","type":"some type","fields":{"some_field":["some_value"]}}`),
	}

	// Create response body with marshalled error response
	bodyBytes, err := json.Marshal(testErrorReponse)
	testhelpers.AssertNoError(t, err)
	res := &http.Response{Status: errorStatus, StatusCode: errorStatusCode, Body: ioutil.NopCloser(bytes.NewReader(bodyBytes))}

	// Test hasResponseError()
	newErr := hasResponseError(res)
	testhelpers.Assert(t, newErr != nil, "Should have error")
	testhelpers.Equals(t, testErrorReponse, newErr)

}

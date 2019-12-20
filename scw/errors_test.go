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

	type testCase struct {
		resStatus     string
		resStatusCode int
		resBody       string
		expectedError SdkError
	}

	run := func(c *testCase) func(t *testing.T) {
		return func(t *testing.T) {
			res := &http.Response{
				Status:     c.resStatus,
				StatusCode: c.resStatusCode,
				Body:       ioutil.NopCloser(strings.NewReader(c.resBody)),
			}

			// Test that hasResponseError converts the response to the expected SdkError.
			newErr := hasResponseError(res)
			testhelpers.Assert(t, newErr != nil, "Should have error")
			testhelpers.Equals(t, c.expectedError, newErr)
		}
	}

	t.Run("invalid_request_error type with fields", run(&testCase{
		resStatus:     "400 Bad Request",
		resStatusCode: http.StatusBadRequest,
		resBody:       `{"fields":{"volumes.5.id":["92 is not a valid UUID."],"volumes.5.name":["required key not provided"]},"message":"Validation Error","type":"invalid_request_error"}`,
		expectedError: &InvalidArgumentsError{
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
		},
	}))

	t.Run("invalid_request_error type with message", run(&testCase{
		resStatus:     "400 Bad Request",
		resStatusCode: http.StatusBadRequest,
		resBody:       `{"message": "server should be running", "type": "invalid_request_error"}`,
		expectedError: &ResponseError{
			Message: "server should be running",
			Type:    "invalid_request_error",
			RawBody: []byte(`{"message": "server should be running", "type": "invalid_request_error"}`),
		},
	}))

	t.Run("conflict type", run(&testCase{
		resStatus:     "409 Conflict",
		resStatusCode: http.StatusConflict,
		resBody:       `{"message": "Group is in use. You cannot delete it.", "type": "conflict"}`,
		expectedError: &ResponseError{
			Message: "group is in use. you cannot delete it.",
			Type:    "conflict",
			RawBody: []byte(`{"message": "Group is in use. You cannot delete it.", "type": "conflict"}`),
		},
	}))
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

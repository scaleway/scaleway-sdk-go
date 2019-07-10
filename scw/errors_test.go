package scw

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
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

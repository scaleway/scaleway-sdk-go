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
	testhelpers.Ok(t, newErr)

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
	)

	// Create expected error response
	testErrorReponse := &ResponseError{
		Message:    errorMessage,
		Type:       errorType,
		Fields:     errorFields,
		StatusCode: errorStatusCode,
	}

	// Create response body with marshalled error response
	bodyBytes, err := json.Marshal(testErrorReponse)
	testhelpers.Ok(t, err)
	res := &http.Response{StatusCode: errorStatusCode, Body: ioutil.NopCloser(bytes.NewReader(bodyBytes))}

	// Test hasResponseError()
	newErr := hasResponseError(res)
	testhelpers.Assert(t, newErr != nil, "Should have error")
	testhelpers.Equals(t, testErrorReponse, newErr)

}

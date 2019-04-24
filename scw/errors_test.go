package scw

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

func TestHasErrorResponseWithStatus200(t *testing.T) {

	res := &http.Response{StatusCode: 200}

	newErr := hasErrorResponse(res)
	testhelpers.Ok(t, newErr)

}

func TestHasErrorResponseWithoutBody(t *testing.T) {

	res := &http.Response{StatusCode: 400}

	newErr := hasErrorResponse(res)
	testhelpers.Assert(t, newErr != nil, "Should have error")

}

func TestHasErrorResponseWithValidError(t *testing.T) {

	var (
		errorMessage    = "some message"
		errorType       = "some type"
		errorFields     = map[string][]string{"some_field": []string{"some_value"}}
		errorStatusCode = 400
	)

	// Create expected error response
	testErrorReponse := &errorResponse{
		Message:    errorMessage,
		Type:       errorType,
		Fields:     errorFields,
		statusCode: errorStatusCode,
	}

	// Create response body with marshalled error response
	bodyBytes, err := json.Marshal(testErrorReponse)
	testhelpers.Ok(t, err)
	res := &http.Response{StatusCode: errorStatusCode, Body: ioutil.NopCloser(bytes.NewReader(bodyBytes))}

	// Test hasErrorResponse()
	newErr := hasErrorResponse(res)
	testhelpers.Assert(t, newErr != nil, "Should have error")
	testhelpers.Equals(t, testErrorReponse, newErr)

}

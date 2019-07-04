package scw

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/scaleway/scaleway-sdk-go/internal/errors"
)

// SdkError is a base interface for all Scaleway SDK errors.
type SdkError interface {
	Error() string
	IsScwSdkError()
}

// ResponseError is an error type for the Scaleway API
type ResponseError struct {
	// Message is a human-friendly error message
	Message string `json:"message"`

	// Type is a string code that defines the kind of error
	Type string `json:"type,omitempty"`

	// Resource is a string code that defines the resource concerned by the error
	Resource string `json:"resource,omitempty"`

	// Fields contains detail about validation error
	Fields map[string][]string `json:"fields,omitempty"`

	// StatusCode is the HTTP status code received
	StatusCode int `json:"-"`

	// Status is the HTTP status received
	Status string `json:"-"`
}

func (e *ResponseError) Error() string {
	s := fmt.Sprintf("scaleway-sdk-go: http error %s", e.Status)

	if e.Resource != "" {
		s = fmt.Sprintf("%s: resource %s", s, e.Resource)
	}

	if e.Message != "" {
		s = fmt.Sprintf("%s: %s", s, e.Message)
	}

	if len(e.Fields) > 0 {
		s = fmt.Sprintf("%s: %v", s, e.Fields)
	}

	return s
}

// IsScwSdkError implement SdkError interface
func (e *ResponseError) IsScwSdkError() {}

func hasResponseError(res *http.Response) SdkError {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}

	newErr := &ResponseError{
		StatusCode: res.StatusCode,
		Status:     res.Status,
	}

	if res.Body == nil {
		return newErr
	}

	err := json.NewDecoder(res.Body).Decode(newErr)
	if err != nil {
		return errors.Wrap(err, "could not parse error response body")
	}

	return newErr
}

package scw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

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

	// Type is a string code that defines the kind of error. This field is only used by instance API
	Type string `json:"type,omitempty"`

	// Resource is a string code that defines the resource concerned by the error. This field is only used by instance API
	Resource string `json:"resource,omitempty"`

	// Fields contains detail about validation error. This field is only used by instance API
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

// hasResponseError throws an error when the HTTP status is not OK
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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.Wrap(err, "cannot read error response body")
	}

	err = json.Unmarshal(body, newErr)
	if err != nil {
		return errors.Wrap(err, "could not parse error response body")
	}
	stdErr := unmarshalStandardError(newErr.Type, body)
	if stdErr != nil {
		return stdErr
	}

	return newErr
}

func unmarshalStandardError(errorType string, body []byte) SdkError {
	var stdErr SdkError

	switch errorType {
	case "invalid_arguments":
		stdErr = &InvalidArgumentsError{}
	default:
		return nil
	}

	err := json.Unmarshal(body, stdErr)
	if err != nil {
		return errors.Wrap(err, "could not parse error %s response body", errorType)
	}

	return stdErr
}

type InvalidArgumentsError struct {
	Details []struct {
		ArgumentName string `json:"argument_name"`
		Reason       string `json:"reason"`
		HelpMessage  string `json:"help_message"`
	} `json:"details"`
}

// IsScwSdkError implements the SdkError interface
func (e *InvalidArgumentsError) IsScwSdkError() {}
func (e *InvalidArgumentsError) Error() string {
	invalidArgs := make([]string, len(e.Details))
	for i, d := range e.Details {
		invalidArgs[i] = d.ArgumentName
		switch d.Reason {
		case "unknown":
			invalidArgs[i] += " is invalid for unexpected reason"
		case "required":
			invalidArgs[i] += " is required"
		case "format":
			invalidArgs[i] += " is wrongly formatted"
		case "constraint":
			invalidArgs[i] += " does not respect constraint"
		}
		if d.HelpMessage != "" {
			invalidArgs[i] += ", " + d.HelpMessage
		}
	}

	return "scaleway-sdk-go: invalid argument(s): " + strings.Join(invalidArgs, "; ")
}

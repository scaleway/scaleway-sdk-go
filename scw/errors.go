package scw

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errorResponse struct {
	// Message is a human-friendly error message
	Message string `json:"message"`

	// Type is a string code that defines the kind of error
	Type string `json:"type,omitempty"`

	// Fields contains detail about validation error
	Fields map[string][]string `json:"fields,omitempty"`

	// statusCode is the HTTP status code received
	statusCode int `json:"-"`
}

func hasErrorResponse(res *http.Response) error {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}
	newErr := &errorResponse{
		statusCode: res.StatusCode,
	}

	if res.Body == nil {
		return newErr
	}

	decoder := json.NewDecoder(res.Body)
	defer res.Body.Close()
	err := decoder.Decode(newErr)
	if err != nil {
		return err
	}

	return newErr
}

func (e *errorResponse) Error() string {
	s := fmt.Sprintf("received status code %d", e.statusCode)

	if e.Type != "" {
		s = fmt.Sprintf("%s: error type is %s", s, e.Type)
	}

	if e.Message != "" {
		s = fmt.Sprintf("%s: %s", s, e.Message)
	}

	if len(e.Fields) > 0 {
		s = fmt.Sprintf("%s: details: %v", s, e.Fields)
	}
	return s
}

package auth

import "net/http"

// Auth implement methods required for authentication.
// Valid authentication are currently a token or no auth.
type Auth interface {
	// Headers returns headers that must be add to the http request
	Headers() http.Header
}

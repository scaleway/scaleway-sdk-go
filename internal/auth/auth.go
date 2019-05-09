package auth

import "net/http"

// Auth implement methods required for authentication.
// Valid authentication are currently a token or no auth.
type Auth interface {
	// Metadata returns non-nil map
	Headers() http.Header
}

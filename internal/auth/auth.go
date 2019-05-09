package auth

import "net/http"

// Auth implement methods required for authentication.
// It can be:
// secret key, no auth
type Auth interface {
	// Metadata returns non-nil map
	Headers() http.Header
}

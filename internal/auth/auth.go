package auth

import "net/http"

// secret key
// no auth
type Auth interface {

	// Metadata returns non-nil map
	Headers() http.Header
}

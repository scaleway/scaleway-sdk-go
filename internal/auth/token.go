package auth

import "net/http"

type token struct {
	accessKey string
	secretKey string
}

// XAuthTokenHeader is Scaleway standard auth header
const XAuthTokenHeader = "x-auth-token"

// NewToken create a token authentication from an
// access key and a secret key
func NewToken(accessKey, secretKey string) *token {
	return &token{accessKey: accessKey, secretKey: secretKey}
}

// Headers returns headers that must be add to the http request
func (t *token) Headers() http.Header {
	headers := http.Header{}
	headers.Set(XAuthTokenHeader, t.secretKey)
	return headers
}

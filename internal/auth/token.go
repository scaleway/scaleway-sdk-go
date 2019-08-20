package auth

import "net/http"

// Token is the pair accessKey + secretKey.
// This type is public because it's an internal package.
type Token struct {
	AccessKey string
	SecretKey string
}

// XAuthTokenHeader is Scaleway standard auth header
const XAuthTokenHeader = "X-Auth-Token"

// NewToken create a token authentication from an
// access key and a secret key
func NewToken(accessKey, secretKey string) *Token {
	return &Token{AccessKey: accessKey, SecretKey: secretKey}
}

// Headers returns headers that must be add to the http request
func (t *Token) Headers() http.Header {
	headers := http.Header{}
	headers.Set(XAuthTokenHeader, t.SecretKey)
	return headers
}

// AnonymizedHeaders returns an anonymized version of Headers()
// This method could be use for logging purpose.
func (t *Token) AnonymizedHeaders() http.Header {
	headers := http.Header{}
	var secret string

	switch {
	case len(t.SecretKey) == 0:
		secret = ""
	case len(t.SecretKey) > 8:
		secret = t.SecretKey[0:8] + "-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	default:
		secret = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	}
	headers.Set(XAuthTokenHeader, secret)
	return headers
}

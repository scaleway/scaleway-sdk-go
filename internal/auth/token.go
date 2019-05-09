package auth

import "net/http"

type token struct {
	accessKey string
	secretKey string
}

// NewToken create a token authentication from an
// access key and a secret key
func NewToken(accessKey, secretKey string) *token {
	return &token{accessKey: accessKey, secretKey: secretKey}
}

func (t *token) Headers() http.Header {
	return http.Header{
		"x-auth-token": []string{t.secretKey},
	}
}

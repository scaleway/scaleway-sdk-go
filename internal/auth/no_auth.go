package auth

import "net/http"

type noAuth struct {
}

func NewNoAuth() *noAuth {
	return &noAuth{}
}

func (t *noAuth) Headers() http.Header {
	return http.Header{}
}

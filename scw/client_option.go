package scw

import (
	"net/http"

	"github.com/scaleway/scaleway-sdk-go/internal/auth"
)

// ClientOption is a function which applies options to a settings object.
type ClientOption func(*settings)

// WithoutAuth option sets the client token to an empty token.
func WithoutAuth() ClientOption {
	return func(s *settings) {
		s.Token = auth.NewNoAuth()
	}
}

// WithAuth option sets the client access key and secret key.
func WithAuth(accessKey, secretKey string) ClientOption {
	return func(s *settings) {
		s.Token = auth.NewToken(accessKey, secretKey)
	}
}

// WithEndpoint option overrides the endpoint URL of the Scaleway API to the given URL.
func WithEndpoint(url string) ClientOption {
	return func(s *settings) {
		s.Url = url
	}
}

// WithInsecure option enables insecure transport on the client.
func WithInsecure() ClientOption {
	return func(s *settings) {
		s.Insecure = true
	}
}

// WithUserAgent option overrides the default user agent of the SDK.
func WithUserAgent(ua string) ClientOption {
	return func(s *settings) {
		s.UserAgent = ua
	}
}

// WithHttpClient options allows passing a custom http.Client which will be used for all requests.
func WithHttpClient(httpClient *http.Client) ClientOption {
	return func(s *settings) {
		s.HttpClient = httpClient
	}
}

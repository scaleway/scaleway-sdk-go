package scw

import (
	"net/http"

	"github.com/scaleway/scaleway-sdk-go/internal/auth"
)

type ClientOption func(*settings)

func WithoutAuth() ClientOption {
	return func(s *settings) {
		s.Token = auth.NewNoAuth()
	}
}

func WithAuth(accessKey, secretKey string) ClientOption {
	return func(s *settings) {
		s.Token = auth.NewToken(accessKey, secretKey)
	}
}

func WithEndpoint(url string) ClientOption {
	return func(s *settings) {
		s.Url = url
	}
}

func WithInsecure() ClientOption {
	return func(s *settings) {
		s.Insecure = true
	}
}

func WithUserAgent(ua string) ClientOption {
	return func(s *settings) {
		s.UserAgent = ua
	}
}

func WithHttpClient(httpClient *http.Client) ClientOption {
	return func(s *settings) {
		s.HttpClient = httpClient
	}
}

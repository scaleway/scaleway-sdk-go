package scw

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/auth"
)

type Client struct {
	httpClient *http.Client
	auth       auth.Auth
	baseUrl    string
	userAgent  string
}

func defaultOptions() []ClientOption {
	return []ClientOption{
		WithEndpoint("https://api.scaleway.com"),
		WithUserAgent(userAgent),
	}
}

func NewClient(opts ...ClientOption) (*Client, error) {
	s := newSettings()

	// apply options
	s.apply(append(defaultOptions(), opts...))

	// validate settings
	err := s.validate()
	if err != nil {
		return nil, err
	}

	// dial the API
	if s.HttpClient == nil {
		s.HttpClient = newHttpClient()
	}

	// insecure mode
	if s.Insecure {
		clientTransport, ok := s.HttpClient.Transport.(*http.Transport)
		if !ok {
			return nil, fmt.Errorf("cannot use insecure mode with HTTP client of type %T", s.HttpClient.Transport)
		}
		if clientTransport.TLSClientConfig == nil {
			clientTransport.TLSClientConfig = &tls.Config{}
		}
		clientTransport.TLSClientConfig.InsecureSkipVerify = true
	}

	return &Client{
		auth:       s.Token,
		httpClient: s.HttpClient,
		baseUrl:    s.Url,
		userAgent:  s.UserAgent,
	}, nil
}

func newHttpClient() *http.Client {
	return &http.Client{
		Timeout: 9 * time.Second,
		Transport: &http.Transport{
			DialContext:           (&net.Dialer{Timeout: 5 * time.Second}).DialContext,
			TLSHandshakeTimeout:   5 * time.Second,
			ResponseHeaderTimeout: 8 * time.Second,
			MaxIdleConnsPerHost:   20,
		},
	}
}

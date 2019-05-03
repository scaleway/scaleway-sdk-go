package scw

import (
	"net/http"

	"github.com/scaleway/scaleway-sdk-go/internal/auth"
)

// ClientOption is a function which applies options to a settings object.
type ClientOption func(*settings)

// WithoutAuth client option sets the client token to an empty token.
func WithoutAuth() ClientOption {
	return func(s *settings) {
		s.Token = auth.NewNoAuth()
	}
}

// WithAuth client option sets the client access key and secret key.
func WithAuth(accessKey, secretKey string) ClientOption {
	return func(s *settings) {
		s.Token = auth.NewToken(accessKey, secretKey)
	}
}

// WithApiUrl client option overrides the API URL of the Scaleway API to the given URL.
func WithApiUrl(apiUrl string) ClientOption {
	return func(s *settings) {
		s.ApiUrl = apiUrl
	}
}

// WithInsecure client option enables insecure transport on the client.
func WithInsecure() ClientOption {
	return func(s *settings) {
		s.Insecure = true
	}
}

// WithUserAgent client option overrides the default user agent of the SDK.
func WithUserAgent(ua string) ClientOption {
	return func(s *settings) {
		s.UserAgent = ua
	}
}

// WithHttpClient client option allows passing a custom http.Client which will be used for all requests.
func WithHttpClient(httpClient *http.Client) ClientOption {
	return func(s *settings) {
		s.HttpClient = httpClient
	}
}

// WithDefaultOrganizationId client option sets the client default organization ID.
//
// It will be used as the default value of the organization_id field in all requests made with this client.
func WithDefaultOrganizationId(organizationId string) ClientOption {
	return func(s *settings) {
		s.DefaultOrganizationId = organizationId
	}
}

// WithDefaultRegion client option sets the client default region.
//
// It will be used as the default value of the region field in all requests made with this client.
func WithDefaultRegion(region Region) ClientOption {
	return func(s *settings) {
		s.DefaultRegion = region
	}
}

// WithDefaultZone client option sets the client default zone.
//
// It will be used as the default value of the zone field in all requests made with this client.
func WithDefaultZone(zone Zone) ClientOption {
	return func(s *settings) {
		s.DefaultZone = zone
	}
}

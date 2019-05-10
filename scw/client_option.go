package scw

import (
	"net/http"

	"github.com/scaleway/scaleway-sdk-go/internal/auth"
	"github.com/scaleway/scaleway-sdk-go/scwconfig"
	"github.com/scaleway/scaleway-sdk-go/utils"
)

// ClientOption is a function which applies options to a settings object.
type ClientOption func(*settings)

// WithoutAuth client option sets the client token to an empty token.
func WithoutAuth() ClientOption {
	return func(s *settings) {
		s.token = auth.NewNoAuth()
	}
}

// WithAuth client option sets the client access key and secret key.
func WithAuth(accessKey, secretKey string) ClientOption {
	return func(s *settings) {
		s.token = auth.NewToken(accessKey, secretKey)
	}
}

// WithAPIURL client option overrides the API URL of the Scaleway API to the given URL.
func WithAPIURL(apiURL string) ClientOption {
	return func(s *settings) {
		s.apiURL = apiURL
	}
}

// WithInsecure client option enables insecure transport on the client.
func WithInsecure() ClientOption {
	return func(s *settings) {
		s.insecure = true
	}
}

// WithUserAgent client option overrides the default user agent of the SDK.
func WithUserAgent(ua string) ClientOption {
	return func(s *settings) {
		s.userAgent = ua
	}
}

// WithHTTPClient client option allows passing a custom http.Client which will be used for all requests.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(s *settings) {
		s.httpClient = httpClient
	}
}

// WithConfig client option configure a client with Scaleway configuration.
func WithConfig(config scwconfig.Config) ClientOption {
	return func(s *settings) {
		// The access key is not used for API authentications.
		accessKey, _ := config.GetAccessKey()
		secretKey, secretKeyExist := config.GetSecretKey()
		if secretKeyExist {
			s.token = auth.NewToken(accessKey, secretKey)
		}

		apiURL, exist := config.GetAPIURL()
		if exist {
			s.apiURL = apiURL
		}

		insecure, exist := config.GetInsecure()
		if exist {
			s.insecure = insecure
		}

		defaultOrganizationID, exist := config.GetDefaultOrganizationID()
		if exist {
			s.defaultOrganizationID = defaultOrganizationID
		}

		defaultRegion, exist := config.GetDefaultRegion()
		if exist {
			s.defaultRegion = defaultRegion
		}

		defaultZone, exist := config.GetDefaultZone()
		if exist {
			s.defaultZone = defaultZone
		}
	}
}

// WithDefaultOrganizationID client option sets the client default organization ID.
//
// It will be used as the default value of the organization_id field in all requests made with this client.
func WithDefaultOrganizationID(organizationID string) ClientOption {
	return func(s *settings) {
		s.defaultOrganizationID = organizationID
	}
}

// WithDefaultRegion client option sets the client default region.
//
// It will be used as the default value of the region field in all requests made with this client.
func WithDefaultRegion(region utils.Region) ClientOption {
	return func(s *settings) {
		s.defaultRegion = region
	}
}

// WithDefaultZone client option sets the client default zone.
//
// It will be used as the default value of the zone field in all requests made with this client.
func WithDefaultZone(zone utils.Zone) ClientOption {
	return func(s *settings) {
		s.defaultZone = zone
	}
}

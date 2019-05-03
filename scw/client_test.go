package scw

import (
	"net/http"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/auth"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

const (
	testApiUrl                = "https://api.example.com/"
	defaultApiUrl             = "https://api.scaleway.com"
	testAccessKey             = "some access key"
	testSecretKey             = "some secret key"
	testDefaultOrganizationId = "some default organization id"
	testDefaultRegion         = RegionFrPar
	testDefaultZone           = ZoneFrPar1
)

func TestNewClientWithDefaults(t *testing.T) {

	options := []ClientOption{
		WithoutAuth(),
		WithInsecure(),
	}

	client, err := NewClient(options...)
	testhelpers.Ok(t, err)

	testhelpers.Equals(t, defaultApiUrl, client.apiUrl)
	testhelpers.Equals(t, auth.NewNoAuth(), client.auth)

}

func TestNewClientWithOptions(t *testing.T) {

	someHTTPClient := &http.Client{}

	options := []ClientOption{
		WithApiUrl(testApiUrl),
		WithAuth(testAccessKey, testSecretKey),
		WithHttpClient(someHTTPClient),
		WithDefaultOrganizationId(testDefaultOrganizationId),
		WithDefaultRegion(testDefaultRegion),
		WithDefaultZone(testDefaultZone),
	}

	client, err := NewClient(options...)
	testhelpers.Ok(t, err)

	testhelpers.Equals(t, testApiUrl, client.apiUrl)
	testhelpers.Equals(t, auth.NewToken(testAccessKey, testSecretKey), client.auth)

	testhelpers.Equals(t, someHTTPClient, client.httpClient)

	testhelpers.Equals(t, testDefaultOrganizationId, client.defaultOrganizationId)
	testhelpers.Equals(t, testDefaultRegion, client.defaultRegion)
	testhelpers.Equals(t, testDefaultZone, client.defaultZone)

}

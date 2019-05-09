package scw

import (
	"net/http"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/auth"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

const (
	testAPIURL                = "https://api.example.com/"
	defaultAPIURL             = "https://api.scaleway.com"
	testAccessKey             = "some access key"
	testSecretKey             = "some secret key"
	testDefaultOrganizationID = "some default organization id"
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

	testhelpers.Equals(t, defaultAPIURL, client.apiURL)
	testhelpers.Equals(t, auth.NewNoAuth(), client.auth)

}

func TestNewClientWithOptions(t *testing.T) {

	someHTTPClient := &http.Client{}

	options := []ClientOption{
		WithAPIURL(testAPIURL),
		WithAuth(testAccessKey, testSecretKey),
		WithHTTPClient(someHTTPClient),
		WithDefaultOrganizationID(testDefaultOrganizationID),
		WithDefaultRegion(testDefaultRegion),
		WithDefaultZone(testDefaultZone),
	}

	client, err := NewClient(options...)
	testhelpers.Ok(t, err)

	testhelpers.Equals(t, testAPIURL, client.apiURL)
	testhelpers.Equals(t, auth.NewToken(testAccessKey, testSecretKey), client.auth)

	testhelpers.Equals(t, someHTTPClient, client.httpClient)

	testhelpers.Equals(t, testDefaultOrganizationID, client.GetDefaultOrganizationID())
	testhelpers.Equals(t, testDefaultRegion, client.GetDefaultRegion())
	testhelpers.Equals(t, testDefaultZone, client.GetDefaultZone())

}

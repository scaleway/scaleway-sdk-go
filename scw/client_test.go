package scw

import (
	"net/http"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/auth"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/utils"
)

const (
	testAPIURL                = "https://api.example.com/"
	defaultAPIURL             = "https://api.scaleway.com"
	testAccessKey             = "ACCESS_KEY"
	testSecretKey             = "7363616C-6577-6573-6862-6F7579616161" // | xxd -ps -r
	testDefaultOrganizationID = "6170692E-7363-616C-6577-61792E636F6D" // | xxd -ps -r
	testDefaultRegion         = utils.RegionFrPar
	testDefaultZone           = utils.ZoneFrPar1
	testInsecure              = true
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

type TestConfig struct{}

func (c *TestConfig) GetAccessKey() (string, bool) {
	return testAccessKey, true
}
func (c *TestConfig) GetSecretKey() (string, bool) {
	return testSecretKey, true
}
func (c *TestConfig) GetAPIURL() (string, bool) {
	return testAPIURL, true
}
func (c *TestConfig) GetInsecure() (bool, bool) {
	return testInsecure, true
}
func (c *TestConfig) GetDefaultOrganizationID() (string, bool) {
	return testDefaultOrganizationID, true
}
func (c *TestConfig) GetDefaultRegion() (utils.Region, bool) {
	return testDefaultRegion, true
}
func (c *TestConfig) GetDefaultZone() (utils.Zone, bool) {
	return testDefaultZone, true
}

func TestNewClientWithOptions(t *testing.T) {

	t.Run("Basic", func(t *testing.T) {
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
	})

	t.Run("With scwconfig", func(t *testing.T) {
		config := &TestConfig{}

		client, err := NewClient(WithConfig(config))
		testhelpers.Ok(t, err)

		testhelpers.Equals(t, auth.NewToken(testAccessKey, testSecretKey), client.auth)
		testhelpers.Equals(t, testAPIURL, client.apiURL)

		clientTransport, ok := client.httpClient.Transport.(*http.Transport)
		testhelpers.Assert(t, ok, "clientTransport must be not nil")
		testhelpers.Assert(t, clientTransport.TLSClientConfig != nil, "TLSClientConfig must be not nil")
		testhelpers.Equals(t, testInsecure, clientTransport.TLSClientConfig.InsecureSkipVerify)

		testhelpers.Equals(t, testDefaultOrganizationID, client.GetDefaultOrganizationID())
		testhelpers.Equals(t, testDefaultRegion, client.GetDefaultRegion())
		testhelpers.Equals(t, testDefaultZone, client.GetDefaultZone())
	})

}

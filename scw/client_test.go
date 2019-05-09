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
	testAccessKey             = "ACCESS_KEY"
	testSecretKey             = "SECRET_KEY"
	testDefaultOrganizationID = "d45a075f-18a1-4e9f-824e-43914a3ae8bd"
	testDefaultRegion         = RegionFrPar
	testDefaultZone           = ZoneFrPar1
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
func (c *TestConfig) GetSecretKey() (secretKey string, exist bool) {
	return testSecretKey, true
}
func (c *TestConfig) GetAPIURL() (apiURL string, exist bool) {
	return testAPIURL, true
}
func (c *TestConfig) GetInsecure() (insecure bool, exist bool) {
	return testInsecure, true
}
func (c *TestConfig) GetDefaultOrganizationID() (defaultOrganizationID string, exist bool) {
	return testDefaultOrganizationID, true
}
func (c *TestConfig) GetDefaultRegion() (defaultRegion string, exist bool) {
	return string(testDefaultRegion), true
}
func (c *TestConfig) GetDefaultZone() (defaultZone string, exist bool) {
	return string(testDefaultZone), true
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
		testhelpers.Equals(t, true, ok)
		testhelpers.NotNil(t, clientTransport.TLSClientConfig)
		testhelpers.Equals(t, testInsecure, clientTransport.TLSClientConfig.InsecureSkipVerify)

		testhelpers.Equals(t, testDefaultOrganizationID, client.GetDefaultOrganizationID())
		testhelpers.Equals(t, testDefaultRegion, client.GetDefaultRegion())
		testhelpers.Equals(t, testDefaultZone, client.GetDefaultZone())
	})

}

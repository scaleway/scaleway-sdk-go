package instance

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// Regression test for #2403.
func TestGetServerResourceNotFoundIncludesLocality(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"type":"not_found","resource":"server","resource_id":"11111111-1111-4111-8111-111111111111"}`))
	}))
	defer server.Close()

	client, err := scw.NewClient(
		scw.WithDefaultZone(scw.ZoneFrPar1),
		scw.WithAPIURL(server.URL),
		scw.WithAuth("SCWXXXXXXXXXXXXXXXXX", "00000000-0000-0000-0000-000000000000"),
	)
	testhelpers.AssertNoError(t, err)

	api := NewAPI(client)

	_, err = api.GetServer(&GetServerRequest{
		Zone:     scw.ZoneFrPar1,
		ServerID: "11111111-1111-4111-8111-111111111111",
	})
	testhelpers.Assert(t, err != nil, "This request should error")

	notFoundErr, ok := err.(*scw.ResourceNotFoundError)
	testhelpers.Assert(t, ok, "Should return a ResourceNotFoundError, got: %T (%v)", err, err)
	testhelpers.Equals(t, string(scw.ZoneFrPar1), notFoundErr.Locality)
}

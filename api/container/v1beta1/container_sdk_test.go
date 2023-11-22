package container

import (
	"testing"

	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/internal/testhelpers"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/internal/testhelpers/httprecorder"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/scw"
)

func TestListContainerNamespaceMultipleRegions(t *testing.T) {
	client, r, err := httprecorder.CreateRecordedScwClient("container-list-regions")
	testhelpers.AssertNoError(t, err)
	defer func() {
		testhelpers.AssertNoError(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	containerAPI := NewAPI(client)

	// Create server
	_, err = containerAPI.ListNamespaces(&ListNamespacesRequest{}, scw.WithRegions(containerAPI.Regions()...))
	testhelpers.Assert(t, err == nil, "This request should not error: %s", err)
}

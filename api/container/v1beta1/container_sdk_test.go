package container

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers/httprecorder"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func TestListContainerNamespaceMultipleRegions(t *testing.T) {
	client, r, err := httprecorder.CreateRecordedScwClient(t)
	testhelpers.AssertNoError(t, err)
	defer func() {
		testhelpers.AssertNoError(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	containerAPI := NewAPI(client)

	// Create server
	_, err = containerAPI.ListNamespaces(&ListNamespacesRequest{}, scw.WithRegions(containerAPI.Regions()...))
	testhelpers.Assert(t, err == nil, "This request should not error: %s", err)
}

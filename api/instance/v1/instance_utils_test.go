package instance

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

func TestAttachIP(t *testing.T) {

	client, r, err := testhelpers.CreateRecordedScwClient()
	testhelpers.Ok(t, err)
	defer r.Stop() // Make sure recorder is stopped once done with it

	instanceAPI := NewApi(client)

	server := "5ee00b7a-0eba-4f5c-84ce-8f0f32177e2a"

	ip, err := instanceAPI.AttachIp(&AttachIpRequest{
		IpId:   "1c91be6a-7f22-42aa-84d9-e83479613db0",
		Zone:   "fr-par-1",
		Server: &server,
	})

	testhelpers.Ok(t, err)
	testhelpers.Equals(t, server, ip.Server.Id)

	ip, err = instanceAPI.DetachIp(&DetachIpRequest{
		IpId: "1c91be6a-7f22-42aa-84d9-e83479613db0",
		Zone: "fr-par-1",
	})

	testhelpers.Ok(t, err)
	testhelpers.Assert(t, nil == ip.Server, "Server object should be nil for detached IP.")

}

package instance

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/httprecorder"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

func TestAttachIP(t *testing.T) {

	client, r, err := httprecorder.CreateRecordedScwClient()
	testhelpers.Ok(t, err)
	defer r.Stop() // Make sure recorder is stopped once done with it

	instanceAPI := NewApi(client)

	const (
		serverID = "5ee00b7a-0eba-4f5c-84ce-8f0f32177e2a"
		ipID     = "1c91be6a-7f22-42aa-84d9-e83479613db0"
		zone     = "fr-par-1"
	)

	ipAttachResponse, err := instanceAPI.AttachIP(&AttachIPRequest{
		IPID:     ipID,
		Zone:     zone,
		ServerID: serverID,
	})

	testhelpers.Ok(t, err)
	testhelpers.Equals(t, serverID, ipAttachResponse.IP.Server.Id)

	ipDetachResponse, err := instanceAPI.DetachIP(&DetachIPRequest{
		IPID: ipID,
		Zone: zone,
	})

	testhelpers.Ok(t, err)
	testhelpers.Assert(t, nil == ipDetachResponse.IP.Server, "Server object should be nil for detached IP.")

}

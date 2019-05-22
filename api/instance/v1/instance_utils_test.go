package instance

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers/httprecorder"
	"github.com/scaleway/scaleway-sdk-go/utils"
)

func TestAttachIP(t *testing.T) {

	client, r, err := httprecorder.CreateRecordedScwClient("utils-test")
	testhelpers.Ok(t, err)
	defer func() {
		testhelpers.Ok(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	instanceAPI := NewAPI(client)

	var (
		serverID     string
		ipID         string
		volumeID     string
		zone         = utils.ZoneFrPar1
		organization = "d429f6a1-c0a6-48cf-8b5a-1f9dfe76ffd3"
		image        = "f974feac-abae-4365-b988-8ec7d1cec10d"
		reverse      = "1.1.1.1"
	)

	// Create server
	createServerResponse, err := instanceAPI.CreateServer(&CreateServerRequest{
		Zone:         zone,
		Name:         "instance_utils_test",
		Organization: organization,
		Image:        image,
	})
	testhelpers.Ok(t, err)
	serverID = createServerResponse.Server.ID
	for _, volume := range createServerResponse.Server.Volumes {
		volumeID = volume.ID
	}

	// Create IP
	createIPResponse, err := instanceAPI.CreateIP(&CreateIPRequest{
		Zone:         zone,
		Organization: organization,
	})
	testhelpers.Ok(t, err)
	ipID = createIPResponse.IP.ID

	// Attach IP
	ipAttachResponse, err := instanceAPI.AttachIP(&AttachIPRequest{
		IPID:     ipID,
		Zone:     zone,
		ServerID: serverID,
	})

	testhelpers.Ok(t, err)
	testhelpers.Equals(t, serverID, ipAttachResponse.IP.Server.ID)

	// Detach IP
	ipDetachResponse, err := instanceAPI.DetachIP(&DetachIPRequest{
		IPID: ipID,
		Zone: zone,
	})

	testhelpers.Ok(t, err)
	testhelpers.Assert(t, nil == ipDetachResponse.IP.Server, "Server object should be nil for detached IP.")

	// Set reverse
	ipSetReverseResponse, err := instanceAPI.UpdateIP(&UpdateIPRequest{
		IPID:    ipID,
		Zone:    zone,
		Reverse: &reverse,
	})
	testhelpers.Ok(t, err)
	testhelpers.Equals(t, reverse, ipSetReverseResponse.IP.Reverse)

	// Omitempty reverse
	ipSetReverseResponse, err = instanceAPI.UpdateIP(&UpdateIPRequest{
		IPID:    ipID,
		Zone:    zone,
		Reverse: nil,
	})
	testhelpers.Ok(t, err)
	testhelpers.Equals(t, reverse, ipSetReverseResponse.IP.Reverse)

	// Unset reverse
	emptyReverse := ""
	ipDeleteReverseResponse, err := instanceAPI.UpdateIP(&UpdateIPRequest{
		IPID:    ipID,
		Zone:    zone,
		Reverse: &emptyReverse,
	})
	testhelpers.Ok(t, err)
	testhelpers.Equals(t, "", ipDeleteReverseResponse.IP.Reverse)

	// Delete IP
	err = instanceAPI.DeleteIP(&DeleteIPRequest{
		Zone: zone,
		IPID: ipID,
	})
	testhelpers.Ok(t, err)

	// Delete Server
	err = instanceAPI.DeleteServer(&DeleteServerRequest{
		Zone:     zone,
		ServerID: serverID,
	})
	testhelpers.Ok(t, err)

	// Delete Volume
	err = instanceAPI.DeleteVolume(&DeleteVolumeRequest{
		Zone:     zone,
		VolumeID: volumeID,
	})
	testhelpers.Ok(t, err)

}

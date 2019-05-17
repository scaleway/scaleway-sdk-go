package instance

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/httprecorder"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/utils"
)

func TestAttachIP(t *testing.T) {

	client, r, err := httprecorder.CreateRecordedScwClient()
	testhelpers.Ok(t, err)
	defer func() {
		testhelpers.Ok(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	instanceAPI := NewApi(client)

	var (
		serverID     = ""
		ipID         = ""
		volumeID     = ""
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
	serverID = createServerResponse.Server.Id
	for _, volume := range createServerResponse.Server.Volumes {
		volumeID = volume.Id
	}

	// Create IP
	createIPResponse, err := instanceAPI.CreateIp(&CreateIpRequest{
		Zone:         zone,
		Organization: organization,
	})
	testhelpers.Ok(t, err)
	ipID = createIPResponse.Ip.Id

	// Attach IP
	ipAttachResponse, err := instanceAPI.AttachIP(&AttachIPRequest{
		IPID:     ipID,
		Zone:     zone,
		ServerID: serverID,
	})

	testhelpers.Ok(t, err)
	testhelpers.Equals(t, serverID, ipAttachResponse.IP.Server.Id)

	// Detach IP
	ipDetachResponse, err := instanceAPI.DetachIP(&DetachIPRequest{
		IPID: ipID,
		Zone: zone,
	})

	testhelpers.Ok(t, err)
	testhelpers.Assert(t, nil == ipDetachResponse.IP.Server, "Server object should be nil for detached IP.")

	// Set reverse
	ipSetReverseResponse, err := instanceAPI.SetReverseForIP(&SetReverseForIPRequest{
		IPID:    ipID,
		Zone:    zone,
		Reverse: reverse,
	})
	testhelpers.Ok(t, err)
	testhelpers.Equals(t, reverse, ipSetReverseResponse.IP.Reverse)

	// Unset reverse
	ipDeleteReverseResponse, err := instanceAPI.DeleteReverseForIP(&DeleteReverseForIPRequest{
		IPID: ipID,
		Zone: zone,
	})
	testhelpers.Ok(t, err)
	testhelpers.Equals(t, "", ipDeleteReverseResponse.IP.Reverse)

	// Delete IP
	err = instanceAPI.DeleteIp(&DeleteIpRequest{
		Zone: zone,
		IpId: ipID,
	})

	// Delete Server
	err = instanceAPI.DeleteServer(&DeleteServerRequest{
		Zone:     zone,
		ServerId: serverID,
	})
	testhelpers.Ok(t, err)

	// Delete Volume
	err = instanceAPI.DeleteVolume(&DeleteVolumeRequest{
		Zone:     zone,
		VolumeId: volumeID,
	})
	testhelpers.Ok(t, err)

}

package instance

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers/httprecorder"
	"github.com/scaleway/scaleway-sdk-go/utils"
)

func TestServerUpdate(t *testing.T) {

	client, r, err := httprecorder.CreateRecordedScwClient()
	testhelpers.Ok(t, err)
	defer func() {
		testhelpers.Ok(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	instanceAPI := NewAPI(client)

	var (
		serverID     string
		volumeID     string
		zone         = utils.ZoneFrPar1
		organization = "d429f6a1-c0a6-48cf-8b5a-1f9dfe76ffd3"
		image        = "f974feac-abae-4365-b988-8ec7d1cec10d"
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

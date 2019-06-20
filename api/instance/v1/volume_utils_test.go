package instance

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers/httprecorder"
	"github.com/scaleway/scaleway-sdk-go/utils"
)

func TestUpdateVolmue(t *testing.T) {

	client, r, err := httprecorder.CreateRecordedScwClient("volume-utils-test")
	testhelpers.Ok(t, err)
	defer func() {
		testhelpers.Ok(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	instanceAPI := NewAPI(client)

	var (
		zone                 = utils.ZoneFrPar1
		organization         = "d429f6a1-c0a6-48cf-8b5a-1f9dfe76ffd3"
		volumeName           = "test volume"
		volumeSize    uint64 = 20000000000
		volumeType           = VolumeTypeLSSD
		newVolumeName        = "some new volume name"

		volumeID string
	)

	// Create volume
	createVolumeResponse, err := instanceAPI.CreateVolume(&CreateVolumeRequest{
		Zone:         zone,
		Name:         volumeName,
		Organization: organization,
		Size:         &volumeSize,
		VolumeType:   volumeType,
	})

	testhelpers.Ok(t, err)

	volumeID = createVolumeResponse.Volume.ID

	// Update volume and test whether successfully updated
	updateVolumeResponse, err := instanceAPI.UpdateVolume(&UpdateVolumeRequest{
		Zone:     zone,
		Name:     &newVolumeName,
		VolumeID: volumeID,
	})

	testhelpers.Ok(t, err)
	testhelpers.Assert(t, updateVolumeResponse.Volume != nil, "Should have volume in response")
	testhelpers.Equals(t, newVolumeName, updateVolumeResponse.Volume.Name)
	testhelpers.Equals(t, volumeSize, updateVolumeResponse.Volume.Size) // check that server is not changed

	// Delete Volume
	err = instanceAPI.DeleteVolume(&DeleteVolumeRequest{
		Zone:     zone,
		VolumeID: volumeID,
	})
	testhelpers.Ok(t, err)

}

package instance

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers/httprecorder"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/recorder"
)

func TestUpdateVolume(t *testing.T) {
	client, r, err := httprecorder.CreateRecordedScwClient(t)
	testhelpers.AssertNoError(t, err)
	defer func() {
		testhelpers.AssertNoError(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	project, ok := client.GetDefaultProjectID()
	if !ok && r.Mode() == recorder.ModeRecordOnly {
		t.Fatal("default project ID is required to record this test")
	}
	instanceAPI := NewAPI(client)

	var (
		zone          = scw.ZoneFrPar1
		volumeName    = "test volume"
		volumeSize    = 20 * scw.GB
		volumeType    = VolumeVolumeTypeLSSD
		newVolumeName = "some new volume name"

		volumeID string
	)

	// Create volume
	createVolumeResponse, err := instanceAPI.CreateVolume(&CreateVolumeRequest{
		Zone:       zone,
		Name:       volumeName,
		Project:    &project,
		Size:       &volumeSize,
		VolumeType: volumeType,
	})

	testhelpers.AssertNoError(t, err)

	volumeID = createVolumeResponse.Volume.ID

	// Update volume and test whether successfully updated
	updateVolumeResponse, err := instanceAPI.UpdateVolume(&UpdateVolumeRequest{
		Zone:     zone,
		Name:     &newVolumeName,
		VolumeID: volumeID,
	})

	testhelpers.AssertNoError(t, err)
	testhelpers.Assert(t, updateVolumeResponse.Volume != nil, "Should have volume in response")
	testhelpers.Equals(t, newVolumeName, updateVolumeResponse.Volume.Name)
	testhelpers.Equals(t, volumeSize, updateVolumeResponse.Volume.Size) // check that server is not changed

	// Delete Volume
	err = instanceAPI.DeleteVolume(&DeleteVolumeRequest{
		Zone:     zone,
		VolumeID: volumeID,
	})
	testhelpers.AssertNoError(t, err)
}

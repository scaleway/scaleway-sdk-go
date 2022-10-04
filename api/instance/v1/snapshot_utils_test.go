package instance

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers/httprecorder"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func TestWaitForSnapshot(t *testing.T) {
	client, r, err := httprecorder.CreateRecordedScwClient("snapshot-wait-test")
	testhelpers.AssertNoError(t, err)
	defer func() {
		testhelpers.AssertNoError(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	instanceAPI := NewAPI(client)
	snapshotName := "backup"
	snapshot, cleanup := createSnapshot(t, instanceAPI, snapshotName)
	defer cleanup()

	res, err := instanceAPI.WaitForSnapshot(&WaitForSnapshotRequest{
		SnapshotID: snapshot.ID,
	})

	testhelpers.AssertNoError(t, err)
	testhelpers.Equals(t, snapshot.ID, res.ID)
	testhelpers.Equals(t, SnapshotStateAvailable, res.State)
	testhelpers.Equals(t, snapshotName, res.Name)
}

// createSnapshot is a helper that create an snapshot.
// It returns the newly created snapshot and a cleanup function
func createSnapshot(t *testing.T, instanceAPI *API, snapshotName string) (*Snapshot, func()) {
	serverRes, err := instanceAPI.CreateServer(&CreateServerRequest{
		CommercialType: "DEV1-M",
		Image:          "ubuntu_focal",
	})
	testhelpers.AssertNoError(t, err)

	// Backup will create a snapshot for each volume + an image base on all snapshots.
	snapshot, err := instanceAPI.CreateSnapshot(&CreateSnapshotRequest{
		Name:     snapshotName,
		VolumeID: &serverRes.Server.Volumes["0"].ID,
	})
	testhelpers.AssertNoError(t, err)

	snapshotRes, err := instanceAPI.GetSnapshot(&GetSnapshotRequest{
		SnapshotID: snapshot.Snapshot.ID,
	})
	testhelpers.AssertNoError(t, err)

	return snapshotRes.Snapshot, func() {
		// Delete all created resources

		err := instanceAPI.DeleteServer(&DeleteServerRequest{
			ServerID: serverRes.Server.ID,
		})
		testhelpers.AssertNoError(t, err)

		err = instanceAPI.DeleteSnapshot(&DeleteSnapshotRequest{
			SnapshotID: snapshotRes.Snapshot.ID,
		})
		testhelpers.AssertNoError(t, err)
	}
}

func TestAPI_UpdateSnapshot(t *testing.T) {
	client, r, err := httprecorder.CreateRecordedScwClient("snapshot-test")
	testhelpers.AssertNoError(t, err)
	defer func() {
		testhelpers.AssertNoError(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	instanceAPI := NewAPI(client)

	var (
		volumeSize = 1 * scw.GB
	)

	createVolume, err := instanceAPI.CreateVolume(&CreateVolumeRequest{
		Name:       "volume_name",
		VolumeType: VolumeVolumeTypeBSSD,
		Size:       &volumeSize,
	})
	testhelpers.AssertNoError(t, err)

	createResponse, err := instanceAPI.CreateSnapshot(&CreateSnapshotRequest{
		Name:     "name",
		VolumeID: &createVolume.Volume.ID,
	})
	testhelpers.AssertNoError(t, err)

	updateResponse, err := instanceAPI.UpdateSnapshot(&UpdateSnapshotRequest{
		SnapshotID: createResponse.Snapshot.ID,
		Name:       scw.StringPtr("new_name"),
		Tags:       scw.StringsPtr([]string{"foo", "bar"}),
	})
	testhelpers.AssertNoError(t, err)
	testhelpers.Equals(t, "new_name", updateResponse.Snapshot.Name)
	testhelpers.Equals(t, []string{"foo", "bar"}, updateResponse.Snapshot.Tags)

	_, err = instanceAPI.WaitForSnapshot(&WaitForSnapshotRequest{
		SnapshotID: createResponse.Snapshot.ID,
	})
	testhelpers.AssertNoError(t, err)

	err = instanceAPI.DeleteSnapshot(&DeleteSnapshotRequest{
		SnapshotID: createResponse.Snapshot.ID,
	})
	testhelpers.AssertNoError(t, err)
}

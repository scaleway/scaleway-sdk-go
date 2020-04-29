package instance

import (
	"strings"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers/httprecorder"
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
		Image:          "ubuntu-bionic",
	})
	testhelpers.AssertNoError(t, err)

	// Backup will create a snapshot for each volume + an image base on all snapshots.
	backupRes, err := instanceAPI.ServerAction(&ServerActionRequest{
		ServerID: serverRes.Server.ID,
		Action:   ServerActionBackup,
		Name:     &snapshotName,
	})
	testhelpers.AssertNoError(t, err)

	tmp := strings.Split(backupRes.Task.HrefResult, "/")
	snapshotID := tmp[2]
	snapshotRes, err := instanceAPI.GetSnapshot(&GetSnapshotRequest{
		SnapshotID: snapshotID,
	})
	testhelpers.AssertNoError(t, err)

	return snapshotRes.Snapshot, func() {
		// Delete all created resources

		err := instanceAPI.DeleteServer(&DeleteServerRequest{
			ServerID: serverRes.Server.ID,
		})
		testhelpers.AssertNoError(t, err)

		err = instanceAPI.DeleteVolume(&DeleteVolumeRequest{
			VolumeID: serverRes.Server.Volumes["0"].ID,
		})
		testhelpers.AssertNoError(t, err)

		err = instanceAPI.DeleteSnapshot(&DeleteSnapshotRequest{
			SnapshotID: snapshotRes.Snapshot.ID,
		})
		testhelpers.AssertNoError(t, err)
	}
}

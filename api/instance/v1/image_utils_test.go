package instance

import (
	"strings"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers/httprecorder"
)

func TestWaitForImage(t *testing.T) {
	client, r, err := httprecorder.CreateRecordedScwClient("image-wait-test")
	testhelpers.AssertNoError(t, err)
	defer func() {
		testhelpers.AssertNoError(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	instanceAPI := NewAPI(client)
	imageName := "backup"
	image, cleanup := createImage(t, instanceAPI, imageName)
	defer cleanup()

	res, err := instanceAPI.WaitForImage(&WaitForImageRequest{
		ImageID: image.ID,
	})

	testhelpers.AssertNoError(t, err)
	testhelpers.Equals(t, image.ID, res.ID)
	testhelpers.Equals(t, ImageStateAvailable, res.State)
	testhelpers.Equals(t, imageName, res.Name)
}

// createImage cis a helper that create an image.
// It return the newly created image and a cleanup function
func createImage(t *testing.T, instanceAPI *API, imageName string) (*Image, func()) {
	serverRes, err := instanceAPI.CreateServer(&CreateServerRequest{
		CommercialType: "DEV1-M",
		Image:          "ubuntu_focal",
	})
	testhelpers.AssertNoError(t, err)

	// Backup will create a snapshot for each volume + an image base on all snapshots.
	backupRes, err := instanceAPI.ServerAction(&ServerActionRequest{
		ServerID: serverRes.Server.ID,
		Action:   ServerActionBackup,
		Name:     &imageName,
	})
	testhelpers.AssertNoError(t, err)

	tmp := strings.Split(backupRes.Task.HrefResult, "/")
	imageID := tmp[2]
	imageRes, err := instanceAPI.GetImage(&GetImageRequest{
		ImageID: imageID,
	})
	testhelpers.AssertNoError(t, err)

	return imageRes.Image, func() {
		// Delete all created resources

		err := instanceAPI.DeleteServer(&DeleteServerRequest{
			ServerID: serverRes.Server.ID,
		})
		testhelpers.AssertNoError(t, err)

		err = instanceAPI.DeleteVolume(&DeleteVolumeRequest{
			VolumeID: serverRes.Server.Volumes["0"].ID,
		})
		testhelpers.AssertNoError(t, err)

		err = instanceAPI.DeleteImage(&DeleteImageRequest{
			ImageID: imageRes.Image.ID,
		})
		testhelpers.AssertNoError(t, err)

		err = instanceAPI.DeleteSnapshot(&DeleteSnapshotRequest{
			SnapshotID: imageRes.Image.RootVolume.ID,
		})
		testhelpers.AssertNoError(t, err)
	}
}

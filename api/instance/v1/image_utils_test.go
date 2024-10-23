package instance_test

import (
	"strings"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers/httprecorder"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func TestWaitForImage(t *testing.T) {
	client, r, err := httprecorder.CreateRecordedScwClient("image-wait-test")
	testhelpers.AssertNoError(t, err)
	defer func() {
		testhelpers.AssertNoError(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	instanceAPI := instance.NewAPI(client)
	imageName := "backup"
	image, cleanup := createImage(t, instanceAPI, imageName)
	defer cleanup()

	res, err := instanceAPI.WaitForImage(&instance.WaitForImageRequest{
		ImageID: image.ID,
	})

	testhelpers.AssertNoError(t, err)
	testhelpers.Equals(t, image.ID, res.ID)
	testhelpers.Equals(t, instance.ImageStateAvailable, res.State)
	testhelpers.Equals(t, imageName, res.Name)
}

// createImage cis a helper that create an image.
// It return the newly created image and a cleanup function
func createImage(t *testing.T, instanceAPI *instance.API, imageName string) (*instance.Image, func()) {
	t.Helper()
	serverRes, err := instanceAPI.CreateServer(&instance.CreateServerRequest{
		CommercialType: "DEV1-M",
		Image:          scw.StringPtr("ubuntu_focal"),
	})
	testhelpers.AssertNoError(t, err)

	// Backup will create a snapshot for each volume + an image base on all snapshots.
	backupRes, err := instanceAPI.ServerAction(&instance.ServerActionRequest{
		ServerID: serverRes.Server.ID,
		Action:   instance.ServerActionBackup,
		Name:     &imageName,
	})
	testhelpers.AssertNoError(t, err)

	tmp := strings.Split(backupRes.Task.HrefResult, "/")
	imageID := tmp[2]
	imageRes, err := instanceAPI.GetImage(&instance.GetImageRequest{
		ImageID: imageID,
	})
	testhelpers.AssertNoError(t, err)

	return imageRes.Image, func() {
		// Delete all created resources

		err := instanceAPI.DeleteServer(&instance.DeleteServerRequest{
			ServerID: serverRes.Server.ID,
		})
		testhelpers.AssertNoError(t, err)

		err = instanceAPI.DeleteVolume(&instance.DeleteVolumeRequest{
			VolumeID: serverRes.Server.Volumes["0"].ID,
		})
		testhelpers.AssertNoError(t, err)

		err = instanceAPI.DeleteImage(&instance.DeleteImageRequest{
			ImageID: imageRes.Image.ID,
		})
		testhelpers.AssertNoError(t, err)

		err = instanceAPI.DeleteSnapshot(&instance.DeleteSnapshotRequest{
			SnapshotID: imageRes.Image.RootVolume.ID,
		})
		testhelpers.AssertNoError(t, err)
	}
}

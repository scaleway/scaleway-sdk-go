package instance

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"

	block "github.com/scaleway/scaleway-sdk-go/api/block/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/api/marketplace/v2"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers/httprecorder"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func TestAPI_GetServerType(t *testing.T) {
	client, r, err := httprecorder.CreateRecordedScwClient("get-server-type")
	testhelpers.AssertNoError(t, err)
	defer func() {
		testhelpers.AssertNoError(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	instanceAPI := NewAPI(client)

	serverType, err := instanceAPI.GetServerType(&GetServerTypeRequest{
		Zone: scw.ZoneFrPar1,
		Name: "GP1-XS",
	})

	testhelpers.AssertNoError(t, err)
	testhelpers.Equals(t, 1*scw.GB, serverType.PerVolumeConstraint.LSSD.MinSize)
	testhelpers.Equals(t, 800*scw.GB, serverType.PerVolumeConstraint.LSSD.MaxSize)
}

func TestAPI_ServerUserData(t *testing.T) {
	client, r, err := httprecorder.CreateRecordedScwClient("server-user-data")
	testhelpers.AssertNoError(t, err)
	defer func() {
		testhelpers.AssertNoError(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	instanceAPI := NewAPI(client)

	key := "hello"
	contentStr := "world"

	serverRes, err := instanceAPI.CreateServer(&CreateServerRequest{
		Zone:           scw.ZoneFrPar1,
		CommercialType: "DEV1-S",
		Name:           namegenerator.GetRandomName("srv"),
		Image:          scw.StringPtr("f974feac-abae-4365-b988-8ec7d1cec10d"),
		Project:        scw.StringPtr("14d2f7ae-9775-414c-9bed-6810e060d500"),
	})
	testhelpers.AssertNoError(t, err)

	content := strings.NewReader(contentStr)
	err = instanceAPI.SetServerUserData(&SetServerUserDataRequest{
		Zone:     scw.ZoneFrPar1,
		ServerID: serverRes.Server.ID,
		Key:      key,
		Content:  content,
	})
	testhelpers.AssertNoError(t, err)

	data, err := instanceAPI.GetServerUserData(&GetServerUserDataRequest{
		Zone:     scw.ZoneFrPar1,
		ServerID: serverRes.Server.ID,
		Key:      key,
	})
	testhelpers.AssertNoError(t, err)

	resUserData, err := io.ReadAll(data)
	testhelpers.AssertNoError(t, err)
	testhelpers.Equals(t, contentStr, string(resUserData))
}

func TestAPI_AllServerUserData(t *testing.T) {
	client, r, err := httprecorder.CreateRecordedScwClient("all-server-user-data")
	testhelpers.AssertNoError(t, err)
	defer func() {
		testhelpers.AssertNoError(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	instanceAPI := NewAPI(client)

	serverRes, err := instanceAPI.CreateServer(&CreateServerRequest{
		Zone:           scw.ZoneFrPar1,
		CommercialType: "DEV1-S",
		Name:           namegenerator.GetRandomName("srv"),
		Image:          scw.StringPtr("f974feac-abae-4365-b988-8ec7d1cec10d"),
		Project:        scw.StringPtr("14d2f7ae-9775-414c-9bed-6810e060d500"),
	})
	testhelpers.AssertNoError(t, err)

	steps := []map[string]string{
		{
			"hello":      "world",
			"scale":      "way",
			"tic":        "tac",
			"cloud-init": "on",
		},
		{
			"scale":      "way",
			"steve":      "wozniak",
			"cloud-init": "off",
		},
		{},
	}

	for _, data := range steps {
		// create user data
		userData := make(map[string]io.Reader, len(data))
		for k, v := range data {
			userData[k] = bytes.NewBufferString(v)
		}

		// set all user data
		err := instanceAPI.SetAllServerUserData(&SetAllServerUserDataRequest{
			Zone:     scw.ZoneFrPar1,
			ServerID: serverRes.Server.ID,
			UserData: userData,
		})
		testhelpers.AssertNoError(t, err)

		// get all user data
		allData, err := instanceAPI.GetAllServerUserData(&GetAllServerUserDataRequest{
			Zone:     scw.ZoneFrPar1,
			ServerID: serverRes.Server.ID,
		})
		testhelpers.AssertNoError(t, err)

		testhelpers.Equals(t, len(data), len(allData.UserData))

		for expectedKey, expectedValue := range data {
			currentReader, exists := allData.UserData[expectedKey]
			testhelpers.Assert(t, exists, "%s key not found in result", expectedKey)

			currentValue, err := io.ReadAll(currentReader)
			testhelpers.AssertNoError(t, err)

			testhelpers.Equals(t, expectedValue, string(currentValue))
		}
	}
}

func TestAPI_CreateServer(t *testing.T) {
	client, r, err := httprecorder.CreateRecordedScwClient("create-server")
	testhelpers.AssertNoError(t, err)
	defer func() {
		testhelpers.AssertNoError(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	instanceAPI := NewAPI(client)

	res, err := instanceAPI.CreateServer(&CreateServerRequest{
		Zone:           scw.ZoneFrPar1,
		CommercialType: "GP1-XS",
		Image:          scw.StringPtr("ubuntu_focal"),
	})

	testhelpers.AssertNoError(t, err)
	// this UUID might change when running the cassette later when the image "ubuntu_focal" got a new version
	testhelpers.Equals(t, "9c41e95b-add2-4ef8-b1b1-af8899748eda", res.Server.Image.ID)
	err = instanceAPI.DeleteServer(&DeleteServerRequest{
		Zone:     scw.ZoneFrPar1,
		ServerID: res.Server.ID,
	})
	testhelpers.AssertNoError(t, err)
}

func TestAPI_CreateServerImageLabelResolution(t *testing.T) {
	client, r, err := httprecorder.CreateRecordedScwClient("create-server-image-label-resolution")
	testhelpers.AssertNoError(t, err)
	defer func() {
		testhelpers.AssertNoError(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	instanceAPI := NewAPI(client)
	serverID := ""
	commercialType := "GP1-XS"
	imageLabel := "ubuntu_noble"
	size := 15 * scw.GB
	zone := scw.ZoneFrPar1

	t.Run("default-root-volume", func(t *testing.T) {
		// Create a server with default settings
		res, err := instanceAPI.CreateServer(&CreateServerRequest{
			Zone:           zone,
			CommercialType: commercialType,
			Image:          &imageLabel,
		})
		testhelpers.AssertNoError(t, err)
		serverID = res.Server.ID

		testhelpers.Equals(t, VolumeVolumeTypeLSSD.String(), res.Server.Volumes["0"].VolumeType.String())

		checkImage(t, client, res.Server.Image.ID, imageLabel, commercialType, marketplace.LocalImageTypeInstanceLocal)
		err = cleanupTestResources(client, serverID)
		testhelpers.AssertNoError(t, err)
	})

	t.Run("sbs-root-volume", func(t *testing.T) {
		volumeType := VolumeVolumeTypeSbsVolume

		// Create server with an SBS root volume
		res, err := instanceAPI.CreateServer(&CreateServerRequest{
			Zone:           zone,
			CommercialType: commercialType,
			Image:          &imageLabel,
			Volumes: map[string]*VolumeServerTemplate{
				"0": {
					Size:       &size,
					VolumeType: volumeType,
				},
			},
		})
		testhelpers.AssertNoError(t, err)
		serverID = res.Server.ID
		testhelpers.Equals(t, volumeType.String(), res.Server.Volumes["0"].VolumeType.String())
		checkImage(t, client, res.Server.Image.ID, imageLabel, commercialType, marketplace.LocalImageTypeInstanceSbs)
		err = cleanupTestResources(client, serverID)
		testhelpers.AssertNoError(t, err)
	})

	t.Run("local-root-volume", func(t *testing.T) {
		volumeType := VolumeVolumeTypeLSSD

		// Create server with a local root volume
		res, err := instanceAPI.CreateServer(&CreateServerRequest{
			Zone:           zone,
			CommercialType: commercialType,
			Image:          &imageLabel,
			Volumes: map[string]*VolumeServerTemplate{
				"0": {
					Size:       &size,
					VolumeType: volumeType,
				},
			},
		})
		testhelpers.AssertNoError(t, err)
		serverID = res.Server.ID

		testhelpers.Equals(t, volumeType.String(), res.Server.Volumes["0"].VolumeType.String())

		checkImage(t, client, res.Server.Image.ID, imageLabel, commercialType, marketplace.LocalImageTypeInstanceLocal)
		err = cleanupTestResources(client, serverID)
		testhelpers.AssertNoError(t, err)
	})

	t.Run("sbs-snapshot-volume", func(t *testing.T) {
		volumeType := VolumeVolumeTypeSbsVolume

		sbsSnapshot := createSBSSnapshot(t, client)

		// Create a server with the sbs snapshot
		res, err := instanceAPI.CreateServer(&CreateServerRequest{
			Zone:           zone,
			CommercialType: commercialType,
			Image:          &imageLabel,
			Volumes: map[string]*VolumeServerTemplate{
				"0": {
					VolumeType:   volumeType,
					BaseSnapshot: &sbsSnapshot.ID,
					Name:         scw.StringPtr("sbs-volume-from-snap"),
				},
			},
		})
		testhelpers.AssertNoError(t, err)
		serverID = res.Server.ID

		testhelpers.Equals(t, volumeType.String(), res.Server.Volumes["0"].VolumeType.String())

		checkImage(t, client, res.Server.Image.ID, imageLabel, commercialType, marketplace.LocalImageTypeInstanceSbs)
		err = cleanupTestResources(client, serverID, sbsSnapshot.ID)
		testhelpers.AssertNoError(t, err)
	})

	t.Run("local-snapshot-volume", func(t *testing.T) {
		volumeType := VolumeVolumeTypeLSSD

		localSnapshot := createLocalSnapshot(t, client)

		// Create a server with the local snapshot
		res, err := instanceAPI.CreateServer(&CreateServerRequest{
			Zone:           zone,
			CommercialType: commercialType,
			Image:          &imageLabel,
			Volumes: map[string]*VolumeServerTemplate{
				"0": {
					VolumeType:   volumeType,
					BaseSnapshot: &localSnapshot.ID,
					Name:         scw.StringPtr("local-volume-from-snap"),
				},
			},
		})
		testhelpers.AssertNoError(t, err)
		serverID = res.Server.ID

		testhelpers.Equals(t, volumeType.String(), res.Server.Volumes["0"].VolumeType.String())

		checkImage(t, client, res.Server.Image.ID, imageLabel, commercialType, marketplace.LocalImageTypeInstanceLocal)
		err = cleanupTestResources(client, serverID, localSnapshot.ID)
		testhelpers.AssertNoError(t, err)
	})

	t.Run("sbs-boot-on-custom-index", func(t *testing.T) {
		blockAPI := block.NewAPI(client)
		localVolumeType := VolumeVolumeTypeLSSD
		blockVolumeType := VolumeVolumeTypeSbsVolume
		boot := true

		// Create sbs volume to boot on
		sbsVolume, err := blockAPI.CreateVolume(&block.CreateVolumeRequest{
			Zone:     zone,
			PerfIops: scw.Uint32Ptr(5000),
			FromEmpty: &block.CreateVolumeRequestFromEmpty{
				Size: size,
			},
		})
		testhelpers.AssertNoError(t, err)

		// Create a server that boots on an additional sbs volume
		res, err := instanceAPI.CreateServer(&CreateServerRequest{
			Zone:           zone,
			CommercialType: commercialType,
			Image:          &imageLabel,
			Volumes: map[string]*VolumeServerTemplate{
				"0": {
					Size:       &size,
					VolumeType: localVolumeType,
					Name:       scw.StringPtr("empty-root-volume"),
				},
				"1": {
					VolumeType: blockVolumeType,
					Boot:       &boot,
					ID:         &sbsVolume.ID,
				},
			},
		})
		testhelpers.AssertNoError(t, err)
		serverID = res.Server.ID

		testhelpers.Equals(t, localVolumeType.String(), res.Server.Volumes["0"].VolumeType.String())
		testhelpers.Equals(t, blockVolumeType.String(), res.Server.Volumes["1"].VolumeType.String())
		testhelpers.Equals(t, false, res.Server.Volumes["0"].Boot)
		testhelpers.Equals(t, true, res.Server.Volumes["1"].Boot)

		checkImage(t, client, res.Server.Image.ID, imageLabel, commercialType, marketplace.LocalImageTypeInstanceSbs)
		err = cleanupTestResources(client, serverID)
		testhelpers.AssertNoError(t, err)
	})

	t.Run("local-boot-on-custom-index", func(t *testing.T) {
		localVolumeType := VolumeVolumeTypeLSSD
		blockVolumeType := VolumeVolumeTypeSbsVolume
		boot := true

		// Create local snapshot for the additional volume to boot on
		localSnapshot := createLocalSnapshot(t, client)

		// Create block snapshot for the root volume
		sbsSnapshot := createSBSSnapshot(t, client)

		// Create a server that boots on an additional local volume
		res, err := instanceAPI.CreateServer(&CreateServerRequest{
			Zone:           zone,
			CommercialType: commercialType,
			Image:          &imageLabel,
			Volumes: map[string]*VolumeServerTemplate{
				"0": {
					VolumeType:   blockVolumeType,
					BaseSnapshot: &sbsSnapshot.ID,
					Name:         scw.StringPtr("empty-root-volume"),
				},
				"1": {
					VolumeType:   localVolumeType,
					Boot:         &boot,
					BaseSnapshot: &localSnapshot.ID,
					Name:         scw.StringPtr("boot-on-index-1-local"),
				},
			},
		})
		testhelpers.AssertNoError(t, err)
		serverID = res.Server.ID

		testhelpers.Equals(t, blockVolumeType.String(), res.Server.Volumes["0"].VolumeType.String())
		testhelpers.Equals(t, localVolumeType.String(), res.Server.Volumes["1"].VolumeType.String())
		testhelpers.Equals(t, false, res.Server.Volumes["0"].Boot)
		testhelpers.Equals(t, true, res.Server.Volumes["1"].Boot)

		checkImage(t, client, res.Server.Image.ID, imageLabel, commercialType, marketplace.LocalImageTypeInstanceLocal)
		err = cleanupTestResources(client, serverID, localSnapshot.ID, sbsSnapshot.ID)
		testhelpers.AssertNoError(t, err)
	})
}

func createLocalSnapshot(t *testing.T, client *scw.Client) *Snapshot {
	t.Helper()

	instanceAPI := NewAPI(client)
	size := 15 * scw.GB
	zone, ok := client.GetDefaultZone()
	if !ok {
		zone = scw.ZoneFrPar1
	}

	// Create a server with a local volume to be snapshot
	tmpServer, err := instanceAPI.CreateServer(&CreateServerRequest{
		Zone:           zone,
		CommercialType: "DEV1-S",
		Image:          scw.StringPtr("ubuntu_focal"),
		Volumes: map[string]*VolumeServerTemplate{
			"0": {
				Size:       &size,
				VolumeType: VolumeVolumeTypeLSSD,
			},
		},
	})
	testhelpers.AssertNoError(t, err)

	// Create local snapshot
	localSnap, err := instanceAPI.CreateSnapshot(&CreateSnapshotRequest{
		Zone:       zone,
		VolumeType: SnapshotVolumeTypeLSSD,
		VolumeID:   &tmpServer.Server.Volumes["0"].ID,
	})
	testhelpers.AssertNoError(t, err)

	// Delete temporary server
	errs := deleteServerAndVolumes(client, tmpServer.Server.ID, zone)
	if errs != nil {
		testhelpers.Assert(t, len(errs) == 0, errors.Join(errs...).Error())
	}

	return localSnap.Snapshot
}

func createSBSSnapshot(t *testing.T, client *scw.Client) *block.Snapshot {
	t.Helper()

	blockAPI := block.NewAPI(client)
	size := 15 * scw.GB
	zone, ok := client.GetDefaultZone()
	if !ok {
		zone = scw.ZoneFrPar1
	}

	// Create sbs volume to be snapshot
	tmpVolume, err := blockAPI.CreateVolume(&block.CreateVolumeRequest{
		Zone:     zone,
		PerfIops: scw.Uint32Ptr(5000),
		FromEmpty: &block.CreateVolumeRequestFromEmpty{
			Size: size,
		},
	})
	testhelpers.AssertNoError(t, err)

	_, err = blockAPI.WaitForVolume(&block.WaitForVolumeRequest{
		Zone:     zone,
		VolumeID: tmpVolume.ID,
	})
	testhelpers.AssertNoError(t, err)

	// Create sbs snapshot
	sbsSnap, err := blockAPI.CreateSnapshot(&block.CreateSnapshotRequest{
		Zone:     zone,
		VolumeID: tmpVolume.ID,
	})
	testhelpers.AssertNoError(t, err)

	// Delete temporary volume
	sbsSnap, err = blockAPI.WaitForSnapshot(&block.WaitForSnapshotRequest{
		Zone:       zone,
		SnapshotID: sbsSnap.ID,
	})
	testhelpers.AssertNoError(t, err)

	err = blockAPI.DeleteVolume(&block.DeleteVolumeRequest{
		Zone:     zone,
		VolumeID: tmpVolume.ID,
	})
	testhelpers.AssertNoError(t, err)

	return sbsSnap
}

func checkImage(t *testing.T, client *scw.Client, actualImageID string, expectedImageLabel string, commercialType string, expectedLocalImageType marketplace.LocalImageType) {
	t.Helper()

	// Get the expected marketplace image for comparison
	marketplaceAPI := marketplace.NewAPI(client)
	actualImage, err := marketplaceAPI.GetLocalImage(&marketplace.GetLocalImageRequest{
		LocalImageID: actualImageID,
	})
	testhelpers.AssertNoError(t, err)
	testhelpers.Equals(t, expectedImageLabel, actualImage.Label)
	testhelpers.Equals(t, expectedLocalImageType, actualImage.Type)
	testhelpers.Equals(t, scw.ZoneFrPar1, actualImage.Zone)
	testhelpers.Equals(t, ArchX86_64.String(), actualImage.Arch)

	isCompatible := false
	for _, compatibleOffer := range actualImage.CompatibleCommercialTypes {
		if compatibleOffer == commercialType {
			isCompatible = true
		}
	}
	testhelpers.Assert(t, isCompatible, "actual image is not compatible with requested commercial type")
}

func deleteServerAndVolumes(client *scw.Client, serverID string, zone scw.Zone) []error {
	instanceAPI := NewAPI(client)
	blockAPI := block.NewAPI(client)
	errs := []error(nil)

	server, err := instanceAPI.GetServer(&GetServerRequest{
		ServerID: serverID,
		Zone:     zone,
	})
	if err != nil {
		return append(errs, fmt.Errorf("- failed to get server: %w\n", err))
	}

	// Delete server first so volumes get detached
	if serverID != "" {
		err = instanceAPI.DeleteServer(&DeleteServerRequest{
			Zone:     scw.ZoneFrPar1,
			ServerID: serverID,
		})
		if err != nil {
			errs = append(errs, fmt.Errorf("- failed to delete server %s: %w\n", serverID, err))
		}
	}

	// Delete volumes
	for _, volume := range server.Server.Volumes {
		switch volume.VolumeType {
		case VolumeServerVolumeTypeLSSD:
			err = instanceAPI.DeleteVolume(&DeleteVolumeRequest{
				VolumeID: volume.ID,
				Zone:     zone,
			})
			if err != nil {
				errs = append(errs, fmt.Errorf("- failed to delete volume %s: %w\n", volume.ID, err))
			}
		case VolumeServerVolumeTypeSbsVolume:
			terminalStatus := block.VolumeStatusAvailable
			_, err = blockAPI.WaitForVolume(&block.WaitForVolumeRequest{
				VolumeID:       volume.ID,
				Zone:           zone,
				TerminalStatus: &terminalStatus,
			})
			if err != nil {
				errs = append(errs, fmt.Errorf("- failed to wait for volume %s: %w\n", volume.ID, err))
			}

			err = blockAPI.DeleteVolume(&block.DeleteVolumeRequest{
				VolumeID: volume.ID,
				Zone:     zone,
			})
			if err != nil {
				errs = append(errs, fmt.Errorf("- failed to delete volume %s: %w\n", volume.ID, err))
			}
		}
	}
	return errs
}

func cleanupTestResources(client *scw.Client, serverID string, snapshotIDs ...string) error {
	instanceAPI := NewAPI(client)
	blockAPI := block.NewAPI(client)
	zone := scw.ZoneFrPar1

	errs := deleteServerAndVolumes(client, serverID, zone)

	// Delete snapshot if needed
	for _, snapshotID := range snapshotIDs {
		err := instanceAPI.DeleteSnapshot(&DeleteSnapshotRequest{
			SnapshotID: snapshotID,
			Zone:       zone,
		})
		if err != nil {
			errBlock := blockAPI.DeleteSnapshot(&block.DeleteSnapshotRequest{
				SnapshotID: snapshotID,
				Zone:       zone,
			})
			if errBlock != nil {
				errs = append(errs, fmt.Errorf("- failed to delete snapshot %s: %w\n", snapshotID, err))
			}
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("cleanup test resources:\n%vThere may be some dangling resources", errors.Join(errs...))
	}
	return nil
}

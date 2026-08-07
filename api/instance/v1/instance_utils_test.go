package instance

import (
	"testing"

	block "github.com/scaleway/scaleway-sdk-go/api/block/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers/httprecorder"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/recorder"
)

func TestInstanceHelpers(t *testing.T) {
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
		serverID string
		ipID     string
		volumeID string
		zone     = scw.ZoneFrPar1
		image    = scw.StringPtr("81b9475d-e1b5-43c2-ac48-4c1a3b640686")
	)

	t.Run("create server", func(t *testing.T) {
		createServerResponse, err := instanceAPI.CreateServer(&CreateServerRequest{
			Zone:           zone,
			Name:           "instance_utils_test",
			Project:        &project,
			Image:          image,
			CommercialType: "PRO2-XXS",
		})
		testhelpers.AssertNoError(t, err)
		serverID = createServerResponse.Server.ID
		for _, volume := range createServerResponse.Server.Volumes {
			volumeID = volume.ID
		}
	})

	t.Run("test ip related functions", func(t *testing.T) {
		// Create IP
		createIPResponse, err := instanceAPI.CreateIP(&CreateIPRequest{
			Zone:    zone,
			Project: &project,
		})
		testhelpers.AssertNoError(t, err)
		ipID = createIPResponse.IP.ID

		// Attach IP
		ipAttachResponse, err := instanceAPI.AttachIP(&AttachIPRequest{
			IP:       ipID,
			Zone:     zone,
			ServerID: serverID,
		})

		testhelpers.AssertNoError(t, err)
		testhelpers.Equals(t, serverID, ipAttachResponse.IP.Server.ID)

		// Detach IP
		ipDetachResponse, err := instanceAPI.DetachIP(&DetachIPRequest{
			IP:   ipID,
			Zone: zone,
		})

		testhelpers.AssertNoError(t, err)
		testhelpers.Assert(t, nil == ipDetachResponse.IP.Server, "Server object should be nil for detached IP.")

		// Delete IP
		err = instanceAPI.DeleteIP(&DeleteIPRequest{
			Zone: zone,
			IP:   ipID,
		})
		testhelpers.AssertNoError(t, err)
	})

	t.Run("Test attach and detach volume", func(t *testing.T) {
		detachVolumeResponse, err := instanceAPI.DetachVolume(&DetachVolumeRequest{
			Zone:     zone,
			VolumeID: volumeID,
		})
		testhelpers.AssertNoError(t, err)

		testhelpers.Assert(t, detachVolumeResponse.Server != nil, "Should have server in response")
		testhelpers.Assert(t, detachVolumeResponse.Server.Volumes != nil, "Should have volumes in response")
		testhelpers.Assert(t, len(detachVolumeResponse.Server.Volumes) == 0, "Server should have zero volumes after detaching")

		attachVolumeResponse, err := instanceAPI.AttachVolume(&AttachVolumeRequest{
			Zone:     zone,
			ServerID: serverID,
			VolumeID: volumeID,
		})
		testhelpers.AssertNoError(t, err)

		testhelpers.Assert(t, attachVolumeResponse.Server != nil, "Should have server in response")
		testhelpers.Assert(t, attachVolumeResponse.Server.Volumes != nil, "Should have volumes in response")
		testhelpers.Assert(t, len(attachVolumeResponse.Server.Volumes) == 1, "Server should have one volumes after attaching")
		testhelpers.Equals(t, volumeID, attachVolumeResponse.Server.Volumes["0"].ID)
	})

	t.Run("teardown: delete server and volume", func(t *testing.T) {
		// Delete Server
		err = instanceAPI.DeleteServer(&DeleteServerRequest{
			Zone:     zone,
			ServerID: serverID,
		})
		testhelpers.AssertNoError(t, err)

		// Delete Volume
		err = instanceAPI.DeleteVolume(&DeleteVolumeRequest{
			Zone:     zone,
			VolumeID: volumeID,
		})
		testhelpers.AssertNoError(t, err)
	})
}

func TestInstanceHelpers_BlockVolume(t *testing.T) {
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
	blockAPI := block.NewAPI(client)

	var (
		serverID  string
		volumeID  string
		volumeID2 string
		zone      = scw.ZoneFrPar1
		image     = scw.StringPtr("81b9475d-e1b5-43c2-ac48-4c1a3b640686")
	)

	t.Run("create server and volume", func(t *testing.T) {
		createVolumeResponse, err := blockAPI.CreateVolume(&block.CreateVolumeRequest{
			Zone:      zone,
			Name:      "instance_utils_test",
			ProjectID: project,
			FromEmpty: &block.CreateVolumeRequestFromEmpty{
				Size: scw.GB * 20,
			},
			Tags: nil,
		})
		testhelpers.AssertNoError(t, err)

		volumeID = createVolumeResponse.ID

		createVolumeResponse, err = blockAPI.CreateVolume(&block.CreateVolumeRequest{
			Zone:      zone,
			Name:      "instance_utils_test2",
			ProjectID: project,
			FromEmpty: &block.CreateVolumeRequestFromEmpty{
				Size: scw.GB * 20,
			},
			Tags: nil,
		})
		testhelpers.AssertNoError(t, err)

		volumeID2 = createVolumeResponse.ID

		createServerResponse, err := instanceAPI.CreateServer(&CreateServerRequest{
			Zone:           zone,
			Name:           "instance_utils_test",
			Project:        &project,
			Image:          image,
			CommercialType: "PRO2-XXS",
			Volumes: map[string]*VolumeServerTemplate{
				"0": {
					ID:         &volumeID,
					VolumeType: VolumeVolumeTypeSbsVolume,
				},
			},
		})
		testhelpers.AssertNoError(t, err)
		serverID = createServerResponse.Server.ID
	})

	t.Run("Test attach and detach volume", func(t *testing.T) {
		detachVolumeResponse, err := instanceAPI.DetachVolume(&DetachVolumeRequest{
			Zone:          zone,
			VolumeID:      volumeID,
			IsBlockVolume: scw.BoolPtr(true),
		})
		testhelpers.AssertNoError(t, err)

		testhelpers.Assert(t, detachVolumeResponse.Server != nil, "Should have server in response")
		testhelpers.Assert(t, detachVolumeResponse.Server.Volumes != nil, "Should have volumes in response")
		testhelpers.Assert(t, len(detachVolumeResponse.Server.Volumes) == 0, "Server should have zero volumes after detaching")

		_, err = instanceAPI.WaitForServer(&WaitForServerRequest{
			Zone:     zone,
			ServerID: serverID,
		})
		testhelpers.AssertNoError(t, err)

		_, err = blockAPI.WaitForVolume(&block.WaitForVolumeRequest{
			VolumeID: volumeID,
			Zone:     zone,
		})
		testhelpers.AssertNoError(t, err)

		attachVolumeResponse, err := instanceAPI.AttachVolume(&AttachVolumeRequest{
			Zone:     zone,
			ServerID: serverID,
			VolumeID: volumeID,
		})
		testhelpers.AssertNoError(t, err)

		testhelpers.Assert(t, attachVolumeResponse.Server != nil, "Should have server in response")
		testhelpers.Assert(t, attachVolumeResponse.Server.Volumes != nil, "Should have volumes in response")
		testhelpers.Assert(t, len(attachVolumeResponse.Server.Volumes) == 1, "Server should have one volumes after attaching")
		testhelpers.Equals(t, volumeID, attachVolumeResponse.Server.Volumes["0"].ID)

		attachVolumeResponse, err = instanceAPI.AttachVolume(&AttachVolumeRequest{
			Zone:     zone,
			ServerID: serverID,
			VolumeID: volumeID2,
		})
		testhelpers.AssertNoError(t, err)

		testhelpers.Assert(t, attachVolumeResponse.Server != nil, "Should have server in response")
		testhelpers.Assert(t, attachVolumeResponse.Server.Volumes != nil, "Should have volumes in response")
		testhelpers.Assert(t, len(attachVolumeResponse.Server.Volumes) == 2, "Server should have two volumes after attaching second volume")
		testhelpers.Equals(t, volumeID2, attachVolumeResponse.Server.Volumes["1"].ID)
	})

	t.Run("teardown: delete server and volume", func(t *testing.T) {
		// Delete Server
		err = instanceAPI.DeleteServer(&DeleteServerRequest{
			Zone:     zone,
			ServerID: serverID,
		})
		testhelpers.AssertNoError(t, err)

		volumeStatusAvailable := block.VolumeStatusAvailable
		_, err = blockAPI.WaitForVolumeAndReferences(&block.WaitForVolumeAndReferencesRequest{
			VolumeID:             volumeID,
			Zone:                 zone,
			VolumeTerminalStatus: &volumeStatusAvailable,
		})
		testhelpers.AssertNoError(t, err)

		// Delete Volume
		err = blockAPI.DeleteVolume(&block.DeleteVolumeRequest{
			Zone:     zone,
			VolumeID: volumeID,
		})
		testhelpers.AssertNoError(t, err)

		_, err = blockAPI.WaitForVolumeAndReferences(&block.WaitForVolumeAndReferencesRequest{
			VolumeID:             volumeID2,
			Zone:                 zone,
			VolumeTerminalStatus: &volumeStatusAvailable,
		})
		testhelpers.AssertNoError(t, err)

		// Delete Volume
		err = blockAPI.DeleteVolume(&block.DeleteVolumeRequest{
			Zone:     zone,
			VolumeID: volumeID2,
		})
		testhelpers.AssertNoError(t, err)
	})
}

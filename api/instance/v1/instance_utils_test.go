package instance

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers/httprecorder"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func TestInstanceHelpers(t *testing.T) {
	client, r, err := httprecorder.CreateRecordedScwClient("utils-test")
	testhelpers.AssertNoError(t, err)
	defer func() {
		testhelpers.AssertNoError(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	instanceAPI := NewAPI(client)

	var (
		serverID    string
		ipID        string
		volumeID    string
		zone        = scw.ZoneFrPar1
		project     = "d429f6a1-c0a6-48cf-8b5a-1f9dfe76ffd3"
		image       = "f974feac-abae-4365-b988-8ec7d1cec10d"
		reverse     = &NullableStringValue{Value: "1.1.1.1"}
		nullReverse = &NullableStringValue{Null: true}
	)

	t.Run("create server", func(t *testing.T) {
		createServerResponse, err := instanceAPI.CreateServer(&CreateServerRequest{
			Zone:    zone,
			Name:    "instance_utils_test",
			Project: &project,
			Image:   image,
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

		// Set reverse
		ipSetReverseResponse, err := instanceAPI.UpdateIP(&UpdateIPRequest{
			IP:      ipID,
			Zone:    zone,
			Reverse: reverse,
		})
		testhelpers.AssertNoError(t, err)
		testhelpers.Equals(t, reverse.Value, *ipSetReverseResponse.IP.Reverse)

		// Omitempty reverse
		ipSetReverseResponse, err = instanceAPI.UpdateIP(&UpdateIPRequest{
			IP:      ipID,
			Zone:    zone,
			Reverse: nil,
		})
		testhelpers.AssertNoError(t, err)
		testhelpers.Equals(t, reverse.Value, *ipSetReverseResponse.IP.Reverse)

		// Unset reverse
		ipDeleteReverseResponse, err := instanceAPI.UpdateIP(&UpdateIPRequest{
			IP:      ipID,
			Zone:    zone,
			Reverse: nullReverse,
		})
		testhelpers.AssertNoError(t, err)
		testhelpers.Equals(t, (*string)(nil), ipDeleteReverseResponse.IP.Reverse)

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

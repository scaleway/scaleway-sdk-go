package instance

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers/httprecorder"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/recorder"
)

func TestServerUpdate(t *testing.T) {
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
		serverID          string
		volumeID          string
		zone              = scw.ZoneFrPar1
		name              = "instance_sdk_server_test"
		dynamicIPRequired = scw.BoolPtr(true)
		commercialType    = "START1-S"
		image             = scw.StringPtr("f974feac-abae-4365-b988-8ec7d1cec10d")
		bootType          = BootTypeLocal
		tags              = []string{"foo", "bar"}
	)

	t.Run("create server", func(t *testing.T) {
		// Create server
		createServerResponse, err := instanceAPI.CreateServer(&CreateServerRequest{
			Zone:              zone,
			Name:              name,
			Project:           &project,
			Image:             image,
			CommercialType:    commercialType,
			Tags:              tags,
			DynamicIPRequired: dynamicIPRequired,
			BootType:          &bootType,
		})
		testhelpers.AssertNoError(t, err)
		serverID = createServerResponse.Server.ID

		testhelpers.Assert(t, createServerResponse.Server != nil, "Should have server in response")
		testhelpers.Assert(t, len(createServerResponse.Server.Volumes) == 1, "should have exactly one volume because we didn't pass volumes map in the requests.")
		for _, volume := range createServerResponse.Server.Volumes {
			volumeID = volume.ID
		}

		testhelpers.Equals(t, name, createServerResponse.Server.Name)
		testhelpers.Equals(t, *image, createServerResponse.Server.Image.ID)
		testhelpers.Equals(t, bootType, createServerResponse.Server.BootType)
		testhelpers.Equals(t, commercialType, createServerResponse.Server.CommercialType)
		testhelpers.Equals(t, tags, createServerResponse.Server.Tags)
		testhelpers.Equals(t, *dynamicIPRequired, createServerResponse.Server.DynamicIPRequired)
		if r.Mode() == recorder.ModeRecordOnly {
			testhelpers.Equals(t, project, createServerResponse.Server.Project)
		}
	})

	t.Run("update server", func(t *testing.T) {
		var (
			newName     = "some_new_name_for_server"
			updatedTags = []string{}
		)

		// Update server
		updateServerResponse, err := instanceAPI.updateServer(&UpdateServerRequest{
			ServerID: serverID,
			Zone:     zone,
			Name:     &newName,
			Tags:     &updatedTags,
		})
		testhelpers.AssertNoError(t, err)
		testhelpers.Assert(t, updateServerResponse.Server != nil, "Should have server in response")

		// Initial values that are not altered in the above request should remaining the same
		testhelpers.Equals(t, *image, updateServerResponse.Server.Image.ID)
		testhelpers.Equals(t, bootType, updateServerResponse.Server.BootType)
		testhelpers.Equals(t, commercialType, updateServerResponse.Server.CommercialType)
		testhelpers.Equals(t, *dynamicIPRequired, updateServerResponse.Server.DynamicIPRequired)
		testhelpers.Assert(t, len(updateServerResponse.Server.Volumes) == 1, "should have exactly one volume because we didn't pass volumes map in the requests.")
		testhelpers.Equals(t, newName, updateServerResponse.Server.Name)
		testhelpers.Equals(t, updatedTags, updateServerResponse.Server.Tags)
		if r.Mode() == recorder.ModeRecordOnly {
			testhelpers.Equals(t, project, updateServerResponse.Server.Project)
		}
	})

	t.Run("remove server volumes", func(t *testing.T) {
		// Remove/detach volumes
		updateServerResponse, err := instanceAPI.updateServer(&UpdateServerRequest{
			ServerID: serverID,
			Zone:     zone,
			Volumes:  &map[string]*VolumeServerTemplate{},
		})
		testhelpers.AssertNoError(t, err)
		testhelpers.Assert(t, updateServerResponse.Server != nil, "Should have server in response")
		testhelpers.Assert(t, len(updateServerResponse.Server.Volumes) == 0, "volume should be detached from server.")
	})

	t.Run("cleanup server and volume", func(t *testing.T) {
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

func TestCreateServerWithIncorrectBody(t *testing.T) {
	client, r, err := httprecorder.CreateRecordedScwClient(t)
	testhelpers.AssertNoError(t, err)
	defer func() {
		testhelpers.AssertNoError(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	instanceAPI := NewAPI(client)

	zone := scw.ZoneFrPar1

	// Create server
	_, err = instanceAPI.CreateServer(&CreateServerRequest{
		Zone: zone,
	})
	testhelpers.Assert(t, err != nil, "This request should error")
}

func TestListServerMultipleZones(t *testing.T) {
	client, r, err := httprecorder.CreateRecordedScwClient(t)
	testhelpers.AssertNoError(t, err)
	defer func() {
		testhelpers.AssertNoError(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	instanceAPI := NewAPI(client)

	// Create server
	_, err = instanceAPI.ListServers(&ListServersRequest{}, scw.WithZones(instanceAPI.Zones()...))
	testhelpers.Assert(t, err == nil, "This request should not error: %s", err)
}

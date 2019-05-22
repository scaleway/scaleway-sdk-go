package instance

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers/httprecorder"
	"github.com/scaleway/scaleway-sdk-go/utils"
)

func TestServerUpdate(t *testing.T) {

	client, r, err := httprecorder.CreateRecordedScwClient("server-test")
	testhelpers.Ok(t, err)
	defer func() {
		testhelpers.Ok(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	instanceAPI := NewAPI(client)

	var (
		serverID          string
		volumeID          string
		zone              = utils.ZoneFrPar1
		name              = "instance_sdk_server_test"
		dynamicIPRequired = true
		commercialType    = "START1-S"
		image             = "f974feac-abae-4365-b988-8ec7d1cec10d"
		enableIPv6        = true
		bootType          = ServerBootTypeLocal
		tags              = []string{"foo", "bar"}
		organization      = "d429f6a1-c0a6-48cf-8b5a-1f9dfe76ffd3"
	)

	t.Run("create server", func(t *testing.T) {
		// Create server
		createServerResponse, err := instanceAPI.CreateServer(&CreateServerRequest{
			Zone:              zone,
			Name:              name,
			Organization:      organization,
			Image:             image,
			EnableIPv6:        enableIPv6,
			BootType:          bootType,
			CommercialType:    commercialType,
			Tags:              tags,
			DynamicIPRequired: dynamicIPRequired,
		})
		testhelpers.Ok(t, err)
		serverID = createServerResponse.Server.ID

		testhelpers.Assert(t, createServerResponse.Server != nil, "Should have server in response")
		testhelpers.Assert(t, 1 == len(createServerResponse.Server.Volumes), "should have exactly one volume because we didn't pass volumes map in the requests.")
		for _, volume := range createServerResponse.Server.Volumes {
			volumeID = volume.ID
		}

		testhelpers.Equals(t, name, createServerResponse.Server.Name)
		testhelpers.Equals(t, organization, createServerResponse.Server.Organization)
		testhelpers.Equals(t, image, createServerResponse.Server.Image.ID)
		testhelpers.Equals(t, enableIPv6, createServerResponse.Server.EnableIPv6)
		testhelpers.Equals(t, bootType, createServerResponse.Server.BootType)
		testhelpers.Equals(t, commercialType, createServerResponse.Server.CommercialType)
		testhelpers.Equals(t, tags, createServerResponse.Server.Tags)
		testhelpers.Equals(t, dynamicIPRequired, createServerResponse.Server.DynamicIPRequired)
	})

	t.Run("update server", func(t *testing.T) {
		var (
			newName     = "some_new_name_for_server"
			updatedTags = []string{}
		)

		// Update server
		updateServerResponse, err := instanceAPI.UpdateServer(&UpdateServerRequest{
			ServerID: serverID,
			Zone:     zone,
			Name:     &newName,
			Tags:     &updatedTags,
		})
		testhelpers.Assert(t, updateServerResponse.Server != nil, "Should have server in response")
		testhelpers.Ok(t, err)

		// Initial values that are not altered in the above request should remaing the same
		testhelpers.Equals(t, organization, updateServerResponse.Server.Organization)
		testhelpers.Equals(t, image, updateServerResponse.Server.Image.ID)
		testhelpers.Equals(t, enableIPv6, updateServerResponse.Server.EnableIPv6)
		testhelpers.Equals(t, bootType, updateServerResponse.Server.BootType)
		testhelpers.Equals(t, commercialType, updateServerResponse.Server.CommercialType)
		testhelpers.Equals(t, dynamicIPRequired, updateServerResponse.Server.DynamicIPRequired)
		testhelpers.Assert(t, 1 == len(updateServerResponse.Server.Volumes), "should have exactly one volume because we didn't pass volumes map in the requests.")

		testhelpers.Equals(t, newName, updateServerResponse.Server.Name)
		testhelpers.Equals(t, updatedTags, updateServerResponse.Server.Tags)
	})

	t.Run("remove server volumes", func(t *testing.T) {
		// Remove/detach volumes
		updateServerResponse, err := instanceAPI.UpdateServer(&UpdateServerRequest{
			ServerID: serverID,
			Zone:     zone,
			Volumes:  &map[string]*VolumeTemplate{},
		})
		testhelpers.Ok(t, err)
		testhelpers.Assert(t, updateServerResponse.Server != nil, "Should have server in response")
		testhelpers.Assert(t, 0 == len(updateServerResponse.Server.Volumes), "volume should be detached from server.")
	})

	t.Run("cleanup server and volume", func(t *testing.T) {
		// Delete Server
		err = instanceAPI.DeleteServer(&DeleteServerRequest{
			Zone:     zone,
			ServerID: serverID,
		})
		testhelpers.Ok(t, err)

		// Delete Volume
		err = instanceAPI.DeleteVolume(&DeleteVolumeRequest{
			Zone:     zone,
			VolumeID: volumeID,
		})
		testhelpers.Ok(t, err)
	})

}

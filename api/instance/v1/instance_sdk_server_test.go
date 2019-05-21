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
		//tags         = []string{"foo", "bar"}
		organization = "d429f6a1-c0a6-48cf-8b5a-1f9dfe76ffd3"
	)

	// Create server
	createServerResponse, err := instanceAPI.CreateServer(&CreateServerRequest{
		Zone:           zone,
		Name:           name,
		Organization:   organization,
		Image:          image,
		EnableIPv6:     enableIPv6,
		BootType:       bootType,
		CommercialType: commercialType,
		//Tags:         &tags,
		DynamicIPRequired: dynamicIPRequired,
	})
	testhelpers.Ok(t, err)
	serverID = createServerResponse.Server.ID
	for _, volume := range createServerResponse.Server.Volumes {
		volumeID = volume.ID
	}

	testhelpers.Equals(t, name, createServerResponse.Server.Name)
	testhelpers.Equals(t, organization, createServerResponse.Server.Organization)
	testhelpers.Equals(t, image, createServerResponse.Server.Image.ID)
	testhelpers.Equals(t, enableIPv6, createServerResponse.Server.EnableIPv6)
	testhelpers.Equals(t, bootType, createServerResponse.Server.BootType)
	testhelpers.Equals(t, commercialType, createServerResponse.Server.CommercialType)
	testhelpers.Equals(t, dynamicIPRequired, createServerResponse.Server.DynamicIPRequired)

	var (
		newName = "some_new_name_for_server"
		//newBoottype = instanceAPI.Bootty
		tags = []string{}
	)

	// Update server
	updateServerResponse, err := instanceAPI.UpdateServer(&UpdateServerRequest{
		ServerID: serverID,
		Zone:     zone,
		Name:     &newName,
		Tags:     &tags,
	})
	testhelpers.Ok(t, err)

	testhelpers.Equals(t, newName, updateServerResponse.Server.Name)
	testhelpers.Equals(t, tags, updateServerResponse.Server.Tags)
	testhelpers.Equals(t, bootType, updateServerResponse.Server.BootType)
	testhelpers.Equals(t, dynamicIPRequired, updateServerResponse.Server.DynamicIPRequired)
	testhelpers.Assert(t, 1 == len(updateServerResponse.Server.Volumes), "should have exactly one volume because we didn't pass volumes map in the requests.")

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

}

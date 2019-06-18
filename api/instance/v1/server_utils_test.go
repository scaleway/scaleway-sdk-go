package instance

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers/httprecorder"
	"github.com/scaleway/scaleway-sdk-go/utils"
)

func TestAPI_GetServerType(t *testing.T) {
	client, r, err := httprecorder.CreateRecordedScwClient("get-server-type")
	testhelpers.Ok(t, err)
	defer func() {
		testhelpers.Ok(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	instanceAPI := NewAPI(client)

	serverType, err := instanceAPI.GetServerType(&GetServerTypeRequest{
		Zone: utils.ZoneFrPar1,
		Name: "GP1-XS",
	})

	testhelpers.Ok(t, err)
	testhelpers.Equals(t, uint64(1000000000), serverType.PerVolumeConstraint.LSSD.MinSize)
	testhelpers.Equals(t, uint64(800000000000), serverType.PerVolumeConstraint.LSSD.MaxSize)

}

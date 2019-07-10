package marketplace

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers/httprecorder"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func TestGetImageByName(t *testing.T) {

	client, r, err := httprecorder.CreateRecordedScwClient("go-vcr")
	testhelpers.AssertNoError(t, err)
	defer func() {
		testhelpers.AssertNoError(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	t.Run("matching input for GetImageByName", func(t *testing.T) {

		// Create SDK objects for Scaleway Instance product
		marketplaceAPI := NewAPI(client)

		imageID, err := marketplaceAPI.GetLocalImageIDByName(&GetLocalImageIDByNameRequest{
			Zone:           scw.ZoneFrPar1,
			CommercialType: "C1",
			ImageName:      "Docker",
		})
		testhelpers.AssertNoError(t, err)

		// Docker C1 at par1: 45a7e942-1fb0-48c0-bbf6-0acb9af24604
		testhelpers.Equals(t, "45a7e942-1fb0-48c0-bbf6-0acb9af24604", imageID)

	})

	t.Run("non-matching name for GetImageByName", func(t *testing.T) {

		// Create SDK objects for Scaleway Instance product
		marketplaceAPI := NewAPI(client)

		_, err := marketplaceAPI.GetLocalImageIDByName(&GetLocalImageIDByNameRequest{
			Zone:           scw.ZoneFrPar1,
			CommercialType: "",
			ImageName:      "foo-bar-image",
		})
		testhelpers.Assert(t, err != nil, "Should have error")

	})

}

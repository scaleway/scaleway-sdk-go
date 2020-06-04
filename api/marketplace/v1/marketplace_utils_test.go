package marketplace

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers/httprecorder"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func TestGetImageByLabel(t *testing.T) {
	client, r, err := httprecorder.CreateRecordedScwClient("go-vcr")
	testhelpers.AssertNoError(t, err)
	defer func() {
		testhelpers.AssertNoError(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	t.Run("matching input for GetLocalImageIDByLabel", func(t *testing.T) {
		// Create SDK objects for Scaleway Instance product
		marketplaceAPI := NewAPI(client)

		imageID, err := marketplaceAPI.GetLocalImageIDByLabel(&GetLocalImageIDByLabelRequest{
			Zone:           scw.ZoneFrPar1,
			CommercialType: "DEV1-S",
			ImageLabel:     "ubuntu_focal",
		})
		testhelpers.AssertNoError(t, err)

		// ubuntu_focal DEV1-S at par1: 365a8b9c-0c6e-4875-a887-dc3213db9e20
		testhelpers.Equals(t, "365a8b9c-0c6e-4875-a887-dc3213db9e20", imageID)
	})

	t.Run("matching input for GetLocalImageIDByLabel with lowercase image label", func(t *testing.T) {
		// Create SDK objects for Scaleway Instance product
		marketplaceAPI := NewAPI(client)

		imageID, err := marketplaceAPI.GetLocalImageIDByLabel(&GetLocalImageIDByLabelRequest{
			Zone:           scw.ZoneFrPar1,
			CommercialType: "dev1-s",
			ImageLabel:     "ubuntu_focal",
		})
		testhelpers.AssertNoError(t, err)

		// ubuntu_focal DEV1-S at par1: 365a8b9c-0c6e-4875-a887-dc3213db9e20
		testhelpers.Equals(t, "365a8b9c-0c6e-4875-a887-dc3213db9e20", imageID)
	})

	t.Run("non-matching label for GetLocalImageIDByLabel", func(t *testing.T) {
		// Create SDK objects for Scaleway Instance product
		marketplaceAPI := NewAPI(client)

		_, err := marketplaceAPI.GetLocalImageIDByLabel(&GetLocalImageIDByLabelRequest{
			Zone:           scw.ZoneFrPar1,
			CommercialType: "DEV1-S",
			ImageLabel:     "foo-bar-image",
		})
		testhelpers.Assert(t, err != nil, "Should have error")
		testhelpers.Equals(t, "scaleway-sdk-go: couldn't find a matching image for the given label (foo-bar-image), zone (fr-par-1) and commercial type (DEV1-S)", err.Error())
	})
}

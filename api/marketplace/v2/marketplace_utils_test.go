package marketplace_test

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/api/marketplace/v2"
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
		marketplaceAPI := marketplace.NewAPI(client)

		image, err := marketplaceAPI.GetLocalImageByLabel(&marketplace.GetLocalImageByLabelRequest{
			Zone:           scw.ZoneFrPar1,
			CommercialType: "DEV1-S",
			ImageLabel:     "ubuntu_focal",
		})
		testhelpers.AssertNoError(t, err)

		// ubuntu_focal DEV1-S at par1: 68cf470e-6c35-4741-bbff-4ce788616461
		testhelpers.Equals(t, "9c41e95b-add2-4ef8-b1b1-af8899748eda", image.ID)
	})

	t.Run("matching input for GetLocalImageIDByLabel with lowercase image label", func(t *testing.T) {
		// Create SDK objects for Scaleway Instance product
		marketplaceAPI := marketplace.NewAPI(client)

		image, err := marketplaceAPI.GetLocalImageByLabel(&marketplace.GetLocalImageByLabelRequest{
			Zone:           scw.ZoneFrPar1,
			CommercialType: "dev1-s",
			ImageLabel:     "ubuntu_focal",
		})
		testhelpers.AssertNoError(t, err)

		// ubuntu_focal DEV1-S at par1: 9c41e95b-add2-4ef8-b1b1-af8899748eda
		testhelpers.Equals(t, "9c41e95b-add2-4ef8-b1b1-af8899748eda", image.ID)
	})

	t.Run("non-matching label for GetLocalImageIDByLabel", func(t *testing.T) {
		// Create SDK objects for Scaleway Instance product
		marketplaceAPI := marketplace.NewAPI(client)

		_, err := marketplaceAPI.GetLocalImageByLabel(&marketplace.GetLocalImageByLabelRequest{
			Zone:           scw.ZoneFrPar1,
			CommercialType: "DEV1-S",
			ImageLabel:     "foo-bar-image",
		})
		testhelpers.Assert(t, err != nil, "Should have error")
		// testhelpers.Equals(t, "scaleway-sdk-go: couldn't find a matching image for the given label (foo-bar-image)", err.Error())
	})
}

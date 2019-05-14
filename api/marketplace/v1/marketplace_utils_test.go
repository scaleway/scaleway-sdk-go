package marketplace

import (
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/scaleway/scaleway-sdk-go/utils"
)

func TestGetImageByName(t *testing.T) {

	// Setup recorder and scw client
	r, err := recorder.NewAsMode("fixtures/marketplace", recorder.ModeReplaying, nil)
	if err != nil {
		testhelpers.Ok(t, err)
	}
	defer r.Stop() // Make sure recorder is stopped once done with it

	httpClient := &http.Client{Transport: r}

	client, err := scw.NewClient(
		scw.WithoutAuth(),
		scw.WithHTTPClient(httpClient),
	)
	if err != nil {
		testhelpers.Ok(t, err)
	}

	t.Run("matching input for GetImageByName", func(t *testing.T) {

		// Create SDK objects for Scaleway Instance product
		marketplaceAPI := NewApi(client)

		imageID, err := marketplaceAPI.GetImageByName("Docker", utils.ZoneFrPar1, "C1")
		testhelpers.Ok(t, err)

		// Docker C1 at par1: 45a7e942-1fb0-48c0-bbf6-0acb9af24604
		testhelpers.Equals(t, "45a7e942-1fb0-48c0-bbf6-0acb9af24604", imageID)

	})

	t.Run("non-matching name for GetImageByName", func(t *testing.T) {

		// Create SDK objects for Scaleway Instance product
		marketplaceAPI := NewApi(client)

		_, err := marketplaceAPI.GetImageByName("foo-bar-image", "", "")
		testhelpers.Assert(t, err != nil, "Should have error")

	})

}

package httprecorder

import (
	"net/http"
	"os"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/scaleway/scaleway-sdk-go/vcr"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/recorder"
)

// IsUpdatingCassette returns true if we are updating cassettes.
func IsUpdatingCassette() bool {
	return os.Getenv("SDK_UPDATE_CASSETTES") == "true"
}

// CreateRecordedScwClient creates a new scw.Client that records all HTTP requests in a cassette.
// This cassette is then replayed whenever tests are executed again. This means that once the
// requests are recorded in the cassette, no more real HTTP request must be made to run the tests.
//
// It is important to call add a `defer recorder.Stop()` so the given cassette files are correctly
// closed and saved after the requests.
//
// To update the cassette files, add  `SDK_UPDATE_CASSETTES=true` to the environment variables.
// When updating cassettes, make sure your Scaleway credentials are set in your config or in the
// variables `SCW_ACCESS_KEY` and `SCW_SECRET_KEY`.
func CreateRecordedScwClient(t *testing.T) (*scw.Client, *recorder.Recorder, error) {
	t.Helper()

	UpdateCassette := IsUpdatingCassette()

	pkgFolder, err := os.Getwd()
	if err != nil {
		t.Fatalf("cannot detect working directory for testing")
	}

	// Setup recorder
	r, err := vcr.NewHTTPRecorder(t, pkgFolder, UpdateCassette)
	if err != nil {
		return nil, nil, err
	}

	// Create new http.Client where transport is the recorder
	httpClient := &http.Client{Transport: r}

	var activeProfile *scw.Profile
	var client *scw.Client

	if UpdateCassette {
		// When updating the recorded test requests, we need the access key and secret key.
		config, err := scw.LoadConfig()
		if err != nil {
			activeProfile = &scw.Profile{}
		} else {
			activeProfile, err = config.GetActiveProfile()
			if err != nil {
				return nil, nil, err
			}
		}

		client, err = scw.NewClient(
			scw.WithHTTPClient(httpClient),
			scw.WithProfile(activeProfile),
			scw.WithEnv(),
			scw.WithDefaultRegion(scw.RegionFrPar),
			scw.WithDefaultZone(scw.ZoneFrPar1),
		)
		if err != nil {
			return nil, nil, err
		}
	} else {
		// No need for auth when using cassette
		client, err = scw.NewClient(
			scw.WithHTTPClient(httpClient),
			scw.WithDefaultRegion(scw.RegionFrPar),
			scw.WithDefaultZone(scw.ZoneFrPar1),
		)
		if err != nil {
			return nil, nil, err
		}
	}

	return client, r, nil
}

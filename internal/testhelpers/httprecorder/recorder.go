package httprecorder

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/scw"
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
func CreateRecordedScwClient(cassetteName string) (*scw.Client, *recorder.Recorder, error) {
	UpdateCassette := IsUpdatingCassette()

	var activeProfile *scw.Profile

	recorderMode := recorder.ModeReplaying
	if UpdateCassette {
		recorderMode = recorder.ModeRecording
		config, err := scw.LoadConfig()
		if err != nil {
			return nil, nil, err
		}
		activeProfile, err = config.GetActiveProfile()
		if err != nil {
			return nil, nil, err
		}
	}

	// Setup recorder and scw client
	r, err := recorder.NewAsMode(fmt.Sprintf("testdata/%s", cassetteName), recorderMode, nil)
	if err != nil {
		return nil, nil, err
	}

	// Add a filter which removes Authorization headers from all requests:
	r.AddFilter(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "x-auth-token")
		delete(i.Request.Headers, "X-Auth-Token")

		if UpdateCassette {
			secretKey := *activeProfile.SecretKey
			if i != nil && strings.Contains(fmt.Sprintf("%v", *i), secretKey) {
				panic(errors.New("found secret key in cassette"))
			}
		}

		return nil
	})

	// Create new http.Client where transport is the recorder
	httpClient := &http.Client{Transport: r}

	var client *scw.Client

	if UpdateCassette {
		// When updating the recoreded test requests, we need the access key and secret key.
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

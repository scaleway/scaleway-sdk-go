package testhelpers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/scaleway/scaleway-sdk-go/scwconfig"
)

// CreateRecordedScwClient creates a new scw.Client that records all HTTP requests in a cassette.
// This cassette is then replayed whenever tests are executed again. This means that once the
// requests are recorded in the cassette, no more real HTTP request must be made to run the tests.
//
// It is important to call add a `defer recorder.Stop()` so the given cassette files are correctly
// closed and saved after the requests.
//
// To update the cassette files, add  `UPDATE` to the environment variables.
// When using `UPDATE`, also the `SCW_ACCESS_KEY` and `SCW_SECRET_KEY` must be set.
func CreateRecordedScwClient() (*scw.Client, *recorder.Recorder, error) {

	_, updateCassette := os.LookupEnv("UPDATE")

	recorderMode := recorder.ModeReplaying
	if updateCassette {
		recorderMode = recorder.ModeRecording
	}

	// Setup recorder and scw client
	r, err := recorder.NewAsMode("testdata/go-vcr", recorderMode, nil)
	if err != nil {
		return nil, nil, err
	}

	// Add a filter which removes Authorization headers from all requests:
	r.AddFilter(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "x-auth-token")
		return nil
	})

	// Create new http.Client where transport is the recorder
	httpClient := &http.Client{Transport: r}

	var client *scw.Client

	if updateCassette {
		// When updating the recoreded test requests, we need the access key and secret key.
		config, err := scwconfig.Load()
		if err != nil {
			return nil, nil, err
		}

		accessKey, exists := config.GetAccessKey()
		if !exists {
			return nil, nil, fmt.Errorf("need access key to update cassette")
		}

		secretKey, exists := config.GetSecretKey()
		if !exists {
			return nil, nil, fmt.Errorf("need secret key to update cassette")
		}

		client, err = scw.NewClient(
			scw.WithAuth(accessKey, secretKey),
			scw.WithHTTPClient(httpClient),
		)
	} else {
		// No need for auth when using casette
		client, err = scw.NewClient(
			scw.WithoutAuth(),
			scw.WithHTTPClient(httpClient),
		)
	}

	return client, r, nil

}

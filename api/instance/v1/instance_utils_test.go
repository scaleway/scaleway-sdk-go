package instance

import (
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func TestAttachIP(t *testing.T) {

	// Setup recorder and scw client
	r, err := recorder.NewAsMode("testdata/go-vcr", recorder.ModeReplaying, nil)
	if err != nil {
		testhelpers.Ok(t, err)
	}
	defer r.Stop() // Make sure recorder is stopped once done with it

	// Add a filter which removes Authorization headers from all requests:
	r.AddFilter(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "x-auth-token")
		return nil
	})

	httpClient := &http.Client{Transport: r}

	client, err := scw.NewClient(
		/*
			scw.WithAuth(
				"",
				"",
			),
		*/
		scw.WithoutAuth(),
		scw.WithHTTPClient(httpClient),
	)
	if err != nil {
		testhelpers.Ok(t, err)
	}

	instanceAPI := NewApi(client)

	server := "5ee00b7a-0eba-4f5c-84ce-8f0f32177e2a"

	/*
		_, err = instanceAPI.UpdateIp(&UpdateIpRequest{
			IpId:   "1c91be6a-7f22-42aa-84d9-e83479613db0",
			Zone:   "fr-par-1",
			Server: &server,
		})

		testhelpers.Ok(t, err)
	*/

	ip, err := instanceAPI.AttachIp(&AttachIpRequest{
		IpId:   "1c91be6a-7f22-42aa-84d9-e83479613db0",
		Zone:   "fr-par-1",
		Server: &server,
	})

	testhelpers.Ok(t, err)
	testhelpers.Equals(t, server, ip.Server.Id)

}

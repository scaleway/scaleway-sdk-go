package e2e

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/api/test/v1"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func TestStandardErrors(t *testing.T) {
	client, _, err := newE2EClient(true)
	testhelpers.AssertNoError(t, err)

	_, err = client.GetHuman(&test.GetHumanRequest{
		HumanID: "b3ba839a-dcf2-4b0a-ac81-fc32370052a0",
	})

	testhelpers.Equals(t, &scw.ResourceNotFound{
		Resource:   "human",
		ResourceID: "b3ba839a-dcf2-4b0a-ac81-fc32370052a0",
	}, err)

	_, err = client.CreateHuman(&test.CreateHumanRequest{
		AltitudeInMeter: -7000000,
	})
	testhelpers.Equals(t, &scw.InvalidArgumentsError{
		Details: []struct {
			ArgumentName string `json:"argument_name"`
			Reason       string `json:"reason"`
			HelpMessage  string `json:"help_message"`
		}{
			{
				ArgumentName: "altitude_in_meter",
				Reason:       "constraint",
				HelpMessage:  "lowest altitude on earth is -6371km",
			},
		},
	}, err)

	var human *test.Human
	for i := 0; i < 10; i++ {
		human, err = client.CreateHuman(&test.CreateHumanRequest{})
		testhelpers.AssertNoError(t, err)
	}

	_, err = client.CreateHuman(&test.CreateHumanRequest{})
	testhelpers.Equals(t, &scw.QuotasExceededError{
		Details: []struct {
			Resource string `json:"resource"`
			Quota    uint32 `json:"quota"`
			Current  uint32 `json:"current"`
		}{
			{
				Resource: "human",
				Quota:    10,
				Current:  10,
			},
		},
	}, err)

	_, err = client.RunHuman(&test.RunHumanRequest{HumanID: human.ID})
	testhelpers.AssertNoError(t, err)

	_, err = client.UpdateHuman(&test.UpdateHumanRequest{HumanID: human.ID})
	testhelpers.Equals(t, &scw.TransientStateError{
		Resource:     "human",
		ResourceID:   human.ID,
		CurrentState: "running",
	}, err)

}

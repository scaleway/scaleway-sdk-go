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
		RawBody:    []byte(`{"message":"resource is not found","resource":"human","resource_id":"b3ba839a-dcf2-4b0a-ac81-fc32370052a0","type":"not_found"}`),
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
		RawBody: []byte(`{"details":[{"argument_name":"altitude_in_meter","help_message":"lowest altitude on earth is -6371km","reason":"constraint"}],"message":"invalid argument(s)","type":"invalid_arguments"}`),
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
		RawBody: []byte(`{"details":[{"current":10,"quota":10,"resource":"human"}],"message":"quota(s) exceeded for this resource","type":"quotas_exceeded"}`),
	}, err)

	_, err = client.RunHuman(&test.RunHumanRequest{HumanID: human.ID})
	testhelpers.AssertNoError(t, err)

	_, err = client.UpdateHuman(&test.UpdateHumanRequest{HumanID: human.ID})

	testhelpers.Equals(t, &scw.TransientStateError{
		Resource:     "human",
		ResourceID:   human.ID,
		CurrentState: "running",
		RawBody:      []byte(`{"current_state":"running","message":"resource is in a transient state","resource":"human","resource_id":"` + human.ID + `","type":"transient_state"}`),
	}, err)

}

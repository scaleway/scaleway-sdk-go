package e2e

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/api/test/v1"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func TestStandardErrors(t *testing.T) {
	client, _, _, err := newE2EClient(true)
	testhelpers.AssertNoError(t, err)

	t.Run("not found", func(t *testing.T) {
		_, err = client.GetHuman(&test.GetHumanRequest{
			HumanID: "b3ba839a-dcf2-4b0a-ac81-fc32370052a0",
		})
		testhelpers.Equals(t, &scw.ResourceNotFoundError{
			Resource:   "human",
			ResourceID: "b3ba839a-dcf2-4b0a-ac81-fc32370052a0",
			RawBody:    []byte(`{"message":"resource is not found","resource":"human","resource_id":"b3ba839a-dcf2-4b0a-ac81-fc32370052a0","type":"not_found"}`),
		}, err)
	})

	t.Run("invalid argument", func(t *testing.T) {
		_, err = client.CreateHuman(&test.CreateHumanRequest{
			AltitudeInMeter: -7000000,
		})
		testhelpers.Equals(t, &scw.InvalidArgumentsError{
			Details: []scw.InvalidArgumentsErrorDetail{
				{
					ArgumentName: "altitude_in_meter",
					Reason:       "constraint",
					HelpMessage:  "lowest altitude on earth is -6371km",
				},
			},
			RawBody: []byte(`{"details":[{"argument_name":"altitude_in_meter","help_message":"lowest altitude on earth is -6371km","reason":"constraint"}],"message":"invalid argument(s)","type":"invalid_arguments"}`),
		}, err)
	})

	t.Run("quotas exceeded", func(t *testing.T) {
		var humans []*test.Human

		for i := 0; i < 10; i++ {
			human, err := client.CreateHuman(&test.CreateHumanRequest{})
			testhelpers.AssertNoError(t, err)
			humans = append(humans, human)
		}

		_, err = client.CreateHuman(&test.CreateHumanRequest{})
		testhelpers.Equals(t, &scw.QuotasExceededError{
			Details: []scw.QuotasExceededErrorDetail{
				{
					Resource: "human",
					Quota:    10,
					Current:  10,
				},
			},
			RawBody: []byte(`{"details":[{"current":10,"quota":10,"resource":"human"}],"message":"quota(s) exceeded for this resource","type":"quotas_exceeded"}`),
		}, err)

		for _, human := range humans {
			_, err := client.DeleteHuman(&test.DeleteHumanRequest{HumanID: human.ID})
			testhelpers.AssertNoError(t, err)
		}
	})

	t.Run("transient state", func(t *testing.T) {
		human, err := client.CreateHuman(&test.CreateHumanRequest{})
		testhelpers.AssertNoError(t, err)
		defer func() {
			_, _ = client.DeleteHuman(&test.DeleteHumanRequest{HumanID: human.ID})
		}()

		_, err = client.RunHuman(&test.RunHumanRequest{HumanID: human.ID})
		testhelpers.AssertNoError(t, err)

		_, err = client.UpdateHuman(&test.UpdateHumanRequest{HumanID: human.ID})
		testhelpers.Equals(t, &scw.TransientStateError{
			Resource:     "human",
			ResourceID:   human.ID,
			CurrentState: "running",
			RawBody:      []byte(`{"current_state":"running","message":"resource is in a transient state","resource":"human","resource_id":"` + human.ID + `","type":"transient_state"}`),
		}, err)
	})

	t.Run("out of stock", func(t *testing.T) {
		_, err = client.CreateHuman(&test.CreateHumanRequest{
			ShoeSize: 60,
		})
		testhelpers.Equals(t, &scw.OutOfStockError{
			Resource: "ShoeSize60",
			RawBody:  []byte(`{"message":"resource is out of stock","resource":"ShoeSize60","type":"out_of_stock"}`),
		}, err)
	})
}

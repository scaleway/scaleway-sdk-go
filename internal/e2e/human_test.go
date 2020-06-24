package e2e

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/api/test/v1"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

var (
	defaultOrganizationID = "6170692e-7363-616c-6577-61792e636f6d" // hint: | xxd -ps -r
)

func newE2EClient(withAuthInClient bool) (*test.API, string, string, error) {
	client, err := scw.NewClient(
		scw.WithDefaultRegion(scw.RegionFrPar),
		scw.WithDefaultOrganizationID(defaultOrganizationID),
		scw.WithUserAgent("sdk-e2e-test"),
	)
	if err != nil {
		return nil, "", "", err
	}
	testClient := test.NewAPI(client)

	registerResponse, err := testClient.Register(&test.RegisterRequest{
		Username: "sidi",
	})
	if err != nil {
		return nil, "", "", err
	}
	if withAuthInClient {
		client, err = scw.NewClient(
			scw.WithDefaultRegion(scw.RegionFrPar),
			scw.WithDefaultOrganizationID(defaultOrganizationID),
			scw.WithAuth(registerResponse.AccessKey, registerResponse.SecretKey),
			scw.WithUserAgent("sdk-e2e-test"),
		)
		testClient = test.NewAPI(client)
	}

	return testClient, registerResponse.AccessKey, registerResponse.SecretKey, err
}

func TestAuthInRequest(t *testing.T) {
	client, accessKey, secretKey, err := newE2EClient(false)
	testhelpers.AssertNoError(t, err)

	requestWithAuth := scw.WithAuthRequest(accessKey, secretKey)
	_, err = client.CreateHuman(&test.CreateHumanRequest{}, requestWithAuth)
	testhelpers.AssertNoError(t, err)
}

func TestHuman(t *testing.T) {
	client, _, _, err := newE2EClient(true)
	testhelpers.AssertNoError(t, err)

	// create
	human, err := client.CreateHuman(&test.CreateHumanRequest{
		Height:               170.5,
		ShoeSize:             35.1,
		AltitudeInMeter:      -12,
		AltitudeInMillimeter: -12050,
		FingersCount:         21,
		HairCount:            9223372036854775808,
		IsHappy:              true,
		EyesColor:            test.EyeColorsAmber,
		OrganizationID:       scw.StringPtr("b3ba839a-dcf2-4b0a-ac81-fc32370052a0"),
	})

	testhelpers.AssertNoError(t, err)
	testhelpers.Equals(t, human.Height, 170.5)
	testhelpers.Equals(t, human.ShoeSize, float32(35.1))
	testhelpers.Equals(t, human.AltitudeInMeter, int32(-12))
	testhelpers.Equals(t, human.AltitudeInMillimeter, int64(-12050))
	testhelpers.Equals(t, human.FingersCount, uint32(21))
	testhelpers.Equals(t, human.HairCount, uint64(9223372036854775808))
	testhelpers.Equals(t, human.IsHappy, true)
	testhelpers.Equals(t, human.EyesColor, test.EyeColorsAmber)
	testhelpers.Equals(t, human.OrganizationID, "b3ba839a-dcf2-4b0a-ac81-fc32370052a0")

	// single parameter update
	human, err = client.UpdateHuman(&test.UpdateHumanRequest{
		HumanID: human.ID,
		IsHappy: scw.BoolPtr(false),
	})

	testhelpers.AssertNoError(t, err)
	testhelpers.Equals(t, human.Height, 170.5)
	testhelpers.Equals(t, human.ShoeSize, float32(35.1))
	testhelpers.Equals(t, human.AltitudeInMeter, int32(-12))
	testhelpers.Equals(t, human.AltitudeInMillimeter, int64(-12050))
	testhelpers.Equals(t, human.FingersCount, uint32(21))
	testhelpers.Equals(t, human.HairCount, uint64(9223372036854775808))
	testhelpers.Equals(t, human.IsHappy, false)
	testhelpers.Equals(t, human.EyesColor, test.EyeColorsAmber)
	testhelpers.Equals(t, human.OrganizationID, "b3ba839a-dcf2-4b0a-ac81-fc32370052a0")

	// update
	human, err = client.UpdateHuman(&test.UpdateHumanRequest{
		HumanID:              human.ID,
		Height:               scw.Float64Ptr(155.666),
		ShoeSize:             scw.Float32Ptr(36.0),
		AltitudeInMeter:      scw.Int32Ptr(2147483647),
		AltitudeInMillimeter: scw.Int64Ptr(2147483647285),
		FingersCount:         scw.Uint32Ptr(20),
		HairCount:            scw.Uint64Ptr(9223372036854775809),
		IsHappy:              scw.BoolPtr(true),
		EyesColor:            test.EyeColorsBlue,
	})

	testhelpers.AssertNoError(t, err)
	testhelpers.Equals(t, human.Height, 155.666)
	testhelpers.Equals(t, human.ShoeSize, float32(36.0))
	testhelpers.Equals(t, human.AltitudeInMeter, int32(2147483647))
	testhelpers.Equals(t, human.AltitudeInMillimeter, int64(2147483647285))
	testhelpers.Equals(t, human.FingersCount, uint32(20))
	testhelpers.Equals(t, human.HairCount, uint64(9223372036854775809))
	testhelpers.Equals(t, human.IsHappy, true)
	testhelpers.Equals(t, human.EyesColor, test.EyeColorsBlue)
	testhelpers.Equals(t, human.OrganizationID, "b3ba839a-dcf2-4b0a-ac81-fc32370052a0")

	// get
	human, err = client.GetHuman(&test.GetHumanRequest{
		HumanID: human.ID,
	})

	testhelpers.AssertNoError(t, err)
	testhelpers.Equals(t, human.Height, 155.666)
	testhelpers.Equals(t, human.ShoeSize, float32(36.0))
	testhelpers.Equals(t, human.AltitudeInMeter, int32(2147483647))
	testhelpers.Equals(t, human.AltitudeInMillimeter, int64(2147483647285))
	testhelpers.Equals(t, human.FingersCount, uint32(20))
	testhelpers.Equals(t, human.HairCount, uint64(9223372036854775809))
	testhelpers.Equals(t, human.IsHappy, true)
	testhelpers.Equals(t, human.EyesColor, test.EyeColorsBlue)
	testhelpers.Equals(t, human.OrganizationID, "b3ba839a-dcf2-4b0a-ac81-fc32370052a0")

	// delete
	_, err = client.DeleteHuman(&test.DeleteHumanRequest{
		HumanID: human.ID,
	})
	testhelpers.AssertNoError(t, err)
}

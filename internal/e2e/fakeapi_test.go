package e2e

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/api/fakeapi/v1"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func newE2EClient(withAuthInClient bool) (*fakeapi.API, string, error) {
	client, err := scw.NewClient(
		scw.WithoutAuth(),
		scw.WithDefaultRegion(scw.RegionFrPar),
	)
	if err != nil {
		return nil, "", err
	}
	fakeAPI := fakeapi.NewAPI(client)

	registerResponse, err := fakeAPI.Register(&fakeapi.RegisterRequest{
		Username: "sidi",
	})
	if err != nil {
		return nil, "", err
	}
	if withAuthInClient {
		client, err = scw.NewClient(
			scw.WithDefaultRegion(scw.RegionFrPar),
			scw.WithAuth("", registerResponse.SecretKey),
		)
		fakeAPI = fakeapi.NewAPI(client)
	}

	return fakeAPI, registerResponse.SecretKey, err
}

func TestAuthInRequest(t *testing.T) {
	client, secretKey, err := newE2EClient(false)
	testhelpers.AssertNoError(t, err)

	requestWithAuth := scw.WithAuthRequest("", secretKey)
	_, err = client.CreateHuman(&fakeapi.CreateHumanRequest{}, requestWithAuth)
	testhelpers.AssertNoError(t, err)
}

func TestHuman(t *testing.T) {
	client, _, err := newE2EClient(true)
	testhelpers.AssertNoError(t, err)

	// create
	human, err := client.CreateHuman(&fakeapi.CreateHumanRequest{
		Height:               170.5,
		ShoeSize:             35.1,
		AltitudeInMeter:      -12,
		AltitudeInMillimeter: -12050,
		FingersCount:         21,
		HairCount:            9223372036854775808,
		IsHappy:              true,
		EyesColor:            fakeapi.EyeColorsAmber,
		ProjectID:            "b3ba839a-dcf2-4b0a-ac81-fc32370052a0",
	})

	testhelpers.AssertNoError(t, err)
	testhelpers.Equals(t, human.Height, 170.5)
	testhelpers.Equals(t, human.ShoeSize, float32(35.1))
	testhelpers.Equals(t, human.AltitudeInMeter, int32(-12))
	testhelpers.Equals(t, human.AltitudeInMillimeter, int64(-12050))
	testhelpers.Equals(t, human.FingersCount, uint32(21))
	testhelpers.Equals(t, human.HairCount, uint64(9223372036854775808))
	testhelpers.Equals(t, human.IsHappy, true)
	testhelpers.Equals(t, human.EyesColor, fakeapi.EyeColorsAmber)
	testhelpers.Equals(t, human.ProjectID, "b3ba839a-dcf2-4b0a-ac81-fc32370052a0")

	// single parameter update
	human, err = client.UpdateHuman(&fakeapi.UpdateHumanRequest{
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
	testhelpers.Equals(t, human.EyesColor, fakeapi.EyeColorsAmber)
	testhelpers.Equals(t, human.ProjectID, "b3ba839a-dcf2-4b0a-ac81-fc32370052a0")

	// update
	human, err = client.UpdateHuman(&fakeapi.UpdateHumanRequest{
		HumanID:              human.ID,
		Height:               scw.Float64Ptr(155.666),
		ShoeSize:             scw.Float32Ptr(36.0),
		AltitudeInMeter:      scw.Int32Ptr(2147483647),
		AltitudeInMillimeter: scw.Int64Ptr(2147483647285),
		FingersCount:         scw.Uint32Ptr(20),
		HairCount:            scw.Uint64Ptr(9223372036854775809),
		IsHappy:              scw.BoolPtr(true),
		EyesColor:            fakeapi.EyeColorsBlue,
	})

	testhelpers.AssertNoError(t, err)
	testhelpers.Equals(t, human.Height, 155.666)
	testhelpers.Equals(t, human.ShoeSize, float32(36.0))
	testhelpers.Equals(t, human.AltitudeInMeter, int32(2147483647))
	testhelpers.Equals(t, human.AltitudeInMillimeter, int64(2147483647285))
	testhelpers.Equals(t, human.FingersCount, uint32(20))
	testhelpers.Equals(t, human.HairCount, uint64(9223372036854775809))
	testhelpers.Equals(t, human.IsHappy, true)
	testhelpers.Equals(t, human.EyesColor, fakeapi.EyeColorsBlue)
	testhelpers.Equals(t, human.ProjectID, "b3ba839a-dcf2-4b0a-ac81-fc32370052a0")

	// get
	human, err = client.GetHuman(&fakeapi.GetHumanRequest{
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
	testhelpers.Equals(t, human.EyesColor, fakeapi.EyeColorsBlue)
	testhelpers.Equals(t, human.ProjectID, "b3ba839a-dcf2-4b0a-ac81-fc32370052a0")

	// delete
	_, err = client.DeleteHuman(&fakeapi.DeleteHumanRequest{
		HumanID: human.ID,
	})
	testhelpers.AssertNoError(t, err)

	// get
	_, err = client.GetHuman(&fakeapi.GetHumanRequest{
		HumanID: human.ID,
	})
	testhelpers.Equals(t, "scaleway-sdk-go: http error 404 Not Found: human not found", err.Error())
}

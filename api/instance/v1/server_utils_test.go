package instance_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers/httprecorder"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func TestAPI_GetServerType(t *testing.T) {
	client, r, err := httprecorder.CreateRecordedScwClient("get-server-type")
	testhelpers.AssertNoError(t, err)
	defer func() {
		testhelpers.AssertNoError(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	instanceAPI := instance.NewAPI(client)

	serverType, err := instanceAPI.GetServerType(&instance.GetServerTypeRequest{
		Zone: scw.ZoneFrPar1,
		Name: "GP1-XS",
	})

	testhelpers.AssertNoError(t, err)
	testhelpers.Equals(t, 1*scw.GB, serverType.PerVolumeConstraint.LSSD.MinSize)
	testhelpers.Equals(t, 800*scw.GB, serverType.PerVolumeConstraint.LSSD.MaxSize)
}

func TestAPI_ServerUserData(t *testing.T) {
	client, r, err := httprecorder.CreateRecordedScwClient("server-user-data")
	testhelpers.AssertNoError(t, err)
	defer func() {
		testhelpers.AssertNoError(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	instanceAPI := instance.NewAPI(client)

	key := "hello"
	contentStr := "world"

	serverRes, err := instanceAPI.CreateServer(&instance.CreateServerRequest{
		Zone:           scw.ZoneFrPar1,
		CommercialType: "DEV1-S",
		Name:           namegenerator.GetRandomName("srv"),
		Image:          scw.StringPtr("f974feac-abae-4365-b988-8ec7d1cec10d"),
		Project:        scw.StringPtr("14d2f7ae-9775-414c-9bed-6810e060d500"),
	})
	testhelpers.AssertNoError(t, err)

	content := strings.NewReader(contentStr)
	err = instanceAPI.SetServerUserData(&instance.SetServerUserDataRequest{
		Zone:     scw.ZoneFrPar1,
		ServerID: serverRes.Server.ID,
		Key:      key,
		Content:  content,
	})
	testhelpers.AssertNoError(t, err)

	data, err := instanceAPI.GetServerUserData(&instance.GetServerUserDataRequest{
		Zone:     scw.ZoneFrPar1,
		ServerID: serverRes.Server.ID,
		Key:      key,
	})
	testhelpers.AssertNoError(t, err)

	resUserData, err := io.ReadAll(data)
	testhelpers.AssertNoError(t, err)
	testhelpers.Equals(t, contentStr, string(resUserData))
}

func TestAPI_AllServerUserData(t *testing.T) {
	client, r, err := httprecorder.CreateRecordedScwClient("all-server-user-data")
	testhelpers.AssertNoError(t, err)
	defer func() {
		testhelpers.AssertNoError(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	instanceAPI := instance.NewAPI(client)

	serverRes, err := instanceAPI.CreateServer(&instance.CreateServerRequest{
		Zone:           scw.ZoneFrPar1,
		CommercialType: "DEV1-S",
		Name:           namegenerator.GetRandomName("srv"),
		Image:          scw.StringPtr("f974feac-abae-4365-b988-8ec7d1cec10d"),
		Project:        scw.StringPtr("14d2f7ae-9775-414c-9bed-6810e060d500"),
	})
	testhelpers.AssertNoError(t, err)

	steps := []map[string]string{
		{
			"hello":      "world",
			"scale":      "way",
			"tic":        "tac",
			"cloud-init": "on",
		},
		{
			"scale":      "way",
			"steve":      "wozniak",
			"cloud-init": "off",
		},
		{},
	}

	for _, data := range steps {
		// create user data
		userData := make(map[string]io.Reader, len(data))
		for k, v := range data {
			userData[k] = bytes.NewBufferString(v)
		}

		// set all user data
		err := instanceAPI.SetAllServerUserData(&instance.SetAllServerUserDataRequest{
			Zone:     scw.ZoneFrPar1,
			ServerID: serverRes.Server.ID,
			UserData: userData,
		})
		testhelpers.AssertNoError(t, err)

		// get all user data
		allData, err := instanceAPI.GetAllServerUserData(&instance.GetAllServerUserDataRequest{
			Zone:     scw.ZoneFrPar1,
			ServerID: serverRes.Server.ID,
		})
		testhelpers.AssertNoError(t, err)

		testhelpers.Equals(t, len(data), len(allData.UserData))

		for expectedKey, expectedValue := range data {
			currentReader, exists := allData.UserData[expectedKey]
			testhelpers.Assert(t, exists, "%s key not found in result", expectedKey)

			currentValue, err := io.ReadAll(currentReader)
			testhelpers.AssertNoError(t, err)

			testhelpers.Equals(t, expectedValue, string(currentValue))
		}
	}
}

func TestAPI_CreateServer(t *testing.T) {
	client, r, err := httprecorder.CreateRecordedScwClient("create-server")
	testhelpers.AssertNoError(t, err)
	defer func() {
		testhelpers.AssertNoError(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	instanceAPI := instance.NewAPI(client)

	res, err := instanceAPI.CreateServer(&instance.CreateServerRequest{
		Zone:           scw.ZoneFrPar1,
		CommercialType: "GP1-XS",
		Image:          scw.StringPtr("ubuntu_focal"),
	})

	testhelpers.AssertNoError(t, err)
	// this UUID might change when running the cassette later when the image "ubuntu_focal" got a new version
	testhelpers.Equals(t, "9c41e95b-add2-4ef8-b1b1-af8899748eda", res.Server.Image.ID)
	err = instanceAPI.DeleteServer(&instance.DeleteServerRequest{
		Zone:     scw.ZoneFrPar1,
		ServerID: res.Server.ID,
	})
	testhelpers.AssertNoError(t, err)
}

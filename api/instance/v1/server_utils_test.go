package instance

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers/httprecorder"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
	"github.com/scaleway/scaleway-sdk-go/utils"
)

func TestAPI_GetServerType(t *testing.T) {
	client, r, err := httprecorder.CreateRecordedScwClient("get-server-type")
	testhelpers.Ok(t, err)
	defer func() {
		testhelpers.Ok(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	instanceAPI := NewAPI(client)

	serverType, err := instanceAPI.GetServerType(&GetServerTypeRequest{
		Zone: utils.ZoneFrPar1,
		Name: "GP1-XS",
	})

	testhelpers.Ok(t, err)
	testhelpers.Equals(t, uint64(1000000000), serverType.PerVolumeConstraint.LSSD.MinSize)
	testhelpers.Equals(t, uint64(800000000000), serverType.PerVolumeConstraint.LSSD.MaxSize)

}

func TestAPI_ServerUserData(t *testing.T) {
	client, r, err := httprecorder.CreateRecordedScwClient("server-user-data")
	testhelpers.Ok(t, err)
	defer func() {
		testhelpers.Ok(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	instanceAPI := NewAPI(client)

	key := "hello"
	contentStr := "world"

	serverRes, err := instanceAPI.CreateServer(&CreateServerRequest{
		Zone:           utils.ZoneFrPar1,
		CommercialType: "DEV1-S",
		Name:           namegenerator.GetRandomName("srv"),
		BootType:       ServerBootTypeLocal,
		Image:          "f974feac-abae-4365-b988-8ec7d1cec10d",
		Organization:   "14d2f7ae-9775-414c-9bed-6810e060d500",
	})
	testhelpers.Ok(t, err)

	content := strings.NewReader(contentStr)
	err = instanceAPI.SetServerUserData(&SetServerUserDataRequest{
		Zone:     utils.ZoneFrPar1,
		ServerID: serverRes.Server.ID,
		Key:      key,
		Content:  content,
	})
	testhelpers.Ok(t, err)

	data, err := instanceAPI.GetServerUserData(&GetServerUserDataRequest{
		Zone:     utils.ZoneFrPar1,
		ServerID: serverRes.Server.ID,
		Key:      key,
	})
	testhelpers.Ok(t, err)

	resUserData, err := ioutil.ReadAll(data)
	testhelpers.Ok(t, err)
	testhelpers.Equals(t, contentStr, string(resUserData))
}

func TestAPI_GetAllServerUserData(t *testing.T) {
	client, r, err := httprecorder.CreateRecordedScwClient("get-all-server-user-data")
	testhelpers.Ok(t, err)
	defer func() {
		testhelpers.Ok(t, r.Stop()) // Make sure recorder is stopped once done with it
	}()

	instanceAPI := NewAPI(client)

	serverRes, err := instanceAPI.CreateServer(&CreateServerRequest{
		Zone:           utils.ZoneFrPar1,
		CommercialType: "DEV1-S",
		Name:           namegenerator.GetRandomName("srv"),
		BootType:       ServerBootTypeLocal,
		Image:          "f974feac-abae-4365-b988-8ec7d1cec10d",
		Organization:   "14d2f7ae-9775-414c-9bed-6810e060d500",
	})
	testhelpers.Ok(t, err)

	data := map[string]string{
		"hello":  "world",
		"scale":  "way",
		"xavier": "niel",
	}

	for k, v := range data {
		err = instanceAPI.SetServerUserData(&SetServerUserDataRequest{
			Zone:     utils.ZoneFrPar1,
			ServerID: serverRes.Server.ID,
			Key:      k,
			Content:  strings.NewReader(v),
		})
		testhelpers.Ok(t, err)
	}

	allData, err := instanceAPI.GetAllServerUserData(&GetAllServerUserDataRequest{
		Zone:     utils.ZoneFrPar1,
		ServerID: serverRes.Server.ID,
	})
	testhelpers.Ok(t, err)

	testhelpers.Equals(t, len(data), len(allData.UserData))

	for expectedKey, expectedValue := range data {
		readerValue, exists := allData.UserData[expectedKey]
		testhelpers.Assert(t, exists, "%s key not found in result", expectedKey)

		currentValue, err := ioutil.ReadAll(readerValue)
		testhelpers.Ok(t, err)

		testhelpers.Equals(t, expectedValue, string(currentValue))
	}
}

package scalewaysdkgo

import (
	"fmt"

	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Example_apiClient() {

	client, err := scw.NewClient(
		scw.WithAuth("ACCESS_KEY", "SECRET_KEY"), // Get your credentials at https://console.scaleway.com/account/credentials
	)
	if err != nil {
		panic(err)
	}

	instanceApi := instance.NewApi(client)

	response, err := instanceApi.ListServers(&instance.ListServersRequest{
		Zone: scw.ZoneFrPar1,
	})
	if err != nil {
		panic(err)
	}

	for _, server := range response.Servers {
		fmt.Println("Server", server.Id, server.Name)
	}

}

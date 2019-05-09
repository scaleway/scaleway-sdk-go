package scalewaysdkgo

import (
	"fmt"

	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/api/lb/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Example_apiClient() {

	// Create a Scaleway client
	client, err := scw.NewClient(
		scw.WithAuth("ACCESS_KEY", "SECRET_KEY"), // Get your credentials at https://console.scaleway.com/account/credentials
	)
	if err != nil {
		// handle error
	}

	// Create SDK objects for specific Scaleway Products
	instance := instance.NewApi(client)
	lb := lb.NewApi(client)

	// Start using the SDKs
	_, _ = instance, lb

}

func Example_listServers() {

	// Create a Scaleway client
	client, err := scw.NewClient(
		scw.WithAuth("ACCESS_KEY", "SECRET_KEY"), // Get your credentials at https://console.scaleway.com/account/credentials
	)
	if err != nil {
		// handle error
	}

	// Create SDK objects for Scaleway Instance product
	instanceAPI := instance.NewApi(client)

	// Call the ListServers method on the Instance SDK
	response, err := instanceAPI.ListServers(&instance.ListServersRequest{
		Zone: scw.ZoneFrPar1,
	})
	if err != nil {
		// handle error
	}

	// Do something with the response...
	fmt.Println(response)
}

func Example_createLoadBalancer() {

	// Create a Scaleway client
	client, err := scw.NewClient(
		scw.WithAuth("ACCESS_KEY", "SECRET_KEY"), // Get your credentials at https://console.scaleway.com/account/credentials
	)
	if err != nil {
		// handle error
	}

	// Create SDK objects for Scaleway Load Balancer product
	lbAPI := lb.NewApi(client)

	// Call the CreateLb method on the LB SDK to create a new load balancer.
	newLb, err := lbAPI.CreateLb(&lb.CreateLbRequest{
		Name:           "My new load balancer",
		Description:    "This is a example of a load balancer",
		OrganizationId: "000a115d-2852-4b0a-9ce8-47f1134ba95a",
		Region:         scw.RegionFrPar,
	})

	if err != nil {
		// handle error
	}

	// Do something with the newly created LB...
	fmt.Println(newLb)

}

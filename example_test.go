package scalewaysdkgo

import (
	"fmt"
	"time"

	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/api/lb/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Example_apiClient() {
	// Create a Scaleway client
	client, err := scw.NewClient(
		scw.WithAuth("ACCESS_KEY", "SECRET_KEY"), // Get your credentials at https://console.scaleway.com/project/credentials
	)
	if err != nil {
		// handle error
	}

	// Create SDK objects for specific Scaleway Products
	instanceAPI := instance.NewAPI(client)
	lbAPI := lb.NewAPI(client)

	// Start using the SDKs
	_, _ = instanceAPI, lbAPI
}

func Example_apiClientWithConfig() {
	// Get Scaleway Config
	config, err := scw.LoadConfig()
	if err != nil {
		// handle error
	}

	// Use active profile
	profile, err := config.GetActiveProfile()
	if err != nil {
		// handle error
	}

	// Create a Scaleway client
	client, err := scw.NewClient(
		scw.WithProfile(profile),
		scw.WithEnv(), // env variable may overwrite profile values
	)
	if err != nil {
		// handle error
	}

	// Create SDK objects for specific Scaleway Products
	instanceAPI := instance.NewAPI(client)
	lbAPI := lb.NewAPI(client)

	// Start using the SDKs
	_, _ = instanceAPI, lbAPI
}

func Example_listServers() {
	// Create a Scaleway client
	client, err := scw.NewClient(
		scw.WithAuth("ACCESS_KEY", "SECRET_KEY"), // Get your credentials at https://console.scaleway.com/project/credentials
	)
	if err != nil {
		// handle error
	}

	// Create SDK objects for Scaleway Instance product
	instanceAPI := instance.NewAPI(client)

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

func Example_listServersWithZones() {
	// Create a Scaleway client
	client, err := scw.NewClient(
		scw.WithAuth("ACCESS_KEY", "SECRET_KEY"), // Get your credentials at https://console.scaleway.com/project/credentials
	)
	if err != nil {
		// handle error
	}

	// Create SDK objects for Scaleway Instance product
	instanceAPI := instance.NewAPI(client)

	// Call the ListServers method on the Instance SDK
	response, err := instanceAPI.ListServers(&instance.ListServersRequest{},
		// Add WithZones option to list servers from multiple zones
		scw.WithZones(scw.ZoneFrPar1, scw.ZoneNlAms1, scw.ZonePlWaw1))
	if err != nil {
		// handle error
	}

	// Do something with the response...
	fmt.Println(response)
}

func Example_createServer() {
	// Create a Scaleway client
	client, err := scw.NewClient(
		scw.WithAuth("ACCESS_KEY", "SECRET_KEY"), // Get your credentials at https://console.scaleway.com/project/credentials
		scw.WithDefaultOrganizationID("ORGANIZATION_ID"),
		scw.WithDefaultZone(scw.ZoneFrPar1),
	)
	if err != nil {
		panic(err)
	}

	// Create SDK objects for Scaleway Instance and marketplace
	instanceAPI := instance.NewAPI(client)

	serverType := "DEV1-S"
	image := "ubuntu_focal"

	// Create a new DEV1-S server
	createRes, err := instanceAPI.CreateServer(&instance.CreateServerRequest{
		Name:              "my-server-01",
		CommercialType:    serverType,
		Image:             image,
		DynamicIPRequired: scw.BoolPtr(true),
	})
	if err != nil {
		panic(err)
	}

	// Start the server and wait until it's ready.
	timeout := 5 * time.Minute
	err = instanceAPI.ServerActionAndWait(&instance.ServerActionAndWaitRequest{
		ServerID: createRes.Server.ID,
		Action:   instance.ServerActionPoweron,
		Timeout:  &timeout,
	})
	if err != nil {
		panic(err)
	}
}

func Example_rebootAllServers() {
	// Create a Scaleway client
	client, err := scw.NewClient(
		scw.WithAuth("ACCESS_KEY", "SECRET_KEY"), // Get your credentials at https://console.scaleway.com/project/credentials
		scw.WithDefaultZone(scw.ZoneFrPar1),
	)
	if err != nil {
		panic(err)
	}

	// Create SDK objects for Scaleway Instance product
	instanceAPI := instance.NewAPI(client)

	// Call the ListServers method of the Instance SDK
	response, err := instanceAPI.ListServers(&instance.ListServersRequest{})
	if err != nil {
		panic(err)
	}

	// For each server if they are running we reboot them using ServerActionAndWait
	timeout := 5 * time.Minute
	for _, server := range response.Servers {
		if server.State == instance.ServerStateRunning {
			fmt.Println("Rebooting server with ID", server.ID)
			err = instanceAPI.ServerActionAndWait(&instance.ServerActionAndWaitRequest{
				ServerID: server.ID,
				Action:   instance.ServerActionReboot,
				Timeout:  &timeout,
			})
			if err != nil {
				panic(err)
			}
		}
	}
	fmt.Println("All servers were successfully rebooted")
}

func Example_createLoadBalancer() {
	// Create a Scaleway client
	client, err := scw.NewClient(
		scw.WithAuth("ACCESS_KEY", "SECRET_KEY"), // Get your credentials at https://console.scaleway.com/project/credentials
	)
	if err != nil {
		// handle error
	}

	// Create SDK objects for Scaleway LoadConfig Balancer product
	lbAPI := lb.NewAPI(client)

	// Call the CreateLb method on the LB SDK to create a new load balancer.
	newLB, err := lbAPI.CreateLB(&lb.CreateLBRequest{
		Name:           "My new load balancer",
		Description:    "This is a example of a load balancer",
		OrganizationID: scw.StringPtr("000a115d-2852-4b0a-9ce8-47f1134ba95a"),
		Region:         scw.RegionFrPar,
	})

	if err != nil {
		// handle error
	}

	// Do something with the newly created LB...
	fmt.Println(newLB)
}

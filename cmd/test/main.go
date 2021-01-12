package main

import (
	"fmt"

	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func orPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	logger.EnableDebugMode()

	config, err := scw.LoadConfig()
	orPanic(err)
	profile, err := config.GetActiveProfile()
	orPanic(err)

	client, err := scw.NewClient(scw.WithProfile(profile))
	orPanic(err)
	api := instance.NewAPI(client)

	res, err := api.ListServers(&instance.ListServersRequest{}, scw.WithZones(scw.ZoneFrPar1, scw.ZoneNlAms1))
	orPanic(err)
	fmt.Printf("%#v", res)
}

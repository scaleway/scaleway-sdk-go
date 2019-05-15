# Scaleway GO SDK

[![GoDoc](https://godoc.org/github.com/scaleway/scaleway-sdk-go?status.svg)](https://godoc.org/github.com/scaleway/scaleway-sdk-go)
[![GoReportCard](https://goreportcard.com/badge/scaleway/scaleway-sdk-go)](https://goreportcard.com/report/github.com/scaleway/scaleway-sdk-go)
[![CircleCI](https://circleci.com/gh/scaleway/scaleway-sdk-go.svg?style=svg)](https://circleci.com/gh/scaleway/scaleway-sdk-go)

**Warning: This is an early release, keep in mind that the API can break**

<p align="center">
  <img height="200" src="docs/static_files/scaleway-logo.png">
</p>

Scaleway is an Iliad Group brand supplying a range of pioneering cloud infrastructure covering a full range of services for professionals: public cloud services with Scaleway, private infrastructure with Scaleway Datacenter and bare-metal cloud services with Online by Scaleway.

## Documentation

- [Godoc](https://godoc.org/github.com/scaleway/scaleway-sdk-go)
- [Developer website](https://developers.scaleway.com) (API documentation)

## Installation

```bash
go get github.com/scaleway/scaleway-sdk-go
```

## Getting Started

```go
package main

import (
	"fmt"

	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/scaleway/scaleway-sdk-go/utils"
)

func main() {

	// Create a Scaleway client
	client, err := scw.NewClient(
		// Get your credentials at https://console.scaleway.com/account/credentials
		scw.WithDefaultOrganizationID("ORGANISATION_ID"),
		scw.WithAuth("ACCESS_KEY", "SECRET_KEY"),
	)
	if err != nil {
		panic(err)
	}

	// Create SDK objects for Scaleway Instance product
	instanceApi := instance.NewApi(client)

	// Call the ListServers method on the Instance SDK
	response, err := instanceApi.ListServers(&instance.ListServersRequest{
		Zone: utils.ZoneFrPar1,
	})
	if err != nil {
		panic(err)
	}

	// Do something with the response...
	for _, server := range response.Servers {
		fmt.Println("Server", server.Id, server.Name)
	}

}
```

## Create And Modifie instance
```go
package main

import (
	"fmt"
	"time"

	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/scaleway/scaleway-sdk-go/utils"
)

func main() {

	// Create a Scaleway client
	client, err := scw.NewClient(
		// Get your credentials at https://console.scaleway.com/account/credentials
		scw.WithDefaultOrganizationID("ORGANISATION_ID"),
		scw.WithAuth("ACCESS_KEY", "SECRET_KEY"),
	)
	if err != nil {
		panic(err)
	}

	// Create SDK objects for Scaleway Instance product
	instanceApi := instance.NewApi(client)

	// Create Server
	var request instance.CreateServerRequest
	request.Zone=utils.ZoneFrPar1
	request.Name="archie.test"
	request.DynamicIpRequired=false
	request.CommercialType="GP1-S"
	request.Image="43a60ffc-6afd-4b09-8a74-b8da67a0f95d"
	request.EnableIpv6=false
	request.DynamicIpRequired=false
	request.BootType=instance.ServerBootTypeLocal
	response,err:=instanceApi.CreateServer(&request)
	if err!=nil{
		fmt.Println("error",err)
	}

	// stop server
	_,err=instanceApi.ServerAction(&instance.ServerActionRequest{
		Zone: utils.ZoneFrPar1,
		ServerId: response.Server.Id,
		Action: instance.ServerActionPoweroff,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Please wait for it to stop")
	<-time.After(2*time.Minute)

	// detache IP
	_,err:=instanceApi.UpdateIp(&instance.UpdateIpRequest{	
		Zone: utils.ZoneFrPar1,
		IpId:server.PublicIp.Id,
		Override:map[string]interface{}{"server":nil},
	})

}
```
## Development

This repository is at its early stage and is still in active development.
If you are looking for a way to contribute please read [CONTRIBUTING.md](CONTRIBUTING.md).

## Reach us

We love feedback.
Feel free to reach us on [Scaleway Slack community](https://slack.scaleway.com/), we are waiting for you on [#opensource](https://scaleway-community.slack.com/app_redirect?channel=opensource).

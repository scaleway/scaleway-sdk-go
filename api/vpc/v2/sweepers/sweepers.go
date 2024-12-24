package sweepers

import (
	"fmt"

	vpcSDK "github.com/scaleway/scaleway-sdk-go/api/vpc/v2"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepVPC(scwClient *scw.Client, region scw.Region) error {
	vpcAPI := vpcSDK.NewAPI(scwClient)

	listVPCs, err := vpcAPI.ListVPCs(&vpcSDK.ListVPCsRequest{Region: region}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing secrets in (%s) in sweeper: %s", region, err)
	}

	for _, v := range listVPCs.Vpcs {
		if v.IsDefault {
			continue
		}
		err := vpcAPI.DeleteVPC(&vpcSDK.DeleteVPCRequest{
			VpcID:  v.ID,
			Region: region,
		})
		if err != nil {
			return fmt.Errorf("error deleting VPC in sweeper: %s", err)
		}
	}

	return nil
}

func SweepPrivateNetwork(scwClient *scw.Client, region scw.Region) error {
	vpcAPI := vpcSDK.NewAPI(scwClient)

	logger.Debugf("sweeper: destroying the private network in (%s)", region)

	listPNResponse, err := vpcAPI.ListPrivateNetworks(&vpcSDK.ListPrivateNetworksRequest{
		Region: region,
	}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing private network in sweeper: %s", err)
	}

	for _, pn := range listPNResponse.PrivateNetworks {
		err := vpcAPI.DeletePrivateNetwork(&vpcSDK.DeletePrivateNetworkRequest{
			Region:           region,
			PrivateNetworkID: pn.ID,
		})
		if err != nil {
			return fmt.Errorf("error deleting private network in sweeper: %s", err)
		}
	}

	return nil
}

func SweepRoute(scwClient *scw.Client, region scw.Region) error {
	vpcAPI := vpcSDK.NewAPI(scwClient)
	vpcRouteAPI := vpcSDK.NewRoutesWithNexthopAPI(scwClient)

	logger.Warningf("sweeper: destroying the route in (%s)", region)

	listRoutesResponse, err := vpcRouteAPI.ListRoutesWithNexthop(&vpcSDK.RoutesWithNexthopAPIListRoutesWithNexthopRequest{
		Region: region,
	}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing route in sweeper: %s", err)
	}

	for _, routeWithNexthop := range listRoutesResponse.Routes {
		if routeWithNexthop.Route != nil {
			err := vpcAPI.DeleteRoute(&vpcSDK.DeleteRouteRequest{
				Region:  region,
				RouteID: routeWithNexthop.Route.ID,
			})
			if err != nil {
				return fmt.Errorf("error deleting route in sweeper: %s", err)
			}
		} else {
			return fmt.Errorf("route is nil in RouteWithNexthop: %v", routeWithNexthop)
		}
	}

	return nil
}

func SweepAllLocalities(scwClient *scw.Client) error {
	for _, region := range (&vpcSDK.API{}).Regions() {
		err := SweepVPC(scwClient, region)
		if err != nil {
			return err
		}

		err = SweepPrivateNetwork(scwClient, region)
		if err != nil {
			return err
		}

		err = SweepRoute(scwClient, region)
		if err != nil {
			return err
		}
	}

	return nil
}

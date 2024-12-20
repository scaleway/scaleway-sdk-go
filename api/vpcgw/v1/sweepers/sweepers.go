package sweepers

import (
	"fmt"

	vpcgwSDK "github.com/scaleway/scaleway-sdk-go/api/vpcgw/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepVPCPublicGateway(scwClient *scw.Client, zone scw.Zone) error {
	api := vpcgwSDK.NewAPI(scwClient)

	listGatewayResponse, err := api.ListGateways(&vpcgwSDK.ListGatewaysRequest{
		Zone: zone,
	}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing public gateway in sweeper: %w", err)
	}

	for _, gateway := range listGatewayResponse.Gateways {
		err := api.DeleteGateway(&vpcgwSDK.DeleteGatewayRequest{
			Zone:      zone,
			GatewayID: gateway.ID,
		})
		if err != nil {
			return fmt.Errorf("error deleting public gateway in sweeper: %w", err)
		}
	}
	return nil
}

func SweepGatewayNetworks(scwClient *scw.Client, zone scw.Zone) error {
	api := vpcgwSDK.NewAPI(scwClient)

	listPNResponse, err := api.ListGatewayNetworks(&vpcgwSDK.ListGatewayNetworksRequest{
		Zone: zone,
	}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing gateway network in sweeper: %s", err)
	}

	for _, gn := range listPNResponse.GatewayNetworks {
		err := api.DeleteGatewayNetwork(&vpcgwSDK.DeleteGatewayNetworkRequest{
			GatewayNetworkID: gn.GatewayID,
			Zone:             zone,
			// Cleanup the dhcp resource related. DON'T CALL THE SWEEPER DHCP
			CleanupDHCP: true,
		})
		if err != nil {
			return fmt.Errorf("error deleting gateway network in sweeper: %s", err)
		}
	}
	return nil
}

func SweepVPCPublicGatewayIP(scwClient *scw.Client, zone scw.Zone) error {
	api := vpcgwSDK.NewAPI(scwClient)

	listIPResponse, err := api.ListIPs(&vpcgwSDK.ListIPsRequest{
		Zone: zone,
	}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing public gateway ip in sweeper: %s", err)
	}

	for _, ip := range listIPResponse.IPs {
		err := api.DeleteIP(&vpcgwSDK.DeleteIPRequest{
			Zone: zone,
			IPID: ip.ID,
		})
		if err != nil {
			return fmt.Errorf("error deleting public gateway ip in sweeper: %s", err)
		}
	}
	return nil
}

func SweepVPCPublicGatewayDHCP(scwClient *scw.Client, zone scw.Zone) error {
	api := vpcgwSDK.NewAPI(scwClient)

	listDHCPsResponse, err := api.ListDHCPs(&vpcgwSDK.ListDHCPsRequest{
		Zone: zone,
	}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing public gateway dhcps in sweeper: %w", err)
	}

	for _, dhcp := range listDHCPsResponse.Dhcps {
		err := api.DeleteDHCP(&vpcgwSDK.DeleteDHCPRequest{
			Zone:   zone,
			DHCPID: dhcp.ID,
		})
		if err != nil {
			return fmt.Errorf("error deleting public gateway dhcp in sweeper: %w", err)
		}
	}

	return nil
}

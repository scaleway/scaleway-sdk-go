package sweepers

import (
	"fmt"

	vpcgwSDK "github.com/scaleway/scaleway-sdk-go/api/vpcgw/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepVPCPublicGateway(scwClient *scw.Client, zone scw.Zone) error {
	api := vpcgwSDK.NewAPI(scwClient)

	listGatewayResponse, err := api.ListGateways(&vpcgwSDK.ListGatewaysRequest{ //nolint:staticcheck
		Zone: zone,
	}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing public gateway in sweeper: %w", err)
	}

	for _, gateway := range listGatewayResponse.Gateways {
		err := api.DeleteGateway(&vpcgwSDK.DeleteGatewayRequest{ //nolint:staticcheck
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

	listPNResponse, err := api.ListGatewayNetworks(&vpcgwSDK.ListGatewayNetworksRequest{ //nolint:staticcheck
		Zone: zone,
	}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing gateway network in sweeper: %s", err)
	}

	for _, gn := range listPNResponse.GatewayNetworks {
		err := api.DeleteGatewayNetwork(&vpcgwSDK.DeleteGatewayNetworkRequest{ //nolint:staticcheck
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

	listIPResponse, err := api.ListIPs(&vpcgwSDK.ListIPsRequest{ //nolint:staticcheck
		Zone: zone,
	}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing public gateway ip in sweeper: %s", err)
	}

	for _, ip := range listIPResponse.IPs {
		err := api.DeleteIP(&vpcgwSDK.DeleteIPRequest{ //nolint:staticcheck
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

	listDHCPsResponse, err := api.ListDHCPs(&vpcgwSDK.ListDHCPsRequest{ //nolint:staticcheck
		Zone: zone,
	}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing public gateway dhcps in sweeper: %w", err)
	}

	for _, dhcp := range listDHCPsResponse.Dhcps {
		err := api.DeleteDHCP(&vpcgwSDK.DeleteDHCPRequest{ //nolint:staticcheck
			Zone:   zone,
			DHCPID: dhcp.ID,
		})
		if err != nil {
			return fmt.Errorf("error deleting public gateway dhcp in sweeper: %w", err)
		}
	}

	return nil
}

func SweepAllLocalities(scwClient *scw.Client) error {
	for _, zone := range (&vpcgwSDK.API{}).Zones() {
		err := SweepVPCPublicGateway(scwClient, zone)
		if err != nil {
			return err
		}
		err = SweepGatewayNetworks(scwClient, zone)
		if err != nil {
			return err
		}
		err = SweepVPCPublicGatewayIP(scwClient, zone)
		if err != nil {
			return err
		}
		err = SweepVPCPublicGatewayDHCP(scwClient, zone)
		if err != nil {
			return err
		}
	}

	return nil
}

package sweepers

import (
	"fmt"

	vpcgwSDK "github.com/scaleway/scaleway-sdk-go/api/vpcgw/v2"
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
		_, err = api.DeleteGateway(&vpcgwSDK.DeleteGatewayRequest{
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
		_, err = api.DeleteGatewayNetwork(&vpcgwSDK.DeleteGatewayNetworkRequest{
			GatewayNetworkID: gn.GatewayID,
			Zone:             zone,
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
		err = api.DeleteIP(&vpcgwSDK.DeleteIPRequest{
			Zone: zone,
			IPID: ip.ID,
		})
		if err != nil {
			return fmt.Errorf("error deleting public gateway ip in sweeper: %s", err)
		}
	}
	return nil
}

func SweepPATRules(scwClient *scw.Client, zone scw.Zone) error {
	api := vpcgwSDK.NewAPI(scwClient)

	listPATRulesResponse, err := api.ListPatRules(&vpcgwSDK.ListPatRulesRequest{
		Zone: zone,
	}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing PAT rules in sweeper: %s", err)
	}

	for _, rule := range listPATRulesResponse.PatRules {
		err = api.DeletePatRule(&vpcgwSDK.DeletePatRuleRequest{
			Zone:      zone,
			PatRuleID: rule.ID,
		})
		if err != nil {
			return fmt.Errorf("error deleting PAT rule in sweeper: %s", err)
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
		err = SweepPATRules(scwClient, zone)
		if err != nil {
			return err
		}
	}

	return nil
}

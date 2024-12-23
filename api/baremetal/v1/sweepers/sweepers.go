package sweepers

import (
	"fmt"

	"github.com/scaleway/scaleway-sdk-go/api/baremetal/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepServers(scwClient *scw.Client, zone scw.Zone) error {
	baremetalAPI := baremetal.NewAPI(scwClient)

	listServers, err := baremetalAPI.ListServers(&baremetal.ListServersRequest{Zone: zone}, scw.WithAllPages())
	if err != nil {
		return err
	}

	for _, server := range listServers.Servers {
		_, err := baremetalAPI.DeleteServer(&baremetal.DeleteServerRequest{
			Zone:     zone,
			ServerID: server.ID,
		})
		if err != nil {
			return fmt.Errorf("error deleting server in sweeper: %s", err)
		}
	}

	return nil
}

func SweepAllLocalities(scwClient *scw.Client) error {
	for _, zone := range (&baremetal.API{}).Zones() {
		err := SweepServers(scwClient, zone)
		if err != nil {
			return err
		}
	}

	return nil
}

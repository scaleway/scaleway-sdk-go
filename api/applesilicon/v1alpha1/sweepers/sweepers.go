package sweepers

import (
	"fmt"

	applesilicon "github.com/scaleway/scaleway-sdk-go/api/applesilicon/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepServer(scwClient *scw.Client, zone scw.Zone) error {
	asAPI := applesilicon.NewAPI(scwClient)

	listServers, err := asAPI.ListServers(&applesilicon.ListServersRequest{Zone: zone}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing apple silicon servers in (%s) in sweeper: %s", zone, err)
	}

	for _, server := range listServers.Servers {
		errDelete := asAPI.DeleteServer(&applesilicon.DeleteServerRequest{
			ServerID: server.ID,
			Zone:     zone,
		})
		if errDelete != nil {
			return fmt.Errorf("error deleting apple silicon server in sweeper: %s", err)
		}
	}

	return nil
}

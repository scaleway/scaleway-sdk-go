package sweepers

import (
	"fmt"

	applesilicon "github.com/scaleway/scaleway-sdk-go/api/applesilicon/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepServer(scwClient *scw.Client, zone scw.Zone, projectScoped bool) error {
	asAPI := applesilicon.NewAPI(scwClient)
	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	var projectID *string = nil
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped {
		projectID = &defaultProjectID
	}
	listServers, err := asAPI.ListServers(&applesilicon.ListServersRequest{Zone: zone, ProjectID: projectID}, scw.WithAllPages())
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

func SweepAllLocalities(scwClient *scw.Client, projectScoped bool) error {
	for _, zone := range (&applesilicon.API{}).Zones() {
		err := SweepServer(scwClient, zone, projectScoped)
		if err != nil {
			return err
		}
	}

	return nil
}

package sweepers

import (
	"fmt"

	"github.com/scaleway/scaleway-sdk-go/api/iot/v1"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepHub(scwClient *scw.Client, region scw.Region, projectScoped bool) error {
	iotAPI := iot.NewAPI(scwClient)
	logger.Warningf("sweeper: destroying the iot hub in (%s)", region)
	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	var projectID *string = nil
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped {
		projectID = &defaultProjectID
	}
	listHubs, err := iotAPI.ListHubs(&iot.ListHubsRequest{Region: region, ProjectID: projectID}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing hubs in (%s) in sweeper: %s", region, err)
	}

	deleteDevices := true
	for _, hub := range listHubs.Hubs {
		err := iotAPI.DeleteHub(&iot.DeleteHubRequest{
			HubID:         hub.ID,
			Region:        hub.Region,
			DeleteDevices: &deleteDevices,
		})
		if err != nil {
			return fmt.Errorf("error deleting hub in sweeper: %s", err)
		}
	}

	return nil
}

func SweepAllLocalities(scwClient *scw.Client, projectScoped bool) error {
	for _, region := range (&iot.API{}).Regions() {
		err := SweepHub(scwClient, region, projectScoped)
		if err != nil {
			return err
		}
	}

	return nil
}

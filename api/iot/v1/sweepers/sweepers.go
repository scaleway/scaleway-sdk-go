package sweepers

import (
	"fmt"

	"github.com/scaleway/scaleway-sdk-go/api/iot/v1"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepHub(scwClient *scw.Client, region scw.Region) error {
	iotAPI := iot.NewAPI(scwClient)
	logger.Warningf("sweeper: destroying the iot hub in (%s)", region)
	listHubs, err := iotAPI.ListHubs(&iot.ListHubsRequest{Region: region}, scw.WithAllPages())
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

package sweepers

import (
	"fmt"

	"github.com/scaleway/scaleway-sdk-go/api/mongodb/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepInstances(scwClient *scw.Client, zone scw.Zone) error {
	mongodbAPI := mongodb.NewAPI(scwClient)
	logger.Warningf("sweeper: destroying the mongodb instance in (%s)", zone)
	extractRegion, err := zone.Region()
	if err != nil {
		return fmt.Errorf("error extract region in (%s) in sweeper: %w", zone, err)
	}
	listInstance, err := mongodbAPI.ListInstances(&mongodb.ListInstancesRequest{
		Region: extractRegion,
	})
	if err != nil {
		return fmt.Errorf("error listing mongodb instance in (%s) in sweeper: %w", zone, err)
	}

	for _, instance := range listInstance.Instances {
		_, err := mongodbAPI.DeleteInstance(&mongodb.DeleteInstanceRequest{
			Region:     extractRegion,
			InstanceID: instance.ID,
		})
		if err != nil {
			return fmt.Errorf("error deleting mongodb instance in sweeper: %w", err)
		}
	}

	return nil
}

package sweepers

import (
	"fmt"

	"github.com/scaleway/scaleway-sdk-go/api/rdb/v1"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepInstance(scwClient *scw.Client, region scw.Region) error {
	rdbAPI := rdb.NewAPI(scwClient)
	logger.Warningf("sweeper: destroying the rdb instance in (%s)", region)
	listInstances, err := rdbAPI.ListInstances(&rdb.ListInstancesRequest{
		Region: region,
	}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing rdb instances in (%s) in sweeper: %s", region, err)
	}

	for _, instance := range listInstances.Instances {
		_, err := rdbAPI.DeleteInstance(&rdb.DeleteInstanceRequest{
			Region:     region,
			InstanceID: instance.ID,
		})
		if err != nil {
			return fmt.Errorf("error deleting rdb instance in sweeper: %s", err)
		}
	}

	return nil
}

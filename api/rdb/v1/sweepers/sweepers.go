package sweepers

import (
	"fmt"

	"github.com/scaleway/scaleway-sdk-go/api/rdb/v1"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepInstance(scwClient *scw.Client, region scw.Region, projectScoped bool) error {
	rdbAPI := rdb.NewAPI(scwClient)
	logger.Warningf("sweeper: destroying the rdb instance in (%s)", region)
	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	var projectID *string = nil
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped {
		projectID = &defaultProjectID
	}
	listInstances, err := rdbAPI.ListInstances(&rdb.ListInstancesRequest{
		Region:    region,
		ProjectID: projectID,
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

func SweepAllLocalities(scwClient *scw.Client, projectScoped bool) error {
	for _, region := range (&rdb.API{}).Regions() {
		err := SweepInstance(scwClient, region, projectScoped)
		if err != nil {
			return err
		}
	}

	return nil
}

package sweepers

import (
	"fmt"

	mongodb "github.com/scaleway/scaleway-sdk-go/api/mongodb/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepSnapshots(scwClient *scw.Client, region scw.Region, projectScoped bool) error {
	mongodbAPI := mongodb.NewAPI(scwClient)
	logger.Warningf("sweeper: destroying mongodb snapshots in (%s)", region)

	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	var projectID *string = nil
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped {
		projectID = &defaultProjectID
	}

	listSnapshots, err := mongodbAPI.ListSnapshots(&mongodb.ListSnapshotsRequest{
		Region:    region,
		ProjectID: projectID,
	}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing mongodb snapshots in (%s) in sweeper: %w", region, err)
	}

	for _, snapshot := range listSnapshots.Snapshots {
		_, err := mongodbAPI.DeleteSnapshot(&mongodb.DeleteSnapshotRequest{
			Region:     region,
			SnapshotID: snapshot.ID,
		})
		if err != nil {
			logger.Warningf("error deleting mongodb snapshot %s in sweeper: %s", snapshot.ID, err)
		}
	}

	return nil
}

func SweepInstances(scwClient *scw.Client, region scw.Region, projectScoped bool) error {
	mongodbAPI := mongodb.NewAPI(scwClient)
	logger.Warningf("sweeper: destroying the mongodb instance in (%s)", region)
	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	var projectID *string = nil
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped {
		projectID = &defaultProjectID
	}
	listInstance, err := mongodbAPI.ListInstances(&mongodb.ListInstancesRequest{
		Region:    region,
		ProjectID: projectID,
	})
	if err != nil {
		return fmt.Errorf("error listing mongodb instance in (%s) in sweeper: %w", region, err)
	}

	for _, instance := range listInstance.Instances {
		_, err := mongodbAPI.DeleteInstance(&mongodb.DeleteInstanceRequest{
			Region:     region,
			InstanceID: instance.ID,
		})
		if err != nil {
			return fmt.Errorf("error deleting mongodb instance in sweeper: %w", err)
		}
	}

	return nil
}

func SweepAllLocalities(scwClient *scw.Client, projectScoped bool) error {
	for _, region := range (&mongodb.API{}).Regions() {
		err := SweepSnapshots(scwClient, region, projectScoped)
		if err != nil {
			return err
		}
		err = SweepInstances(scwClient, region, projectScoped)
		if err != nil {
			return err
		}
	}

	return nil
}

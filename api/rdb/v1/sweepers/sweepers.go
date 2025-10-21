package sweepers

import (
	"fmt"

	"github.com/scaleway/scaleway-sdk-go/api/rdb/v1"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepBackups(scwClient *scw.Client, region scw.Region) error {
	rdbAPI := rdb.NewAPI(scwClient)
	logger.Warningf("sweeper: destroying rdb backups in (%s)", region)

	listBackups, err := rdbAPI.ListDatabaseBackups(&rdb.ListDatabaseBackupsRequest{
		Region: region,
	}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing rdb backups in (%s) in sweeper: %s", region, err)
	}

	for _, backup := range listBackups.DatabaseBackups {
		_, err := rdbAPI.DeleteDatabaseBackup(&rdb.DeleteDatabaseBackupRequest{
			Region:           region,
			DatabaseBackupID: backup.ID,
		})
		if err != nil {
			logger.Warningf("error deleting rdb backup %s in sweeper: %s", backup.ID, err)
		}
	}

	return nil
}

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

func SweepAllLocalities(scwClient *scw.Client) error {
	for _, region := range (&rdb.API{}).Regions() {
		err := SweepBackups(scwClient, region)
		if err != nil {
			return err
		}
		err = SweepInstance(scwClient, region)
		if err != nil {
			return err
		}
	}

	return nil
}

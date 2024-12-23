package sweepers

import (
	"fmt"

	sdbSDK "github.com/scaleway/scaleway-sdk-go/api/serverless_sqldb/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepDatabase(scwClient *scw.Client, region scw.Region) error {
	sdbAPI := sdbSDK.NewAPI(scwClient)
	logger.Warningf("sweeper: destroying the serverless sql database in (%s)", region)
	listServerlessSQLDBDatabases, err := sdbAPI.ListDatabases(
		&sdbSDK.ListDatabasesRequest{
			Region: region,
		}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing database in (%s) in sweeper: %s", region, err)
	}

	for _, database := range listServerlessSQLDBDatabases.Databases {
		_, err := sdbAPI.DeleteDatabase(&sdbSDK.DeleteDatabaseRequest{
			DatabaseID: database.ID,
			Region:     region,
		})
		if err != nil {
			return fmt.Errorf("error deleting database in sweeper: %s", err)
		}
	}

	return nil
}

func SweepAllLocalities(scwClient *scw.Client) error {
	for _, region := range (&sdbSDK.API{}).Regions() {
		err := SweepDatabase(scwClient, region)
		if err != nil {
			return err
		}
	}

	return nil
}

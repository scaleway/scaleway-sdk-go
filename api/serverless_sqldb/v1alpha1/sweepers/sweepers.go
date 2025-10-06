package sweepers

import (
	"fmt"

	sdbSDK "github.com/scaleway/scaleway-sdk-go/api/serverless_sqldb/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepDatabase(scwClient *scw.Client, region scw.Region, projectScoped bool) error {
	sdbAPI := sdbSDK.NewAPI(scwClient)
	logger.Warningf("sweeper: destroying the serverless sql database in (%s)", region)
	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	var projectID *string = nil
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped {
		projectID = &defaultProjectID
	}
	listServerlessSQLDBDatabases, err := sdbAPI.ListDatabases(
		&sdbSDK.ListDatabasesRequest{
			Region:    region,
			ProjectID: *projectID,
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

func SweepAllLocalities(scwClient *scw.Client, projectScoped bool) error {
	for _, region := range (&sdbSDK.API{}).Regions() {
		err := SweepDatabase(scwClient, region, projectScoped)
		if err != nil {
			return err
		}
	}

	return nil
}

package sweepers

import (
	"fmt"

	jobs "github.com/scaleway/scaleway-sdk-go/api/jobs/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepJobDefinition(scwClient *scw.Client, region scw.Region, projectScoped bool) error {
	jobsAPI := jobs.NewAPI(scwClient)
	logger.Warningf("sweeper: destroying the jobs definitions in (%s)", region)
	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	var projectID *string = nil
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped {
		projectID = &defaultProjectID
	}
	listJobDefinitions, err := jobsAPI.ListJobDefinitions(
		&jobs.ListJobDefinitionsRequest{
			Region:    region,
			ProjectID: projectID,
		}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing definition in (%s) in sweeper: %s", region, err)
	}

	for _, definition := range listJobDefinitions.JobDefinitions {
		err := jobsAPI.DeleteJobDefinition(&jobs.DeleteJobDefinitionRequest{
			JobDefinitionID: definition.ID,
			Region:          region,
		})
		if err != nil {
			return fmt.Errorf("error deleting definition in sweeper: %s", err)
		}
	}

	return nil
}

func SweepAllLocalities(scwClient *scw.Client, projectScoped bool) error {
	for _, region := range (&jobs.API{}).Regions() {
		err := SweepJobDefinition(scwClient, region, projectScoped)
		if err != nil {
			return err
		}
	}

	return nil
}

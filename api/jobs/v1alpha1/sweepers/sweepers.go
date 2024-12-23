package sweepers

import (
	"fmt"

	jobs "github.com/scaleway/scaleway-sdk-go/api/jobs/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepJobDefinition(scwClient *scw.Client, region scw.Region) error {
	jobsAPI := jobs.NewAPI(scwClient)
	logger.Warningf("sweeper: destroying the jobs definitions in (%s)", region)
	listJobDefinitions, err := jobsAPI.ListJobDefinitions(
		&jobs.ListJobDefinitionsRequest{
			Region: region,
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

package sweepers

import (
	"fmt"

	"github.com/scaleway/scaleway-sdk-go/api/inference/v1"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepDeployment(scwClient *scw.Client, region scw.Region, projectScoped bool) error {
	inferenceAPI := inference.NewAPI(scwClient)
	logger.Warningf("sweeper: destroying the inference deployments in (%s)", region)
	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	var projectID *string = nil
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped {
		projectID = &defaultProjectID
	}
	listDeployments, err := inferenceAPI.ListDeployments(
		&inference.ListDeploymentsRequest{
			Region:    region,
			ProjectID: projectID,
		}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing deployment in (%s) in sweeper: %s", region, err)
	}

	for _, deployment := range listDeployments.Deployments {
		_, err := inferenceAPI.DeleteDeployment(&inference.DeleteDeploymentRequest{
			DeploymentID: deployment.ID,
			Region:       region,
		})
		if err != nil {
			return fmt.Errorf("error deleting deployment in sweeper: %s", err)
		}
	}

	return nil
}

func SweepAllLocalities(scwClient *scw.Client, projectScoped bool) error {
	for _, locality := range (&inference.API{}).Regions() {
		err := SweepDeployment(scwClient, locality, projectScoped)
		if err != nil {
			return err
		}
	}

	return nil
}

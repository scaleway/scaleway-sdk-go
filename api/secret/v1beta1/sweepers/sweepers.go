package sweepers

import (
	"fmt"

	secretSDK "github.com/scaleway/scaleway-sdk-go/api/secret/v1beta1"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepSecret(scwClient *scw.Client, region scw.Region, projectScoped bool) error {
	secretAPI := secretSDK.NewAPI(scwClient)

	logger.Warningf("sweeper: deleting the secrets in (%s)", region)

	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	var projectID *string = nil
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped {
		projectID = &defaultProjectID
	}

	listSecrets, err := secretAPI.ListSecrets(&secretSDK.ListSecretsRequest{Region: region, ProjectID: projectID}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing secrets in (%s) in sweeper: %s", region, err)
	}

	for _, se := range listSecrets.Secrets {
		err := secretAPI.DeleteSecret(&secretSDK.DeleteSecretRequest{
			SecretID: se.ID,
			Region:   region,
		})
		if err != nil {
			return fmt.Errorf("error deleting secret in sweeper: %s", err)
		}
	}

	return nil
}

func SweepAllLocalities(scwClient *scw.Client, projectScoped bool) error {
	for _, region := range (&secretSDK.API{}).Regions() {
		err := SweepSecret(scwClient, region, projectScoped)
		if err != nil {
			return err
		}
	}
	return nil
}

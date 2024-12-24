package sweepers

import (
	"fmt"

	secretSDK "github.com/scaleway/scaleway-sdk-go/api/secret/v1beta1"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepSecret(scwClient *scw.Client, region scw.Region) error {
	secretAPI := secretSDK.NewAPI(scwClient)

	logger.Warningf("sweeper: deleting the secrets in (%s)", region)

	listSecrets, err := secretAPI.ListSecrets(&secretSDK.ListSecretsRequest{Region: region}, scw.WithAllPages())
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

func SweepAllLocalities(scwClient *scw.Client) error {
	for _, region := range (&secretSDK.API{}).Regions() {
		err := SweepSecret(scwClient, region)
		if err != nil {
			return err
		}
	}
	return nil
}

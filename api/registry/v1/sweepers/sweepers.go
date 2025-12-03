package sweepers

import (
	"fmt"

	"github.com/scaleway/scaleway-sdk-go/api/registry/v1"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepNamespace(scwClient *scw.Client, region scw.Region, projectScoped bool) error {
	registryAPI := registry.NewAPI(scwClient)
	logger.Warningf("sweeper: destroying the registry namespaces in (%s)", region)
	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	var projectID *string = nil
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped {
		projectID = &defaultProjectID
	}
	listNamespaces, err := registryAPI.ListNamespaces(
		&registry.ListNamespacesRequest{Region: region, ProjectID: projectID}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing namespaces in (%s) in sweeper: %s", region, err)
	}

	for _, ns := range listNamespaces.Namespaces {
		_, err := registryAPI.DeleteNamespace(&registry.DeleteNamespaceRequest{
			NamespaceID: ns.ID,
			Region:      region,
		})
		if err != nil {
			return fmt.Errorf("error deleting namespace in sweeper: %s", err)
		}
	}

	return nil
}

func SweepAllLocalities(scwClient *scw.Client, projectScoped bool) error {
	for _, region := range (&registry.API{}).Regions() {
		err := SweepNamespace(scwClient, region, projectScoped)
		if err != nil {
			return err
		}
	}

	return nil
}

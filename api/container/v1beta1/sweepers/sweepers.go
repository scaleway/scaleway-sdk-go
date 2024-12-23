package sweepers

import (
	"fmt"

	container "github.com/scaleway/scaleway-sdk-go/api/container/v1beta1"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepTrigger(scwClient *scw.Client, region scw.Region) error {
	containerAPI := container.NewAPI(scwClient)

	logger.Warningf("sweeper: destroying the container triggers in (%s)", region)
	listTriggers, err := containerAPI.ListTriggers(
		&container.ListTriggersRequest{
			Region: region,
		}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing trigger in (%s) in sweeper: %s", region, err)
	}

	for _, trigger := range listTriggers.Triggers {
		_, err := containerAPI.DeleteTrigger(&container.DeleteTriggerRequest{
			TriggerID: trigger.ID,
			Region:    region,
		})
		if err != nil {
			return fmt.Errorf("error deleting trigger in sweeper: %s", err)
		}
	}

	return nil
}

func SweepContainer(scwClient *scw.Client, region scw.Region) error {
	containerAPI := container.NewAPI(scwClient)
	logger.Warningf("sweeper: destroying the container in (%s)", region)
	listNamespaces, err := containerAPI.ListContainers(
		&container.ListContainersRequest{
			Region: region,
		}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing containers in (%s) in sweeper: %s", region, err)
	}

	for _, cont := range listNamespaces.Containers {
		_, err := containerAPI.DeleteContainer(&container.DeleteContainerRequest{
			ContainerID: cont.ID,
			Region:      region,
		})
		if err != nil {
			return fmt.Errorf("error deleting container in sweeper: %s", err)
		}
	}

	return nil
}

func SweepNamespace(scwClient *scw.Client, region scw.Region) error {
	containerAPI := container.NewAPI(scwClient)
	logger.Warningf("sweeper: destroying the container namespaces in (%s)", region)
	listNamespaces, err := containerAPI.ListNamespaces(
		&container.ListNamespacesRequest{
			Region: region,
		}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing namespaces in (%s) in sweeper: %s", region, err)
	}

	for _, ns := range listNamespaces.Namespaces {
		_, err := containerAPI.DeleteNamespace(&container.DeleteNamespaceRequest{
			NamespaceID: ns.ID,
			Region:      region,
		})
		if err != nil {
			return fmt.Errorf("error deleting namespace in sweeper: %s", err)
		}
	}

	return nil
}

func SweepAllLocalities(scwClient *scw.Client) error {
	for _, region := range (&container.API{}).Regions() {
		err := SweepTrigger(scwClient, region)
		if err != nil {
			return err
		}
		err = SweepContainer(scwClient, region)
		if err != nil {
			return err
		}
		err = SweepNamespace(scwClient, region)
		if err != nil {
			return err
		}
	}
	return nil
}

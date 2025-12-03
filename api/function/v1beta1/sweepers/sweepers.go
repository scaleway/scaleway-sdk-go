package sweepers

import (
	"fmt"

	functionSDK "github.com/scaleway/scaleway-sdk-go/api/function/v1beta1"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepTriggers(scwClient *scw.Client, region scw.Region, projectScoped bool) error {
	functionAPI := functionSDK.NewAPI(scwClient)
	logger.Warningf("sweeper: destroying the function triggers in (%s)", region)

	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	var projectID *string = nil
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped {
		projectID = &defaultProjectID
	}
	listTriggers, err := functionAPI.ListTriggers(
		&functionSDK.ListTriggersRequest{
			Region:    region,
			ProjectID: projectID,
		}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing trigger in (%s) in sweeper: %s", region, err)
	}

	for _, trigger := range listTriggers.Triggers {
		_, err := functionAPI.DeleteTrigger(&functionSDK.DeleteTriggerRequest{
			TriggerID: trigger.ID,
			Region:    region,
		})
		if err != nil {
			return fmt.Errorf("error deleting trigger in sweeper: %s", err)
		}
	}

	return nil
}

func SweepNamespaces(scwClient *scw.Client, region scw.Region, projectScoped bool) error {
	functionAPI := functionSDK.NewAPI(scwClient)
	logger.Debugf("sweeper: destroying the function namespaces in (%s)", region)
	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	var projectID *string = nil
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped {
		projectID = &defaultProjectID
	}
	listNamespaces, err := functionAPI.ListNamespaces(
		&functionSDK.ListNamespacesRequest{
			Region:    region,
			ProjectID: projectID,
		}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing namespaces in (%s) in sweeper: %s", region, err)
	}

	for _, ns := range listNamespaces.Namespaces {
		_, err := functionAPI.DeleteNamespace(&functionSDK.DeleteNamespaceRequest{
			NamespaceID: ns.ID,
			Region:      region,
		})
		if err != nil {
			logger.Debugf("sweeper: error (%s)", err)

			return fmt.Errorf("error deleting namespace in sweeper: %s", err)
		}
	}

	return nil
}

func SweepFunctions(scwClient *scw.Client, region scw.Region, projectScoped bool) error {
	functionAPI := functionSDK.NewAPI(scwClient)
	logger.Debugf("sweeper: destroying the function in (%s)", region)
	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	var projectID *string = nil
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped {
		projectID = &defaultProjectID
	}
	listFunctions, err := functionAPI.ListFunctions(
		&functionSDK.ListFunctionsRequest{
			Region:    region,
			ProjectID: projectID,
		}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing functions in (%s) in sweeper: %s", region, err)
	}

	for _, f := range listFunctions.Functions {
		_, err := functionAPI.DeleteFunction(&functionSDK.DeleteFunctionRequest{
			FunctionID: f.ID,
			Region:     region,
		})
		if err != nil {
			return fmt.Errorf("error deleting functions in sweeper: %s", err)
		}
	}

	return nil
}

func SweepCrons(scwClient *scw.Client, region scw.Region, projectScoped bool) error {
	functionAPI := functionSDK.NewAPI(scwClient)
	logger.Warningf("sweeper: destroying the function cron in (%s)", region)
	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped {
		listFunctions, err := functionAPI.ListFunctions(
			&functionSDK.ListFunctionsRequest{
				Region:    region,
				ProjectID: &defaultProjectID,
			}, scw.WithAllPages())
		if err != nil {
			return fmt.Errorf("error listing functions in (%s) in sweeper: %s", region, err)
		}
		for _, function := range listFunctions.Functions {
			listCron, err := functionAPI.ListCrons(
				&functionSDK.ListCronsRequest{
					Region:     region,
					FunctionID: function.ID,
				}, scw.WithAllPages())
			if err != nil {
				return fmt.Errorf("error listing cron in (%s) in sweeper: %s", region, err)
			}

			for _, cron := range listCron.Crons {
				_, err := functionAPI.DeleteCron(&functionSDK.DeleteCronRequest{
					CronID: cron.ID,
					Region: region,
				})
				if err != nil {
					return fmt.Errorf("error deleting cron in sweeper: %s", err)
				}
			}
		}
	} else {
		listCron, err := functionAPI.ListCrons(
			&functionSDK.ListCronsRequest{
				Region: region,
			}, scw.WithAllPages())
		if err != nil {
			return fmt.Errorf("error listing cron in (%s) in sweeper: %s", region, err)
		}

		for _, cron := range listCron.Crons {
			_, err := functionAPI.DeleteCron(&functionSDK.DeleteCronRequest{
				CronID: cron.ID,
				Region: region,
			})
			if err != nil {
				return fmt.Errorf("error deleting cron in sweeper: %s", err)
			}
		}
	}

	return nil
}

func SweepAllLocalities(scwClient *scw.Client, projectScoped bool) error {
	for _, region := range (&functionSDK.API{}).Regions() {
		err := SweepTriggers(scwClient, region, projectScoped)
		if err != nil {
			return err
		}
		err = SweepNamespaces(scwClient, region, projectScoped)
		if err != nil {
			return err
		}
		err = SweepFunctions(scwClient, region, projectScoped)
		if err != nil {
			return err
		}
		err = SweepCrons(scwClient, region, projectScoped)
		if err != nil {
			return err
		}
	}

	return nil
}

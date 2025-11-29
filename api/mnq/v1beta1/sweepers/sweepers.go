package sweepers

import (
	"fmt"
	"strings"

	accountSDK "github.com/scaleway/scaleway-sdk-go/api/account/v3"
	mnq "github.com/scaleway/scaleway-sdk-go/api/mnq/v1beta1"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepSQSCredentials(scwClient *scw.Client, region scw.Region, projectScoped bool) error {
	mnqAPI := mnq.NewSqsAPI(scwClient)
	projectID, _ := scwClient.GetDefaultProjectID()
	logger.Warningf("sweeper: destroying the mnq sqs credentials in (%s)", region)
	sqsInfo, err := mnqAPI.GetSqsInfo(&mnq.SqsAPIGetSqsInfoRequest{Region: region})
	if err != nil {
		return fmt.Errorf("error getting sns info in sweeper: %s", err)
	}
	if sqsInfo.Status == mnq.SqsInfoStatusDisabled {
		logger.Infof("sqs is disabled, skipping")
		return nil
	}
	listSqsCredentials, err := mnqAPI.ListSqsCredentials(
		&mnq.SqsAPIListSqsCredentialsRequest{
			Region:    region,
			ProjectID: scw.StringPtr(projectID),
		}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing sqs credentials in (%s) in sweeper: %s", region, err)
	}

	for _, credentials := range listSqsCredentials.SqsCredentials {
		err := mnqAPI.DeleteSqsCredentials(&mnq.SqsAPIDeleteSqsCredentialsRequest{
			SqsCredentialsID: credentials.ID,
			Region:           region,
		})
		if err != nil {
			return fmt.Errorf("error deleting sqs credentials in sweeper: %s", err)
		}
	}

	return nil
}

func SweepSQS(scwClient *scw.Client, region scw.Region, projectScoped bool) error {
	accountAPI := accountSDK.NewProjectAPI(scwClient)
	mnqAPI := mnq.NewSqsAPI(scwClient)

	logger.Warningf("sweeper: destroying the mnq sqss in (%s)", region)
	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped && exists && defaultProjectID != "" {
		sqsInfo, err := mnqAPI.GetSqsInfo(&mnq.SqsAPIGetSqsInfoRequest{Region: region, ProjectID: defaultProjectID})
		if err != nil {
			return fmt.Errorf("error getting sqs info in sweeper: %s", err)
		}
		if sqsInfo.Status == mnq.SqsInfoStatusDisabled {
			logger.Infof("sqs is disabled, skipping")
			return nil
		}
		_, err = mnqAPI.DeactivateSqs(&mnq.SqsAPIDeactivateSqsRequest{
			Region:    region,
			ProjectID: defaultProjectID,
		})
		if err != nil {
			return err
		}
	} else {
		listProjects, err := accountAPI.ListProjects(&accountSDK.ProjectAPIListProjectsRequest{}, scw.WithAllPages())
		if err != nil {
			return fmt.Errorf("failed to list projects: %w", err)
		}
		for _, project := range listProjects.Projects {
			if !strings.HasPrefix(project.Name, "tf_tests") {
				continue
			}

			_, err := mnqAPI.DeactivateSqs(&mnq.SqsAPIDeactivateSqsRequest{
				Region:    region,
				ProjectID: project.ID,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func SweepSNSCredentials(scwClient *scw.Client, region scw.Region, projectScoped bool) error {
	mnqAPI := mnq.NewSnsAPI(scwClient)
	logger.Warningf("sweeper: destroying the mnq sns credentials in (%s)", region)
	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	var projectID *string = nil
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped && exists && defaultProjectID != "" {
		projectID = &defaultProjectID
	}
	infoProjectID := ""
	if projectID != nil {
		infoProjectID = *projectID
	}
	snsInfo, err := mnqAPI.GetSnsInfo(&mnq.SnsAPIGetSnsInfoRequest{Region: region, ProjectID: infoProjectID})
	if err != nil {
		return fmt.Errorf("error getting sns info in sweeper: %s", err)
	}
	if snsInfo.Status == mnq.SnsInfoStatusDisabled {
		logger.Infof("sns is disabled, skipping")
		return nil
	}
	listSnsCredentials, err := mnqAPI.ListSnsCredentials(
		&mnq.SnsAPIListSnsCredentialsRequest{
			Region:    region,
			ProjectID: projectID,
		}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing sns credentials in (%s) in sweeper: %s", region, err)
	}

	for _, credentials := range listSnsCredentials.SnsCredentials {
		err := mnqAPI.DeleteSnsCredentials(&mnq.SnsAPIDeleteSnsCredentialsRequest{
			SnsCredentialsID: credentials.ID,
			Region:           region,
		})
		if err != nil {
			return fmt.Errorf("error deleting sns credentials in sweeper: %s", err)
		}
	}

	return nil
}

func SweepSNS(scwClient *scw.Client, region scw.Region, projectScoped bool) error {
	accountAPI := accountSDK.NewProjectAPI(scwClient)
	mnqAPI := mnq.NewSnsAPI(scwClient)

	logger.Warningf("sweeper: destroying the mnq sns in (%s)", region)
	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped && exists && defaultProjectID != "" {
		snsInfo, err := mnqAPI.GetSnsInfo(&mnq.SnsAPIGetSnsInfoRequest{Region: region, ProjectID: defaultProjectID})
		if err != nil {
			return fmt.Errorf("error getting sns info in sweeper: %s", err)
		}
		if snsInfo.Status == mnq.SnsInfoStatusDisabled {
			logger.Infof("sns is disabled, skipping")
			return nil
		}
		_, err = mnqAPI.DeactivateSns(&mnq.SnsAPIDeactivateSnsRequest{
			Region:    region,
			ProjectID: defaultProjectID,
		})
		if err != nil {
			logger.Debugf("sweeper: error (%s)", err)
			return err
		}
	} else {
		listProjects, err := accountAPI.ListProjects(&accountSDK.ProjectAPIListProjectsRequest{}, scw.WithAllPages())
		if err != nil {
			return fmt.Errorf("failed to list projects: %w", err)
		}
		for _, project := range listProjects.Projects {
			if !strings.HasPrefix(project.Name, "tf_tests") {
				continue
			}

			_, err := mnqAPI.DeactivateSns(&mnq.SnsAPIDeactivateSnsRequest{
				Region:    region,
				ProjectID: project.ID,
			})
			if err != nil {
				logger.Debugf("sweeper: error (%s)", err)
				return err
			}
		}
	}

	return nil
}

func SweepNatsAccount(scwClient *scw.Client, region scw.Region, projectScoped bool) error {
	mnqAPI := mnq.NewNatsAPI(scwClient)
	logger.Warningf("sweeper: destroying the mnq nats accounts in (%s)", region)
	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	var projectID *string = nil
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped && exists && defaultProjectID != "" {
		projectID = &defaultProjectID
	}
	listNatsAccounts, err := mnqAPI.ListNatsAccounts(
		&mnq.NatsAPIListNatsAccountsRequest{
			Region:    region,
			ProjectID: projectID,
		}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing nats account in (%s) in sweeper: %s", region, err)
	}

	for _, account := range listNatsAccounts.NatsAccounts {
		err := mnqAPI.DeleteNatsAccount(&mnq.NatsAPIDeleteNatsAccountRequest{
			NatsAccountID: account.ID,
			Region:        region,
		})
		if err != nil {
			logger.Warningf("sweeper: error (%s)", err)

			return fmt.Errorf("error deleting nats account in sweeper: %s", err)
		}
	}

	return nil
}

func SweepAllSNS(scwClient *scw.Client, projectScoped bool) error {
	for _, region := range (&mnq.SnsAPI{}).Regions() {
		err := SweepSNSCredentials(scwClient, region, projectScoped)
		if err != nil {
			return err
		}

		err = SweepSNS(scwClient, region, projectScoped)
		if err != nil {
			return err
		}
	}

	return nil
}

func SweepAllSQS(scwClient *scw.Client, projectScoped bool) error {
	for _, region := range (&mnq.SqsAPI{}).Regions() {
		err := SweepSQSCredentials(scwClient, region, projectScoped)
		if err != nil {
			return err
		}
		err = SweepSQS(scwClient, region, projectScoped)
		if err != nil {
			return err
		}
	}

	return nil
}

func SweepAllNats(scwClient *scw.Client, projectScoped bool) error {
	for _, region := range (&mnq.NatsAPI{}).Regions() {
		err := SweepNatsAccount(scwClient, region, projectScoped)
		if err != nil {
			return err
		}
	}
	return nil
}

func SweepAllLocalities(scwClient *scw.Client, projectScoped bool) error {
	err := SweepAllSNS(scwClient, projectScoped)
	if err != nil {
		return err
	}

	err = SweepAllSQS(scwClient, projectScoped)
	if err != nil {
		return err
	}

	err = SweepAllNats(scwClient, projectScoped)
	if err != nil {
		return err
	}
	return nil
}

package sweepers

import (
	"fmt"
	"strings"

	accountSDK "github.com/scaleway/scaleway-sdk-go/api/account/v3"
	"github.com/scaleway/scaleway-sdk-go/api/cockpit/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepToken(scwClient *scw.Client) error {
	accountAPI := accountSDK.NewProjectAPI(scwClient)
	cockpitAPI := cockpit.NewRegionalAPI(scwClient)

	listProjects, err := accountAPI.ListProjects(&accountSDK.ProjectAPIListProjectsRequest{}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("failed to list projects: %w", err)
	}

	for _, project := range listProjects.Projects {
		if !strings.HasPrefix(project.Name, "tf_tests") {
			continue
		}

		listTokens, err := cockpitAPI.ListTokens(&cockpit.RegionalAPIListTokensRequest{
			ProjectID: project.ID,
		}, scw.WithAllPages())
		if err != nil {
			return fmt.Errorf("failed to list tokens: %w", err)
		}

		for _, token := range listTokens.Tokens {
			err = cockpitAPI.DeleteToken(&cockpit.RegionalAPIDeleteTokenRequest{
				TokenID: token.ID,
			})
			if err != nil {
				return fmt.Errorf("failed to delete token: %w", err)
			}
		}
	}

	return nil
}

func SweepGrafanaUser(scwClient *scw.Client) error {
	accountAPI := accountSDK.NewProjectAPI(scwClient)
	cockpitAPI := cockpit.NewGlobalAPI(scwClient)

	listProjects, err := accountAPI.ListProjects(&accountSDK.ProjectAPIListProjectsRequest{}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("failed to list projects: %w", err)
	}

	for _, project := range listProjects.Projects {
		if !strings.HasPrefix(project.Name, "tf_tests") {
			continue
		}

		listGrafanaUsers, err := cockpitAPI.ListGrafanaUsers(&cockpit.GlobalAPIListGrafanaUsersRequest{
			ProjectID: project.ID,
		}, scw.WithAllPages())
		if err != nil {
			return fmt.Errorf("failed to list grafana users: %w", err)
		}

		for _, grafanaUser := range listGrafanaUsers.GrafanaUsers {
			err = cockpitAPI.DeleteGrafanaUser(&cockpit.GlobalAPIDeleteGrafanaUserRequest{
				ProjectID:     project.ID,
				GrafanaUserID: grafanaUser.ID,
			})
			if err != nil {
				return fmt.Errorf("failed to delete grafana user: %w", err)
			}
		}
	}

	return nil
}

func SweepSource(scwClient *scw.Client, region scw.Region) error {
	accountAPI := accountSDK.NewProjectAPI(scwClient)
	cockpitAPI := cockpit.NewRegionalAPI(scwClient)

	listProjects, err := accountAPI.ListProjects(&accountSDK.ProjectAPIListProjectsRequest{}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("failed to list projects: %w", err)
	}

	for _, project := range listProjects.Projects {
		if !strings.HasPrefix(project.Name, "tf_tests") {
			continue
		}

		listDatasources, err := cockpitAPI.ListDataSources(&cockpit.RegionalAPIListDataSourcesRequest{
			ProjectID: project.ID,
			Region:    region,
		}, scw.WithAllPages())
		if err != nil {
			return fmt.Errorf("failed to list sources: %w", err)
		}

		for _, datsource := range listDatasources.DataSources {
			err = cockpitAPI.DeleteDataSource(&cockpit.RegionalAPIDeleteDataSourceRequest{
				DataSourceID: datsource.ID,
				Region:       region,
			})
			if err != nil {
				return fmt.Errorf("failed to delete cockpit source: %w", err)
			}
		}
	}

	return nil
}

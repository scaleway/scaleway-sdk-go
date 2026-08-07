package sweepers

import (
	"fmt"
	"strings"

	accountSDK "github.com/scaleway/scaleway-sdk-go/api/account/v3"
	"github.com/scaleway/scaleway-sdk-go/api/cockpit/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepToken(scwClient *scw.Client, region scw.Region, projectScoped bool) error {
	accountAPI := accountSDK.NewProjectAPI(scwClient)
	cockpitAPI := cockpit.NewRegionalAPI(scwClient)

	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped {
		listTokens, err := cockpitAPI.ListTokens(&cockpit.RegionalAPIListTokensRequest{
			ProjectID: defaultProjectID,
			Region:    region,
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
	} else {
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
				Region:    region,
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
	}

	return nil
}

func SweepGrafanaUser(scwClient *scw.Client, projectScoped bool) error {
	accountAPI := accountSDK.NewProjectAPI(scwClient)
	cockpitAPI := cockpit.NewGlobalAPI(scwClient)

	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}

	if projectScoped {
		//nolint:staticcheck // ListGrafanaUsers is deprecated but still required for sweeper functionality
		listGrafanaUsers, err := cockpitAPI.ListGrafanaUsers(&cockpit.GlobalAPIListGrafanaUsersRequest{
			ProjectID: defaultProjectID,
		}, scw.WithAllPages())
		if err != nil {
			return fmt.Errorf("failed to list grafana users: %w", err)
		}

		for _, grafanaUser := range listGrafanaUsers.GrafanaUsers {
			//nolint:staticcheck // DeleteGrafanaUser is deprecated but still required for sweeper functionality
			err = cockpitAPI.DeleteGrafanaUser(&cockpit.GlobalAPIDeleteGrafanaUserRequest{
				ProjectID:     defaultProjectID,
				GrafanaUserID: grafanaUser.ID,
			})
			if err != nil {
				return fmt.Errorf("failed to delete grafana user: %w", err)
			}
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

			//nolint:staticcheck // DeleteGrafanaUser is deprecated but still required for sweeper functionality
			listGrafanaUsers, err := cockpitAPI.ListGrafanaUsers(&cockpit.GlobalAPIListGrafanaUsersRequest{
				ProjectID: project.ID,
			}, scw.WithAllPages())
			if err != nil {
				return fmt.Errorf("failed to list grafana users: %w", err)
			}

			for _, grafanaUser := range listGrafanaUsers.GrafanaUsers {
				//nolint:staticcheck // DeleteGrafanaUser is deprecated but still required for sweeper functionality
				err = cockpitAPI.DeleteGrafanaUser(&cockpit.GlobalAPIDeleteGrafanaUserRequest{
					ProjectID:     project.ID,
					GrafanaUserID: grafanaUser.ID,
				})
				if err != nil {
					return fmt.Errorf("failed to delete grafana user: %w", err)
				}
			}
		}
	}

	return nil
}

func SweepSource(scwClient *scw.Client, region scw.Region, projectScoped bool) error {
	accountAPI := accountSDK.NewProjectAPI(scwClient)
	cockpitAPI := cockpit.NewRegionalAPI(scwClient)
	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped {
		listDatasources, err := cockpitAPI.ListDataSources(&cockpit.RegionalAPIListDataSourcesRequest{
			ProjectID: defaultProjectID,
			Region:    region,
		}, scw.WithAllPages())
		if err != nil {
			return fmt.Errorf("failed to list sources: %w", err)
		}

		for _, datasource := range listDatasources.DataSources {
			if datasource.Origin == cockpit.DataSourceOriginScaleway {
				continue
			}
			err = cockpitAPI.DeleteDataSource(&cockpit.RegionalAPIDeleteDataSourceRequest{
				DataSourceID: datasource.ID,
				Region:       region,
			})
			if err != nil {
				return fmt.Errorf("failed to delete cockpit source: %w", err)
			}
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
	}

	return nil
}

func SweepAllLocalities(scwClient *scw.Client, projectScoped bool) error {
	for _, region := range (&cockpit.RegionalAPI{}).Regions() {
		if err := SweepToken(scwClient, region, projectScoped); err != nil {
			return err
		}
	}
	if err := SweepGrafanaUser(scwClient, projectScoped); err != nil {
		return err
	}
	for _, region := range (&cockpit.RegionalAPI{}).Regions() {
		if err := SweepSource(scwClient, region, projectScoped); err != nil {
			return err
		}
	}

	return nil
}

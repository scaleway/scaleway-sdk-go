package sweepers

import (
	"fmt"

	"github.com/scaleway/scaleway-sdk-go/api/account/v3"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepProjects(scwClient *scw.Client) error {
	accountAPI := account.NewProjectAPI(scwClient)

	req := &account.ProjectAPIListProjectsRequest{}
	listProjects, err := accountAPI.ListProjects(req, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("failed to list projects: %w", err)
	}
	for _, project := range listProjects.Projects {
		// Do not delete default project
		if project.ID == req.OrganizationID || !testhelpers.IsTestResource(project.Name) {
			continue
		}
		err = accountAPI.DeleteProject(&account.ProjectAPIDeleteProjectRequest{
			ProjectID: project.ID,
		})
		if err != nil {
			return fmt.Errorf("failed to delete project: %w", err)
		}
	}
	return nil
}

func SweepAll(scwClient *scw.Client) error {
	if err := SweepProjects(scwClient); err != nil {
		return err
	}
	return nil
}

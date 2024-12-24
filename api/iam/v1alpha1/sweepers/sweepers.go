package sweepers

import (
	"errors"
	"fmt"

	iam "github.com/scaleway/scaleway-sdk-go/api/iam/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepUser(scwClient *scw.Client) error {
	api := iam.NewAPI(scwClient)

	orgID, exists := scwClient.GetDefaultOrganizationID()
	if !exists {
		return errors.New("missing organizationID")
	}

	listUsers, err := api.ListUsers(&iam.ListUsersRequest{
		OrganizationID: &orgID,
	})
	if err != nil {
		return fmt.Errorf("failed to list users: %w", err)
	}
	for _, user := range listUsers.Users {
		if !testhelpers.IsTestResource(user.Email) {
			continue
		}
		err = api.DeleteUser(&iam.DeleteUserRequest{
			UserID: user.ID,
		})
		if err != nil {
			return fmt.Errorf("failed to delete user: %w", err)
		}
	}

	return nil
}

func SweepSSHKey(scwClient *scw.Client) error {
	iamAPI := iam.NewAPI(scwClient)

	logger.Warningf("sweeper: destroying the SSH keys")

	listSSHKeys, err := iamAPI.ListSSHKeys(&iam.ListSSHKeysRequest{}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing SSH keys in sweeper: %s", err)
	}

	for _, sshKey := range listSSHKeys.SSHKeys {
		if !testhelpers.IsTestResource(sshKey.Name) {
			continue
		}
		err := iamAPI.DeleteSSHKey(&iam.DeleteSSHKeyRequest{
			SSHKeyID: sshKey.ID,
		})
		if err != nil {
			return fmt.Errorf("error deleting SSH key in sweeper: %s", err)
		}
	}

	return nil
}

func SweepPolicy(scwClient *scw.Client) error {
	api := iam.NewAPI(scwClient)

	orgID, exists := scwClient.GetDefaultOrganizationID()
	if !exists {
		return errors.New("missing organizationID")
	}

	listPols, err := api.ListPolicies(&iam.ListPoliciesRequest{
		OrganizationID: orgID,
	})
	if err != nil {
		return fmt.Errorf("failed to list policies: %w", err)
	}
	for _, pol := range listPols.Policies {
		if !testhelpers.IsTestResource(pol.Name) {
			continue
		}
		err = api.DeletePolicy(&iam.DeletePolicyRequest{
			PolicyID: pol.ID,
		})
		if err != nil {
			return fmt.Errorf("failed to delete policy: %w", err)
		}
	}

	return nil
}

func SweepGroup(scwClient *scw.Client) error {
	api := iam.NewAPI(scwClient)

	orgID, exists := scwClient.GetDefaultOrganizationID()
	if !exists {
		return errors.New("missing organizationID")
	}

	listApps, err := api.ListGroups(&iam.ListGroupsRequest{
		OrganizationID: orgID,
	})
	if err != nil {
		return fmt.Errorf("failed to list groups: %w", err)
	}
	for _, group := range listApps.Groups {
		if !testhelpers.IsTestResource(group.Name) {
			continue
		}
		err = api.DeleteGroup(&iam.DeleteGroupRequest{
			GroupID: group.ID,
		})
		if err != nil {
			return fmt.Errorf("failed to delete group: %w", err)
		}
	}
	return nil
}

func SweepApplication(scwClient *scw.Client) error {
	api := iam.NewAPI(scwClient)

	orgID, exists := scwClient.GetDefaultOrganizationID()
	if !exists {
		return errors.New("missing organizationID")
	}

	listApps, err := api.ListApplications(&iam.ListApplicationsRequest{
		OrganizationID: orgID,
	})
	if err != nil {
		return fmt.Errorf("failed to list applications: %w", err)
	}
	for _, app := range listApps.Applications {
		if !testhelpers.IsTestResource(app.Name) {
			continue
		}

		err = api.DeleteApplication(&iam.DeleteApplicationRequest{
			ApplicationID: app.ID,
		})
		if err != nil {
			return fmt.Errorf("failed to delete application: %w", err)
		}
	}

	return nil
}

func SweepAPIKey(scwClient *scw.Client) error {
	api := iam.NewAPI(scwClient)

	logger.Debugf("sweeper: destroying the api keys")

	orgID, exists := scwClient.GetDefaultOrganizationID()
	if !exists {
		return errors.New("missing organizationID")
	}

	listAPIKeys, err := api.ListAPIKeys(&iam.ListAPIKeysRequest{
		OrganizationID: &orgID,
	}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("failed to list api keys: %w", err)
	}
	for _, key := range listAPIKeys.APIKeys {
		if !testhelpers.IsTestResource(key.Description) {
			continue
		}
		err = api.DeleteAPIKey(&iam.DeleteAPIKeyRequest{
			AccessKey: key.AccessKey,
		})
		if err != nil {
			return fmt.Errorf("failed to delete api key: %w", err)
		}
	}

	return nil
}

func SweepAll(scwClient *scw.Client) error {
	if err := SweepUser(scwClient); err != nil {
		return err
	}
	if err := SweepSSHKey(scwClient); err != nil {
		return err
	}
	if err := SweepPolicy(scwClient); err != nil {
		return err
	}
	if err := SweepGroup(scwClient); err != nil {
		return err
	}
	if err := SweepApplication(scwClient); err != nil {
		return err
	}
	if err := SweepAPIKey(scwClient); err != nil {
		return err
	}

	return nil
}

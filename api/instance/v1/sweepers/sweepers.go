package sweepers

import (
	"fmt"

	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepVolumes(scwClient *scw.Client, zone scw.Zone, projectScoped bool) error {
	instanceAPI := instance.NewAPI(scwClient)
	logger.Warningf("sweeper: destroying the volumes in (%s)", zone)
	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	var projectID *string = nil
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped {
		projectID = &defaultProjectID
	}

	listVolumesResponse, err := instanceAPI.ListVolumes(&instance.ListVolumesRequest{
		Zone:    zone,
		Project: projectID,
	}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing volumes in sweeper: %s", err)
	}

	for _, volume := range listVolumesResponse.Volumes {
		if volume.Server == nil {
			err := instanceAPI.DeleteVolume(&instance.DeleteVolumeRequest{
				Zone:     zone,
				VolumeID: volume.ID,
			})
			if err != nil {
				return fmt.Errorf("error deleting volume in sweeper: %s", err)
			}
		}
	}
	return nil
}

func SweepSnapshots(scwClient *scw.Client, zone scw.Zone, projectScoped bool) error {
	api := instance.NewAPI(scwClient)
	logger.Warningf("sweeper: destroying instance snapshots in (%+v)", zone)
	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	var projectID *string = nil
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped {
		projectID = &defaultProjectID
	}

	listSnapshotsResponse, err := api.ListSnapshots(&instance.ListSnapshotsRequest{
		Zone:    zone,
		Project: projectID,
	}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing instance snapshots in sweeper: %w", err)
	}

	for _, snapshot := range listSnapshotsResponse.Snapshots {
		err := api.DeleteSnapshot(&instance.DeleteSnapshotRequest{
			Zone:       zone,
			SnapshotID: snapshot.ID,
		})
		if err != nil {
			return fmt.Errorf("error deleting instance snapshot in sweeper: %w", err)
		}
	}

	return nil
}

func SweepServers(scwClient *scw.Client, zone scw.Zone, projectScoped bool) error {
	instanceAPI := instance.NewAPI(scwClient)
	logger.Warningf("sweeper: destroying the instance server in (%s)", zone)
	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	var projectID *string = nil
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped {
		projectID = &defaultProjectID
	}
	listServers, err := instanceAPI.ListServers(&instance.ListServersRequest{Zone: zone, Project: projectID}, scw.WithAllPages())
	if err != nil {
		logger.Warningf("error listing servers in (%s) in sweeper: %s", zone, err)
		return nil
	}

	for _, srv := range listServers.Servers {
		switch srv.State {
		case instance.ServerStateStopped, instance.ServerStateStoppedInPlace:
			err := instanceAPI.DeleteServer(&instance.DeleteServerRequest{
				Zone:     zone,
				ServerID: srv.ID,
			})
			if err != nil {
				return fmt.Errorf("error deleting server in sweeper: %w", err)
			}

		case instance.ServerStateRunning:
			_, err := instanceAPI.ServerAction(&instance.ServerActionRequest{
				Zone:     zone,
				ServerID: srv.ID,
				Action:   instance.ServerActionTerminate,
			})
			if err != nil {
				return fmt.Errorf("error terminating server in sweeper: %w", err)
			}
		}
	}

	return nil
}

func SweepSecurityGroups(scwClient *scw.Client, zone scw.Zone, projectScoped bool) error {
	instanceAPI := instance.NewAPI(scwClient)
	logger.Warningf("sweeper: destroying the security groups in (%s)", zone)
	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	var projectID *string = nil
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped {
		projectID = &defaultProjectID
	}

	listResp, err := instanceAPI.ListSecurityGroups(&instance.ListSecurityGroupsRequest{
		Zone:    zone,
		Project: projectID,
	}, scw.WithAllPages())
	if err != nil {
		logger.Warningf("error listing security groups in sweeper: %s", err)
		return nil
	}

	for _, securityGroup := range listResp.SecurityGroups {
		err = instanceAPI.DeleteSecurityGroup(&instance.DeleteSecurityGroupRequest{
			Zone:            zone,
			SecurityGroupID: securityGroup.ID,
		})
		if err != nil {
			return fmt.Errorf("error deleting security groups in sweeper: %s", err)
		}
	}

	return nil
}

func SweepPlacementGroup(scwClient *scw.Client, zone scw.Zone, projectScoped bool) error {
	instanceAPI := instance.NewAPI(scwClient)
	logger.Warningf("sweeper: destroying the instance placement group in (%s)", zone)
	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	var projectID *string = nil
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped {
		projectID = &defaultProjectID
	}
	listPlacementGroups, err := instanceAPI.ListPlacementGroups(&instance.ListPlacementGroupsRequest{
		Zone:    zone,
		Project: projectID,
	}, scw.WithAllPages())
	if err != nil {
		logger.Warningf("error listing placement groups in (%s) in sweeper: %s", zone, err)
		return nil
	}

	for _, pg := range listPlacementGroups.PlacementGroups {
		err := instanceAPI.DeletePlacementGroup(&instance.DeletePlacementGroupRequest{
			Zone:             zone,
			PlacementGroupID: pg.ID,
		})
		if err != nil {
			return fmt.Errorf("error deleting placement group in sweeper: %s", err)
		}
	}

	return nil
}

func SweepIP(scwClient *scw.Client, zone scw.Zone, projectScoped bool) error {
	instanceAPI := instance.NewAPI(scwClient)
	logger.Warningf("sweeper: destroying instance ips in (%+v)", zone)
	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	var projectID *string = nil
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped {
		projectID = &defaultProjectID
	}

	listIPs, err := instanceAPI.ListIPs(&instance.ListIPsRequest{Zone: zone, Project: projectID}, scw.WithAllPages())
	if err != nil {
		logger.Warningf("error listing ips in (%s) in sweeper: %s", zone, err)
		return nil
	}

	for _, ip := range listIPs.IPs {
		err := instanceAPI.DeleteIP(&instance.DeleteIPRequest{
			IP:   ip.ID,
			Zone: zone,
		})
		if err != nil {
			return fmt.Errorf("error deleting ip in sweeper: %s", err)
		}
	}

	return nil
}

func SweepImages(scwClient *scw.Client, zone scw.Zone, projectScoped bool) error {
	api := instance.NewAPI(scwClient)
	logger.Warningf("sweeper: destroying instance images in (%+v)", zone)
	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	var projectID *string = nil
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped {
		projectID = &defaultProjectID
	}

	listImagesResponse, err := api.ListImages(&instance.ListImagesRequest{
		Zone:    zone,
		Public:  scw.BoolPtr(false),
		Project: projectID,
	}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing instance images in sweeper: %w", err)
	}

	for _, image := range listImagesResponse.Images {
		err := api.DeleteImage(&instance.DeleteImageRequest{
			Zone:    zone,
			ImageID: image.ID,
		})
		if err != nil {
			return fmt.Errorf("error deleting instance image in sweeper: %w", err)
		}
	}

	return nil
}

func SweepAllLocalities(scwClient *scw.Client, projectScoped bool) error {
	for _, zone := range (&instance.API{}).Zones() {
		err := SweepVolumes(scwClient, zone, projectScoped)
		if err != nil {
			return err
		}
		err = SweepSnapshots(scwClient, zone, projectScoped)
		if err != nil {
			return err
		}
		err = SweepServers(scwClient, zone, projectScoped)
		if err != nil {
			return err
		}
		err = SweepSecurityGroups(scwClient, zone, projectScoped)
		if err != nil {
			return err
		}
		err = SweepPlacementGroup(scwClient, zone, projectScoped)
		if err != nil {
			return err
		}
		err = SweepIP(scwClient, zone, projectScoped)
		if err != nil {
			return err
		}
		err = SweepImages(scwClient, zone, projectScoped)
		if err != nil {
			return err
		}
	}

	return nil
}

package sweepers

import (
	"fmt"

	block "github.com/scaleway/scaleway-sdk-go/api/block/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepVolumes(scwClient *scw.Client, zone scw.Zone) error {
	blockAPI := block.NewAPI(scwClient)

	listVolumes, err := blockAPI.ListVolumes(
		&block.ListVolumesRequest{
			Zone: zone,
		}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing volume in (%s) in sweeper: %s", zone, err)
	}

	for _, volume := range listVolumes.Volumes {
		err := blockAPI.DeleteVolume(&block.DeleteVolumeRequest{
			VolumeID: volume.ID,
			Zone:     zone,
		})
		if err != nil {
			return fmt.Errorf("error deleting volume in sweeper: %s", err)
		}
	}

	return nil
}

func SweepSnapshots(scwClient *scw.Client, zone scw.Zone) error {
	blockAPI := block.NewAPI(scwClient)

	listSnapshots, err := blockAPI.ListSnapshots(
		&block.ListSnapshotsRequest{
			Zone: zone,
		}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing snapshot in (%s) in sweeper: %s", zone, err)
	}

	for _, snapshot := range listSnapshots.Snapshots {
		err := blockAPI.DeleteSnapshot(&block.DeleteSnapshotRequest{
			SnapshotID: snapshot.ID,
			Zone:       zone,
		})
		if err != nil {
			return fmt.Errorf("error deleting snapshot in sweeper: %s", err)
		}
	}

	return nil
}

func SweepAllLocalities(scwClient *scw.Client) error {
	for _, zone := range (&block.API{}).Zones() {
		err := SweepVolumes(scwClient, zone)
		if err != nil {
			return err
		}
		err = SweepSnapshots(scwClient, zone)
		if err != nil {
			return err
		}
	}

	return nil
}

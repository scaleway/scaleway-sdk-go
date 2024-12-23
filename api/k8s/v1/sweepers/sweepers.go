package sweepers

import (
	"fmt"

	"github.com/scaleway/scaleway-sdk-go/api/k8s/v1"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepCluster(scwClient *scw.Client, region scw.Region) error {
	k8sAPI := k8s.NewAPI(scwClient)

	logger.Warningf("sweeper: destroying the k8s cluster in (%s)", region)
	listClusters, err := k8sAPI.ListClusters(&k8s.ListClustersRequest{Region: region}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing clusters in (%s) in sweeper: %s", region, err)
	}

	for _, cluster := range listClusters.Clusters {
		// remove pools
		listPools, err := k8sAPI.ListPools(&k8s.ListPoolsRequest{
			Region:    region,
			ClusterID: cluster.ID,
		}, scw.WithAllPages())
		if err != nil {
			return fmt.Errorf("error listing pool in (%s) in sweeper: %s", region, err)
		}

		for _, pool := range listPools.Pools {
			_, err := k8sAPI.DeletePool(&k8s.DeletePoolRequest{
				Region: region,
				PoolID: pool.ID,
			})
			if err != nil {
				return fmt.Errorf("error deleting pool in sweeper: %s", err)
			}
		}
		_, err = k8sAPI.DeleteCluster(&k8s.DeleteClusterRequest{
			Region:                  region,
			ClusterID:               cluster.ID,
			WithAdditionalResources: true,
		})
		if err != nil {
			return fmt.Errorf("error deleting cluster in sweeper: %s", err)
		}
	}

	return nil
}

func SweepAllLocalities(scwClient *scw.Client) error {
	for _, region := range (&k8s.API{}).Regions() {
		err := SweepCluster(scwClient, region)
		if err != nil {
			return err
		}
	}
	return nil
}

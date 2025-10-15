package sweepers

import (
	"fmt"

	"github.com/scaleway/scaleway-sdk-go/api/redis/v1"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepCluster(scwClient *scw.Client, zone scw.Zone, projectScoped bool) error {
	redisAPI := redis.NewAPI(scwClient)
	logger.Warningf("sweeper: destroying the redis cluster in (%s)", zone)
	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	var projectID *string = nil
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped {
		projectID = &defaultProjectID
	}
	listClusters, err := redisAPI.ListClusters(&redis.ListClustersRequest{
		Zone:      zone,
		ProjectID: projectID,
	}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing redis clusters in (%s) in sweeper: %w", zone, err)
	}

	for _, cluster := range listClusters.Clusters {
		_, err := redisAPI.DeleteCluster(&redis.DeleteClusterRequest{
			Zone:      zone,
			ClusterID: cluster.ID,
		})
		if err != nil {
			return fmt.Errorf("error deleting redis cluster in sweeper: %w", err)
		}
	}

	return nil
}

func SweepAllLocalities(scwClient *scw.Client, projectScoped bool) error {
	for _, zone := range (&redis.API{}).Zones() {
		err := SweepCluster(scwClient, zone, projectScoped)
		if err != nil {
			return err
		}
	}

	return nil
}

package sweepers

import (
	"fmt"

	webhostingSDK "github.com/scaleway/scaleway-sdk-go/api/webhosting/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepWebHosting(scwClient *scw.Client, region scw.Region) error {
	webhsotingAPI := webhostingSDK.NewAPI(scwClient)

	listHostings, err := webhsotingAPI.ListHostings(&webhostingSDK.ListHostingsRequest{Region: region}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing hostings in (%s) in sweeper: %s", region, err)
	}

	for _, hosting := range listHostings.Hostings {
		_, err := webhsotingAPI.DeleteHosting(&webhostingSDK.DeleteHostingRequest{
			HostingID: hosting.ID,
			Region:    region,
		})
		if err != nil {
			return fmt.Errorf("error deleting hosting in sweeper: %s", err)
		}
	}

	return nil
}

package sweepers

import (
	"fmt"

	"github.com/scaleway/scaleway-sdk-go/api/ipam/v1"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepIP(scwClient *scw.Client, region scw.Region, projectScoped bool) error {
	ipamAPI := ipam.NewAPI(scwClient)

	logger.Warningf("sweeper: deleting the IPs in (%s)", region)
	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	var projectID *string = nil
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped {
		projectID = &defaultProjectID
	}

	listIPs, err := ipamAPI.ListIPs(&ipam.ListIPsRequest{Region: region, ProjectID: projectID}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing ips in (%s) in sweeper: %s", region, err)
	}

	for _, v := range listIPs.IPs {
		err := ipamAPI.ReleaseIP(&ipam.ReleaseIPRequest{
			IPID:   v.ID,
			Region: region,
		})
		if err != nil {
			return fmt.Errorf("error releasing IP in sweeper: %s", err)
		}
	}

	return nil
}

func SweepAllLocalities(scwClient *scw.Client, projectScoped bool) error {
	for _, region := range (&ipam.API{}).Regions() {
		err := SweepIP(scwClient, region, projectScoped)
		if err != nil {
			return err
		}
	}

	return nil
}

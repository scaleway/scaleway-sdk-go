package sweepers

import (
	"fmt"

	"github.com/scaleway/scaleway-sdk-go/api/lb/v1"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepLB(scwClient *scw.Client, zone scw.Zone, projectScoped bool) error {
	lbAPI := lb.NewZonedAPI(scwClient)

	logger.Warningf("sweeper: destroying the lbs in (%s)", zone)
	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	var projectID *string = nil
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped {
		projectID = &defaultProjectID
	}
	listLBs, err := lbAPI.ListLBs(&lb.ZonedAPIListLBsRequest{
		Zone:      zone,
		ProjectID: projectID,
	}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing lbs in (%s) in sweeper: %s", zone, err)
	}

	for _, l := range listLBs.LBs {
		_, err := lbAPI.WaitForLbInstances(&lb.ZonedAPIWaitForLBInstancesRequest{
			Zone: zone,
			LBID: l.ID,
		})
		if err != nil {
			return fmt.Errorf("error waiting for lb in sweeper: %s", err)
		}
		err = lbAPI.DeleteLB(&lb.ZonedAPIDeleteLBRequest{
			LBID:      l.ID,
			ReleaseIP: true,
			Zone:      zone,
		})
		if err != nil {
			return fmt.Errorf("error deleting lb in sweeper: %s", err)
		}
	}

	return nil
}

func SweepIP(scwClient *scw.Client, zone scw.Zone, projectScoped bool) error {
	lbAPI := lb.NewZonedAPI(scwClient)
	logger.Warningf("sweeper: destroying the lb ips in zone (%s)", zone)

	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	var projectID *string = nil
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped {
		projectID = &defaultProjectID
	}
	listIPs, err := lbAPI.ListIPs(&lb.ZonedAPIListIPsRequest{Zone: zone, ProjectID: projectID}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing lb ips in (%s) in sweeper: %s", zone, err)
	}

	for _, ip := range listIPs.IPs {
		if ip.LBID == nil {
			err := lbAPI.ReleaseIP(&lb.ZonedAPIReleaseIPRequest{
				Zone: zone,
				IPID: ip.ID,
			})
			if err != nil {
				return fmt.Errorf("error deleting lb ip in sweeper: %s", err)
			}
		}
	}

	return nil
}

func SweepAllLocalities(scwClient *scw.Client, projectScoped bool) error {
	for _, zone := range (&lb.ZonedAPI{}).Zones() {
		err := SweepLB(scwClient, zone, projectScoped)
		if err != nil {
			return err
		}
		err = SweepIP(scwClient, zone, projectScoped)
		if err != nil {
			return err
		}
	}

	return nil
}

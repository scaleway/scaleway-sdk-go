package sweepers

import (
	"fmt"

	flexibleip "github.com/scaleway/scaleway-sdk-go/api/flexibleip/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepFlexibleIP(scwClient *scw.Client, zone scw.Zone, projectScoped bool) error {
	fipAPI := flexibleip.NewAPI(scwClient)
	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	var projectID *string = nil
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped {
		projectID = &defaultProjectID
	}

	listIPs, err := fipAPI.ListFlexibleIPs(&flexibleip.ListFlexibleIPsRequest{Zone: zone, ProjectID: projectID}, scw.WithAllPages())
	if err != nil {
		logger.Warningf("error listing ips in (%s) in sweeper: %s", zone, err)
		return nil
	}

	for _, ip := range listIPs.FlexibleIPs {
		err := fipAPI.DeleteFlexibleIP(&flexibleip.DeleteFlexibleIPRequest{
			FipID: ip.ID,
			Zone:  zone,
		})
		if err != nil {
			return fmt.Errorf("error deleting ip in sweeper: %s", err)
		}
	}

	return nil
}

func SweepAllLocalities(scwClient *scw.Client, projectScoped bool) error {
	for _, zone := range (&flexibleip.API{}).Zones() {
		err := SweepFlexibleIP(scwClient, zone, projectScoped)
		if err != nil {
			return err
		}
	}

	return nil
}

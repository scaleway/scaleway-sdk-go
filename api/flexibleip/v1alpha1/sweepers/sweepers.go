package sweepers

import (
	"fmt"

	flexibleip "github.com/scaleway/scaleway-sdk-go/api/flexibleip/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepFlexibleIP(scwClient *scw.Client, zone scw.Zone) error {
	fipAPI := flexibleip.NewAPI(scwClient)

	listIPs, err := fipAPI.ListFlexibleIPs(&flexibleip.ListFlexibleIPsRequest{Zone: zone}, scw.WithAllPages())
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

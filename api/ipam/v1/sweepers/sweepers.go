package sweepers

import (
	"fmt"

	"github.com/scaleway/scaleway-sdk-go/api/ipam/v1"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepIP(scwClient *scw.Client, region scw.Region) error {
	ipamAPI := ipam.NewAPI(scwClient)

	logger.Warningf("sweeper: deleting the IPs in (%s)", region)

	listIPs, err := ipamAPI.ListIPs(&ipam.ListIPsRequest{Region: region}, scw.WithAllPages())
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

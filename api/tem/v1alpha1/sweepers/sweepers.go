package sweepers

import (
	"fmt"

	temSDK "github.com/scaleway/scaleway-sdk-go/api/tem/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepDomain(scwClient *scw.Client, region scw.Region, skippedDomain string) error {
	temAPI := temSDK.NewAPI(scwClient)
	logger.Warningf("sweeper: revoking the tem domains in (%s)", region)

	listDomains, err := temAPI.ListDomains(&temSDK.ListDomainsRequest{Region: region}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing domains in (%s) in sweeper: %s", region, err)
	}

	for _, ns := range listDomains.Domains {
		if ns.Name == skippedDomain {
			logger.Debugf("sweeper: skipping deletion of domain %s", ns.Name)
			continue
		}
		_, err := temAPI.RevokeDomain(&temSDK.RevokeDomainRequest{
			DomainID: ns.ID,
			Region:   region,
		})
		if err != nil {
			return fmt.Errorf("error revoking domain in sweeper: %s", err)
		}
	}

	return nil
}

package sweepers

import (
	"fmt"

	webhostingSDK "github.com/scaleway/scaleway-sdk-go/api/webhosting/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SweepWebHosting(scwClient *scw.Client, region scw.Region, projectScoped bool) error {
	webHostingAPI := webhostingSDK.NewHostingAPI(scwClient)
	defaultProjectID, exists := scwClient.GetDefaultProjectID()
	var projectID *string = nil
	if projectScoped && (!exists || (defaultProjectID == "")) {
		return fmt.Errorf("failed to get the default project id for a project scoped sweep")
	}
	if projectScoped {
		projectID = &defaultProjectID
	}

	listHostings, err := webHostingAPI.ListHostings(&webhostingSDK.HostingAPIListHostingsRequest{Region: region, ProjectID: projectID}, scw.WithAllPages())
	if err != nil {
		return fmt.Errorf("error listing hostings in (%s) in sweeper: %s", region, err)
	}

	for _, hosting := range listHostings.Hostings {
		_, err := webHostingAPI.DeleteHosting(&webhostingSDK.HostingAPIDeleteHostingRequest{
			HostingID: hosting.ID,
			Region:    region,
		})
		if err != nil {
			return fmt.Errorf("error deleting hosting in sweeper: %s", err)
		}
	}

	return nil
}

func SweepAllLocalities(scwClient *scw.Client, projectScoped bool) error {
	for _, region := range (&webhostingSDK.HostingAPI{}).Regions() {
		err := SweepWebHosting(scwClient, region, projectScoped)
		if err != nil {
			return err
		}
	}
	return nil
}

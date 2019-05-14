package marketplace

import (
	"fmt"
	"strings"

	"github.com/scaleway/scaleway-sdk-go/utils"
)

// GetImageByName search for an image with the given name (exact match) in the given region
// it returns the latest version of this specific image.
func (s *Api) GetImageByName(imageName string, zone utils.Zone, commercialType string) (string, error) {

	listImageRequest := &ListImagesRequest{}
	listImageResponse, err := s.ListImages(listImageRequest)
	if err != nil {
		return "", err
	}

	images := listImageResponse.Images
	_ = images

	for _, image := range images {

		// Match name of the image
		if strings.ToLower(image.Name) == strings.ToLower(imageName) {

			// Look in all local images of all the versions for a match of the zone and commercial type
			for _, version := range image.Versions {
				for _, localImage := range version.LocalImages {

					// Check if in correct zone
					if localImage.Zone != zone {
						continue
					}

					// Check if compatible with wanted commercial type
					for _, compatibleCommercialType := range localImage.CompatibleCommercialTypes {
						if compatibleCommercialType == commercialType {
							return localImage.Id, nil
						}
					}
				}
			}
		}

	}

	return "", fmt.Errorf("couldn't find a matching image for the given name (%s), zone (%s) and commercial type (%s)", imageName, zone, commercialType)
}

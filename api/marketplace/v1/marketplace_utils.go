package marketplace

import (
	"fmt"

	"github.com/scaleway/scaleway-sdk-go/scw"

	"github.com/scaleway/scaleway-sdk-go/utils"
)

// getLocalImage returns the correct local version of an image matching
// the current zone and the compatible commercial type
func (version *Version) getLocalImage(zone utils.Zone, commercialType string) (*LocalImage, error) {

	for _, localImage := range version.LocalImages {

		// Check if in correct zone
		if localImage.Zone != zone {
			continue
		}

		// Check if compatible with wanted commercial type
		for _, compatibleCommercialType := range localImage.CompatibleCommercialTypes {
			if compatibleCommercialType == commercialType {
				return localImage, nil
			}
		}
	}

	return nil, fmt.Errorf("couldn't find compatible local image for this image version (%s)", version.ID)

}

// getLatestVersion returns the current/latests version on an image,
// or an error in case the image doesn't have a public version.
func (image *Image) getLatestVersion() (*Version, error) {

	for _, version := range image.Versions {
		if version.ID == image.CurrentPublicVersion {
			return version, nil
		}
	}

	return nil, fmt.Errorf("latest version could not be found for image %s", image.Name)
}

// GetLocalImageIDByNameRequest is used by FindLocalImageIDByName
type GetLocalImageIDByNameRequest struct {
	ImageName      string
	Zone           utils.Zone
	CommercialType string
}

// GetLocalImageIDByName search for an image with the given name (exact match) in the given region
// it returns the latest version of this specific image.
func (s *API) GetLocalImageIDByName(req *GetLocalImageIDByNameRequest) (string, error) {

	listImageRequest := &ListImagesRequest{}
	listImageResponse, err := s.ListImages(listImageRequest, scw.WithAllPages())
	if err != nil {
		return "", err
	}

	images := listImageResponse.Images
	_ = images

	for _, image := range images {

		// Match name of the image
		if image.Name == req.ImageName {

			latestVersion, err := image.getLatestVersion()
			if err != nil {
				return "", fmt.Errorf("couldn't find a matching image for the given name (%s), zone (%s) and commercial type (%s): %s", req.ImageName, req.Zone, req.CommercialType, err)
			}

			localImage, err := latestVersion.getLocalImage(req.Zone, req.CommercialType)
			if err != nil {
				return "", fmt.Errorf("couldn't find a matching image for the given name (%s), zone (%s) and commercial type (%s): %s", req.ImageName, req.Zone, req.CommercialType, err)
			}

			return localImage.ID, nil
		}

	}

	return "", fmt.Errorf("couldn't find a matching image for the given name (%s), zone (%s) and commercial type (%s)", req.ImageName, req.Zone, req.CommercialType)
}

// UnsafeSetTotalCount should not be used
// Internal usage only
func (r *ListImagesResponse) UnsafeSetTotalCount(totalCount int) {
	r.TotalCount = uint32(totalCount)
}

// UnsafeSetTotalCount should not be used
// Internal usage only
func (r *ListVersionsResponse) UnsafeSetTotalCount(totalCount int) {
	r.TotalCount = uint32(totalCount)
}

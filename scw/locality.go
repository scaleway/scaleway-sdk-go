package scw

import (
	"encoding/json"

	"github.com/scaleway/scaleway-sdk-go/logger"
)

// Zone is an availability zone
type Zone string

const (
	// ZoneFrPar1 represents the fr-par-1 zone
	ZoneFrPar1 = Zone("fr-par-1")
	// ZoneFrPar2 represents the fr-par-2 zone
	ZoneFrPar2 = Zone("fr-par-2")
	// ZoneNlAms1 represents the nl-ams-1 zone
	ZoneNlAms1 = Zone("nl-ams-1")
)

var (
	// AllZones is an array that list all zones
	AllZones = []Zone{
		ZoneFrPar1,
		ZoneFrPar2,
		ZoneNlAms1,
	}
)

// Exists checks whether a zone exists
func (zone *Zone) Exists() bool {
	for _, z := range AllZones {
		if z == *zone {
			return true
		}
	}
	return false
}

// Region is a geographical location
type Region string

const (
	// RegionFrPar represents the fr-par region
	RegionFrPar = Region("fr-par")
	// RegionNlAms represents the nl-ams region
	RegionNlAms = Region("nl-ams")
)

var (
	// AllRegions is an array that list all regions
	AllRegions = []Region{
		RegionFrPar,
		RegionNlAms,
	}
)

// Exists checks whether a region exists
func (region *Region) Exists() bool {
	for _, r := range AllRegions {
		if r == *region {
			return true
		}
	}
	return false
}

// GetZones is a function that returns the zones for the specified region
func (region Region) GetZones() []Zone {
	switch region {
	case RegionFrPar:
		return []Zone{ZoneFrPar1, ZoneFrPar2}
	case RegionNlAms:
		return []Zone{ZoneNlAms1}
	default:
		return []Zone{}
	}
}

// ParseZone parse a string value into a Zone object
func ParseZone(zone string) (Zone, error) {
	switch zone {
	case "par1":
		// would be triggered by API market place
		// logger.Warningf("par1 is a deprecated name for zone, use fr-par-1 instead")
		return ZoneFrPar1, nil
	case "ams1":
		// would be triggered by API market place
		// logger.Warningf("ams1 is a deprecated name for zone, use nl-ams-1 instead")
		return ZoneNlAms1, nil
	default:
		newZone := Zone(zone)
		if !newZone.Exists() {
			logger.Warningf("%s is an unknown zone", newZone)
		}
		return newZone, nil
	}
}

// UnmarshalJSON implements the Unmarshaler interface for a Zone.
// this to call ParseZone on the string input and return the correct Zone object.
func (zone *Zone) UnmarshalJSON(input []byte) error {

	// parse input value as string
	var stringValue string
	err := json.Unmarshal(input, &stringValue)
	if err != nil {
		return err
	}

	// parse string as Zone
	*zone, err = ParseZone(stringValue)
	if err != nil {
		return err
	}
	return nil
}

// ParseRegion parse a string value into a Zone object
func ParseRegion(region string) (Region, error) {
	switch region {
	case "par1":
		// would be triggered by API market place
		// logger.Warningf("par1 is a deprecated name for region, use fr-par instead")
		return RegionFrPar, nil
	case "ams1":
		// would be triggered by API market place
		// logger.Warningf("ams1 is a deprecated name for region, use nl-ams instead")
		return RegionNlAms, nil
	default:
		newRegion := Region(region)
		if !newRegion.Exists() {
			logger.Warningf("%s is an unknown region", newRegion)
		}
		return newRegion, nil
	}
}

// UnmarshalJSON implements the Unmarshaler interface for a Region.
// this to call ParseRegion on the string input and return the correct Region object.
func (region *Region) UnmarshalJSON(input []byte) error {
	// parse input value as string
	var stringValue string
	err := json.Unmarshal(input, &stringValue)
	if err != nil {
		return err
	}

	// parse string as Region
	*region, err = ParseRegion(stringValue)
	if err != nil {
		return err
	}
	return nil
}

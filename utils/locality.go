package utils

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

// GetZones is a function that returns the zones for the specified region
func (r Region) GetZones() []Zone {
	switch r {
	case RegionFrPar:
		return []Zone{ZoneFrPar1, ZoneFrPar2}
	case RegionNlAms:
		return []Zone{ZoneNlAms1}
	default:
		return []Zone{}
	}
}

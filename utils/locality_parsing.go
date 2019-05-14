package utils

import (
	"encoding/json"
)

// ParseZone parse a string value into a Zone object
func ParseZone(zone string) Zone {
	switch zone {
	case "par1":
		// LEGACY
		return ZoneFrPar1
	case "ams1":
		// LEGACY
		return ZoneNlAms1
	default:
		return Zone(zone)
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
	*zone = ParseZone(stringValue)
	return nil
}

// ParseRegion parse a string value into a Zone object
func ParseRegion(region string) Region {
	switch region {
	case "par1":
		// LEGACY
		return RegionFrPar
	case "ams1":
		// LEGACY
		return RegionNlAms
	default:
		return Region(region)
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
	*region = ParseRegion(stringValue)
	return nil
}

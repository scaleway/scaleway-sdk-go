package scw

// Region is a geographical location
type Region string

const (
	RegionFrPar = Region("fr-par")
	RegionNlAms = Region("nl-ams")
)

// Zone is an availability zone
type Zone string

const (
	ZoneFrPar1 = Zone("fr-par-1")
	ZoneNlAms1 = Zone("nl-ams-1")
)

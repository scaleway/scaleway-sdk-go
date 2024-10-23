package scw_test

import (
	"encoding/json"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func TestParseZone(t *testing.T) {
	tests := []struct {
		input    string
		err      error
		expected scw.Zone
	}{
		{
			input:    "fr-par-1",
			expected: scw.ZoneFrPar1,
		},
		{
			input:    "pl-waw-1",
			expected: scw.ZonePlWaw1,
		},
		{
			input:    "pl-waw-2",
			expected: scw.ZonePlWaw2,
		},
		{
			input:    "pl-waw-3",
			expected: scw.ZonePlWaw3,
		},
		{
			input:    "nl-ams-2",
			expected: scw.ZoneNlAms2,
		},
		{
			input:    "nl-ams-3",
			expected: scw.ZoneNlAms3,
		},
		{
			input:    "par1",
			expected: scw.ZoneFrPar1,
		},
		{
			input:    "ams1",
			expected: scw.ZoneNlAms1,
		},
		{
			input:    "xx-xxx-1",
			expected: "xx-xxx-1",
		},
		{
			input:    "fr-par",
			expected: "",
			err:      errors.New("bad zone format, available zones are: fr-par-1, fr-par-2, fr-par-3, nl-ams-1, nl-ams-2, nl-ams-3, pl-waw-1, pl-waw-2, pl-waw-3"),
		},
		{
			input:    "fr-par-n",
			expected: "",
			err:      errors.New("bad zone format, available zones are: fr-par-1, fr-par-2, fr-par-3, nl-ams-1, nl-ams-2, nl-ams-3, pl-waw-1, pl-waw-2, pl-waw-3"),
		},
		{
			input:    "fr-par-0",
			expected: "",
			err:      errors.New("bad zone format, available zones are: fr-par-1, fr-par-2, fr-par-3, nl-ams-1, nl-ams-2, nl-ams-3, pl-waw-1, pl-waw-2, pl-waw-3"),
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			z, err := scw.ParseZone(test.input)
			testhelpers.Equals(t, test.err, err)
			testhelpers.Equals(t, test.expected, z)
		})
	}
}

func TestZoneJSONUnmarshall(t *testing.T) {
	t.Run("test with zone", func(t *testing.T) {
		input := `{"Test": "par1"}`
		value := struct{ Test scw.Zone }{}

		err := json.Unmarshal([]byte(input), &value)
		testhelpers.AssertNoError(t, err)

		testhelpers.Equals(t, scw.ZoneFrPar1, value.Test)
	})

	t.Run("test with region", func(t *testing.T) {
		input := `{"Test": "par1"}`
		value := struct{ Test scw.Region }{}

		err := json.Unmarshal([]byte(input), &value)
		testhelpers.AssertNoError(t, err)

		testhelpers.Equals(t, scw.RegionFrPar, value.Test)
	})
}

func TestParseRegion(t *testing.T) {
	tests := []struct {
		input    string
		err      error
		expected scw.Region
	}{
		{
			input:    "fr-par",
			expected: scw.RegionFrPar,
		},
		{
			input:    "par1",
			expected: scw.RegionFrPar,
		},
		{
			input:    "ams1",
			expected: scw.RegionNlAms,
		},
		{
			input:    "pl-waw",
			expected: scw.RegionPlWaw,
		},
		{
			input:    "xx-xxx",
			expected: "xx-xxx",
		},
		{
			input:    "fr-par-1",
			expected: "",
			err:      errors.New("bad region format, available regions are: fr-par, nl-ams, pl-waw"),
		},
		{
			input:    "fr-pa1",
			expected: "",
			err:      errors.New("bad region format, available regions are: fr-par, nl-ams, pl-waw"),
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			r, err := scw.ParseRegion(test.input)
			testhelpers.Equals(t, test.err, err)
			testhelpers.Equals(t, test.expected, r)
		})
	}
}

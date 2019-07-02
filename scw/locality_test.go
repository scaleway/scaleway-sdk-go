package scw

import (
	"encoding/json"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

func TestParseZone(t *testing.T) {

	tests := []struct {
		input    string
		expected Zone
	}{
		{
			input:    "fr-par-1",
			expected: ZoneFrPar1,
		},
		{
			input:    "par1",
			expected: ZoneFrPar1,
		},
		{
			input:    "ams1",
			expected: ZoneNlAms1,
		},
	}

	for _, test := range tests {
		z, err := ParseZone(test.input)
		testhelpers.Ok(t, err)
		testhelpers.Equals(t, test.expected, z)
	}

}

func TestZoneJSONUnmarshall(t *testing.T) {

	t.Run("test with zone", func(t *testing.T) {

		input := `{"Test": "par1"}`
		value := struct{ Test Zone }{}

		err := json.Unmarshal([]byte(input), &value)
		testhelpers.Ok(t, err)

		testhelpers.Equals(t, ZoneFrPar1, value.Test)

	})

	t.Run("test with region", func(t *testing.T) {

		input := `{"Test": "par1"}`
		value := struct{ Test Region }{}

		err := json.Unmarshal([]byte(input), &value)
		testhelpers.Ok(t, err)

		testhelpers.Equals(t, RegionFrPar, value.Test)

	})

}

func TestParseRegion(t *testing.T) {

	tests := []struct {
		input    string
		expected Region
	}{
		{
			input:    "fr-par",
			expected: RegionFrPar,
		},
		{
			input:    "par1",
			expected: RegionFrPar,
		},
		{
			input:    "ams1",
			expected: RegionNlAms,
		},
	}

	for _, test := range tests {
		r, err := ParseRegion(test.input)
		testhelpers.Ok(t, err)
		testhelpers.Equals(t, test.expected, r)
	}

}

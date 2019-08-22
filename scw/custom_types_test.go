package scw

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

func TestSize_String(t *testing.T) {
	cases := []struct {
		size Size
		want string
	}{
		{size: 42 * MB, want: "42000000"},
		{size: 42 * B, want: "42"},
	}

	for _, c := range cases {
		t.Run(c.want, func(t *testing.T) {
			testhelpers.Equals(t, c.want, c.size.String())
		})
	}
}

func TestTimeSeries_MarshallJSON(t *testing.T) {
	cases := []struct {
		name string
		ts   *TimeSeries
		want string
		err  error
	}{
		{
			name: "basic",
			ts: &TimeSeries{
				Name: "cpu_usage",
				Points: []*TimeSeriesPoint{
					{
						Timestamp: time.Date(2019, time.August, 8, 15, 00, 00, 0, time.UTC),
						Value:     0.2,
					},
					{
						Timestamp: time.Date(2019, time.August, 8, 15, 01, 00, 0, time.UTC),
						Value:     10.6,
					},
				},
				Metadata: map[string]string{
					"node": "a77e0ce3",
				},
			},
			want: `{"name":"cpu_usage","points":[["2019-08-08T15:00:00Z",0.2],["2019-08-08T15:01:00Z",10.6]],"metadata":{"node":"a77e0ce3"}}`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := json.Marshal(c.ts)

			testhelpers.Equals(t, c.err, err)
			if c.err == nil {
				testhelpers.Equals(t, c.want, string(got))
			}
		})
	}
}

func TestTimeSeries_UnmarshallJSON(t *testing.T) {
	cases := []struct {
		name string
		json string
		want *TimeSeries
		err  error
	}{
		{
			name: "basic",
			json: `
				{
				  "name": "cpu_usage",
				  "points": [
					["2019-08-08T15:00:00Z", 0.2],
					["2019-08-08T15:01:00Z", 10.6]
				  ],
				  "metadata": {
					"node": "a77e0ce3"
				  }
				}
			`,
			want: &TimeSeries{
				Name: "cpu_usage",
				Points: []*TimeSeriesPoint{
					{
						Timestamp: time.Date(2019, time.August, 8, 15, 00, 00, 0, time.UTC),
						Value:     0.2,
					},
					{
						Timestamp: time.Date(2019, time.August, 8, 15, 01, 00, 0, time.UTC),
						Value:     10.6,
					},
				},
				Metadata: map[string]string{
					"node": "a77e0ce3",
				},
			},
		},
		{
			name: "with timestamp error",
			json: `{"name":"cpu_usage","points":[["2019/08/08T15-00-00Z",0.2]]}`,
			err:  fmt.Errorf("2019/08/08T15-00-00Z timestamp is not in RFC 3339 format"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ts := &TimeSeries{}
			err := json.Unmarshal([]byte(c.json), ts)

			testhelpers.Equals(t, c.err, err)
			if c.err == nil {
				testhelpers.Equals(t, c.want, ts)
			}
		})
	}
}

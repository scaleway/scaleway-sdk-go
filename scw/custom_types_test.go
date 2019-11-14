package scw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
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

func TestTimeSeries_UnmarshalJSON(t *testing.T) {
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

func TestFile_UnmarshalJSON(t *testing.T) {

	type testCase struct {
		json        string
		name        string
		contentType string
		content     []byte
	}

	run := func(c *testCase) func(t *testing.T) {
		return func(t *testing.T) {
			f := File{}
			err := json.Unmarshal([]byte(c.json), &f)
			testhelpers.AssertNoError(t, err)
			testhelpers.Equals(t, c.name, f.Name)
			testhelpers.Equals(t, c.contentType, f.ContentType)
			s, err := ioutil.ReadAll(f.Content)
			testhelpers.AssertNoError(t, err)
			testhelpers.Equals(t, c.content, s)
		}
	}

	t.Run("empty", run(&testCase{
		json:    `{}`,
		content: []byte{},
	}))

	t.Run("strint_content", run(&testCase{
		json:        `{"name": "test", "content_type":"text/plain", "content": "dGVzdDQyCg=="}`,
		name:        "test",
		contentType: "text/plain",
		content:     []byte("test42\n"),
	}))

	t.Run("binary_content", run(&testCase{
		json:        `{"name": "test", "content_type":"text/plain", "content": "AAAACg=="}`,
		name:        "test",
		contentType: "text/plain",
		content:     []byte("\x00\x00\x00\n"),
	}))
}

func TestIPNet_MarshallJSON(t *testing.T) {
	cases := []struct {
		name    string
		ipRange IPNet
		want    string
		err     error
	}{
		{
			name:    "ip",
			ipRange: IPNet{IPNet: net.IPNet{IP: net.IPv4(42, 42, 42, 42), Mask: net.CIDRMask(32, 32)}},
			want:    `"42.42.42.42/32"`,
		},
		{
			name:    "network",
			ipRange: IPNet{IPNet: net.IPNet{IP: net.IPv4(42, 42, 42, 42), Mask: net.CIDRMask(16, 32)}},
			want:    `"42.42.42.42/16"`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := json.Marshal(c.ipRange)

			testhelpers.Equals(t, c.err, err)
			if c.err == nil {
				testhelpers.Equals(t, c.want, string(got))
			}
		})
	}
}

func TestIPNet_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		name string
		json string
		want IPNet
		err  string
	}{
		{
			name: "IPv4 with CIDR",
			json: `"42.42.42.42/32"`,
			want: IPNet{IPNet: net.IPNet{IP: net.IPv4(42, 42, 42, 42), Mask: net.CIDRMask(32, 32)}},
		},
		{
			name: "IPv4 with network",
			json: `"192.0.2.1/24"`,
			want: IPNet{IPNet: net.IPNet{IP: net.IPv4(192, 0, 2, 0), Mask: net.CIDRMask(24, 32)}},
		},
		{
			name: "IPv6 with network",
			json: `"2001:db8:abcd:8000::/50"`,
			want: IPNet{IPNet: net.IPNet{IP: net.ParseIP("2001:db8:abcd:8000::"), Mask: net.CIDRMask(50, 128)}},
		},
		{
			name: "IPv4 alone",
			json: `"42.42.42.42"`,
			want: IPNet{IPNet: net.IPNet{IP: net.IPv4(42, 42, 42, 42), Mask: net.CIDRMask(32, 32)}},
		},
		{
			name: "IPv6 alone",
			json: `"2001:db8:abcd:8000::"`,
			want: IPNet{IPNet: net.IPNet{IP: net.ParseIP("2001:db8:abcd:8000::"), Mask: net.CIDRMask(128, 128)}},
		},
		{
			name: "invalid CIDR error",
			json: `"invalidvalue"`,
			err:  "invalid CIDR address: invalidvalue",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ipNet := &IPNet{}
			err := json.Unmarshal([]byte(c.json), ipNet)
			if err != nil {
				testhelpers.Equals(t, c.err, err.Error())
			}

			testhelpers.Equals(t, c.want.String(), ipNet.String())
		})
	}
}

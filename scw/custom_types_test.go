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

func TestMoney_NewMoneyFromFloat(t *testing.T) {
	cases := []struct {
		value     float64
		currency  string
		precision int
		want      *Money
	}{
		{
			value:     0.0,
			currency:  "EUR",
			precision: 0,
			want: &Money{
				CurrencyCode: "EUR",
				Units:        0,
				Nanos:        0,
			},
		},
		{
			value:     1.0,
			currency:  "EUR",
			precision: 3,
			want: &Money{
				CurrencyCode: "EUR",
				Units:        1,
				Nanos:        0,
			},
		},
		{
			value:     1.3,
			currency:  "EUR",
			precision: 3,
			want: &Money{
				CurrencyCode: "EUR",
				Units:        1,
				Nanos:        300000000,
			},
		},
		{
			value:     1.333,
			currency:  "EUR",
			precision: 2,
			want: &Money{
				CurrencyCode: "EUR",
				Units:        1,
				Nanos:        330000000,
			},
		},
		{
			value:     1.04,
			currency:  "EUR",
			precision: 1,
			want: &Money{
				CurrencyCode: "EUR",
				Units:        1,
				Nanos:        0,
			},
		},
		{
			value:     1.05,
			currency:  "EUR",
			precision: 1,
			want: &Money{
				CurrencyCode: "EUR",
				Units:        1,
				Nanos:        100000000,
			},
		},
		{
			value:     1.123456789,
			currency:  "EUR",
			precision: 9,
			want: &Money{
				CurrencyCode: "EUR",
				Units:        1,
				Nanos:        123456789,
			},
		},
		{
			value:     1.999999999,
			currency:  "EUR",
			precision: 9,
			want: &Money{
				CurrencyCode: "EUR",
				Units:        1,
				Nanos:        999999999,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.want.String(), func(t *testing.T) {
			testhelpers.Equals(t, c.want, NewMoneyFromFloat(c.value, c.currency, c.precision))
		})
	}
}

func TestMoney_String(t *testing.T) {
	cases := []struct {
		money *Money
		want  string
	}{
		{
			money: &Money{
				CurrencyCode: "EUR",
				Units:        10,
			},
			want: "€ 10.00",
		},
		{
			money: &Money{
				CurrencyCode: "USD",
				Units:        10,
				Nanos:        1,
			},
			want: "$ 10.000000001",
		},
		{
			money: &Money{
				CurrencyCode: "EUR",
				Nanos:        100000000,
			},
			want: "€ 0.10",
		},
		{
			money: &Money{
				CurrencyCode: "EUR",
				Nanos:        500000,
			},
			want: "€ 0.0005",
		},
		{
			money: &Money{
				CurrencyCode: "EUR",
				Nanos:        333000000,
			},
			want: "€ 0.333",
		},
		{
			money: &Money{
				CurrencyCode: "EUR",
				Nanos:        123456789,
			},
			want: "€ 0.123456789",
		},
		{
			money: &Money{
				CurrencyCode: "?",
			},
			want: "? 0.00",
		},
	}

	for _, c := range cases {
		t.Run(c.want, func(t *testing.T) {
			testhelpers.Equals(t, c.want, c.money.String())
		})
	}
}

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
		{
			name:    "network with ip",
			ipRange: IPNet{IPNet: net.IPNet{IP: net.IPv4(192, 168, 1, 42), Mask: net.CIDRMask(24, 32)}},
			want:    `"192.168.1.42/24"`,
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
			want: IPNet{IPNet: net.IPNet{IP: net.IPv4(192, 0, 2, 1), Mask: net.CIDRMask(24, 32)}},
		},
		{
			name: "IPv4 with network 2",
			json: `"192.168.1.42/24"`,
			want: IPNet{IPNet: net.IPNet{IP: net.IPv4(192, 168, 1, 42), Mask: net.CIDRMask(24, 32)}},
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

func TestDuration_MarshallJSON(t *testing.T) {
	cases := []struct {
		name     string
		duration *Duration
		want     string
		err      error
	}{
		{
			name:     "small seconds",
			duration: &Duration{Seconds: 3, Nanos: 0},
			want:     `"3.000000000s"`,
		},
		{
			name:     "small seconds, small nanos",
			duration: &Duration{Seconds: 3, Nanos: 12e7},
			want:     `"3.120000000s"`,
		},
		{
			name:     "small seconds, big nanos",
			duration: &Duration{Seconds: 3, Nanos: 123456789},
			want:     `"3.123456789s"`,
		},
		{
			name:     "big seconds, big nanos",
			duration: &Duration{Seconds: 345679384, Nanos: 123456789},
			want:     `"345679384.123456789s"`,
		},
		{
			name:     "negative small seconds",
			duration: &Duration{Seconds: -3, Nanos: 0},
			want:     `"-3.000000000s"`,
		},
		{
			name:     "negative small seconds, small nanos",
			duration: &Duration{Seconds: -3, Nanos: -12e7},
			want:     `"-3.120000000s"`,
		},
		{
			name:     "negative small seconds, big nanos",
			duration: &Duration{Seconds: -3, Nanos: -123456789},
			want:     `"-3.123456789s"`,
		},
		{
			name:     "negative big seconds, big nanos",
			duration: &Duration{Seconds: -345679384, Nanos: -123456789},
			want:     `"-345679384.123456789s"`,
		},
		{
			name:     "negative big seconds, big nanos",
			duration: &Duration{},
			want:     `"0.000000000s"`,
		},
		{
			name:     "null duration",
			duration: nil,
			want:     `null`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := json.Marshal(c.duration)

			testhelpers.Equals(t, c.err, err)
			if c.err == nil {
				testhelpers.Equals(t, c.want, string(got))
			}
		})
	}
}

func TestDuration_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		name string
		json string
		want *Duration
		err  string
	}{
		{
			name: "error negative nanos",
			json: `{"duration":"a.12s"}`,
			want: nil,
			err:  "scaleway-sdk-go: invalid units: strconv.ParseInt: parsing \"a\": invalid syntax",
		},
		{
			name: "error negative nanos",
			json: `{"duration":"3.-12s"}`,
			want: nil,
			err:  "scaleway-sdk-go: invalid nanos: strconv.ParseUint: parsing \"-12000000\": invalid syntax",
		},
		{
			name: "null",
			json: `{"duration":null}`,
			want: nil,
		},
		{
			name: "small seconds",
			json: `{"duration":"3.00s"}`,
			want: &Duration{Seconds: 3, Nanos: 0},
		},
		{
			name: "small seconds, small nanos",
			json: `{"duration":"3.12s"}`,
			want: &Duration{Seconds: 3, Nanos: 12e7},
		},
		{
			name: "bug seconds",
			json: `{"duration":"987654321.00s"}`,
			want: &Duration{Seconds: 987654321, Nanos: 0},
		},
		{
			name: "big seconds, big nanos",
			json: `{"duration":"987654321.123456789s"}`,
			want: &Duration{Seconds: 987654321, Nanos: 123456789},
		},
		{
			name: "negative small seconds",
			json: `{"duration":"-3.00s"}`,
			want: &Duration{Seconds: -3, Nanos: 0},
		},
		{
			name: "negative small seconds, small nanos",
			json: `{"duration":"-3.12s"}`,
			want: &Duration{Seconds: -3, Nanos: -12e7},
		},
		{
			name: "negative bug seconds",
			json: `{"duration":"-987654321.00s"}`,
			want: &Duration{Seconds: -987654321, Nanos: 0},
		},
		{
			name: "negative big seconds, big nanos",
			json: `{"duration":"-987654321.123456789s"}`,
			want: &Duration{Seconds: -987654321, Nanos: -123456789},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var testType struct {
				Duration *Duration
			}
			err := json.Unmarshal([]byte(c.json), &testType)
			if err != nil {
				testhelpers.Equals(t, c.err, err.Error())
			} else {
				testhelpers.Equals(t, c.want, testType.Duration)
			}
		})
	}
}

func TestJSONObject_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		name string
		json string
		want *JSONObject
		err  error
	}{
		{
			name: "basic",
			json: `
				{
					"test": "scw"
				}
			`,
			want: &JSONObject{
				"test": "scw",
			},
		},
		{
			name: "multi-types",
			json: `
			{
				"firstName": "John",
				"lastName": "Smith",
				"isAlive": true,
				"age": 23,
				"address": {
					"city": "Paris",
					"country": "FR"
				}
			}
			`,
			want: &JSONObject{
				"firstName": "John",
				"lastName":  "Smith",
				"isAlive":   true,
				"age":       float64(23),
				"address": map[string]interface{}{
					"city":    "Paris",
					"country": "FR",
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ts, err := DecodeJSONObject(c.json, NoEscape)
			testhelpers.Equals(t, c.err, err)
			if c.err == nil {
				testhelpers.Equals(t, *c.want, ts)
			}
		})
	}
}

func TestJSONObject_MarshalJSON(t *testing.T) {
	cases := []struct {
		name      string
		jsonValue *JSONObject
		want      string
		err       error
	}{
		{
			name: "basic",
			jsonValue: &JSONObject{
				"test": "scw",
			},
			want: `{"test":"scw"}`,
		},
		{
			name: "multi-types",
			want: `{"address":{"city":"Paris","country":"FR"},"age":23,"firstName":"John","isAlive":true,"lastName":"Smith"}`,
			jsonValue: &JSONObject{
				"firstName": "John",
				"lastName":  "Smith",
				"isAlive":   true,
				"age":       float64(23),
				"address": map[string]interface{}{
					"city":    "Paris",
					"country": "FR",
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := EncodeJSONObject(*c.jsonValue, NoEscape)
			testhelpers.Equals(t, c.err, err)
			if c.err == nil {
				testhelpers.Equals(t, c.want, string(got))
			}
		})
	}
}

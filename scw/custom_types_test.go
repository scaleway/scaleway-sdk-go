package scw_test

import (
	"encoding/json"
	"errors"
	"io"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func TestMoney_NewMoneyFromFloat(t *testing.T) {
	cases := []struct {
		value     float64
		currency  string
		precision int
		want      *scw.Money
	}{
		{
			value:     0.0,
			currency:  "EUR",
			precision: 0,
			want: &scw.Money{
				CurrencyCode: "EUR",
				Units:        0,
				Nanos:        0,
			},
		},
		{
			value:     1.0,
			currency:  "EUR",
			precision: 3,
			want: &scw.Money{
				CurrencyCode: "EUR",
				Units:        1,
				Nanos:        0,
			},
		},
		{
			value:     1.3,
			currency:  "EUR",
			precision: 3,
			want: &scw.Money{
				CurrencyCode: "EUR",
				Units:        1,
				Nanos:        300000000,
			},
		},
		{
			value:     1.333,
			currency:  "EUR",
			precision: 2,
			want: &scw.Money{
				CurrencyCode: "EUR",
				Units:        1,
				Nanos:        330000000,
			},
		},
		{
			value:     1.04,
			currency:  "EUR",
			precision: 1,
			want: &scw.Money{
				CurrencyCode: "EUR",
				Units:        1,
				Nanos:        0,
			},
		},
		{
			value:     1.05,
			currency:  "EUR",
			precision: 1,
			want: &scw.Money{
				CurrencyCode: "EUR",
				Units:        1,
				Nanos:        100000000,
			},
		},
		{
			value:     1.123456789,
			currency:  "EUR",
			precision: 9,
			want: &scw.Money{
				CurrencyCode: "EUR",
				Units:        1,
				Nanos:        123456789,
			},
		},
		{
			value:     1.999999999,
			currency:  "EUR",
			precision: 9,
			want: &scw.Money{
				CurrencyCode: "EUR",
				Units:        1,
				Nanos:        999999999,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.want.String(), func(t *testing.T) {
			testhelpers.Equals(t, c.want, scw.NewMoneyFromFloat(c.value, c.currency, c.precision))
		})
	}
}

func TestMoney_String(t *testing.T) {
	cases := []struct {
		money *scw.Money
		want  string
	}{
		{
			money: &scw.Money{
				CurrencyCode: "EUR",
				Units:        10,
			},
			want: "€ 10.00",
		},
		{
			money: &scw.Money{
				CurrencyCode: "USD",
				Units:        10,
				Nanos:        1,
			},
			want: "$ 10.000000001",
		},
		{
			money: &scw.Money{
				CurrencyCode: "EUR",
				Nanos:        100000000,
			},
			want: "€ 0.10",
		},
		{
			money: &scw.Money{
				CurrencyCode: "EUR",
				Nanos:        500000,
			},
			want: "€ 0.0005",
		},
		{
			money: &scw.Money{
				CurrencyCode: "EUR",
				Nanos:        333000000,
			},
			want: "€ 0.333",
		},
		{
			money: &scw.Money{
				CurrencyCode: "EUR",
				Nanos:        123456789,
			},
			want: "€ 0.123456789",
		},
		{
			money: &scw.Money{
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
		size scw.Size
		want string
	}{
		{size: 42 * scw.MB, want: "42000000"},
		{size: 42 * scw.B, want: "42"},
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
		ts   *scw.TimeSeries
		want string
		err  error
	}{
		{
			name: "basic",
			ts: &scw.TimeSeries{
				Name: "cpu_usage",
				Points: []*scw.TimeSeriesPoint{
					{
						Timestamp: time.Date(2019, time.August, 8, 15, 0, 0, 0, time.UTC),
						Value:     0.2,
					},
					{
						Timestamp: time.Date(2019, time.August, 8, 15, 1, 0, 0, time.UTC),
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
		want *scw.TimeSeries
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
			want: &scw.TimeSeries{
				Name: "cpu_usage",
				Points: []*scw.TimeSeriesPoint{
					{
						Timestamp: time.Date(2019, time.August, 8, 15, 0o0, 0o0, 0, time.UTC),
						Value:     0.2,
					},
					{
						Timestamp: time.Date(2019, time.August, 8, 15, 0o1, 0o0, 0, time.UTC),
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
			err:  errors.New("2019/08/08T15-00-00Z timestamp is not in RFC 3339 format"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ts := &scw.TimeSeries{}
			err := json.Unmarshal([]byte(c.json), ts)

			testhelpers.Equals(t, c.err, err)
			if c.err == nil {
				testhelpers.Equals(t, c.want, ts)
			}
		})
	}
}

func TestFile_MarshalJSON(t *testing.T) {
	cases := []struct {
		name string
		file *scw.File
		want string
		err  error
	}{
		{
			name: "basic",
			file: &scw.File{
				Name:        "example.txt",
				ContentType: "text/plain",
				Content:     strings.NewReader("Hello, world!"),
			},
			want: `{"name":"example.txt","content_type":"text/plain","content":"Hello, world!"}`,
		},
		{
			name: "empty",
			file: &scw.File{},
			want: `{"name":"","content_type":"","content":""}`,
		},
		{
			name: "nil content",
			file: &scw.File{
				Name:        "example.txt",
				ContentType: "text/plain",
				Content:     nil,
			},
			want: `{"name":"example.txt","content_type":"text/plain","content":""}`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := json.Marshal(c.file)

			testhelpers.Equals(t, c.err, err)
			if c.err == nil {
				testhelpers.Equals(t, c.want, string(got))
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
			t.Helper()
			f := scw.File{}
			err := json.Unmarshal([]byte(c.json), &f)
			testhelpers.AssertNoError(t, err)
			testhelpers.Equals(t, c.name, f.Name)
			testhelpers.Equals(t, c.contentType, f.ContentType)
			s, err := io.ReadAll(f.Content)
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
		ipRange scw.IPNet
		want    string
		err     error
	}{
		{
			name:    "ip",
			ipRange: scw.IPNet{IPNet: net.IPNet{IP: net.IPv4(42, 42, 42, 42), Mask: net.CIDRMask(32, 32)}},
			want:    `"42.42.42.42/32"`,
		},
		{
			name:    "network",
			ipRange: scw.IPNet{IPNet: net.IPNet{IP: net.IPv4(42, 42, 42, 42), Mask: net.CIDRMask(16, 32)}},
			want:    `"42.42.42.42/16"`,
		},
		{
			name:    "network with ip",
			ipRange: scw.IPNet{IPNet: net.IPNet{IP: net.IPv4(192, 168, 1, 42), Mask: net.CIDRMask(24, 32)}},
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
		want scw.IPNet
		err  string
	}{
		{
			name: "IPv4 with CIDR",
			json: `"42.42.42.42/32"`,
			want: scw.IPNet{IPNet: net.IPNet{IP: net.IPv4(42, 42, 42, 42), Mask: net.CIDRMask(32, 32)}},
		},
		{
			name: "IPv4 with network",
			json: `"192.0.2.1/24"`,
			want: scw.IPNet{IPNet: net.IPNet{IP: net.IPv4(192, 0, 2, 1), Mask: net.CIDRMask(24, 32)}},
		},
		{
			name: "IPv4 with network 2",
			json: `"192.168.1.42/24"`,
			want: scw.IPNet{IPNet: net.IPNet{IP: net.IPv4(192, 168, 1, 42), Mask: net.CIDRMask(24, 32)}},
		},
		{
			name: "IPv6 with network",
			json: `"2001:db8:abcd:8000::/50"`,
			want: scw.IPNet{IPNet: net.IPNet{IP: net.ParseIP("2001:db8:abcd:8000::"), Mask: net.CIDRMask(50, 128)}},
		},
		{
			name: "IPv4 alone",
			json: `"42.42.42.42"`,
			want: scw.IPNet{IPNet: net.IPNet{IP: net.IPv4(42, 42, 42, 42), Mask: net.CIDRMask(32, 32)}},
		},
		{
			name: "IPv6 alone",
			json: `"2001:db8:abcd:8000::"`,
			want: scw.IPNet{IPNet: net.IPNet{IP: net.ParseIP("2001:db8:abcd:8000::"), Mask: net.CIDRMask(128, 128)}},
		},
		{
			name: "invalid CIDR error",
			json: `"invalidvalue"`,
			err:  "invalid CIDR address: invalidvalue",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ipNet := &scw.IPNet{}
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
		duration *scw.Duration
		want     string
		err      error
	}{
		{
			name:     "small seconds",
			duration: &scw.Duration{Seconds: 3, Nanos: 0},
			want:     `"3.000000000s"`,
		},
		{
			name:     "small seconds, small nanos",
			duration: &scw.Duration{Seconds: 3, Nanos: 12e7},
			want:     `"3.120000000s"`,
		},
		{
			name:     "small seconds, big nanos",
			duration: &scw.Duration{Seconds: 3, Nanos: 123456789},
			want:     `"3.123456789s"`,
		},
		{
			name:     "big seconds, big nanos",
			duration: &scw.Duration{Seconds: 345679384, Nanos: 123456789},
			want:     `"345679384.123456789s"`,
		},
		{
			name:     "negative small seconds",
			duration: &scw.Duration{Seconds: -3, Nanos: 0},
			want:     `"-3.000000000s"`,
		},
		{
			name:     "negative small seconds, small nanos",
			duration: &scw.Duration{Seconds: -3, Nanos: -12e7},
			want:     `"-3.120000000s"`,
		},
		{
			name:     "negative small seconds, big nanos",
			duration: &scw.Duration{Seconds: -3, Nanos: -123456789},
			want:     `"-3.123456789s"`,
		},
		{
			name:     "negative big seconds, big nanos",
			duration: &scw.Duration{Seconds: -345679384, Nanos: -123456789},
			want:     `"-345679384.123456789s"`,
		},
		{
			name:     "negative big seconds, big nanos",
			duration: &scw.Duration{},
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
		want *scw.Duration
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
			want: &scw.Duration{Seconds: 3, Nanos: 0},
		},
		{
			name: "small seconds, small nanos",
			json: `{"duration":"3.12s"}`,
			want: &scw.Duration{Seconds: 3, Nanos: 12e7},
		},
		{
			name: "bug seconds",
			json: `{"duration":"987654321.00s"}`,
			want: &scw.Duration{Seconds: 987654321, Nanos: 0},
		},
		{
			name: "big seconds, big nanos",
			json: `{"duration":"987654321.123456789s"}`,
			want: &scw.Duration{Seconds: 987654321, Nanos: 123456789},
		},
		{
			name: "negative small seconds",
			json: `{"duration":"-3.00s"}`,
			want: &scw.Duration{Seconds: -3, Nanos: 0},
		},
		{
			name: "negative small seconds, small nanos",
			json: `{"duration":"-3.12s"}`,
			want: &scw.Duration{Seconds: -3, Nanos: -12e7},
		},
		{
			name: "negative bug seconds",
			json: `{"duration":"-987654321.00s"}`,
			want: &scw.Duration{Seconds: -987654321, Nanos: 0},
		},
		{
			name: "negative big seconds, big nanos",
			json: `{"duration":"-987654321.123456789s"}`,
			want: &scw.Duration{Seconds: -987654321, Nanos: -123456789},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var testType struct {
				Duration *scw.Duration
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

func TestDuration_ToTimeDuration(t *testing.T) {
	cases := []struct {
		name     string
		duration *scw.Duration
		want     time.Duration
	}{
		{
			name:     "nil duration",
			duration: nil,
			want:     time.Duration(0),
		},
		{
			name:     "zero duration",
			duration: &scw.Duration{Seconds: 0, Nanos: 0},
			want:     time.Duration(0),
		},
		{
			name:     "seconds only",
			duration: &scw.Duration{Seconds: 10, Nanos: 0},
			want:     time.Duration(10) * time.Second,
		},
		{
			name:     "nanoseconds only",
			duration: &scw.Duration{Seconds: 0, Nanos: 500},
			want:     time.Duration(500),
		},
		{
			name:     "seconds and nanoseconds",
			duration: &scw.Duration{Seconds: 10, Nanos: 500},
			want:     time.Duration(10)*time.Second + time.Duration(500),
		},
		{
			name:     "negative seconds",
			duration: &scw.Duration{Seconds: -10, Nanos: 0},
			want:     time.Duration(-10) * time.Second,
		},
		{
			name:     "negative nanoseconds",
			duration: &scw.Duration{Seconds: 0, Nanos: -500},
			want:     time.Duration(-500),
		},
		{
			name:     "negative seconds and nanoseconds",
			duration: &scw.Duration{Seconds: -10, Nanos: -500},
			want:     time.Duration(-10)*time.Second + time.Duration(-500),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.duration.ToTimeDuration()
			if got == nil {
				if c.want != 0 {
					t.Errorf("got nil, want %s", c.want)
				}
			} else if *got != c.want {
				t.Errorf("got %s, want %s", *got, c.want)
			}
		})
	}
}

func TestDuration_FromTimeDuration(t *testing.T) {
	cases := []struct {
		name     string
		duration time.Duration
		want     *scw.Duration
	}{
		{
			name:     "zero duration",
			want:     &scw.Duration{Seconds: 0, Nanos: 0},
			duration: time.Duration(0),
		},
		{
			name:     "seconds only",
			want:     &scw.Duration{Seconds: 10, Nanos: 0},
			duration: time.Duration(10) * time.Second,
		},
		{
			name:     "nanoseconds only",
			want:     &scw.Duration{Seconds: 0, Nanos: 500},
			duration: time.Duration(500),
		},
		{
			name:     "seconds and nanoseconds",
			want:     &scw.Duration{Seconds: 10, Nanos: 500},
			duration: time.Duration(10)*time.Second + time.Duration(500),
		},
		{
			name:     "negative seconds",
			want:     &scw.Duration{Seconds: -10, Nanos: 0},
			duration: time.Duration(-10) * time.Second,
		},
		{
			name:     "negative nanoseconds",
			want:     &scw.Duration{Seconds: 0, Nanos: -500},
			duration: time.Duration(-500),
		},
		{
			name:     "negative seconds and nanoseconds",
			want:     &scw.Duration{Seconds: -10, Nanos: -500},
			duration: time.Duration(-10)*time.Second + time.Duration(-500),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := scw.NewDurationFromTimeDuration(c.duration)
			if got == nil {
				t.Errorf("got nil, want %v", c.want)
			} else if got.Seconds != c.want.Seconds && got.Nanos != c.want.Nanos {
				t.Errorf("got %v, want %v", *got, c.want)
			}
		})
	}
}

func TestJSONObject_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		name string
		json string
		want *scw.JSONObject
		err  error
	}{
		{
			name: "basic",
			json: `
				{
					"test": "scw"
				}
			`,
			want: &scw.JSONObject{
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
			want: &scw.JSONObject{
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
			ts, err := scw.DecodeJSONObject(c.json, scw.NoEscape)
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
		jsonValue *scw.JSONObject
		want      string
		err       error
	}{
		{
			name: "basic",
			jsonValue: &scw.JSONObject{
				"test": "scw",
			},
			want: `{"test":"scw"}`,
		},
		{
			name: "multi-types",
			want: `{"address":{"city":"Paris","country":"FR"},"age":23,"firstName":"John","isAlive":true,"lastName":"Smith"}`,
			jsonValue: &scw.JSONObject{
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
			got, err := scw.EncodeJSONObject(*c.jsonValue, scw.NoEscape)
			testhelpers.Equals(t, c.err, err)
			if c.err == nil {
				testhelpers.Equals(t, c.want, got)
			}
		})
	}
}

func TestDecimal(t *testing.T) {
	d := scw.Decimal("1.22")
	testhelpers.Equals(t, "1.22", d.String())

	dPtr := new(scw.Decimal)
	testhelpers.Equals(t, "", dPtr.String())

	*dPtr = "1.22"
	testhelpers.Equals(t, "1.22", dPtr.String())
}

func TestDecimal_MarshalJSON(t *testing.T) {
	d := Decimal("1.22")
	testhelpers.Equals(t, "1.22", d.String())

	value, err := json.Marshal(d)
	testhelpers.AssertNoError(t, err)
	testhelpers.Equals(t, "{\"value\":\"1.22\"}", string(value))
}

func TestDecimal_UnmarshalJSON(t *testing.T) {
	value := "{\"value\":\"1.22\"}"
	d := new(Decimal)
	err := json.Unmarshal([]byte(value), d)
	testhelpers.AssertNoError(t, err)
	testhelpers.Equals(t, "1.22", d.String())
}

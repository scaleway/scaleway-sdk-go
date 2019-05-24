package async

import (
	"fmt"
	"testing"
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

const flakiness = time.Millisecond * 500

type value struct {
	doneIterations int
	totalDuration  time.Duration
}

func getMock(iterations int, sleepTime time.Duration) func() (interface{}, error, bool) {
	cpt := iterations
	var startTime time.Time

	return func() (interface{}, error, bool) {
		if cpt == iterations {
			startTime = time.Now()
		}
		cpt -= 1

		// fake working time
		time.Sleep(sleepTime)

		v := &value{
			doneIterations: iterations - cpt,
			totalDuration:  time.Since(startTime),
		}
		return v, nil, cpt == 0
	}
}

func getFibonacciInterval() func() <-chan time.Time {
	x, y := 0, 1
	return func() <-chan time.Time {
		x, y = y, x+y
		return time.After(time.Duration(x) * time.Second)
	}
}

func TestWaitSync(t *testing.T) {
	testsCases := []struct {
		name     string
		config   *WaitSyncConfig
		expValue interface{}
		expErr   error
	}{
		{
			name: "With default timeout and interval",
			config: &WaitSyncConfig{
				get: getMock(2, 0),
			},
			expValue: &value{
				doneIterations: 2,
				totalDuration:  time.Second,
			},
		},
		{
			name: "With useless timeout",
			config: &WaitSyncConfig{
				get:     getMock(2, time.Second),
				timeout: 4 * time.Second,
			},
			expValue: &value{
				doneIterations: 2,
				totalDuration:  3 * time.Second,
			},
		},
		{
			name: "Should timeout",
			config: &WaitSyncConfig{
				get:     getMock(2, 2*time.Second),
				timeout: time.Second,
			},
			expValue: nil,
			expErr:   fmt.Errorf("timeout after 1s"),
		},
		{
			name: "With interval",
			config: &WaitSyncConfig{
				get: getMock(2, 0),
				interval: func() <-chan time.Time {
					return time.After(time.Second * 2)
				},
			},
			expValue: &value{
				doneIterations: 2,
				totalDuration:  2 * time.Second,
			},
		},
		{
			name: "With fibonacci interval",
			config: &WaitSyncConfig{
				get:      getMock(5, 0),
				interval: getFibonacciInterval(),
			},
			expValue: &value{
				doneIterations: 5,
				totalDuration:  7 * time.Second,
			},
		},
		{
			name: "Should timeout with interval",
			config: &WaitSyncConfig{
				get:     getMock(2, time.Second),
				timeout: 2 * time.Second,
				interval: func() <-chan time.Time {
					return time.After(time.Second * 2)
				},
			},
			expValue: nil,
			expErr:   fmt.Errorf("timeout after 2s"),
		},
	}
	for _, c := range testsCases {
		c := c // do not remove me
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			terminalValue, err := WaitSync(c.config)

			testhelpers.Equals(t, c.expErr, err)

			if c.expValue != nil {
				exp := c.expValue.(*value)
				acc := terminalValue.(*value)
				testhelpers.Equals(t, exp.doneIterations, acc.doneIterations)

				ok := exp.totalDuration > acc.totalDuration-flakiness && exp.totalDuration < acc.totalDuration+flakiness
				testhelpers.Assert(t, ok, "totalDuration don't match the target: (acc: %v, exp: %v)", acc.totalDuration, exp.totalDuration)
			}
		})
	}
}

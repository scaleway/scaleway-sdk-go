package async_test

import (
	"errors"
	"testing"
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

const flakiness = 500 * time.Millisecond

type value struct {
	doneIterations int
	totalDuration  time.Duration
}

func getMock(iterations int, sleepTime time.Duration) func() (interface{}, bool, error) {
	cpt := iterations
	var startTime time.Time

	return func() (interface{}, bool, error) {
		if cpt == iterations {
			startTime = time.Now()
		}
		cpt--

		// fake working time
		time.Sleep(sleepTime)

		v := &value{
			doneIterations: iterations - cpt,
			totalDuration:  time.Since(startTime),
		}
		return v, cpt == 0, nil
	}
}

func TestWaitSync(t *testing.T) {
	t.Parallel()
	testsCases := []struct {
		name     string
		config   *async.WaitSyncConfig
		expValue interface{}
		expErr   error
	}{
		{
			name: "With default timeout and interval",
			config: &async.WaitSyncConfig{
				Get: getMock(2, 0),
			},
			expValue: &value{
				doneIterations: 2,
				totalDuration:  time.Second,
			},
		},
		{
			name: "With useless timeout",
			config: &async.WaitSyncConfig{
				Get:     getMock(2, time.Second),
				Timeout: 4 * time.Second,
			},
			expValue: &value{
				doneIterations: 2,
				totalDuration:  3 * time.Second,
			},
		},
		{
			name: "Should timeout",
			config: &async.WaitSyncConfig{
				Get:     getMock(2, 2*time.Second),
				Timeout: time.Second,
			},
			expValue: nil,
			expErr:   errors.New("timeout after 1s"),
		},
		{
			name: "With interval",
			config: &async.WaitSyncConfig{
				Get:              getMock(2, 0),
				IntervalStrategy: async.LinearIntervalStrategy(2 * time.Second),
			},
			expValue: &value{
				doneIterations: 2,
				totalDuration:  2 * time.Second,
			},
		},
		{
			name: "With fibonacci interval",
			config: &async.WaitSyncConfig{
				Get:              getMock(5, 0),
				IntervalStrategy: async.FibonacciIntervalStrategy(time.Second, 1),
			},
			expValue: &value{
				doneIterations: 5,
				totalDuration:  7 * time.Second,
			},
		},
		{
			name: "Should timeout with interval",
			config: &async.WaitSyncConfig{
				Get:              getMock(2, time.Second),
				Timeout:          2 * time.Second,
				IntervalStrategy: async.LinearIntervalStrategy(2 * time.Second),
			},
			expValue: nil,
			expErr:   errors.New("timeout after 2s"),
		},
	}
	for _, c := range testsCases {
		c := c // do not remove me
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			terminalValue, err := async.WaitSync(c.config)

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

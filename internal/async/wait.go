package async

import (
	"fmt"
	"time"
)

var (
	defaultInterval = time.Second
	defaultTimeout  = time.Minute * 5
)

type IntervalStrategy func() <-chan time.Time

// WaitSyncConfig defines the waiting options.
type WaitSyncConfig struct {
	// This method will be called from another goroutine.
	Get              func() (value any, isTerminal bool, err error)
	IntervalStrategy IntervalStrategy
	Timeout          time.Duration
}

// LinearIntervalStrategy defines a linear interval duration.
func LinearIntervalStrategy(interval time.Duration) IntervalStrategy {
	return func() <-chan time.Time {
		return time.After(interval)
	}
}

// FibonacciIntervalStrategy defines an interval duration who follow the Fibonacci sequence.
func FibonacciIntervalStrategy(base time.Duration, factor float32) IntervalStrategy {
	var x, y float32 = 0, 1

	return func() <-chan time.Time {
		x, y = y, x+(y*factor)
		return time.After(time.Duration(x) * base)
	}
}

// WaitSync waits and returns when a given stop condition is true or if an error occurs.
func WaitSync(config *WaitSyncConfig) (terminalValue any, err error) {
	// initialize configuration
	if config.IntervalStrategy == nil {
		config.IntervalStrategy = LinearIntervalStrategy(defaultInterval)
	}

	if config.Timeout == 0 {
		config.Timeout = defaultTimeout
	}

	resultValue := make(chan any, 1)
	resultErr := make(chan error, 1)
	timeoutDone := make(chan struct{})

	go func() {
		for {
			// get the payload
			value, stopCondition, err := config.Get()
			// send the payload
			if err != nil {
				resultErr <- err
				return
			}
			if stopCondition {
				resultValue <- value
				return
			}

			// waiting for an interval before next get() call or a timeout
			select {
			case <-timeoutDone:
				return
			case <-config.IntervalStrategy():
				// sleep
			}
		}
	}()

	// waiting for a result or a timeout
	select {
	case val := <-resultValue:
		return val, nil
	case err := <-resultErr:
		return nil, err
	case <-time.After(config.Timeout):
		// Notify that the timeout is reached to stop the goroutine.
		close(timeoutDone)
		// Return immediately without waiting for the goroutine to stop, because it will be stopped at the next interval.
		// We accept that the goroutine may leak for a short time.
		return nil, fmt.Errorf("timeout after %v", config.Timeout)
	}
}

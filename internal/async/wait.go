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
	get      func() (value interface{}, err error, isTerminal bool)
	interval IntervalStrategy
	timeout  time.Duration
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
func WaitSync(config *WaitSyncConfig) (terminalValue interface{}, err error) {
	// initialize configuration
	if config.interval == nil {
		config.interval = LinearIntervalStrategy(defaultInterval)
	}
	if config.timeout == 0 {
		config.timeout = defaultTimeout
	}

	resultValue := make(chan interface{})
	resultErr := make(chan error)
	timeout := make(chan bool)

	go func() {
		for {
			// get the payload
			value, err, stopCondition := config.get()

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
			case <-timeout:
				return
			case <-config.interval():
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
	case <-time.After(config.timeout):
		timeout <- true
		return nil, fmt.Errorf("timeout after %v", config.timeout)
	}
}

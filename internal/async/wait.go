package async

import (
	"fmt"
	"time"
)

var (
	defaultInterval = time.Second
	defaultTimeout  = time.Minute * 5
)

// WaitOptions defines the waiting options.
type WaitSyncConfig struct {
	get      func() (value interface{}, err error, stopCondition bool)
	interval func() <-chan time.Time
	timeout  time.Duration
}

// WaitSync waits and returns when a given stop condition is true or if an error occurs.
//
// In this sync implementation, the timeout could be longer than expected.
func WaitSync(config *WaitSyncConfig) (terminalValue interface{}, err error) {
	// initialize configuration
	if config.interval == nil {
		config.interval = func() <-chan time.Time {
			return time.After(defaultInterval)
		}
	}
	if config.timeout == 0 {
		config.timeout = defaultTimeout
	}

	type payload struct {
		value interface{}
		err   error
	}

	result := make(chan payload, 1)
	timeout := make(chan bool, 1)

	go func() {
		for {
			// get the payload
			value, err, stopCondition := config.get()
			p := payload{value, err}

			// send the payload
			if err != nil || stopCondition {
				result <- p
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
	case res := <-result:
		if res.err != nil {
			return nil, err
		}
		return res.value, nil
	case <-time.After(config.timeout):
		timeout <- true
		return nil, fmt.Errorf("timeout after %v", config.timeout)
	}
}

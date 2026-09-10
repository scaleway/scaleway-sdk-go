package scw

import (
	"context"
	"crypto/tls"
	"net/http"
	"strconv"
	"time"
)

const defaultMaxRetries = 5

// RateLimitTransport implements http.RoundTripper
type RateLimitTransport struct {
	// Base contains a default RoundTripper that we use in our custom
	// RoundTrip method
	Base http.RoundTripper

	// State stores the current limit state values
	State *RateLimitState

	// MaxRetries is the maximum number of 429 retries before returning
	// the response to the caller. Defaults to 5 if zero.
	MaxRetries int
}

// base returns the transport : safety net in case RateLimitTransport was
// created manually without a Base.
func (t *RateLimitTransport) base() http.RoundTripper {
	if t.Base != nil {
		return t.Base
	}

	return http.DefaultTransport
}

// RoundTrip intercepts each SDK request
func (t *RateLimitTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	maxRetries := t.MaxRetries
	if maxRetries == 0 {
		maxRetries = defaultMaxRetries
	}

	for retry := 0; ; retry++ {
		// Proactive strategy: check if we should wait
		wait := t.State.GetWaitDuration()
		if wait > 0 {
			if err := sleepWithContext(req.Context(), wait); err != nil {
				return nil, err
			}
		}

		resp, err := t.base().RoundTrip(req)
		if err != nil {
			return resp, err
		}

		// State update
		remainingStr := resp.Header.Get("X-Ratelimit-Remaining")
		resetStr := resp.Header.Get("X-Ratelimit-Reset")

		if remainingStr != "" && resetStr != "" {
			remaining, _ := strconv.Atoi(remainingStr)
			resetSec, _ := strconv.Atoi(resetStr)
			t.State.Update(remaining, resetSec)
		}

		// Reactive strategy: 429 handling
		if resp.StatusCode == http.StatusTooManyRequests && retry < maxRetries {
			// This is necessary since the body of a successful HTTP request is consumed
			if req.GetBody != nil {
				req.Body, err = req.GetBody()
				if err != nil {
					return nil, err
				}
			}

			retryAfterStr := resp.Header.Get("Retry-After")
			if retryAfterSec, err := strconv.Atoi(retryAfterStr); err == nil {
				err = resp.Body.Close()
				if err != nil {
					return nil, err
				}

				retryWait := time.Duration(retryAfterSec) * time.Second
				if err := sleepWithContext(req.Context(), retryWait); err != nil {
					return nil, err
				}

				continue
			}
		}

		// Success, or 404, 500... or exceeded t.MaxRetries of 429s
		return resp, nil
	}
}

func (t *RateLimitTransport) SetInsecureTransport() {
	// Need to cast to *http.Transport to access TLSClientConfig.
	// If the cast fails, probably doesn't have a TLSClientConfig anyway.
	transport, ok := t.Base.(*http.Transport)
	if !ok {
		return
	}

	if transport.TLSClientConfig == nil {
		transport.TLSClientConfig = &tls.Config{}
	}
	transport.TLSClientConfig.InsecureSkipVerify = true
}

// sleepWithContext pauses the goroutine, while still listening to context's cancellation.
func sleepWithContext(ctx context.Context, d time.Duration) error {
	timer := time.NewTimer(d)
	defer timer.Stop()

	select {
	case <-timer.C:
		// Timer went up
		return nil
	case <-ctx.Done():
		// Request was cancelled
		return ctx.Err()
	}
}

package scw

import (
	"context"
	"crypto/tls"
	"net/http"
	"strconv"
	"time"
)

const (
	defaultMaxRetries     = 5
	defaultAttemptTimeout = 30 * time.Second

	// minBackoffSeconds is the minimum time to wait after a 429, independently
	// from the RetryAfter header
	minBackoffSeconds = 1
	// maxBackoffSeconds is the maximum time to wait after a 429
	maxBackoffSeconds = 30
)

// RateLimitTransport implements http.RoundTripper
type RateLimitTransport struct {
	// Base contains a default RoundTripper that we use in our custom
	// RoundTrip method
	Base http.RoundTripper

	// State stores the current limit state values
	State *RateLimitState

	// MaxRetries is the maximum number of 429 retries before returning
	// the response to the caller. Set to 0 to disable retries entirely.
	// The SDK default is 5 (set in newHTTPClient); callers can override it
	// by accessing the transport via client.Transport.(*RateLimitTransport).
	MaxRetries int

	// AttemptTimeout is the per-HTTP-attempt timeout. Defaults to 30s if
	// zero. Retry sleeps use the caller's context, not this budget, so a
	// long Retry-After does not starve the actual HTTP call.
	AttemptTimeout time.Duration
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

	attemptTimeout := t.AttemptTimeout
	if attemptTimeout == 0 {
		attemptTimeout = defaultAttemptTimeout
	}

	// The caller's context owns the total time cancellation and is not consumed
	// by the sleeps.
	parentCtx := req.Context()

	for retry := 0; ; retry++ {
		// Proactive strategy: check if we should wait. Uses parentCtx so
		// the wait respects caller cancellation but not the per-attempt
		// deadline.
		if wait := t.State.GetWaitDuration(); wait > 0 {
			if err := sleepWithContext(parentCtx, wait); err != nil {
				return nil, err
			}
		}

		// Run the request with a fresh per-attempt deadline.
		attemptCtx, cancel := context.WithTimeout(parentCtx, attemptTimeout)
		resp, err := t.base().RoundTrip(req.WithContext(attemptCtx))
		cancel()
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
			// Only retry when the body can be safely replayed
			if req.GetBody == nil && req.Body != nil {
				return resp, nil
			}

			if req.GetBody != nil {
				req.Body, err = req.GetBody()
				if err != nil {
					return nil, err
				}
			}

			// Per RFC 9110, Retry-After may also be an HTTP-date, but Scaleway's API
			// only specified integers.
			retryAfterStr := resp.Header.Get("Retry-After")
			if retryAfterSec, err := strconv.Atoi(retryAfterStr); err == nil {
				err = resp.Body.Close()
				if err != nil {
					return nil, err
				}

				// In case Retry-After: 0, we still want to backoff the next try a bit
				if floor := minRetryBackoffSeconds(retry); retryAfterSec < floor {
					retryAfterSec = floor
				}

				// Update state to proactively wait in the next loop iteration
				t.State.Update(0, retryAfterSec)

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

// minRetryBackoffSeconds returns the minimum backoff in seconds for a given
// retry attempt. Grows exponentially (1, 2, 4, 8, 16, capped at 30) so repeated
// 429s with Retry-After: 0 don't hammer the endpoint.
func minRetryBackoffSeconds(retry int) int {
	s := minBackoffSeconds << retry
	if s > maxBackoffSeconds {
		s = maxBackoffSeconds
	}
	return s
}

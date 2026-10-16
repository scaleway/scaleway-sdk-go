package scw

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/scaleway/scaleway-sdk-go/logger"
)

const (
	defaultMaxRetries = 5

	// minBackoffSeconds is the minimum time to wait after a 429, independently
	// from the RetryAfter header
	minBackoffSeconds = 1
	// maxBackoffSeconds is the maximum time to wait after a 429
	maxBackoffSeconds = 30

	// Per-attempt phase timeouts, applied to the base *http.Transport.
	// We bind these on the transport rather than via context.WithTimeout so
	// that cancelling the attempt context does not tear down resp.Body before
	// the caller reads it.
	dialTimeout           = 30 * time.Second
	tlsHandshakeTimeout   = 10 * time.Second
	responseHeaderTimeout = 30 * time.Second
	expectContinueTimeout = 1 * time.Second
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
}

// RoundTrip intercepts each SDK request
func (t *RateLimitTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	maxRetries := t.MaxRetries

	// The caller's context owns the total time cancellation and is not consumed
	// by the sleeps. Per-attempt phase timeouts live on the base *http.Transport
	// (dial, TLS, response-header), so we never wrap the request in a context
	// whose cancellation would tear down resp.Body before the caller reads it.
	parentCtx := req.Context()

	for retry := 0; ; retry++ {
		// Proactive strategy: check if we should wait. Uses parentCtx so
		// the wait respects caller cancellation.
		if wait := t.State.GetWaitDuration(); wait > 0 {
			if err := sleepWithContext(parentCtx, wait); err != nil {
				return nil, err
			}
		}

		resp, err := t.Base.RoundTrip(req)
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
					_ = resp.Body.Close()
					return nil, err
				}
			}

			// Per RFC 9110, Retry-After may also be an HTTP-date, but Scaleway's API
			// only specified integers.
			retryAfterStr := resp.Header.Get("Retry-After")
			if retryAfterSec, err := strconv.Atoi(retryAfterStr); err == nil {
				logger.Debugf(
					"client: got 429 Too Many Requests, retrying (attempt %d/%d) in %ds\n",
					retry+1, maxRetries, retryAfterSec,
				)

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

	setInsecureSkipVerify(transport)
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
	return min(minBackoffSeconds<<retry, maxBackoffSeconds)
}

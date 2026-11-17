package scw

import (
	"context"
	"math"
	"net/http"
	"strconv"
	"strings"
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

	// defaultClientTimeout bounds the whole Do call: attempts, retry sleeps
	// and the response body read. It must comfortably exceed the worst
	// legitimate rate-limit sequence (5 retries x 30s backoff plus attempt
	// time, about 3 minutes).
	defaultClientTimeout = 5 * time.Minute
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

		// State update. Malformed values are skipped, never zero-filled — a garbage
		// header must not stall the client.
		remaining, remainingOK := parseRateLimitValue(resp.Header.Get("X-Ratelimit-Remaining"))
		resetSec, resetOK := parseRateLimitValue(resp.Header.Get("X-Ratelimit-Reset"))
		if remainingOK && resetOK {
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

			// Parse Retry-After: should be an integer according to Scaleway's API.
			// Defaults to 0 in case of parsing error.
			// Note: Per RFC 9110, Retry-After may also be an HTTP-date, but we don't
			// handle that case.
			retryAfterSec := 0
			if v, err := strconv.Atoi(resp.Header.Get("Retry-After")); err == nil {
				retryAfterSec = v
			}
			retryAfterSec = clampRetryAfterSeconds(retryAfterSec, retry)

			err = resp.Body.Close()
			if err != nil {
				return resp, err
			}

			logger.Debugf(
				"client: got 429 Too Many Requests, retrying (attempt %d/%d) in %ds\n",
				retry+1, maxRetries, retryAfterSec,
			)

			t.State.Update(0, retryAfterSec)

			continue
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

// clampRetryAfterSeconds returns the number of seconds to actually wait before
// a 429 retry.
// - retryAfterSec: the time to wait the server asked for
// - retry: the current attempt number
// It guarantees the wait is never too small (hammering) and never too large (a
// bogus Retry-After: 86400 stalling the whole client).
func clampRetryAfterSeconds(retryAfterSec, retry int) int {
	if floor := minBackoffSeconds << min(retry, 5); retryAfterSec < floor {
		retryAfterSec = floor
	}

	return min(retryAfterSec, maxBackoffSeconds)
}

// parseRateLimitValue parses an X-Ratelimit-Remaining or X-Ratelimit-Reset
// header value. We return the most restrictive (minimum) entry.
func parseRateLimitValue(value string) (int, bool) {
	minValue := math.MaxInt
	found := false
	for _, entry := range strings.Split(value, ",") {
		if i := strings.IndexByte(entry, ';'); i >= 0 {
			entry = entry[:i] // strip draft parameters, e.g. "50;w=1"
		}

		n, err := strconv.Atoi(strings.TrimSpace(entry))
		if err != nil || n < 0 {
			continue
		}

		minValue = min(minValue, n)
		found = true
	}

	return minValue, found
}

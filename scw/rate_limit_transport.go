package scw

import (
	"context"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
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

	// responseHeaderTimeout is the per-attempt timeout for time-to-first-header
	responseHeaderTimeout = 30 * time.Second

	// defaultClientTimeout bounds the whole Do call: attempts, retry sleeps
	// and the response body read. It must comfortably exceed the worst
	// legitimate rate-limit sequence (5 retries x 30s backoff plus attempt
	// time, about 3 minutes).
	defaultClientTimeout = 5 * time.Minute
)

// rateLimitState stores the current limit state
type rateLimitState struct {
	// mu protects rateLimitState from concurrent edits by different
	// goroutines
	mu sync.Mutex

	// remaining is the number of requests still available for this time window
	remaining int

	// resetTime is the time of window change
	resetTime time.Time
}

// update updates the limit state with the headers values
func (s *rateLimitState) update(remaining int, resetInSeconds int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Cap resetInSeconds to guard against a reset value in an unexpected format
	resetInSeconds = min(resetInSeconds, maxBackoffSeconds)
	s.remaining = remaining
	s.resetTime = time.Now().Add(time.Duration(resetInSeconds) * time.Second)
}

// getWaitDuration computes the time to wait before the next request
func (s *rateLimitState) getWaitDuration() time.Duration {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.remaining > 0 {
		return 0
	}

	now := time.Now()
	if s.resetTime.After(now) {
		return s.resetTime.Sub(now)
	}

	return 0
}

// RateLimitTransport implements http.RoundTripper
type RateLimitTransport struct {
	Base http.RoundTripper

	// State stores the current limit state values
	State *rateLimitState

	// MaxRetries is the maximum number of 429 retries before returning
	// the response to the caller. Defaults to 5. Can be overridden by accessing
	// the transport via client.Transport.(*RateLimitTransport).
	MaxRetries int
}

// RoundTrip intercepts each SDK request
func (t *RateLimitTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Use the caller's context to respect its cancellation
	parentCtx := req.Context()

	for retry := 0; ; retry++ {
		// Proactive strategy: check if we should wait
		if wait := t.State.getWaitDuration(); wait > 0 {
			if err := sleepWithContext(parentCtx, wait); err != nil {
				return nil, err
			}
		}

		resp, err := t.Base.RoundTrip(req)
		if err != nil {
			return resp, err
		}

		remaining, resetSec, ok := parseRateLimitHeaders(
			resp.Header.Get("X-Ratelimit-Remaining"),
			resp.Header.Get("X-Ratelimit-Reset"),
		)
		if ok {
			t.State.update(remaining, resetSec)
		}

		// Reactive strategy: 429 handling
		if resp.StatusCode == http.StatusTooManyRequests && retry < t.MaxRetries {
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
				retry+1, t.MaxRetries, retryAfterSec,
			)

			t.State.update(0, retryAfterSec)

			continue
		}

		// Success, or 404, 500... or exceeded t.MaxRetries of 429s
		return resp, nil
	}
}

func (t *RateLimitTransport) SetInsecureTransport() {
	transport, ok := t.Base.(*http.Transport)
	if !ok {
		// Probably doesn't have a TLSClientConfig anyway
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

// parseRateLimitHeaders parses paired X-Ratelimit-Remaining and X-Ratelimit-Reset
// header values.
// It returns the most restrictive remaining count and the latest reset (in
// seconds) among exhausted buckets, or ok=false if no pair parses.
func parseRateLimitHeaders(remainingStr, resetStr string) (remaining, resetSec int, ok bool) {
	remainingList := strings.Split(remainingStr, ",")
	resetList := strings.Split(resetStr, ",")

	remaining = math.MaxInt
	for i, remEntry := range remainingList {
		if i >= len(resetList) {
			break
		}

		rem, remOK := parseRateLimitEntry(remEntry)
		reset, resetOK := parseRateLimitEntry(resetList[i])
		if !remOK || !resetOK {
			continue
		}

		remaining = min(remaining, rem)
		if rem == 0 {
			resetSec = max(resetSec, reset)
		}
		ok = true
	}

	return remaining, resetSec, ok
}

// parseRateLimitEntry parses one comma-list entry of an X-Ratelimit-* header.
func parseRateLimitEntry(entry string) (int, bool) {
	if i := strings.IndexByte(entry, ';'); i >= 0 {
		entry = entry[:i] // strip draft parameters, e.g. ";w=1"
	}

	n, err := strconv.Atoi(strings.TrimSpace(entry))
	if err != nil || n < 0 {
		return 0, false
	}

	return n, true
}

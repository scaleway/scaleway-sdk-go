package scw

import (
	"context"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// RateLimitState stores the current limit state
type RateLimitState struct {
	mu        sync.RWMutex
	remaining int
	resetAt   time.Time
}

// Update updates the limit state with the headers values
func (s *RateLimitState) Update(remaining int, resetInSeconds int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.remaining = remaining
	s.resetAt = time.Now().Add(time.Duration(resetInSeconds) * time.Second)
}

// WaitDuration computes the time to wait before the next request
func (s *RateLimitState) WaitDuration() time.Duration {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.remaining > 0 {
		return 0
	}

	now := time.Now()
	if now.Before(s.resetAt) {
		return s.resetAt.Sub(now)
	}

	return 0
}

// RateLimitTransport implements http.RoundTripper
type RateLimitTransport struct {
	Base  http.RoundTripper
	State *RateLimitState
}

// base returns the transport
func (t *RateLimitTransport) base() http.RoundTripper {
	if t.Base != nil {
		return t.Base
	}
	return http.DefaultTransport
}

// RoundTrip intercepts each SDK request
func (t *RateLimitTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for {
		// Proactive strategy: check if we should wait
		wait := t.State.WaitDuration()
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
		if resp.StatusCode == http.StatusTooManyRequests {
			retryAfterStr := resp.Header.Get("Retry-After")
			if retryAfterSec, err := strconv.Atoi(retryAfterStr); err == nil {
				resp.Body.Close()
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

		// Success, or 404, 500...
		return resp, nil
	}
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

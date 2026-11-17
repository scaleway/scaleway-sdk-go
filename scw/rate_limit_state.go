package scw

import (
	"sync"
	"time"
)

// RateLimitState stores the current limit state
type RateLimitState struct {
	// mu protects RateLimitState from concurrent edits by different
	// goroutines
	mu sync.Mutex

	// remaining is the number of requests still available for this time window
	remaining int

	// resetTime is the time of window change
	resetTime time.Time
}

// Update updates the limit state with the headers values
func (s *RateLimitState) Update(remaining int, resetInSeconds int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Cap resetInSeconds to guard against a reset value in an unexpected format
	resetInSeconds = min(resetInSeconds, maxBackoffSeconds)
	s.remaining = remaining
	s.resetTime = time.Now().Add(time.Duration(resetInSeconds) * time.Second)
}

// GetWaitDuration computes the time to wait before the next request
func (s *RateLimitState) GetWaitDuration() time.Duration {
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

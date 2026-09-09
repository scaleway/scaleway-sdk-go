package scw

import (
	"sync"
	"time"
)

// RateLimitState stores the current limit state
type RateLimitState struct {
	// mu protects RateLimitState from concurrent edits by different
	// goroutines
	mu sync.RWMutex

	// remaining is the number of requests still available for this time window
	remaining int

	// resetAt is the time to wait before the next time window
	resetAt time.Time
}

// Update updates the limit state with the headers values
func (s *RateLimitState) Update(remaining int, resetInSeconds int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.remaining = remaining
	s.resetAt = time.Now().Add(time.Duration(resetInSeconds) * time.Second)
}

// GetWaitDuration computes the time to wait before the next request
func (s *RateLimitState) GetWaitDuration() time.Duration {
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

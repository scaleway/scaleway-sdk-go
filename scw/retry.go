package scw

import (
	"bytes"
	"context"
	stderrors "errors"
	"io"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	sdkerrors "github.com/scaleway/scaleway-sdk-go/errors"
	"github.com/scaleway/scaleway-sdk-go/logger"
)

// RetryPolicy defines how the client retries transient HTTP failures.
type RetryPolicy struct {
	// MaxAttempts is the total number of attempts, including the initial one.
	// A value of 1 disables retries. Values <= 0 default to 1.
	MaxAttempts int

	// MinDelay is the initial backoff delay between two attempts.
	// Defaults to 100ms.
	MinDelay time.Duration

	// MaxDelay caps the backoff delay. Defaults to 30s.
	MaxDelay time.Duration

	// RetryBudget is the total number of retries shared across all requests
	// made through the same client. A value <= 0 disables the budget (retries
	// are only constrained by MaxAttempts). The budget is refilled over time
	// at a rate of RetryRefillRate tokens per second, up to RetryBudget.
	RetryBudget int

	// RetryRefillRate is the number of retry tokens added back to the budget
	// per second. Defaults to RetryBudget/10.
	RetryRefillRate float64
}

// DefaultRetryPolicy returns a sane RetryPolicy used when none is explicitly
// provided. It is not enabled by default: callers must opt in via
// WithRetryPolicy.
func DefaultRetryPolicy() *RetryPolicy {
	return &RetryPolicy{
		MaxAttempts: 5,
		MinDelay:    100 * time.Millisecond,
		MaxDelay:    30 * time.Second,
	}
}

// defaultRetryableStatusCodes is the set of HTTP status codes that are
// considered transient and worth retrying.
var defaultRetryableStatusCodes = map[int]struct{}{
	http.StatusTooManyRequests:     {},
	http.StatusServiceUnavailable:  {},
	http.StatusGatewayTimeout:      {},
	http.StatusBadGateway:          {},
	http.StatusRequestTimeout:      {},
	http.StatusConflict:            {},
	http.StatusInternalServerError: {},
}

// isRetryableHTTPStatus reports whether the given status code is considered
// transient.
func isRetryableHTTPStatus(statusCode int) bool {
	_, ok := defaultRetryableStatusCodes[statusCode]
	return ok
}

// isRetryableNetworkError reports whether a network error returned by the
// underlying http.Client is worth retrying.
func isRetryableNetworkError(err error) bool {
	if err == nil {
		return false
	}
	// Context cancellation/deadline errors are not retryable: the caller
	// explicitly asked to stop.
	if stderrors.Is(err, context.Canceled) || stderrors.Is(err, context.DeadlineExceeded) {
		return false
	}
	return true
}

// retryBudget is a thread-safe token bucket shared across all requests made
// through a single client. Each retry consumes one token; tokens are refilled
// over time so that bursts of failures do not exhaust the budget forever.
type retryBudget struct {
	mu         sync.Mutex
	tokens     float64
	maxTokens  float64
	refillRate float64 // tokens per second
	lastRefill time.Time
}

// newRetryBudget builds a retryBudget from a RetryPolicy. Returns nil if the
// policy has no budget configured.
func newRetryBudget(p *RetryPolicy) *retryBudget {
	if p == nil || p.RetryBudget <= 0 {
		return nil
	}
	refill := p.RetryRefillRate
	if refill <= 0 {
		refill = float64(p.RetryBudget) / 10.0
	}
	return &retryBudget{
		tokens:     float64(p.RetryBudget),
		maxTokens:  float64(p.RetryBudget),
		refillRate: refill,
		lastRefill: time.Now(),
	}
}

// tryAcquire consumes one token. It returns false if no token is available.
func (b *retryBudget) tryAcquire() bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.refill()
	if b.tokens < 1 {
		return false
	}
	b.tokens--
	return true
}

// refill adds tokens based on elapsed time since the last refill.
func (b *retryBudget) refill() {
	now := time.Now()
	elapsed := now.Sub(b.lastRefill).Seconds()
	if elapsed > 0 {
		b.tokens = math.Min(b.maxTokens, b.tokens+elapsed*b.refillRate)
		b.lastRefill = now
	}
}

// parseRetryAfter parses the Retry-After header. It supports both delta
// seconds and HTTP-date formats. Returns 0 and false if the header is absent
// or unparseable.
func parseRetryAfter(h http.Header) (time.Duration, bool) {
	val := h.Get("Retry-After")
	if val == "" {
		return 0, false
	}
	// Delta seconds.
	if secs, err := strconv.Atoi(val); err == nil {
		if secs < 0 {
			secs = 0
		}
		return time.Duration(secs) * time.Second, true
	}
	// HTTP-date.
	if t, err := http.ParseTime(val); err == nil {
		d := time.Until(t)
		if d < 0 {
			d = 0
		}
		return d, true
	}
	return 0, false
}

// computeBackoff returns the delay to wait before the next attempt. It uses an
// exponential backoff with full jitter, capped by MaxDelay. If the server
// provided a Retry-After header, it takes precedence (but is still capped).
func (p *RetryPolicy) computeBackoff(attempt int, res *http.Response) time.Duration {
	if attempt < 0 {
		attempt = 0
	}
	// Honor Retry-After when present.
	if res != nil {
		if d, ok := parseRetryAfter(res.Header); ok {
			if p.MaxDelay > 0 && d > p.MaxDelay {
				return p.MaxDelay
			}
			return d
		}
	}

	minDelay := p.MinDelay
	if minDelay <= 0 {
		minDelay = 100 * time.Millisecond
	}
	maxDelay := p.MaxDelay
	if maxDelay <= 0 {
		maxDelay = 30 * time.Second
	}
	if maxDelay < minDelay {
		maxDelay = minDelay
	}

	// Exponential backoff: min * 2^attempt, capped at max.
	backoff := minDelay << uint(attempt)
	if backoff > maxDelay || backoff < 0 {
		backoff = maxDelay
	}
	// Full jitter.
	if backoff > 0 {
		backoff = time.Duration(rand.Int63n(int64(backoff)))
	}
	return backoff
}

// shouldRetry reports whether an attempt should be retried according to the
// policy, the resulting error and the HTTP response.
func (p *RetryPolicy) shouldRetry(err error, res *http.Response) bool {
	if err != nil {
		return isRetryableNetworkError(err)
	}
	if res == nil {
		return false
	}
	return isRetryableHTTPStatus(res.StatusCode)
}

// doWithRetry executes fn with retry/backoff according to the policy. fn is
// expected to perform a single HTTP attempt and return either an error or a
// response. The response body, if any, is returned to the caller untouched.
//
// The body of the *http.Request passed to fn must be re-readable between
// attempts: callers should ensure req.Body is an io.ReadSeeker or buffer it
// before calling doWithRetry.
func (p *RetryPolicy) doWithRetry(ctx context.Context, budget *retryBudget, fn func(req *http.Request) (*http.Response, error), req *http.Request) (*http.Response, error) {
	maxAttempts := p.MaxAttempts
	if maxAttempts <= 0 {
		maxAttempts = 1
	}

	var lastErr error
	var lastRes *http.Response

	for attempt := 0; attempt < maxAttempts; attempt++ {
		// Reset a seekable body before every attempt (except the first one,
		// which has not been consumed yet).
		if attempt > 0 {
			if seeker, ok := req.Body.(io.Seeker); ok {
				if _, err := seeker.Seek(0, io.SeekStart); err != nil {
					return nil, sdkerrors.Wrap(err, "could not rewind request body for retry")
				}
			}
		}

		res, err := fn(req)
		lastErr = err
		lastRes = res

		if !p.shouldRetry(err, res) {
			return res, err
		}

		// Last attempt: do not wait.
		if attempt == maxAttempts-1 {
			break
		}

		// Check the shared retry budget.
		if budget != nil && !budget.tryAcquire() {
			logger.Debugf("retry: budget exhausted, stopping after %d attempts\n", attempt+1)
			break
		}

		// Drain and close the response body before retrying so that the
		// underlying connection can be reused. This must happen before the
		// backoff sleep to release the connection promptly.
		if res != nil && res.Body != nil {
			_, _ = io.Copy(io.Discard, res.Body)
			_ = res.Body.Close()
		}

		delay := p.computeBackoff(attempt, res)

		logger.Debugf("retry: attempt %d failed, retrying in %s\n", attempt+1, delay)

		select {
		case <-time.After(delay):
		case <-ctx.Done():
			return res, ctx.Err()
		}
	}

	return lastRes, lastErr
}

// bufferBody reads the request body into memory and replaces req.Body with a
// bytes.Reader so that it can be re-read on each retry attempt. It is a no-op
// for nil bodies and for bodies that already implement io.ReadSeeker.
func bufferBody(req *http.Request) error {
	if req.Body == nil {
		return nil
	}
	if _, ok := req.Body.(io.Seeker); ok {
		return nil
	}
	buf, err := io.ReadAll(req.Body)
	if err != nil {
		return sdkerrors.Wrap(err, "could not buffer request body for retry")
	}
	if err := req.Body.Close(); err != nil {
		return sdkerrors.Wrap(err, "could not close request body")
	}
	req.Body = newReadSeekCloser(bytes.NewReader(buf))
	req.ContentLength = int64(len(buf))
	return nil
}

// readSeekCloser wraps a *bytes.Reader to implement io.ReadSeekCloser so that
// the request body can be rewound between retry attempts.
type readSeekCloser struct {
	*bytes.Reader
}

func newReadSeekCloser(r *bytes.Reader) io.ReadSeekCloser {
	return &readSeekCloser{Reader: r}
}

func (readSeekCloser) Close() error { return nil }

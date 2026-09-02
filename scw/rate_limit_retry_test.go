package scw

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

func TestRateLimitTransportThrottles(t *testing.T) {
	var calls int32
	// Server that records how many concurrent/total requests it receives.
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	// Build a client with a very low rate: 2 rps, burst 1.
	client, err := NewClient(
		WithAPIURL(srv.URL),
		WithoutAuth(),
		WithRateLimit(2, 1),
	)
	testhelpers.AssertNoError(t, err)

	start := time.Now()
	for i := 0; i < 4; i++ {
		req := &ScalewayRequest{Method: http.MethodGet, Path: "/"}
		err := client.Do(req, nil)
		testhelpers.AssertNoError(t, err)
	}
	elapsed := time.Since(start)

	// 4 requests at 2 rps with burst 1 should take roughly 1s
	// (1 immediate, then ~0.5s, ~0.5s, ~0.5s). Allow some slack.
	testhelpers.Assert(t, elapsed >= 900*time.Millisecond, "expected rate limiting to delay requests, got %s", elapsed)
	testhelpers.Equals(t, int32(4), atomic.LoadInt32(&calls))
}

func TestNewClientWithRateLimitWrapsTransport(t *testing.T) {
	client, err := NewClient(
		WithoutAuth(),
		WithRateLimit(1, 1),
	)
	testhelpers.AssertNoError(t, err)

	stdClient, ok := client.httpClient.(*http.Client)
	testhelpers.Assert(t, ok, "httpClient should be a *http.Client")
	_, ok = stdClient.Transport.(*rateLimitTransport)
	testhelpers.Assert(t, ok, "transport should be a rateLimitTransport")
}

func TestNewClientWithoutRateLimitDoesNotWrap(t *testing.T) {
	client, err := NewClient(WithoutAuth())
	testhelpers.AssertNoError(t, err)

	stdClient, ok := client.httpClient.(*http.Client)
	testhelpers.Assert(t, ok, "httpClient should be a *http.Client")
	_, ok = stdClient.Transport.(*rateLimitTransport)
	testhelpers.Assert(t, !ok, "transport should not be a rateLimitTransport")
}

func TestRetryPolicyRetriesOn5xx(t *testing.T) {
	var attempts int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt32(&attempts, 1)
		if n < 3 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer srv.Close()

	client, err := NewClient(
		WithAPIURL(srv.URL),
		WithoutAuth(),
		WithRetryPolicy(&RetryPolicy{
			MaxAttempts: 5,
			MinDelay:    time.Millisecond,
			MaxDelay:    10 * time.Millisecond,
		}),
	)
	testhelpers.AssertNoError(t, err)

	req := &ScalewayRequest{Method: http.MethodGet, Path: "/"}
	err = client.Do(req, nil)
	testhelpers.AssertNoError(t, err)
	testhelpers.Equals(t, int32(3), atomic.LoadInt32(&attempts))
}

func TestRetryPolicyStopsAfterMaxAttempts(t *testing.T) {
	var attempts int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&attempts, 1)
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer srv.Close()

	client, err := NewClient(
		WithAPIURL(srv.URL),
		WithoutAuth(),
		WithRetryPolicy(&RetryPolicy{
			MaxAttempts: 3,
			MinDelay:    time.Millisecond,
			MaxDelay:    10 * time.Millisecond,
		}),
	)
	testhelpers.AssertNoError(t, err)

	req := &ScalewayRequest{Method: http.MethodGet, Path: "/"}
	err = client.Do(req, nil)
	testhelpers.Assert(t, err != nil, "expected an error after exhausting retries")
	testhelpers.Equals(t, int32(3), atomic.LoadInt32(&attempts))
}

func TestRetryPolicyDoesNotRetryOn4xx(t *testing.T) {
	var attempts int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&attempts, 1)
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer srv.Close()

	client, err := NewClient(
		WithAPIURL(srv.URL),
		WithoutAuth(),
		WithRetryPolicy(&RetryPolicy{
			MaxAttempts: 5,
			MinDelay:    time.Millisecond,
			MaxDelay:    10 * time.Millisecond,
		}),
	)
	testhelpers.AssertNoError(t, err)

	req := &ScalewayRequest{Method: http.MethodGet, Path: "/"}
	err = client.Do(req, nil)
	testhelpers.Assert(t, err != nil, "expected a 400 error")
	testhelpers.Equals(t, int32(1), atomic.LoadInt32(&attempts))
}

func TestRetryPolicyRetriesWithBody(t *testing.T) {
	var attempts int32
	var lastBody string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt32(&attempts, 1)
		body, _ := io.ReadAll(r.Body)
		lastBody = string(body)
		if n < 2 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	client, err := NewClient(
		WithAPIURL(srv.URL),
		WithoutAuth(),
		WithRetryPolicy(&RetryPolicy{
			MaxAttempts: 3,
			MinDelay:    time.Millisecond,
			MaxDelay:    10 * time.Millisecond,
		}),
	)
	testhelpers.AssertNoError(t, err)

	req := &ScalewayRequest{Method: http.MethodPost, Path: "/"}
	testhelpers.AssertNoError(t, req.SetBody(map[string]string{"hello": "world"}))

	err = client.Do(req, nil)
	testhelpers.AssertNoError(t, err)
	testhelpers.Equals(t, int32(2), atomic.LoadInt32(&attempts))
	testhelpers.Equals(t, `{"hello":"world"}`, lastBody)
}

func TestRetryPolicyHonorsContextCancellation(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer srv.Close()

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // already cancelled

	client, err := NewClient(
		WithAPIURL(srv.URL),
		WithoutAuth(),
		WithRetryPolicy(&RetryPolicy{
			MaxAttempts: 5,
			MinDelay:    time.Second,
			MaxDelay:    2 * time.Second,
		}),
	)
	testhelpers.AssertNoError(t, err)

	req := &ScalewayRequest{Method: http.MethodGet, Path: "/"}
	req.ctx = ctx
	err = client.Do(req, nil)
	testhelpers.Assert(t, err != nil, "expected an error due to cancelled context")
}

func TestRetryBudgetExhaustion(t *testing.T) {
	var attempts int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&attempts, 1)
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer srv.Close()

	// Budget of 1 retry token: the first request consumes 1 retry, the
	// second request should not be allowed to retry.
	client, err := NewClient(
		WithAPIURL(srv.URL),
		WithoutAuth(),
		WithRetryPolicy(&RetryPolicy{
			MaxAttempts:     5,
			MinDelay:        time.Millisecond,
			MaxDelay:        10 * time.Millisecond,
			RetryBudget:     1,
			RetryRefillRate: 0, // no refill during the test
		}),
	)
	testhelpers.AssertNoError(t, err)

	// First request: 1 initial attempt + 1 retry (budget allows 1).
	err = client.Do(&ScalewayRequest{Method: http.MethodGet, Path: "/"}, nil)
	testhelpers.Assert(t, err != nil, "expected an error")
	firstAttempts := atomic.LoadInt32(&attempts)
	testhelpers.Equals(t, int32(2), firstAttempts)

	// Second request: budget exhausted, only 1 attempt.
	err = client.Do(&ScalewayRequest{Method: http.MethodGet, Path: "/"}, nil)
	testhelpers.Assert(t, err != nil, "expected an error")
	secondAttempts := atomic.LoadInt32(&attempts) - firstAttempts
	testhelpers.Equals(t, int32(1), secondAttempts)
}

func TestParseRetryAfterSeconds(t *testing.T) {
	h := http.Header{}
	h.Set("Retry-After", "2")
	d, ok := parseRetryAfter(h)
	testhelpers.Assert(t, ok, "expected ok")
	testhelpers.Equals(t, 2*time.Second, d)
}

func TestParseRetryAfterAbsent(t *testing.T) {
	h := http.Header{}
	_, ok := parseRetryAfter(h)
	testhelpers.Assert(t, !ok, "expected not ok for absent header")
}

func TestRetryPolicyComputeBackoffRetryAfterTakesPrecedence(t *testing.T) {
	p := &RetryPolicy{MaxAttempts: 3, MinDelay: time.Millisecond, MaxDelay: 10 * time.Second}
	res := &http.Response{Header: http.Header{}}
	res.Header.Set("Retry-After", "5")
	d := p.computeBackoff(0, res)
	testhelpers.Equals(t, 5*time.Second, d)
}

func TestRetryPolicyComputeBackoffCapped(t *testing.T) {
	p := &RetryPolicy{MaxAttempts: 3, MinDelay: time.Second, MaxDelay: 5 * time.Second}
	res := &http.Response{Header: http.Header{}}
	d := p.computeBackoff(10, res)
	testhelpers.Assert(t, d <= 5*time.Second, "backoff should be capped at MaxDelay, got %s", d)
}

func TestBufferBodySeekable(t *testing.T) {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "http://example.com", bytes.NewReader([]byte("hello")))
	testhelpers.AssertNoError(t, bufferBody(req))

	// Consume the body.
	body, _ := io.ReadAll(req.Body)
	testhelpers.Equals(t, "hello", string(body))

	// Rewind for retry.
	seeker, ok := req.Body.(io.Seeker)
	testhelpers.Assert(t, ok, "body should be seekable after buffering")
	_, err := seeker.Seek(0, io.SeekStart)
	testhelpers.AssertNoError(t, err)

	body2, _ := io.ReadAll(req.Body)
	testhelpers.Equals(t, "hello", string(body2))
}

func TestIsRetryableHTTPStatus(t *testing.T) {
	testhelpers.Assert(t, isRetryableHTTPStatus(429), "429 should be retryable")
	testhelpers.Assert(t, isRetryableHTTPStatus(503), "503 should be retryable")
	testhelpers.Assert(t, !isRetryableHTTPStatus(400), "400 should not be retryable")
	testhelpers.Assert(t, !isRetryableHTTPStatus(404), "404 should not be retryable")
}

func TestIsRetryableNetworkError(t *testing.T) {
	testhelpers.Assert(t, isRetryableNetworkError(errors.New("connection reset")), "transient error should be retryable")
	testhelpers.Assert(t, !isRetryableNetworkError(context.Canceled), "context.Canceled should not be retryable")
	testhelpers.Assert(t, !isRetryableNetworkError(context.DeadlineExceeded), "context.DeadlineExceeded should not be retryable")
	testhelpers.Assert(t, !isRetryableNetworkError(nil), "nil should not be retryable")
}

func TestRetryPolicyClosesIntermediateResponseBody(t *testing.T) {
	var attempts int32
	var closedBodies int32

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt32(&attempts, 1)
		if n < 3 {
			// Write a body so there's something to drain.
			w.WriteHeader(http.StatusServiceUnavailable)
			_, _ = w.Write([]byte("transient error body"))
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	client, err := NewClient(
		WithAPIURL(srv.URL),
		WithoutAuth(),
		WithRetryPolicy(&RetryPolicy{
			MaxAttempts: 5,
			MinDelay:    time.Millisecond,
			MaxDelay:    10 * time.Millisecond,
		}),
	)
	testhelpers.AssertNoError(t, err)

	// Wrap the httpClient to track body closes.
	stdClient := client.httpClient.(*http.Client)
	original := stdClient.Transport
	stdClient.Transport = &bodyCloseTrackingTransport{rt: original, closed: &closedBodies}

	req := &ScalewayRequest{Method: http.MethodGet, Path: "/"}
	err = client.Do(req, nil)
	testhelpers.AssertNoError(t, err)
	testhelpers.Equals(t, int32(3), atomic.LoadInt32(&attempts))
	// All 3 response bodies should be closed: 2 intermediate (drained in the
	// retry loop) + 1 final (closed by Client.do's defer).
	testhelpers.Equals(t, int32(3), atomic.LoadInt32(&closedBodies))
}

type bodyCloseTrackingTransport struct {
	rt     http.RoundTripper
	closed *int32
}

func (t *bodyCloseTrackingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	res, err := t.rt.RoundTrip(req)
	if res != nil && res.Body != nil {
		originalBody := res.Body
		res.Body = &closeTrackingBody{ReadCloser: originalBody, closed: t.closed}
	}
	return res, err
}

type closeTrackingBody struct {
	io.ReadCloser
	closed *int32
}

func (b *closeTrackingBody) Close() error {
	atomic.AddInt32(b.closed, 1)
	return b.ReadCloser.Close()
}

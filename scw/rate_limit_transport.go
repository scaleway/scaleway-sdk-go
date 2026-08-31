package scw

import (
	"net/http"

	"github.com/scaleway/scaleway-sdk-go/logger"
	"golang.org/x/time/rate"
)

// rateLimitTransport is an http.RoundTripper that throttles outgoing requests
// using a token bucket limiter. It composes with other transports such as the
// requestLoggerTransport and is compatible with setInsecureMode as it does not
// replace the *http.Client type.
type rateLimitTransport struct {
	rt      http.RoundTripper
	limiter *rate.Limiter
}

// newRateLimitTransport wraps the given RoundTripper with a rate limiter.
// If rt is nil, http.DefaultTransport is used.
func newRateLimitTransport(rt http.RoundTripper, rps float64, burst int) *rateLimitTransport {
	if rt == nil {
		rt = http.DefaultTransport
	}
	return &rateLimitTransport{
		rt:      rt,
		limiter: rate.NewLimiter(rate.Limit(rps), burst),
	}
}

// RoundTrip waits until a token is available in the bucket before delegating
// the request to the underlying transport. The request context is respected:
// if it is cancelled while waiting, the limiter returns an error.
func (t *rateLimitTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if err := t.limiter.Wait(req.Context()); err != nil {
		return nil, err
	}
	return t.rt.RoundTrip(req)
}

// setRateLimit wraps the transport of the given *http.Client with a
// rateLimitTransport. If the client is not a standard *http.Client, a warning
// is logged and the call is a no-op, mirroring the behaviour of
// setInsecureMode and setRequestLogging.
func setRateLimit(c httpClient, rps float64, burst int) {
	standardHTTPClient, ok := c.(*http.Client)
	if !ok {
		logger.Warningf("client: cannot use rate limit with HTTP client of type %T", c)
		return
	}

	// Do not wrap twice. If the transport is already a rateLimitTransport,
	// just replace its limiter.
	if existing, ok := standardHTTPClient.Transport.(*rateLimitTransport); ok {
		existing.limiter.SetLimit(rate.Limit(rps))
		existing.limiter.SetBurst(burst)
		return
	}

	standardHTTPClient.Transport = newRateLimitTransport(standardHTTPClient.Transport, rps, burst)
}

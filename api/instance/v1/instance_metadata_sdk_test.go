package instance

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"
)

// stubRoundTripper is a minimal http.RoundTripper used to test the metadata
// client without performing real network calls.
type stubRoundTripper struct {
	// bodies maps request URL to the response body returned for that URL.
	bodies map[string]string
	// err, when non-nil, is returned for every request.
	err      error
	requests []*http.Request
}

func (s *stubRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	s.requests = append(s.requests, req)
	if s.err != nil {
		return nil, s.err
	}

	body, ok := s.bodies[req.URL.String()]
	if !ok {
		return &http.Response{
			StatusCode: http.StatusNotFound,
			Body:       http.NoBody,
			Header:     make(http.Header),
		}, nil
	}

	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(body)),
		Header:     make(http.Header),
	}, nil
}

func newStubClient(rt http.RoundTripper) *http.Client {
	return &http.Client{Transport: rt}
}

// TestGetMetadataWithContext_Success verifies that a successful metadata
// fetch returns the decoded payload.
func TestGetMetadataWithContext_Success(t *testing.T) {
	payload := Metadata{
		ID:             "server-id",
		Hostname:       "test-host",
		CommercialType: "DEV1-S",
	}

	raw, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("failed to marshal payload: %v", err)
	}

	rt := &stubRoundTripper{
		bodies: map[string]string{
			"http://169.254.42.42":                  "{}",
			"http://169.254.42.42/conf?format=json": string(raw),
			"http://[fd00:42::42]":                  "{}",
			"http://[fd00:42::42]/conf?format=json": string(raw),
		},
	}

	meta := NewMetadataAPI(WithMetadataHTTPClient(newStubClient(rt)))
	got, err := meta.GetMetadataWithContext(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got.ID != payload.ID || got.Hostname != payload.Hostname {
		t.Errorf("got %+v, want %+v", got, payload)
	}
}

// TestGetMetadataWithContext_TransportError verifies that a transport-level
// error during the metadata fetch is surfaced as an error and does not panic.
func TestGetMetadataWithContext_TransportError(t *testing.T) {
	rt := &stubRoundTripper{
		bodies: map[string]string{
			"http://169.254.42.42": "{}",
			"http://[fd00:42::42]": "{}",
		},
		err: errors.New("connection refused"),
	}

	meta := NewMetadataAPI(WithMetadataHTTPClient(newStubClient(rt)))
	_, err := meta.GetMetadataWithContext(context.Background())
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}

// TestGetMetadataURLWithContext_PanicSafety ensures that when the HTTP client
// returns an error (resp is nil), the URL probe loop does not panic.
func TestGetMetadataURLWithContext_PanicSafety(t *testing.T) {
	rt := &stubRoundTripper{
		err: errors.New("connection refused"),
	}

	meta := NewMetadataAPI(WithMetadataHTTPClient(newStubClient(rt)))

	// Should not panic and should fall back to the v4 address.
	url := meta.getMetadataURLWithContext(context.Background())
	if url != metadataAPIv4 {
		t.Errorf("expected fallback to %s, got %s", metadataAPIv4, url)
	}
}

// TestGetMetadataURLWithContext_SelectsReachable verifies that the probe
// selects the first address that returns HTTP 200.
func TestGetMetadataURLWithContext_SelectsReachable(t *testing.T) {
	rt := &stubRoundTripper{
		bodies: map[string]string{
			"http://169.254.42.42": "{}",
			"http://[fd00:42::42]": "{}",
		},
	}

	meta := NewMetadataAPI(WithMetadataHTTPClient(newStubClient(rt)))
	url := meta.getMetadataURLWithContext(context.Background())

	if url != metadataAPIv4 {
		t.Errorf("expected %s, got %s", metadataAPIv4, url)
	}

	if meta.MetadataURL == nil || *meta.MetadataURL != metadataAPIv4 {
		t.Errorf("expected MetadataURL to be cached as %s", metadataAPIv4)
	}
}

// TestGetMetadataURLWithContext_Cached verifies that a previously resolved
// MetadataURL is reused without further HTTP probes.
func TestGetMetadataURLWithContext_Cached(t *testing.T) {
	rt := &stubRoundTripper{
		bodies: map[string]string{},
	}

	cached := "http://example.invalid"
	meta := NewMetadataAPI(WithMetadataHTTPClient(newStubClient(rt)))
	meta.MetadataURL = &cached

	url := meta.getMetadataURLWithContext(context.Background())

	if url != cached {
		t.Errorf("expected cached %s, got %s", cached, url)
	}

	if len(rt.requests) != 0 {
		t.Errorf("expected no HTTP requests, got %d", len(rt.requests))
	}
}

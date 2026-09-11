package instance

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
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

// canBindPrivilegedPort skips the test if the process cannot bind to a
// privileged port (< 1024). The userdata API requires requests to originate
// from a privileged source port.
func canBindPrivilegedPort(t *testing.T) {
	t.Helper()

	var lc net.ListenConfig
	ln, err := lc.Listen(context.Background(), "tcp", ":1")
	if err != nil {
		t.Skipf("skipping: cannot bind to privileged port: %v", err)
	}

	if err := ln.Close(); err != nil {
		t.Fatalf("failed to close listener: %v", err)
	}
}

// TestNewUserDataHTTPClient verifies that the helper produces a working
// HTTP client when bound to an ephemeral port (port 0).
func TestNewUserDataHTTPClient(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	}))
	defer srv.Close()

	addr, err := net.ResolveTCPAddr("tcp", ":0")
	if err != nil {
		t.Fatalf("failed to resolve tcp address: %v", err)
	}

	client := newUserDataHTTPClient(addr)
	if client == nil || client.Transport == nil {
		t.Fatal("expected non-nil client and transport")
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, srv.URL, nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			t.Errorf("failed to close response body: %v", err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	if string(body) != "ok" {
		t.Errorf("expected 'ok', got %q", body)
	}
}

// TestGetUserDataWithContext_EmptyKey verifies that an empty key returns an
// error without making an HTTP request.
func TestGetUserDataWithContext_EmptyKey(t *testing.T) {
	meta := NewMetadataAPI()

	_, err := meta.GetUserDataWithContext(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty key, got nil")
	}
}

// TestSetUserDataWithContext_EmptyKey verifies that an empty key returns an
// error without making an HTTP request.
func TestSetUserDataWithContext_EmptyKey(t *testing.T) {
	meta := NewMetadataAPI()
	err := meta.SetUserDataWithContext(context.Background(), "", []byte("value"))
	if err == nil {
		t.Fatal("expected error for empty key, got nil")
	}
}

// TestDeleteUserDataWithContext_EmptyKey verifies that an empty key returns
// an error without making an HTTP request.
func TestDeleteUserDataWithContext_EmptyKey(t *testing.T) {
	meta := NewMetadataAPI()

	err := meta.DeleteUserDataWithContext(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty key, got nil")
	}
}

// TestListUserDataWithContext_Success verifies that ListUserDataWithContext
// decodes the userdata list returned by the metadata service.
func TestListUserDataWithContext_Success(t *testing.T) {
	canBindPrivilegedPort(t)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/user_data" {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(UserData{UserData: []string{"key1", "key2"}}); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	}))
	defer srv.Close()

	url := srv.URL
	meta := NewMetadataAPI()
	meta.MetadataURL = &url

	res, err := meta.ListUserDataWithContext(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(res.UserData) != 2 || res.UserData[0] != "key1" || res.UserData[1] != "key2" {
		t.Errorf("unexpected userdata: %+v", res)
	}
}

// TestGetUserDataWithContext_Success verifies that GetUserDataWithContext
// returns the raw bytes for a given key.
func TestGetUserDataWithContext_Success(t *testing.T) {
	canBindPrivilegedPort(t)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/user_data/mykey" {
			_, _ = w.Write([]byte("myvalue"))
			return
		}
		http.NotFound(w, r)
	}))
	defer srv.Close()

	url := srv.URL
	meta := NewMetadataAPI()
	meta.MetadataURL = &url

	got, err := meta.GetUserDataWithContext(context.Background(), "mykey")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if string(got) != "myvalue" {
		t.Errorf("expected 'myvalue', got %q", got)
	}
}

// TestSetUserDataWithContext_Success verifies that SetUserDataWithContext
// sends a PATCH request with the correct body.
func TestSetUserDataWithContext_Success(t *testing.T) {
	canBindPrivilegedPort(t)

	var receivedBody []byte
	var receivedMethod string

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/user_data/mykey" {
			receivedMethod = r.Method
			receivedBody, _ = io.ReadAll(r.Body)
			w.WriteHeader(http.StatusOK)
			return
		}
		http.NotFound(w, r)
	}))
	defer srv.Close()

	url := srv.URL
	meta := NewMetadataAPI()
	meta.MetadataURL = &url

	err := meta.SetUserDataWithContext(context.Background(), "mykey", []byte("myvalue"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if receivedMethod != http.MethodPatch {
		t.Errorf("expected method %s, got %s", http.MethodPatch, receivedMethod)
	}

	if string(receivedBody) != "myvalue" {
		t.Errorf("expected body 'myvalue', got %q", receivedBody)
	}
}

// TestDeleteUserDataWithContext_Success verifies that DeleteUserDataWithContext
// sends a DELETE request for the given key.
func TestDeleteUserDataWithContext_Success(t *testing.T) {
	canBindPrivilegedPort(t)

	var receivedMethod string

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/user_data/mykey" {
			receivedMethod = r.Method
			w.WriteHeader(http.StatusOK)
			return
		}
		http.NotFound(w, r)
	}))
	defer srv.Close()

	url := srv.URL
	meta := NewMetadataAPI()
	meta.MetadataURL = &url

	err := meta.DeleteUserDataWithContext(context.Background(), "mykey")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if receivedMethod != http.MethodDelete {
		t.Errorf("expected method %s, got %s", http.MethodDelete, receivedMethod)
	}
}

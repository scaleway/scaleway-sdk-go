package scw

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/scaleway/scaleway-sdk-go/internal/auth"
)

type ScalewayRequest struct {
	Method  string
	Path    string
	Headers http.Header
	Query   url.Values
	Body    io.Reader
	Ctx     context.Context
}

func (c *Client) Do(req *ScalewayRequest, opts ...RequestOption) (*http.Response, error) {
	if req == nil {
		return nil, fmt.Errorf("request must be non-nil")
	}

	// build url
	url, err := req.getURL(c.baseUrl)
	if err != nil {
		return nil, err
	}

	// build request
	request, err := http.NewRequest(req.Method, url.String(), req.Body)
	if err != nil {
		return nil, err
	}

	request.Header = req.getAllHeaders(c.auth, c.userAgent)

	// execute request
	resp, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	err = hasErrorResponse(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// getAllHeaders constructs a http.Header object and aggregates all headers into the object.
func (req *ScalewayRequest) getAllHeaders(token auth.Auth, userAgent string) http.Header {
	allHeaders := token.Headers()
	allHeaders.Set("User-Agent", userAgent)
	if req.Body != nil {
		allHeaders.Set("content-type", "application/json")
	}
	for key, value := range req.Headers {
		for _, v := range value {
			allHeaders.Set(key, v)
		}
	}
	return allHeaders
}

// getURL constructs a URL based on the base url and the client.
func (req *ScalewayRequest) getURL(baseURL string) (*url.URL, error) {
	url, err := url.Parse(baseURL + req.Path)
	if err != nil {
		return nil, fmt.Errorf("invalid url %s: %s", baseURL+req.Path, err)
	}
	url.RawQuery = req.Query.Encode()

	return url, nil
}

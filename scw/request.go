package scw

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"

	"github.com/scaleway/scaleway-sdk-go/internal/auth"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/utils"
)

// ScalewayRequest contains all the contents related to performing a request on the Scaleway API.
type ScalewayRequest struct {
	Method  string
	Path    string
	Headers http.Header
	Query   url.Values
	Body    io.Reader
	Ctx     context.Context

	Pagination RequestPagination
}

// Do performs an HTTP request based on the ScalewayRequest object.
// RequestOptions are executed on the ScalewayRequest.
func (c *Client) Do(req *ScalewayRequest, res interface{}, opts ...RequestOption) error {
	if req == nil {
		return fmt.Errorf("request must be non-nil")
	}

	// PSEUDO CODE
	if IsFullPaginationEnabled {
		c.doFullPagination(req, res, opts...)
	}

	// apply request options
	for _, opt := range opts {
		opt(req)
	}

	// build url
	url, err := req.getURL(c.apiURL)
	if err != nil {
		return err
	}
	logger.Debugf("creating %s request on %s", req.Method, url.String())

	// build request
	request, err := http.NewRequest(req.Method, url.String(), req.Body)
	if err != nil {
		return err
	}

	request.Header = req.getAllHeaders(c.auth, c.userAgent)

	if req.Ctx != nil {
		request = request.WithContext(req.Ctx)
	}

	if logger.ShouldLog(logger.LogLevelDebug) {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			logger.Warningf("cannot dump outgoing request: %s", err)
		} else {
			logger.Debugf("dumping http request:\n" + string(dump))
		}
	}

	// execute request
	httpRes, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}

	err = hasResponseError(httpRes)
	if err != nil {
		return err
	}

	defer httpRes.Body.Close()
	err = json.NewDecoder(httpRes.Body).Decode(&res)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) doFullPagination(req *ScalewayRequest, res interface{}, opts ...RequestOption) error {
	req.Pagination.Page = 1        // Init pagination
	err := c.Do(req, res, opts...) // Do first request
	if err != nil {
		return err
	}

	total := reflect.ValueOf(res).FieldByName("TotalCount").Uint() // We could all use tag to fiend correct field
	slice := reflect.ValueOf(res).FieldByName("????").Interface().([]interface{})

	// Loop all page
	for i := uint64(0); i < total; i++ {
		r := reflect.New(reflect.ValueOf(res).Type())
		req.Pagination.Page = 1      // Init pagination
		err := c.Do(req, r, opts...) // Do first request
		if err != nil {
			return err
		}
		slice = append(slice, reflect.ValueOf(res).FieldByName("????").Interface().([]interface{})...)
	}
	reflect.ValueOf(res).FieldByName("????").Set(reflect.ValueOf(slice))
	return nil
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

// SetBody json marshal the given body and write the json content type
// to the request. It also catches when body is a file.
func (req *ScalewayRequest) SetBody(body interface{}) error {
	var contentType string
	var content io.Reader

	switch b := body.(type) {
	case *utils.File:
		contentType = b.ContentType
		content = b.Content
	default:
		buf, err := json.Marshal(body)
		if err != nil {
			return err
		}
		contentType = "application/json"
		content = bytes.NewReader(buf)
	}

	req.Headers.Set("Content-Type", contentType)
	req.Body = content

	return nil
}

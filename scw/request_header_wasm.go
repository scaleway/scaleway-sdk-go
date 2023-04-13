//go:build wasm && js

package scw

// getAllHeaders constructs a http.Header object and aggregates all headers into the object.
func (req *ScalewayRequest) getAllHeaders(token auth.Auth, userAgent string, anonymized bool) http.Header {
	var allHeaders http.Header
	if anonymized {
		allHeaders = token.AnonymizedHeaders()
	} else {
		allHeaders = token.Headers()
	}

	allHeaders.Set("X-SCW-CLI", userAgent)
	if req.Body != nil {
		allHeaders.Set("Content-Type", "application/json")
	}
	for key, value := range req.Headers {
		allHeaders.Del(key)
		for _, v := range value {
			allHeaders.Add(key, v)
		}
	}

	return allHeaders
}

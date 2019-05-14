package marshaler

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/scaleway/scaleway-sdk-go/utils"
)

// Body is an HTTP Body with a content-type
type Body struct {
	io.Reader
	contentType string
}

// ContentType return the content type of the body
func (b *Body) ContentType() string {
	return b.contentType
}

// MarshalBody marshal a request to the adequate body
func MarshalBody(body interface{}) (*Body, error) {
	res := &Body{}

	switch b := body.(type) {
	case *utils.File:
		res.contentType = b.ContentType
		res.Reader = b.Content
	default:
		buf, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		res.contentType = "application/json"
		res.Reader = bytes.NewReader(buf)
	}

	return res, nil
}

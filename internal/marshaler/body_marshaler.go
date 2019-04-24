package marshaler

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/scaleway/scaleway-sdk-go/utils"
)

type Body struct {
	io.ReadCloser
	contentType string
}

func (b *Body) ContentType() string {
	return b.contentType
}

func MarshalBody(body interface{}) (*Body, error) {
	res := &Body{}

	switch b := body.(type) {
	case *utils.File:
		res.contentType = b.ContentType
		res.ReadCloser = ioutil.NopCloser(b.Content)
	default:
		buf, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		res.contentType = "application/json"
		res.ReadCloser = ioutil.NopCloser(bytes.NewBuffer(buf))
	}

	return res, nil
}

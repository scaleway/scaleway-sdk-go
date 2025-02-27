package audit_trail

import (
	"github.com/scaleway/scaleway-sdk-go/errors"
)

func (r *ListEventsResponse) UnsafeAppend(res interface{}) (*string, error) {
	results, ok := res.(*ListEventsResponse)
	if !ok {
		return nil, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Events = append(r.Events, results.Events...)

	return results.NextPageToken, nil
}

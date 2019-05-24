package scw

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

type lister interface {
	UnsafeGetTotalCount() int
	UnsafeAppend(interface{}) (int, error)
}

// doListAll collects all pages of a List request and aggregate all results on a single response.
func (c *Client) doListAll(req *ScalewayRequest, res interface{}) (err error) {

	// check for lister interface
	if response, isLister := res.(lister); isLister {

		pageCount := math.MaxUint32
		for page := 1; page <= pageCount; page++ {
			// set current page
			req.Query.Set("page", strconv.Itoa(page))

			// request the next page
			nextPage := newPage(response)
			err := c.Do(req, nextPage)
			if err != nil {
				return err
			}

			// append results
			pageSize, err := response.UnsafeAppend(nextPage)
			if err != nil {
				return err
			}

			// set total count on first request
			if pageCount == math.MaxUint32 {
				totalCount := nextPage.(lister).UnsafeGetTotalCount()
				pageCount = (totalCount + pageSize - 1) / pageSize
			}
		}
		return nil
	}

	return fmt.Errorf("%T does not support pagination", res)
}

// newPage returns a variable set to the zero value of the given type
func newPage(v interface{}) interface{} {
	// reflect.New always create a pointer, that's why we use reflect.Indirect before
	return reflect.New(reflect.Indirect(reflect.ValueOf(v)).Type()).Interface()
}

package scw

import (
	"fmt"
	"math"
	"net/url"
	"reflect"
	"strconv"
)

type lister interface {
	UnsafeGetTotalCount() int
	UnsafeAppend(interface{}) (int, error)
}

// doListAll collect all pages of a List request and aggregate all results on a single response.
func (c *Client) doListAll(req *ScalewayRequest, res interface{}) (err error) {

	// restore original query parameters at the end of the function
	originalQuery := url.Values{}
	for k, v := range req.Query {
		originalQuery[k] = v
	}
	defer func() {
		req.Query = originalQuery
	}()

	// check for lister interface
	switch response := res.(type) {
	case lister:
		n := 0
		page := 1 // start at page 1
		totalCount := math.MaxUint32
		for i := 0; i < totalCount; i += n {
			// set current page
			req.Query.Set("page", strconv.Itoa(page))
			page++

			// request the next page
			nextPage := newPage(response)
			err := c.Do(req, nextPage)
			if err != nil {
				return err
			}

			// append results
			n, err = response.UnsafeAppend(nextPage)
			if err != nil {
				return err
			}

			// set total count on first request
			if totalCount == math.MaxUint32 {
				totalCount = nextPage.(lister).UnsafeGetTotalCount()
			}
		}
		return nil
	}

	return fmt.Errorf("%T does not support pagination", res)
}

// newPage return a variable set to the zero value of the given type
func newPage(v interface{}) interface{} {
	// reflect.New always create a pointer, that's why we use reflect.Indirect before
	return reflect.New(reflect.Indirect(reflect.ValueOf(v)).Type()).Interface()
}

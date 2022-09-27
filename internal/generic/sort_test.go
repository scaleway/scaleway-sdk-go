package generic

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

func Test_SortSliceByField(t *testing.T) {
	type Elem struct {
		Field string
	}
	elems := []*Elem{
		{"2"},
		{"1"},
		{"3"},
	}
	SortSliceByField(elems, "Field", func(i interface{}, i2 interface{}) bool {
		return i.(string) < i2.(string)
	})
	testhelpers.Assert(t, elems[0].Field == "1", "slice is not sorted")
	testhelpers.Assert(t, elems[1].Field == "2", "slice is not sorted")
	testhelpers.Assert(t, elems[2].Field == "3", "slice is not sorted")
}

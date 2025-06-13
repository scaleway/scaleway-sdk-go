package scw_test

import (
	"testing"
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

var (
	testString                 = "some string"
	testBytes                  = []byte{0, 1, 2}
	testBool                   = true
	testInt32    int32         = 42
	testInt64    int64         = 43
	testUInt32   uint32        = 44
	testUInt64   uint64        = 45
	testFloat32  float32       = 46
	testFloat64  float64       = 47
	testDuration time.Duration = 48
	testTime                   = time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC)
	testSize                   = 3 * scw.GB
)

func TestStringPtr(t *testing.T) {
	pointer := scw.StringPtr(testString)
	slice := []string{testString}
	sliceOfPointers := scw.StringSlicePtr(slice)
	pointerToSlice := scw.StringsPtr(slice)

	// value to pointer value
	testhelpers.Assert(t, pointer != nil, "Pointer should have value")
	testhelpers.Equals(t, testString, *pointer)

	// slice of values to slice of pointers to value
	testhelpers.Equals(t, 1, len(sliceOfPointers))
	testhelpers.Assert(t, sliceOfPointers[0] != nil, "Pointer should have value")
	testhelpers.Equals(t, testString, *sliceOfPointers[0])

	// slice of value to pointer to slice of values
	testhelpers.Assert(t, pointerToSlice != nil, "Pointer should have value")
	testhelpers.Equals(t, slice, *pointerToSlice)
}

func TestBytesPtr(t *testing.T) {
	pointer := scw.BytesPtr(testBytes)
	slice := [][]byte{testBytes}
	sliceOfPointers := scw.BytesSlicePtr(slice)

	// value to pointer value
	testhelpers.Assert(t, pointer != nil, "Pointer should have value")
	testhelpers.Equals(t, testBytes, *pointer)

	// slice of values to slice of pointers to value
	testhelpers.Equals(t, 1, len(sliceOfPointers))
	testhelpers.Assert(t, sliceOfPointers[0] != nil, "Pointer should have value")
	testhelpers.Equals(t, testBytes, *sliceOfPointers[0])
}

func TestBoolPtr(t *testing.T) {
	pointer := scw.BoolPtr(testBool)
	slice := []bool{testBool}
	sliceOfPointers := scw.BoolSlicePtr(slice)

	// value to pointer value
	testhelpers.Assert(t, pointer != nil, "Pointer should have value")
	testhelpers.Equals(t, testBool, *pointer)

	// slice of values to slice of pointers to value
	testhelpers.Equals(t, 1, len(sliceOfPointers))
	testhelpers.Assert(t, sliceOfPointers[0] != nil, "Pointer should have value")
	testhelpers.Equals(t, testBool, *sliceOfPointers[0])
}

func TestInt32Ptr(t *testing.T) {
	pointer := scw.Int32Ptr(testInt32)
	slice := []int32{testInt32}
	sliceOfPointers := scw.Int32SlicePtr(slice)

	// value to pointer value
	testhelpers.Assert(t, pointer != nil, "Pointer should have value")
	testhelpers.Equals(t, testInt32, *pointer)

	// slice of values to slice of pointers to value
	testhelpers.Equals(t, 1, len(sliceOfPointers))
	testhelpers.Assert(t, sliceOfPointers[0] != nil, "Pointer should have value")
	testhelpers.Equals(t, testInt32, *sliceOfPointers[0])
}

func TestInt64Ptr(t *testing.T) {
	pointer := scw.Int64Ptr(testInt64)
	slice := []int64{testInt64}
	sliceOfPointers := scw.Int64SlicePtr(slice)

	// value to pointer value
	testhelpers.Assert(t, pointer != nil, "Pointer should have value")
	testhelpers.Equals(t, testInt64, *pointer)

	// slice of values to slice of pointers to value
	testhelpers.Equals(t, 1, len(sliceOfPointers))
	testhelpers.Assert(t, sliceOfPointers[0] != nil, "Pointer should have value")
	testhelpers.Equals(t, testInt64, *sliceOfPointers[0])
}

func TestUint32Ptr(t *testing.T) {
	pointer := scw.Uint32Ptr(testUInt32)
	slice := []uint32{testUInt32}
	sliceOfPointers := scw.Uint32SlicePtr(slice)

	// value to pointer value
	testhelpers.Assert(t, pointer != nil, "Pointer should have value")
	testhelpers.Equals(t, testUInt32, *pointer)

	// slice of values to slice of pointers to value
	testhelpers.Equals(t, 1, len(sliceOfPointers))
	testhelpers.Assert(t, sliceOfPointers[0] != nil, "Pointer should have value")
	testhelpers.Equals(t, testUInt32, *sliceOfPointers[0])
}

func TestUint64Ptr(t *testing.T) {
	pointer := scw.Uint64Ptr(testUInt64)
	slice := []uint64{testUInt64}
	sliceOfPointers := scw.Uint64SlicePtr(slice)

	// value to pointer value
	testhelpers.Assert(t, pointer != nil, "Pointer should have value")
	testhelpers.Equals(t, testUInt64, *pointer)

	// slice of values to slice of pointers to value
	testhelpers.Equals(t, 1, len(sliceOfPointers))
	testhelpers.Assert(t, sliceOfPointers[0] != nil, "Pointer should have value")
	testhelpers.Equals(t, testUInt64, *sliceOfPointers[0])
}

func TestFloat32Ptr(t *testing.T) {
	pointer := scw.Float32Ptr(testFloat32)
	slice := []float32{testFloat32}
	sliceOfPointers := scw.Float32SlicePtr(slice)

	// value to pointer value
	testhelpers.Assert(t, pointer != nil, "Pointer should have value")
	testhelpers.Equals(t, testFloat32, *pointer)

	// slice of values to slice of pointers to value
	testhelpers.Equals(t, 1, len(sliceOfPointers))
	testhelpers.Assert(t, sliceOfPointers[0] != nil, "Pointer should have value")
	testhelpers.Equals(t, testFloat32, *sliceOfPointers[0])
}

func TestFloat64Ptr(t *testing.T) {
	pointer := scw.Float64Ptr(testFloat64)
	slice := []float64{testFloat64}
	sliceOfPointers := scw.Float64SlicePtr(slice)

	// value to pointer value
	testhelpers.Assert(t, pointer != nil, "Pointer should have value")
	testhelpers.Equals(t, testFloat64, *pointer)

	// slice of values to slice of pointers to value
	testhelpers.Equals(t, 1, len(sliceOfPointers))
	testhelpers.Assert(t, sliceOfPointers[0] != nil, "Pointer should have value")
	testhelpers.Equals(t, testFloat64, *sliceOfPointers[0])
}

func TestDurationPtr(t *testing.T) {
	pointer := scw.TimeDurationPtr(testDuration)

	// value to pointer value
	testhelpers.Assert(t, pointer != nil, "Pointer should have value")
	testhelpers.Equals(t, testDuration, *pointer)
}

func TestTimePtr(t *testing.T) {
	pointer := scw.TimePtr(testTime)

	// value to pointer value
	testhelpers.Assert(t, pointer != nil, "Pointer should have value")
	testhelpers.Equals(t, testTime, *pointer)
}

func TestSizePtr(t *testing.T) {
	pointer := scw.SizePtr(testSize)

	// value to pointer value
	testhelpers.Assert(t, pointer != nil, "Pointer should have value")
	testhelpers.Equals(t, testSize, *pointer)
}

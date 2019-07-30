package scw

import (
	"fmt"

	"github.com/dustin/go-humanize"
)

type Size uint64

const (
	Byte Size = 1 << (iota * 10)
	KiByte
	MiByte
	GiByte
	TiByte
	PiByte
	EiByte
)

// String returns the string representation of a Size.
func (s Size) String() string {
	return fmt.Sprintf("%d", s)
}

// ParseSize parses a string representation of a Size and
// returns the corresponding value.
//
// ParseSize("42 MB") -> 42000000, nil
// ParseSize("42 mib") -> 44040192, nil
func ParseSize(str string) (Size, error) {
	bytes, err := humanize.ParseBytes(str)
	return Size(bytes), err
}

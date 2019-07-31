package scw

import (
	"fmt"
)

type Size uint64

const (
	B  Size = 1
	KB      = 1000 * B
	MB      = 1000 * KB
	GB      = 1000 * MB
	TB      = 1000 * GB
	PB      = 1000 * TB
)

// String returns the string representation of a Size.
func (s Size) String() string {
	return fmt.Sprintf("%d", s)
}

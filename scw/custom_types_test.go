package scw

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

func TestSize_String(t *testing.T) {
	cases := []struct {
		size Size
		want string
	}{
		{size: 42 * MB, want: "42000000"},
		{size: 42 * B, want: "42"},
	}

	for _, c := range cases {
		t.Run(c.want, func(t *testing.T) {
			testhelpers.Equals(t, c.want, c.size.String())
		})
	}
}

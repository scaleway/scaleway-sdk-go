package scw

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

func TestParseSize(t *testing.T) {
	cases := []struct {
		str  string
		want Size
	}{
		{str: "42", want: Size(42 * B)},
		{str: "42 KB", want: 42 * KB},
		{str: "42 MB", want: 42 * MB},
		{str: "42 GB", want: 42 * GB},
		{str: "42 TB", want: 42 * TB},
		{str: "42 PB", want: 42 * PB},
	}

	for _, c := range cases {
		t.Run(c.str, func(t *testing.T) {
			got, _ := ParseSize(c.str)
			testhelpers.Equals(t, c.want, got)
		})
	}
}

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

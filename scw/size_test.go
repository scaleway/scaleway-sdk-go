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
		{str: "42 MB", want: Size(42000000)},
		{str: "42 mib", want: Size(44040192)},
		{str: "42", want: Size(42)},
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
		{size: Size(42000000), want: "42000000"},
		{size: Size(42), want: "42"},
	}

	for _, c := range cases {
		t.Run(c.want, func(t *testing.T) {
			testhelpers.Equals(t, c.want, c.size.String())
		})
	}
}

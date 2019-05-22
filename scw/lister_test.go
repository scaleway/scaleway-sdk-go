package scw

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

func TestNewPage(t *testing.T) {
	type fakeType struct {
		plop int
	}

	testhelpers.Equals(t, &fakeType{}, newPage(&fakeType{3}))
}

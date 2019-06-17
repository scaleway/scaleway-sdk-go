// Source: github.com/docker/docker/pkg/namesgenerator

package namesgenerator

import (
	"strings"
	"testing"
)

func TestNameFormat(t *testing.T) {
	name := GetRandomName(0)
	if strings.Count(name, "-") != 2 {
		t.Fatalf("Generated name does not contain exactly 2 hyphen")
	}
	if strings.ContainsAny(name, "0123456789") {
		t.Fatalf("Generated name contains numbers!")
	}
}

func TestNameRetries(t *testing.T) {
	name := GetRandomName(1)
	if strings.Count(name, "-") != 2 {
		t.Fatalf("Generated name does not contain exactly 2 hyphen")
	}
	if !strings.ContainsAny(name, "0123456789") {
		t.Fatalf("Generated name doesn't contain a number")
	}

}

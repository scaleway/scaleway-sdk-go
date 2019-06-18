// Source: github.com/docker/docker/pkg/namesgenerator

package namegenerator

import (
	"strings"
	"testing"
)

func TestNameFormat(t *testing.T) {
	name := GetRandomName()
	if strings.Count(name, "-") != 1 {
		t.Fatalf("Generated name does not contain exactly 1 hyphen")
	}
	if strings.ContainsAny(name, "0123456789") {
		t.Fatalf("Generated name contains numbers!")
	}
}

func TestNameFormatWithPrefix(t *testing.T) {
	name := GetRandomName("scw")
	if strings.Count(name, "-") != 2 {
		t.Fatalf("Generated name does not contain exactly 2 hyphens")
	}
	if !strings.HasPrefix(name, "scw-") {
		t.Fatalf("Generated name must begin with \"tf-scw-\"")
	}
	if strings.ContainsAny(name, "0123456789") {
		t.Fatalf("Generated name contains numbers!")
	}
}

func TestNameFormatWithPrefixes(t *testing.T) {
	name := GetRandomName("tf", "scw")
	if strings.Count(name, "-") != 3 {
		t.Fatalf("Generated name does not contain exactly 3 hyphens")
	}
	if !strings.HasPrefix(name, "tf-scw-") {
		t.Fatalf("Generated name must begin with \"tf-scw-\"")
	}
	if strings.ContainsAny(name, "0123456789") {
		t.Fatalf("Generated name contains numbers!")
	}
}

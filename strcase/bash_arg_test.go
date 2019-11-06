package strcase

import (
	"testing"
)

func TestToKebabCase(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"testCase", "test-case"},
		{"TestCase", "test-case"},
		{"Test Case", "test-case"},
		{" Test Case", "test-case"},
		{"Test Case ", "test-case"},
		{" Test Case ", "test-case"},
		{"test", "test"},
		{"test_case", "test-case"},
		{"Test", "test"},
		{"", ""},
		{"ManyManyWords", "many-many-words"},
		{"manyManyWords", "many-many-words"},
		{"AnyKind of_string", "any-kind-of-string"},
		{"numbers2and55with000", "numbers2and55with000"},
		{"JSONData", "json-data"},
		{"userID", "user-id"},
		{"AAAbbb", "aa-abbb"},
	}
	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			got := toKebab(c.in)
			if got != c.want {
				t.Errorf("toKebab(%q) == %q, want %q", c.in, got, c.want)
			}
		})
	}
}

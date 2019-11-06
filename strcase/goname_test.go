package strcase

import "testing"

func TestLowerCaseFirstLetterOrAcronyms(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"t", "t"},
		{"Test Case", "test Case"},
		{"test Case", "test Case"},
		{"TEST CASE", "tEST CASE"},
		{"tEST CASE", "tEST CASE"},
		{"#EST CASE", "#EST CASE"},
		{"APITest", "apiTest"},
		{"AVATATest", "aVATATest"},
		{"TestStuff", "testStuff"},
	}
	for _, c := range cases {
		result := lowerCaseFirstLetterOrAcronyms(c.in)
		if result != c.want {
			t.Errorf("lowerCaseFirstLetterOrAcronyms(%q) == %q, want %q", c.in, result, c.want)
		}
	}
}

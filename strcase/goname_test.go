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

func TestTitleFirstWord(t *testing.T) {
	cases := [][]string{
		{"", ""},
		{"t", "T"},
		{"Test Case", "Test Case"},
		{"test Case", "Test Case"},
		{"TEST CASE", "TEST CASE"},
		{"tEST CASE", "TEST CASE"},
		{"#EST CASE", "#EST CASE"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := TitleFirstWord(in)
		if result != out {
			t.Error("'" + result + "' != '" + out + "'")
		}
	}
}

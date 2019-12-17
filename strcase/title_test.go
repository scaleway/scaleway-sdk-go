package strcase

import "testing"

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

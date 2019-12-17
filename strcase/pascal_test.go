package strcase

import "testing"

func TestToPascal(t *testing.T) {
	cases := [][]string{
		{"test_case", "TestCase"},
		{"test", "Test"},
		{"TestCase", "TestCase"},
		{" test  case ", "TestCase"},
		{"", ""},
		{"many_many_words", "ManyManyWords"},
		{"AnyKind of_string", "AnyKindOfString"},
		{"odd-fix", "OddFix"},
		{"numbers2And55with000", "Numbers2And55With000"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := ToPascal(in)
		if result != out {
			t.Error("'" + result + "' != '" + out + "'")
		}
	}
}

package strcase

import (
	"testing"
)

func TestToSnake(t *testing.T) {
	cases := [][2]string{
		{"testCase", "test_case"},
		{"TestCase", "test_case"},
		{"Test Case", "test_case"},
		{" Test Case", "test_case"},
		{"Test Case ", "test_case"},
		{" Test Case ", "test_case"},
		{"test", "test"},
		{"test_case", "test_case"},
		{"Test", "test"},
		{"", ""},
		{"ManyManyWords", "many_many_words"},
		{"manyManyWords", "many_many_words"},
		{"AnyKind of_string", "any_kind_of_string"},
		{"numbers2and55with000", "numbers2and55with000"},
		{"ip-v6", "ip_v6"},
		{"ipV6", "ip_v6"},
		{"IPV6", "ipv6"},
		{"ipv6", "ipv6"},
		{"JSONData", "json_data"},
		{"userID", "user_id"},
		{"AAAbbb", "aa_abbb"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := ToSnake(in)
		if result != out {
			t.Error("'" + result + "' != '" + out + "'")
		}
	}
}

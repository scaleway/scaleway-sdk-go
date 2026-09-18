package strcase

import "strings"

// Converts a string to CamelCase
func toPascalInitCase(s string, initCase bool) string {
	s = addWordBoundariesToNumbers(s)
	s = strings.Trim(s, " ")
	var n strings.Builder
	capNext := initCase
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			n.WriteString(string(v))
		}
		if v >= '0' && v <= '9' {
			n.WriteString(string(v))
		}
		if v >= 'a' && v <= 'z' {
			if capNext {
				n.WriteString(strings.ToUpper(string(v)))
			} else {
				n.WriteString(string(v))
			}
		}
		if v == '_' || v == ' ' || v == '-' {
			capNext = true
		} else {
			capNext = false
		}
	}
	return n.String()
}

func ToPascal(s string) string {
	return toPascalInitCase(s, true)
}

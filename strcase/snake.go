package strcase

import (
	"strings"
)

// Converts a string to snake_case
func ToSnake(s string) string {
	s = strings.Trim(s, " ")
	n := ""
	for i, v := range s {
		// treat acronyms as words, eg for JSONData -> JSON is a whole word
		nextCaseIsChanged := false
		if i+1 < len(s) {
			next := s[i+1]
			if (isUpperLetter(v) && isLowerLetter(int32(next))) || (isLowerLetter(v) && isUpperLetter(int32(next))) {
				nextCaseIsChanged = true
			}
		}

		if i > 0 && n[len(n)-1] != '_' && nextCaseIsChanged {
			// add underscore if next letter case type is changed
			if isUpperLetter(v) {
				n += "_" + string(v)
			} else if isLowerLetter(v) {
				n += string(v) + "_"
			}
		} else if v == ' ' || v == '-' {
			// replace spaces and dashes with underscores
			n += "_"
		} else {
			n = n + string(v)
		}
	}
	n = strings.ToLower(n)
	return n
}

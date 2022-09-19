package strcase

import "strings"

var customBashNames = map[string]string{
	"aclid":  "acl-id",
	"ipid":   "ip-id",
	"lbid":   "lb-id",
	"dhcpid": "dhcp-id",
}

// ToBashArg returns the Bash public name of the given string.
func ToBashArg(s string) string {
	s = ToPublicGoName(s)
	if customBashName, exists := customBashNames[strings.ToLower(s)]; exists {
		return customBashName
	}
	for _, initialism := range customInitialisms {
		// catch this kind of pattern: ExampleIDs ==> ExampleIds ==> example-ids
		s = strings.Replace(s, initialism[0], strings.Title(strings.ToLower(initialism[0])), -1)
	}
	return toKebab(s)
}

// toKebab converts a string to kebab-case.
func toKebab(s string) string {
	return toDelimited(s, '-')
}

// toDelimited converts a string to delimited lowercase.
func toDelimited(s string, del uint8) string {
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

		if i > 0 && n[len(n)-1] != del && nextCaseIsChanged {
			// add delimiter if next letter case type is changed
			if isUpperLetter(v) {
				n += string(del) + string(v)
			} else if isLowerLetter(v) {
				n += string(v) + string(del)
			}
		} else if v == ' ' || v == '-' || v == '_' {
			// replace spaces and dashes with delimiter
			n += string(del)
		} else {
			n = n + string(v)
		}
	}
	n = strings.ToLower(n)
	return n
}

func isUpperLetter(c int32) bool {
	return c >= 'A' && c <= 'Z'
}

func isLowerLetter(c int32) bool {
	return c >= 'a' && c <= 'z'
}

package strcase

import (
	"strings"
	"unicode"
)

// ToPrivateGoName returns the Go public name of the given string.
func ToPublicGoName(s string) string {
	return toGoName(TitleFirstWord(s))
}

// ToPrivateGoName returns the Go private name of the given string.
func ToPrivateGoName(s string) string {
	return toGoName(lowerCaseFirstLetterOrAcronyms(s))
}

// toGoName returns a different name if it should be different.
func toGoName(name string) (should string) {
	name = strings.Replace(name, " ", "_", -1)
	name = strings.Replace(name, "-", "_", -1)

	// Fast path for simple cases: "_" and all lowercase.
	if name == "_" {
		return name
	}
	allLower := true
	for _, r := range name {
		if !unicode.IsLower(r) {
			allLower = false
			break
		}
	}
	if allLower {
		return name
	}

	// Split camelCase at any lower->upper transition, and split on underscores.
	// Check each word for common initialisms.
	runes := []rune(name)
	w, i := 0, 0 // index of start of word, scan
	for i+1 <= len(runes) {
		eow := false // whether we hit the end of a word
		if i+1 == len(runes) {
			eow = true
		} else if runes[i+1] == '_' {
			// underscore; shift the remainder forward over any run of underscores
			eow = true
			n := 1
			for i+n+1 < len(runes) && runes[i+n+1] == '_' {
				n++
			}

			// Leave at most one underscore if the underscore is between two digits
			if i+n+1 < len(runes) && unicode.IsDigit(runes[i]) && unicode.IsDigit(runes[i+n+1]) {
				n--
			}

			copy(runes[i+1:], runes[i+n+1:])
			runes = runes[:len(runes)-n]
		} else if unicode.IsLower(runes[i]) && !unicode.IsLower(runes[i+1]) {
			// lower->non-lower
			eow = true
		}
		i++
		if !eow {
			continue
		}

		// [w,i) is a word.
		word := string(runes[w:i])
		u := strings.ToUpper(word)
		if commonInitialisms[u] {
			// Keep consistent case, which is lowercase only at the start.
			if w == 0 && unicode.IsLower(runes[w]) {
				u = strings.ToLower(u)
			}
			// All the common initialisms are ASCII,
			// so we can replace the bytes exactly.
			copy(runes[w:], []rune(u))
		} else if specialCase, exist := customInitialisms[u]; exist {
			if w == 0 && unicode.IsLower(runes[w]) {
				u = specialCase[1]
			} else {
				u = specialCase[0]
			}

			copy(runes[w:], []rune(u))
		} else if w > 0 && strings.ToLower(word) == word {
			// already all lowercase, and not the first word, so uppercase the first character.
			runes[w] = unicode.ToUpper(runes[w])
		}
		w = i
	}
	return string(runes)
}

// commonInitialisms is a set of common initialisms.
// Only add entries that are highly unlikely to be non-initialisms.
// For instance, "ID" is fine (Freudian code is rare), but "AND" is not.
var commonInitialisms = map[string]bool{
	"ACL":   true,
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DHCP":  true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LB":    true,
	"LHS":   true,
	"MNQ":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSD":   true,
	"SSH":   true,
	"SNS":   true,
	"SQS":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XMPP":  true,
	"XSRF":  true,
	"XSS":   true,
}

// customInitialisms is a set of common initialisms we use at Scaleway.
// value[0] is the uppercase replacement
// value[1] is the lowercase replacement
var customInitialisms = map[string][2]string{
	"ACLS":  {"ACLs", "acls"},
	"APIS":  {"APIs", "apis"},
	"CPUS":  {"CPUs", "cpus"},
	"IDS":   {"IDs", "ids"},
	"IPS":   {"IPs", "ips"},
	"IPV":   {"IPv", "ipv"}, // handle IPV4 && IPV6
	"LBS":   {"LBs", "lbs"},
	"UIDS":  {"UIDs", "uids"},
	"UUIDS": {"UUIDs", "uuids"},
	"URIS":  {"URIs", "uris"},
	"URLS":  {"URLs", "urls"},
}

// TitleFirstWord upper case the first letter of a string.
func TitleFirstWord(s string) string {
	if s == "" {
		return s
	}

	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])

	return string(r)
}

// UntitleFirstWord lower case the first letter of a string.
func UntitleFirstWord(s string) string {
	if s == "" {
		return s
	}

	r := []rune(s)

	firstWord := strings.Split(s, " ")[0]
	_, isCommonInitialism := commonInitialisms[firstWord]
	_, isCustomInitialism := customInitialisms[firstWord]
	if !isCommonInitialism && !isCustomInitialism {
		r[0] = unicode.ToLower(r[0])
	}

	return string(r)
}

// lowerCaseFirstLetterOrAcronyms lower case the first letter of a string.
func lowerCaseFirstLetterOrAcronyms(s string) string {
	r := []rune(s)
	if len(r) == 0 {
		return ""
	}

	for i := 0; len(r) > i && unicode.IsUpper(r[i]); i++ {
		word := string(r[:i+1])
		if u := strings.ToUpper(word); commonInitialisms[u] {
			copy(r[0:], []rune(strings.ToLower(u)))
			break
		}
	}
	r[0] = unicode.ToLower(r[0])

	return string(r)
}

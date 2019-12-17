package strcase

import "strings"

// Converts a string to kebab-case
func ToKebab(s string) string {
	return strings.Replace(ToSnake(s), "_", "-", -1)
}

// Converts a string to kebab-case
func ToSpace(s string) string {
	return strings.Replace(ToSnake(s), "_", " ", -1)
}

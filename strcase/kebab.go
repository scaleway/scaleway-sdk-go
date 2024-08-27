package strcase

import "strings"

// Converts a string to kebab-case
func ToKebab(s string) string {
	return strings.ReplaceAll(ToSnake(s), "_", "-")
}

// Converts a string to kebab-case
func ToSpace(s string) string {
	return strings.ReplaceAll(ToSnake(s), "_", " ")
}

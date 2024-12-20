package testhelpers

import "strings"

// IsTestResource returns true if given resource identifier is from terraform test
// identifier should be resource name but some resource don't have names
// return true if identifier match regex "tf[-_]test"
// common used prefixes are "tf_tests", "tf_test", "tf-tests", "tf-test"
func IsTestResource(identifier string) bool {
	return len(identifier) >= len("tf_test") &&
		strings.HasPrefix(identifier, "tf") &&
		(identifier[2] == '_' || identifier[2] == '-') &&
		identifier[3:7] == "test"
}

package exercices

import (
	"strings"
)

// Count counts the number of non-overlapping instances of substr in s
func Count(str string, substr string) int {
	return strings.Count(str, substr)
}

// ToUpper returns s with all Unicode letters mapped to their upper case.
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

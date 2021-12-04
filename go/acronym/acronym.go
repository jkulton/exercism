// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) (result string) {
	re := regexp.MustCompile(`[^A-Za-z\- ]`)
	s = strings.ReplaceAll(s, "-", " ")
	s = re.ReplaceAllLiteralString(s, "")

	words := strings.Split(s, " ")

	for _, word := range words {
		if len(word) > 0 {
			char := string(word[0])
			result += strings.ToUpper(char)
		}
	}

	return result
}

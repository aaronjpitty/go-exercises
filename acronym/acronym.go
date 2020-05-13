package acronym

import (
	"strings"
)

// Abbreviate function to get the acronym of words
func Abbreviate(s string) string {

	// replace all hyphens with space
	spaces := strings.ReplaceAll(s, "-", " ")
	// replace all underscors with nothing, (thus removing underscores)
	unders := strings.ReplaceAll(spaces, "_", "")
	// split a string into substrings, removing any space characters, including newlines
	words := strings.Fields(unders)

	res := ""

	for _, word := range words {
		// slice each word in words to give us just the individual letters
		res += word[0:1]
	}

	return strings.ToUpper(res)
}

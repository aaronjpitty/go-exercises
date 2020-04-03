package bob

import (
	"regexp"
	"strings"
)

// Run outside of the function to speed up regex compilation
var containsLetters = regexp.MustCompile(`\pL`)
var containsLowerCase = regexp.MustCompile(`\p{Ll}`)

// Hey function which responds correctly
func Hey(remark string) string {
	str := strings.TrimSpace(remark)
	letters := containsLetters.MatchString(str)
	lowers := containsLowerCase.MatchString(str)
	yelled := letters && !lowers
	question := strings.HasSuffix(str, "?")

	switch {
	case yelled && question:
		return "Calm down, I know what I'm doing!"
	case yelled:
		return "Whoa, chill out!"
	case question:
		return "Sure."
	case str == "":
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}

}

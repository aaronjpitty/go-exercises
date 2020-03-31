package bob

import (
	"regexp"
	"strings"
)

// Hey function which responds correctly
func Hey(remark string) string {
	remark = strings.Trim(remark, " \t\n\r")

	re, _ := regexp.Compile("[a-zA-Z]+")

	// Check if the string was all capitals
	yelled := strings.ToUpper(remark) == remark && (len(re.FindAllString(remark, -1)) > 0)
	
	// Check if the string had a question mark
	question := strings.HasSuffix(remark, "?")

	if yelled {
		if question {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	}

	if question {
		return "Sure."
	}

	if remark == "" {
		return "Fine. Be that way!"
	}

	return "Whatever."

}
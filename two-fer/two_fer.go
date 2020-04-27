package twofer

import (
	"fmt"
)

// ShareWith function. If string not empty, use one phrase, else, fall back to my name.
func ShareWith(name string) string {

	if len(name) > 0 {
		return fmt.Sprintf("One for %s, one for me.", name)
	} else if name == "" {
		name = "aaron pitty"
	}

	return "One for you, one for me."

}

package twofer

import (
	"fmt"
)

func ShareWith(name string) string {

	if len(name) > 0 {
		return fmt.Sprintf("One for %v, one for me.", name)
	} else {
		return "One for you, one for me."
	}

}
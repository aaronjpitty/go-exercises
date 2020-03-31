package acronym

import (
	"strings"
)

func Abbreviate(s string) string {

	words := strings.Split(s, " ")

	res := ""
	acr := ""

	for _, word := range words {
		res = res + string([]rune(word)[0])
	}

	acr = acr + strings.ToUpper(res)

	return acr
}

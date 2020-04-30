package scrabble

import "strings"

// GetScore function using case to set value for each letter
func GetScore(r rune) int {
	switch r {
	case 'A', 'E', 'I', 'L', 'N', 'O', 'R', 'S', 'T', 'U':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'M', 'C', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	default:
		return 0
	}
}

// Score takes a word and returns the value of that word
// according to scrabble scoring rules
func Score(w string) int {
	score := 0
	for _, r := range strings.ToUpper(w) {
		score += GetScore(r)
	}
	return score
}

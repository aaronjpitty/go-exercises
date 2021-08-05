package reverse

// Uses runes to reverse a string
//func Reverse(s string) string {
//	n := len(s)
//	runes := make([]rune, n)
//	for _, rune := range s {
//		n--
//		runes[n] = rune
//	}
//	return string(runes[n:])
//}


//func Reverse(s string) string {
//	size := len(s)
//	buf := make([]byte, size)
//	for start := 0; start < size; {
//		r, n := utf8.DecodeRuneInString(s[start:])
//		start += n
//		utf8.EncodeRune(buf[size-start:], r)
//	}
//	return string(buf)
//}

// Reverse: This one seems to be the most efficient method
func Reverse(s string) (result string) {
	for _,v := range s {
		result = string(v) + result
	}
	return
}
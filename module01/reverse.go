package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	// solution 4 for non english(non ASCII) chars handling. Called "rune"
	var ret string
	for _, r := range word {
		ret = string(r) + ret
	}
	return ret
	// solution 3
	// var ret string
	// for i := 0; i < len(word); i++ {
	// 	ret = string(word[i]) + ret
	// }
	// return ret

	// solution 2
	// var str strings.Builder
	// for i := len(word) - 1; i >= 0; i-- {
	// 	str.WriteByte(word[i])
	// }
	// return str.String()

	// solution 1
	// var ret string
	// for i := len(word) - 1; i >= 0; i-- {
	// 	ret = ret + string(word[i])
	// }
	// return ret
}

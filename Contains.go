package hangman_classic

func Contains(s string, e string) bool {
	for _, a := range s {
		if isLower(a) {
			a = putUpper(a)
		}
		if isLower(rune(e[0])) {
			e = string(putUpper(rune(e[0]))) + e[1:]
		}
		if string(a) == e {
			return true
		}
	}
	return false
}

package hangman_classic

func Contains(s string, e string) bool {
	for _, a := range s {
		if string(a) == e {
			return true
		}
	}
	return false
}

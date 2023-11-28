package core

func Contains(s string, e string) (bool, int) {
	for i, a := range s {
		if isLower(a) {
			a = putUpper(a)
		}
		if isLower(rune(e[0])) {
			e = string(putUpper(rune(e[0]))) + e[1:]
		}
		if string(a) == e {
			return true, i
		}
	}
	return false, -1
}

func ContainsArray(s []string, e string) (bool, int) {
	for i, a := range s {
		if a == e {
			return true, i
		}
	}
	return false, -1
}

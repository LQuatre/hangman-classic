package hangman_classic

import "os"

func EmptyFile(path string) {
	var RemplaceFile *os.File
	RemplaceFile, _ = os.Create(path)
	RemplaceFile.Close()
}

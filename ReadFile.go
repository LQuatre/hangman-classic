package hangman_classic

import (
	"os"
	"strings"
)

func ReadFile(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		//Do something
	}
	lines := strings.Split(string(content), "\n")
	return lines
}

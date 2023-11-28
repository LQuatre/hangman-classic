package core

import (
	"os"
	"strings"
)

func ReadFile(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	return lines
}

func Readable(filename string) bool {
	content, err := os.ReadFile(filename)
	if err != nil {
		return false
	}
	return len(content) > 0
}

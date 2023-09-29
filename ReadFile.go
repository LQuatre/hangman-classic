package hangman_classic

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	// Get the file size
	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	str := string(bs)
	return strings.Split(str, "\n")
}

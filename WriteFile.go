package hangman_classic

import (
	"fmt"
	"os"
)

func WriteFile(filename string, data []string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for i := 0; i < len(data); i++ {
		if i == len(data)-1 {
			file.WriteString(data[i])
			break
		}
		file.WriteString(data[i] + "\n")
	}
}

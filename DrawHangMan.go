package hangman_classic

import "fmt"

// files : "assets/Hangman Positions/hangman.txt"

func DrawHangMang(data HangManData, state int) {
	HangMan := ReadFile("assets/Hangman Positions/hangman.txt")
	for index, character := range string {
		if index >= state*8 && index < (state+1)*8 {
			fmt.Println(character)
		}
	}
}

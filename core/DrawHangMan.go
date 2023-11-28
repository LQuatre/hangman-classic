package core

import "fmt"

// files : "assets/Hangman Positions/hangman.txt"

func DrawHangMang(state int) {
	HangMan := ReadFile("assets/Hangman Positions/hangman.txt")
	returned := ""
	for index, HangMan := range HangMan {
		if index >= state*8 && index < (state+1)*8 {
			returned += HangMan + "\n"
		}
	}
	fmt.Println(returned)
}

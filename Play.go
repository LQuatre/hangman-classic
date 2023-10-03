package hangman_classic

import (
	"fmt"
)

type HangManData struct {
	Word             string     // Word composed of '_', ex: H_ll_
	ToFind           string     // Final word chosen by the program at the beginning. It is the word to find
	WordLength       int        // Length of the word
	Attempts         int        // Number of attempts left
	HangmanPositions [10]string // It can be the array where the positions parsed in "hangman.txt" are stored
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func isLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func putUpper(s rune) rune {
	if s == '_' {
		return s
	}
	if isLower(s) {
		return s - 32
	}
	return s
}

func putLower(s rune) rune {
	if s == '_' {
		return s
	}
	if isUpper(s) {
		return s + 32
	}
	return s
}

func Play(data HangManData) {
	// X Le programme lit l'entrée standard pour suggérer une lettre.
	fmt.Println("Good Luck, you have 10 attempts.")
	for _, v := range data.Word {
		fmt.Print(string(putUpper(v)) + " ")
	}
	fmt.Print("\n\n")

	fmt.Print("Choose: ")

	var read string
	fmt.Scanln(&read)

	for data.Attempts < 10 {

		// X Si la lettre n'est pas présente, il affichera un message d'erreur et le nombre de tentatives diminuera (10->9->...0).
		if !Contains(data.ToFind, read) {
			data.Attempts++
			fmt.Println("Not present in the word, ", 10-data.Attempts, " attempts remaining.")
			DrawHangMang(data.Attempts)
		} else {
			// X Si la lettre est présente, il révélera toutes les lettres correspondantes dans le mot.
			for i := 0; i < data.WordLength; i++ {
				guestUint8 := read[0]
				guestRune := rune(guestUint8)
				if putUpper(rune(data.ToFind[i])) == putUpper(guestRune) {
					data.Word = data.Word[:i] + string(guestRune) + data.Word[i+1:]
				}
			}
			for _, v := range data.Word {
				fmt.Print(string(putUpper(v)) + " ")
			}
			fmt.Print("\n\n")
		}

		if data.Attempts == 10 {
			fmt.Println("You lose !")
			fmt.Println("The word was: ", data.ToFind)
			break
		}

		// X Le programme continue jusqu'à ce que le mot soit trouvé, ou que le nombre de tentatives soit égal à 0.
		if Contains(data.Word, "_") {
			fmt.Print("Choose: ")
			fmt.Scanln(&read)
		} else {
			fmt.Println("Congrats !")
			break
		}
	}

}

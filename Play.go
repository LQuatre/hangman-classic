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

func Play(data HangManData) {
	// X Le programme lit l'entrée standard pour suggérer une lettre.
	fmt.Println("Welcome to Hangman !")
	fmt.Println("You have 10 attempts to guess the word.")
	fmt.Println("The word has", data.WordLength, "letters.")
	fmt.Println("The word is : ", data.Word)
	fmt.Println("Please enter a letter : ")

	var read string
	fmt.Scanln(&read)

	for data.Attempts > 0 {

		// X Si la lettre n'est pas présente, il affichera un message d'erreur et le nombre de tentatives diminuera (10->9->...0).
		if !Contains(data.ToFind, read) {
			data.Attempts--

			fmt.Println("The letter is not in the word. You have", data.Attempts, "attempts left.")
			fmt.Println("The word is : ", data.Word)
		} else {
			// X Si la lettre est présente, il révélera toutes les lettres correspondantes dans le mot.
			fmt.Println("The letter is in the word.")
			for i := 0; i < data.WordLength; i++ {
				if data.ToFind[i] == read[0] {
					data.Word = data.Word[:i] + read + data.Word[i+1:]
				}
			}
			fmt.Println("The word is : ", data.Word)
		}

		if data.Attempts == 0 {
			fmt.Println("You lose !")
			break
		}

		// X Le programme continue jusqu'à ce que le mot soit trouvé, ou que le nombre de tentatives soit égal à 0.
		if Contains(data.Word, "_") {
			fmt.Println("Please enter a letter : ")
			fmt.Scanln(&read)
		} else {
			fmt.Println("You win !")
			break
		}
	}

}

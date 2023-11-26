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
	Ranking          map[string]int
	PlayerName       string
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

func letterByteTo1to26number(letter byte) (int, bool) {
	if letter >= 'A' && letter <= 'Z' {
		return int(letter - 'A' + 1), true
	} else if letter >= 'a' && letter <= 'z' {
		return int(letter - 'a' + 1), false
	}
	return 0, false
}

var myWord []string

func Play(data HangManData) {
	asciiLetters := ReadFile("assets/Ascii-letters/standard.txt")
	// remove \n from every line
	for i := 0; i < len(asciiLetters); i++ {
		asciiLetters[i] = asciiLetters[i][:len(asciiLetters[i])-1]
	}

	myWord = []string{"", "", "", "", "", "", "", "", ""}

	// X Le programme affichera le mot à trouver s en ascii art.
	for i := 0; i < data.WordLength; i++ {
		if data.Word[i] == '_' {
			myLetter := asciiLetters[569:577]
			//	remove \n from every line
			for i := 0; i < 8; i++ {
				myWord[i] += myLetter[i]
			}
		} else {
			letterNumber, isUpper := letterByteTo1to26number(data.Word[i])
			lineStart := 0
			var myLetter []string
			if isUpper {
				lineStart = 298 + (letterNumber-1)*9
			} else {
				lineStart = 586 + (letterNumber-1)*9
			}
			myLetter = asciiLetters[lineStart : lineStart+9]
			//	remove \n from every line
			for i := 0; i < 9; i++ {
				myWord[i] += myLetter[i]
			}
		}
	}

	for i := 0; i < 9; i++ {
		fmt.Println(myWord[i])
	}

	// X Le programme lit l'entrée standard pour suggérer une lettre.
	fmt.Println("Good Luck, you have ", 10-data.Attempts, " attempts remaining.")

	var read string
	for data.Attempts < 10 {
		// X Le programme continue jusqu'à ce que le mot soit trouvé, ou que le nombre de tentatives soit égal à 0.
		if Contains(data.Word, "_") {
			fmt.Print("Choose: ")
			fmt.Scanln(&read)
		} else {
			data.Word = data.ToFind
			fmt.Println("Congrats !")
			saveRank(data)
			break
		}

		if len(read) > 1 {
			if (read == "STOP" || read == "stop" || read == "Stop" || read == "QUIT" || read == "quit" || read == "Quit") {
				fmt.Println("You stopped the game.")
				saveGame(data)
				break
			}

			if read == data.ToFind {
				data.Word = data.ToFind
				fmt.Println("Congrats !")
				saveRank(data)
				break
			} else {
				DrawHangMang(data.Attempts)
				data.Attempts += 2
				fmt.Println("Wrong word, ", 10-data.Attempts, " attempts remaining.")
			}
		} else {
			// X Si la lettre n'est pas présente, il affichera un message d'erreur et le nombre de tentatives diminuera (10->9->...0).
			if !Contains(data.ToFind, read) {
				DrawHangMang(data.Attempts)
				data.Attempts++
				fmt.Println("Not present in the word, ", 10-data.Attempts, " attempts remaining.")
			} else if read == "" {
				fmt.Println("Please enter a valid letter.")
			} else {
				// X Si la lettre est présente, il révélera toutes les lettres correspondantes dans le mot.
				for i := 0; i < data.WordLength; i++ {
					guestUint8 := read[0]
					guestRune := rune(guestUint8)
					if putUpper(rune(data.ToFind[i])) == putUpper(guestRune) {
						data.Word = data.Word[:i] + string(guestRune) + data.Word[i+1:]
					}
				}
			}
		}

		myWord = []string{"", "", "", "", "", "", "", "", ""}

		// X Le programme affichera le mot à trouver s en ascii art.
		for i := 0; i < data.WordLength; i++ {
			if data.Word[i] == '_' {
				underscore := asciiLetters[569:577]
				//	remove \n from every line
				for i := 0; i < 8; i++ {
					myWord[i] += underscore[i]
				}
			} else {
				letterNumber, isUpper := letterByteTo1to26number(data.Word[i])
				lineStart := 0
				var myLetter []string
				if isUpper {
					lineStart = 298 + (letterNumber-1)*9
				} else {
					lineStart = 586 + (letterNumber-1)*9
				}
				myLetter = asciiLetters[lineStart : lineStart+9]
				//	remove \n from every line
				for i := 0; i < 9; i++ {
					myWord[i] += myLetter[i]
				}
			}
		}

		for i := 0; i < 9; i++ {
			fmt.Println(myWord[i])
		}

		if data.Attempts == 10 {
			fmt.Println("You lose !")
			fmt.Println("The word was: ", data.ToFind)
			break
		}
	}
}

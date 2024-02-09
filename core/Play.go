package core

import (
	"fmt"
	"strconv"
)

type HangManData struct {
	Word             string     // Word composed of '_', ex: H_ll_
	ToFind           string     // Final word chosen by the program at the beginning. It is the word to find
	WordLength       int        // Length of the word
	Attempts         int        // Number of attempts left
	HangmanPositions [10]string // It can be the array where the positions parsed in "hangman.txt" are stored
	Ranking          map[string]int
	PlayerName       string
	IsAscii          bool
	Ascii            []string
}

func ClearTerminal() {
	fmt.Print("\033[H\033[2J")
}

func MoveCursor(x int, y int) {
	fmt.Printf("\033[%d;%dH", x, y)
}

func Print(str string, data HangManData) {
	ClearTerminal()
	for i := 0; i < len(str); i++ {
		for j := 0; j < 9; j++ {
			fmt.Println(data.Ascii[(int(rune(str[i]))-32)*9+j])
			MoveCursor(j+1, i*11+1)
		}
	}
	MoveCursor(10, 1)
}

func isLower(char rune) bool {
	return char >= 'a' && char <= 'z'
}

func putUpper(char rune) rune {
	return char - ('a' - 'A')
}

func Play(data HangManData) {
	// Afficher le mot initial
	displayWord(data)

	var info string
	info += "Welcome to Hangman !\n"
	info += "Good Luck, you have " + strconv.Itoa(10-data.Attempts) + " attempts remaining.\n"
	print(info)

	var read string
	for data.Attempts < 10 {
		// Vérifier si le jeu est terminé
		if gameEnded(data) {
			break
		}

		// Lire l'entrée du joueur
		readInput(&read)

		// Traiter l'entrée du joueur
		processInput(&data, &info, read)

		// Afficher le mot mis à jour et le pendu
		displayWord(data)
		DrawHangMang(data.Attempts)

		// Afficher les informations
		print(info)
		info = ""
	}
	print("Game over ! The word was " + data.ToFind + ".\n")
}

func isLetterPresent(data HangManData, letter string) bool {
	result, _ := Contains(data.ToFind, letter)
	return result
}

func updateWord(data HangManData, letter string) {
	for i := 0; i < data.WordLength; i++ {
		if string(data.ToFind[i]) == letter {
			data.Word = data.Word[:i] + letter + data.Word[i+1:]
		}
	}
}

// Fonction pour afficher le mot avec les lettres découvertes
func displayWord(data HangManData) {
	if data.IsAscii == true {
		Print(data.Word, data)
	} else {
		fmt.Println(data.Word)
	}
}

// Fonction pour vérifier si le jeu est terminé
func gameEnded(data HangManData) bool {
	if present, _ := Contains(data.Word, "_"); !present {
		data.Word = data.ToFind
		fmt.Println("Congrats !")
		saveRank(data)
		return true
	}
	return false
}

// Fonction pour lire l'entrée du joueur
func readInput(read *string) {
	fmt.Print("Choose: ")
	fmt.Scanln(read)
	ClearTerminal()
}

// Fonction pour traiter l'entrée du joueur
func processInput(data *HangManData, info *string, read string) {
	if len(read) > 1 {
		// Traiter l'entrée si elle est un mot
		processWordInput(data, info, read)
	} else {
		// Traiter l'entrée si elle est une lettre
		processLetterInput(data, info, read)
	}
}

// Fonction pour traiter l'entrée si elle est un mot
func processWordInput(data *HangManData, info *string, read string) {
	if read == "STOP" || read == "stop" || read == "Stop" || read == "QUIT" || read == "quit" || read == "Quit" {
		*info += "Game stopped.\n"
		saveGame(*data)
		return
	}

	if read == data.ToFind {
		data.Word = data.ToFind
		*info += "Congrats !\n"
		saveRank(*data)
	} else {
		data.Attempts += 2
		*info += "Wrong word, " + strconv.Itoa(10-data.Attempts) + " attempts remaining.\n"
	}
}

// Fonction pour traiter l'entrée si elle est une lettre
func processLetterInput(data *HangManData, info *string, read string) {
	if read == "" {
		*info += "Please enter a valid letter.\n"
		return
	}

	if !isLetterPresent(*data, read) {
		data.Attempts++
		*info += "Wrong letter, " + strconv.Itoa(10-data.Attempts) + " attempts remaining.\n"
	} else {
		updateWord(*data, read)
		*info += "Present letter, " + strconv.Itoa(10-data.Attempts) + " attempts remaining.\n"
		for i := 0; i < data.WordLength; i++ {
			if string(data.ToFind[i]) == read {
				data.Word = data.Word[:i] + read + data.Word[i+1:]
			}
		}
	}
}

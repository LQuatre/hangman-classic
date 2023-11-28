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
	var info string
	info += "Welcome to Hangman !\n"

	if data.IsAscii == true {
		Print(data.Word, data)
	} else {
		fmt.Println(data.Word)
	}

	info = "Good Luck, you have " + strconv.Itoa(10-data.Attempts) + " attempts remaining.\n"

	DrawHangMang(data.Attempts)
	print(info)
	info = ""

	var read string
	for data.Attempts < 10 {
		if present, _ := Contains(data.Word, "_"); present {
			fmt.Print("Choose: ")
			fmt.Scanln(&read)
			ClearTerminal()
		} else {
			data.Word = data.ToFind
			fmt.Println("Congrats !")
			saveRank(data)
			break
		}

		if len(read) > 1 {
			if read == "STOP" || read == "stop" || read == "Stop" || read == "QUIT" || read == "quit" || read == "Quit" {
				info += "Game stopped.\n"
				saveGame(data)
				break
			}

			if read == data.ToFind {
				data.Word = data.ToFind
				info += "Congrats !\n"
				saveRank(data)
				break
			} else {
				data.Attempts += 2
				info += "Wrong word, " + strconv.Itoa(10-data.Attempts) + " attempts remaining.\n"
			}
		} else {
			if present, _ := Contains(data.ToFind, read); !present {
				data.Attempts++
				info += "Wrong letter, " + strconv.Itoa(10-data.Attempts) + " attempts remaining.\n"
			} else if read == "" {
				info += "Please enter a valid letter.\n"
			} else {
				for i := 0; i < data.WordLength; i++ {
					if string(data.ToFind[i]) == read {
						data.Word = data.Word[:i] + read + data.Word[i+1:]
					}
				}
				info += "Present letter, " + strconv.Itoa(10-data.Attempts) + " attempts remaining.\n"
			}
		}

		if data.IsAscii == true {
			Print(data.Word, data)
		} else {
			fmt.Println(data.Word)
		}
		DrawHangMang(data.Attempts)
		print(info)
		info = ""

		if data.Attempts == 10 {
			fmt.Println("You lose !")
			fmt.Println("The word was: ", data.ToFind)
			break
		}
	}
}

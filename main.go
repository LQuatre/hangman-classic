package main

import (
	"fmt"
	"os"
	"strings"
)

func stringToInt(s string) int {
	result := 0
	for _, a := range s {
		result = result*10 + int(a-'0')
	}
	return result
}

func main() {
	hangman_classic.ClearTerminal()
	MyHangManData := hangman_classic.HangManData{}
	args := os.Args[1:]
	rankingFile := hangman_classic.ReadFile("assets/Save/ranking.txt")
	var ranking [][3]string // PlayerName : Attempts left : Word
	for _, line := range rankingFile {
		if line != "" {
			ranking = append(ranking, [3]string{strings.Split(line, ":")[0], strings.Split(line, ":")[1], strings.Split(line, ":")[2]})
		}
	}
	fmt.Println("Do you want to see the ranking ? (y/n)")
	var answer string
	fmt.Scanln(&answer)
	for answer != "y" && answer != "n" {
		fmt.Println("Please enter a valid answer.")
		fmt.Scanln(&answer)
	}
	if answer == "y" {
		fmt.Println("Voici le classement :")
		fmt.Println("Name : Attempts left : Word")
		var sortedRanking [][3]string // PlayerName : Attempts left : Word
		for i := 0; i < len(ranking); i++ {
			if len(sortedRanking) == 0 {
				sortedRanking = append(sortedRanking, ranking[i])
			} else {
				for j := 0; j < len(sortedRanking); j++ {
					if stringToInt(ranking[i][1]) >
						stringToInt(sortedRanking[j][1]) {
						sortedRanking = append(sortedRanking[:j], append([][3]string{ranking[i]}, sortedRanking[j:]...)...)
						break
					} else if j == len(sortedRanking)-1 {
						sortedRanking = append(sortedRanking, ranking[i])
						break
					}
				}
			}
		}
		for i := 0; i < len(sortedRanking); i++ {
			if i < 10 {
				if i == 0 {
					fmt.Println("🥇 " + sortedRanking[i][0] + " with " + sortedRanking[i][1] + " attempts left and the word " + sortedRanking[i][2])
				} else if i == 1 {
					fmt.Println("🥈 " + sortedRanking[i][0] + " with " + sortedRanking[i][1] + " attempts left and the word " + sortedRanking[i][2])
				} else if i == 2 {
					fmt.Println("🥉 " + sortedRanking[i][0] + " with " + sortedRanking[i][1] + " attempts left and the word " + sortedRanking[i][2])
				} else {
					fmt.Println("🏅 " + sortedRanking[i][0] + " with " + sortedRanking[i][1] + " attempts left and the word " + sortedRanking[i][2])
				}
			}
		}
	}
	fmt.Println("Do you want to continue ? (y/n)")
	fmt.Scanln(&answer)
	for answer != "y" && answer != "n" {
		fmt.Println("Please enter a valid answer.")
		fmt.Scanln(&answer)
	}
	if answer != "y" {
		return
	}

	if len(args) == 0 {
		fmt.Println("Enter arguments")
		return
	} else if len(args) == 1 {
		if hangman_classic.Readable("assets/Dictionary/" + args[0]) {
			hangman_classic.NewGame(args, MyHangManData)
		} else {
			fmt.Println("1 - Please provide a valid argument.")
		}
	} else if present, pos := hangman_classic.ContainsArray(args, "--startWith"); present {
		if hangman_classic.Readable("assets/Save/" + args[pos+1]) {
			saved := hangman_classic.ReadFile("assets/Save/" + args[2])
			if len(saved) == 0 || saved[0] == "" {
				hangman_classic.NewGame(args, MyHangManData)
				return
			}
			MyHangManData.ToFind = saved[0]
			MyHangManData.Word = saved[1]
			MyHangManData.PlayerName = saved[2]
			MyHangManData.WordLength = len(MyHangManData.Word)
			MyHangManData.Attempts = stringToInt(saved[3])
			hangman_classic.Play(hangman_classic.HangManData(MyHangManData))
			hangman_classic.EmptyFile("assets/Save/" + args[2])
		} else {
			fmt.Println("2 - Please provide a valid argument.")
		}
	} else if present, pos := hangman_classic.ContainsArray(args, "--letterFile"); present {
		if hangman_classic.Readable("assets/Ascii-letters/" + args[pos+1]) {
			MyHangManData.IsAscii = true
			MyHangManData.Ascii = hangman_classic.ReadFile("assets/Ascii-letters/" + args[pos+1])
			hangman_classic.NewGame(args, MyHangManData)
		} else {
			fmt.Println("3 - Please provide a valid argument.")
		}
	}
}

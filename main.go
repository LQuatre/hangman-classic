package main

import (
	"fmt"
	"github.com/LQuatre/hangman-classic"
	"os"
	"strings"
)

func main() {
	core.ClearTerminal()
	MyHangManData := core.HangManData{}
	ranking := loadRanking()

	fmt.Println("Do you want to see the ranking ? (y/n)")
	answer := getValidAnswer()

	if answer == "y" {
		displayRanking(ranking)
	}

	fmt.Println("Do you want to continue ? (y/n)")
	answer = getValidAnswer()
	if answer != "y" {
		return
	}

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Enter arguments")
		return
	}

	if len(args) == 1 {
		if core.Readable("assets/Dictionary/" + args[0]) {
			core.NewGame(args, MyHangManData)
		} else {
			fmt.Println("Please provide a valid argument.")
		}
		return
	}

	if present, pos := core.ContainsArray(args, "--startWith"); present {
		handleStartWith(args, pos, MyHangManData)
		return
	}

	if present, pos := core.ContainsArray(args, "--letterFile"); present {
		handleLetterFile(args, pos, MyHangManData)
		return
	}
}

func loadRanking() [][3]string {
	rankingFile := core.ReadFile("assets/Save/ranking.txt")
	var ranking [][3]string // PlayerName : Attempts left : Word
	for _, line := range rankingFile {
		if line != "" {
			ranking = append(ranking, [3]string{strings.Split(line, ":")[0], strings.Split(line, ":")[1], strings.Split(line, ":")[2]})
		}
	}
	return ranking
}

func getValidAnswer() string {
	var answer string
	fmt.Scanln(&answer)
	for answer != "y" && answer != "n" {
		fmt.Println("Please enter a valid answer.")
		fmt.Scanln(&answer)
	}
	return answer
}

func displayRanking(ranking [][3]string) {
	fmt.Println("Voici le classement :")
	fmt.Println("Name : Attempts left : Word")
	sortedRanking := sortRanking(ranking)
	for i, player := range sortedRanking {
		if i < 10 {
			medal := ""
			switch i {
			case 0:
				medal = "🥇 "
			case 1:
				medal = "🥈 "
			case 2:
				medal = "🥉 "
			default:
				medal = "🏅 "
			}
			fmt.Printf("%s%s with %s attempts left and the word %s\n", medal, player[0], player[1], player[2])
		}
	}
}

func sortRanking(ranking [][3]string) [][3]string {
	var sortedRanking [][3]string // PlayerName : Attempts left : Word
	for _, player := range ranking {
		if len(sortedRanking) == 0 {
			sortedRanking = append(sortedRanking, player)
		} else {
			for j, sortedPlayer := range sortedRanking {
				if stringToInt(player[1]) > stringToInt(sortedPlayer[1]) {
					sortedRanking = append(sortedRanking[:j], append([][3]string{player}, sortedRanking[j:]...)...)
					break
				} else if j == len(sortedRanking)-1 {
					sortedRanking = append(sortedRanking, player)
					break
				}
			}
		}
	}
	return sortedRanking
}

func stringToInt(s string) int {
	result := 0
	for _, a := range s {
		result = result*10 + int(a-'0')
	}
	return result
}

func handleStartWith(args []string, pos int, MyHangManData core.HangManData) {
	if core.Readable("assets/Save/" + args[pos+1]) {
		saved := core.ReadFile("assets/Save/" + args[2])
		if len(saved) == 0 || saved[0] == "" {
			core.NewGame(args, MyHangManData)
			return
		}
		MyHangManData.ToFind = saved[0]
		MyHangManData.Word = saved[1]
		MyHangManData.PlayerName = saved[2]
		MyHangManData.WordLength = len(MyHangManData.Word)
		MyHangManData.Attempts = stringToInt(saved[3])
		core.Play(MyHangManData)
		core.EmptyFile("assets/Save/" + args[2])
	} else {
		fmt.Println("Please provide a valid argument.")
	}
}

func handleLetterFile(args []string, pos int, MyHangManData core.HangManData) {
	if core.Readable("assets/Ascii-letters/" + args[pos+1]) {
		MyHangManData.IsAscii = true
		MyHangManData.Ascii = core.ReadFile("assets/Ascii-letters/" + args[pos+1])
		core.NewGame(args, MyHangManData)
	} else {
		fmt.Println("Please provide a valid argument.")
	}
}

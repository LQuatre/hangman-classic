package hangman_classic

import (
	"fmt"
	"strconv"
)

func saveRank(data HangManData) {
	saveFile := ReadFile("assets/Save/ranking.txt")
	var attempts string
	attempts = strconv.Itoa(10 - data.Attempts)
	var rankLine string
	fmt.Println("Your word was " + data.ToFind)
	fmt.Println("You had " + attempts + " attempts left")
	fmt.Println("Do you want to save your score ? (y/n)")
	var answer string
	fmt.Scanln(&answer)
	for answer != "y" && answer != "n" {
		fmt.Println("Please enter a valid answer.")
		fmt.Scanln(&answer)
	}
	if answer != "y" {
		return
	}
	rankLine = data.PlayerName + ":" + attempts + ":" + data.Word
	saveFile = append(saveFile, rankLine)
	WriteFile("assets/Save/ranking.txt", saveFile)
	fmt.Println("Your score has been saved !")
}

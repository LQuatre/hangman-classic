package main

import (
	"fmt"
	"hangman_classic"
	"math/rand"
	"os"
	"strings"
)

func main() {
	MyHangManData := hangman_classic.HangManData{}
	// X Vous aurez 10 tentatives pour terminer le jeu.

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Please provide a dictionary file.")
		return
	}

	// O Tout d'abord, le programme choisira au hasard un mot dans le fichier.
	openFile := hangman_classic.ReadFile("assets/Dictionnary/" + args[0])
	lenFile := len(openFile)
	random := rand.Intn(lenFile)
	MyHangManData.ToFind = openFile[random]

	// X Le programme révèlera n lettres aléatoires dans le mot, où n est le len(word) / 2 - 1.
	lettersList := strings.Split(MyHangManData.ToFind, "")
	MyHangManData.WordLength = len(lettersList)
	numberOfLettersToReveal := MyHangManData.WordLength/2 - 1
	LettersToReveal := lettersList[:numberOfLettersToReveal]

	MyHangManData.Word = ""
	for i := 0; i < MyHangManData.WordLength; i++ {
		MyHangManData.Word += "_"
	}
	for i := 0; i < numberOfLettersToReveal; i++ {
		MyHangManData.Word = MyHangManData.Word[:i] + LettersToReveal[i] + MyHangManData.Word[i+1:]
	}

	hangman_classic.Play(hangman_classic.HangManData(MyHangManData))

}

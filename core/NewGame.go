package core

import (
	"math/rand"
	"strings"
)

func NewGame(args []string, MyHangManData HangManData) {
	openFile := ReadFile("assets/Dictionary/" + args[0])
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
	Play(HangManData(MyHangManData))
}

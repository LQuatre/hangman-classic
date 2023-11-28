package core

func IntToString(i int) string {
	var result string
	for i > 0 {
		result = string(i%10+48) + result
		i /= 10
	}
	return result
}

func saveGame(data HangManData) {
	var saveFile []string
	saveFile = append(saveFile, data.ToFind)
	saveFile = append(saveFile, data.Word)
	saveFile = append(saveFile, data.PlayerName)
	saveFile = append(saveFile, IntToString(data.Attempts))
	WriteFile("assets/Save/save.txt", saveFile)
}

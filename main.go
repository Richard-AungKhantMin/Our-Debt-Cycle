package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	//initalizing the payees
	initUsersPayees()

	//check history, read history
	latestFile, err := getLatestTxtFile("historyFiles")
	if latestFile != "" {
		isErrNil("Failed to find the latest history. ", err)
	}
	latestFilePath := filepath.Join("historyFiles", latestFile)
	readHistory(latestFilePath)

	//main stuffs
	Options()
	readHistory(os.Args[1])
	cleanUpAmounts()
	writeResults()

	if len(printingResult) == 0 {
		rabbit := readFile("rabbit.txt")
		fmt.Println(strings.Join(rabbit, "\n"))
		fmt.Println("\033[32mCongratulation!!! You guys are free of debts.\033[0m")
	} else {
		fmt.Println(strings.Join(printingResult, "\n"))
	}

}

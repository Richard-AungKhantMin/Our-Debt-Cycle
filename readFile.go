package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func logToFile(fileName, message string) {
	// Open file in append mode, create if not exists
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Write message to file
	_, err = file.WriteString(message + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func readFile(fileName string) []string {

	//Opening the file
	file, err := os.Open(fileName)
	isErrNil("Error reading the file", err)
	defer file.Close()

	//Reading each line
	tempFile := bufio.NewScanner(file)
	lines := make([]string, 0)
	for tempFile.Scan() {
		lines = append(lines, tempFile.Text())
	}
	err = tempFile.Err()
	isErrNil("Error reading each lines", err)

	return lines
}

func readLatestHistory() {

	//check history, read history
	latestFile, err := getLatestTxtFile("historyFiles")
	if latestFile != "" {
		isErrNil("Failed to find the latest history. ", err)
	}
	latestFilePath := filepath.Join("historyFiles", latestFile)
	readHistory(latestFilePath)

}

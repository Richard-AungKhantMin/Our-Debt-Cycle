package main

import (
	"fmt"
	"os"
)

func logToFile(message string) {
	// Open file in append mode, create if not exists
	file, err := os.OpenFile("history.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

func main() {

	logToFile(os.Args[1])

}

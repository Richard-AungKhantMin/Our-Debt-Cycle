package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	//initalizing the payees
	initUsersPayees()
	readLatestHistory()

	switch len(os.Args) {
	case 2:
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

	case 1:
		http.HandleFunc("/", HomeHandler)
		http.HandleFunc("/debt", debtsHandler)
		fmt.Println("The sever  has started on http://localhost:8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	default:
		log.Fatal("You could either use 'go run .' for web sever or 'go run . fileName.txt' for terminal interface.")
	}

}

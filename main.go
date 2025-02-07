package main

import (
	"fmt"
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
		/* 	http.HandleFunc("/", functions.HomeHandler) */
		log.Fatal("The sever  has started")
	default:
		log.Fatal("You could either use 'go run .' for web sever or 'go run . fileName.txt' for terminal interface.")
	}




	// Set up HTTP routes
	http.HandleFunc("/api/transaction", addTransactionHandler)
	http.HandleFunc("/api/balances", getBalancesHandler)
	http.HandleFunc("/api/cleanup", cleanupHandler)
	
	// Serve frontend files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}




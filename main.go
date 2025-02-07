package main

import (
	"fmt"
	"strings"
)

func main() {

	initUsersPayees()
	Options()
	readHistory()
	cleanUpAmounts()
	writeResults()

	fmt.Println(strings.Join(printingResult, "\n"))
}

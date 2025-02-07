package main

import (
	"errors"
	"log"
	"strings"
)

func initUsers() {
	for _, mainUser := range allUsers {
		for _, otherUser := range allUsers {

			if otherUser.Username != mainUser.Username {
				mainUser.ToPay = append(mainUser.ToPay, &Payee{Username: otherUser.Username, Amount: 0})
			}
		}
	}
}

func readHistory() {

	lines := readFile("history.txt")

	for _, eachLine := range lines {

		splitted := strings.Fields(eachLine)
		if len(splitted) != 2 {
			log.Fatal("The info is not in 'Payer-Payee Amount' format")
		}
		names := strings.Split(splitted[0], "-")
		if len(names) != 2 {
			log.Fatal("The info is not in 'Payer-Payee Amount' format")
		}

		payerName, payeeName := names[0], names[1]
		amount := splitted[1]

		Payers, err := getPayers(payerName, payeeName)
		isErrNil("Error: ", err)

	}

}

func addAmounts(Payers []User)

func getPayers(payerName, payeeName string) ([]User, error) {

	payers := make([]User, 0)
	var err error

	switch {
	case strings.ToLower(payerName) == "all":
		for _, eachUser := range allUsers {
			if notMe(eachUser.Username, payeeName) {
				payers = append(payers, *eachUser)
			}
		}

	case doTheyExist(payerName):
		OneUser := getUser(payerName)
		if OneUser == nil {
			err = errors.New("Couldn't fetch the user")
		}
		payers = append(payers)

	default:
		err = errors.New("Coundn't find the Payer's name")
	}

	return payers, err
}

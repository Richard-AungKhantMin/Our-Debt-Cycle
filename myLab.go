package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func initUsersPayees() {
	for _, mainUser := range allUsers {
		for _, otherUser := range allUsers {

			if otherUser.Username != mainUser.Username {
				mainUser.ToPay = append(mainUser.ToPay, &Payee{Username: otherUser.Username, Amount: 0})
			}
		}
	}
}

func readHistory(fileName string) {

	lines := readFile(fileName)

	for _, eachLine := range lines {

		if eachLine == "" {
			continue
		}

		splitted := strings.Fields(eachLine)
		if len(splitted) != 2 {
			log.Fatal("The info is not in 'Payer-Payee Amount' format")
		}
		names := strings.Split(splitted[0], "-")
		if len(names) != 2 {
			log.Fatal("The info is not in 'Payer-Payee Amount' format")
		}

		payerName, payeeName := names[0], names[1]

		amount, err := strconv.ParseFloat(splitted[1], 64)
		isErrNil("The amount is not a number ", err)
		amount = calculateAmounts(payerName, amount)

		Payers, err := getPayers(payerName, payeeName)
		isErrNil("Error: ", err)

		err = addAmounts(Payers, payeeName, amount)
		isErrNil("Error: ", err)
	}

}

func cleanUpAmounts() {

	for _, eachUser1 := range allUsers {
		for _, eachPayee1 := range eachUser1.ToPay {

			anotherUser, err := getUser(eachPayee1.Username)
			isErrNil("Error: ", err)

			for _, eachPayee2 := range anotherUser.ToPay {
				if eachPayee2.Username == eachUser1.Username {
					cancelOutAmounts(eachPayee1, eachPayee2)
				}
			}

		}
	}

}

func cancelOutAmounts(payee1, payee2 *Payee) {

	if payee1.Amount > payee2.Amount {
		payee1.Amount = payee1.Amount - payee2.Amount
		payee2.Amount = 0.0
	} else {
		payee2.Amount = payee2.Amount - payee1.Amount
		payee1.Amount = 0.0
	}

}

func addAmounts(Payers []*User, payeeName string, amount float64) error {

	var paid bool
	for _, eachPayer := range Payers {
		for _, eachPayee := range eachPayer.ToPay {
			if eachPayee.Username == payeeName {
				eachPayee.Amount = eachPayee.Amount + amount
				paid = true
			}
		}
	}

	if !paid {
		return errors.New("Couldn't find the payee!")
	}

	return nil
}

func calculateAmounts(payerName string, amount float64) float64 {

	if strings.ToLower(payerName) == "all" {
		amount = amount / float64(3)
	}

	return amount
}

func getPayers(payerName, payeeName string) ([]*User, error) {

	payers := make([]*User, 0)
	var err error

	switch {
	case strings.ToLower(payerName) == "all":
		for _, eachUser := range allUsers {
			if notMe(eachUser.Username, payeeName) {
				payers = append(payers, eachUser)
			}
		}

	case doTheyExist(payerName):
		OneUser, err := getUser(payerName)
		isErrNil("Error: ", err)
		payers = append(payers, OneUser)
	default:
		err = errors.New("Coundn't find the Payer's name")
	}

	return payers, err
}

func totalDebts(printingResult *[]string) {

	for _, eachUser := range allUsers {
		for _, eachPayee := range eachUser.ToPay {
			eachUser.TotalAmount = eachUser.TotalAmount + eachPayee.Amount
		}
		*printingResult = append(*printingResult, fmt.Sprintf("%s has total debt of %.2f€", eachUser.Username, eachUser.TotalAmount))
	}

}

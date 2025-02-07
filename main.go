package main

import "fmt"

func main() {

	initUsers()
	for _, user := range allUsers {
		fmt.Printf("%s's ToPay List:\n", user.Username)
		for _, payee := range user.ToPay {
			fmt.Printf("- %s: $%.2f\n", payee.Username, payee.Amount)
		}
	}

}

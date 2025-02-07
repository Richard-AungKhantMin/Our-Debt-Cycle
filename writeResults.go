package main

import "fmt"

func writeResults() {

	for _, eachUser := range allUsers {
		for _, eachPayee := range eachUser.ToPay {
			if eachPayee.Amount != 0 {
				logToFile(fmt.Sprintf("%s-%s %.2f", eachUser.Username, eachPayee.Username, eachPayee.Amount))
				printingResult = append(printingResult, fmt.Sprintf("%s has to pay %s %.2f€.", eachUser.Username, eachPayee.Username, eachPayee.Amount))
			}
		}

	}

}

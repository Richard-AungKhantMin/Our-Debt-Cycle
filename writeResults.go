package main

import (
	"fmt"
	"path/filepath"
	"time"
)

func writeResults() {

	currentTime := time.Now()

	formattedName := currentTime.Format("Jan-02-2006_15-04-05") + ".txt"
	resultFileName := filepath.Join("historyFiles", formattedName)

	for _, eachUser := range allUsers {
		for _, eachPayee := range eachUser.ToPay {
			if eachPayee.Amount != 0 {
				logToFile(resultFileName, fmt.Sprintf("%s-%s %.2f", eachUser.Username, eachPayee.Username, eachPayee.Amount))
				printingResult = append(printingResult, fmt.Sprintf("%s has to pay %s %.2f€.", eachUser.Username, eachPayee.Username, eachPayee.Amount))
			}
		}

	}

}

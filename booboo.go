package main

import (
	"fmt"
)

func notMe(name, me string) bool {
	if name == me {
		return false
	}
	return true
}

func doTheyExist(name string) bool {
	for _, thatUser := range allUsers {
		if thatUser.Username == name {
			return true
		}
	}
	return false
}

func isErrNil(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
	}
}

func getUser(name string) (*User, error) {
	for _, thatUser := range allUsers {
		if thatUser.Username == name {
			return thatUser, nil
		}
	}
	return nil, fmt.Errorf("Failed to fetch the user: %s", name)
}

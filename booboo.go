package main

import "log"

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
		log.Fatal(msg, err)
	}
}

func getUser(name string) *User {
	for _, thatUser := range allUsers {
		if thatUser.Username == name {
			return thatUser
		}
	}
	return nil
}

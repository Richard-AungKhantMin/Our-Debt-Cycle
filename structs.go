package main

type User struct {
	Username string
	ToPay    []Payee
}

type Payee struct {
	Username string
	Amount   float64
}

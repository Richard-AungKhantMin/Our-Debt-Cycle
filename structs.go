package main

type User struct {
	Username    string
	ToPay       []*Payee
	TotalAmount float64
}

type Payee struct {
	Username string
	Amount   float64
}

var (
	Richie = &User{Username: "Richie"}
	Milli  = &User{Username: "Milli"}
	Chan   = &User{Username: "Chan"}
)

var allUsers = []*User{
	Richie,
	Milli,
	Chan,
}

var (
	printingResult = make([]string, 0)
)

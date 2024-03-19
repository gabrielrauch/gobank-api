package main

import "math/rand"

type Account struct {
	ID        int     `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Agency    int64   `json:"agency"`
	Balance   float64 `json:"balance"`
}

func newAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        rand.Intn(1000),
		FirstName: firstName,
		LastName:  lastName,
		Agency:    rand.Int63n(1000),
		Balance:   rand.ExpFloat64(),
	}
}

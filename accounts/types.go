package accounts

import "github.com/bank/clients"

type CurrentAccount struct {
	Client  clients.Client
	Agency  int
	balance float64
}

type SavingAccount struct {
	Client  clients.Client
	Agency  int
	balance float64
}
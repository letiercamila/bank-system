package accounts

import (
	"fmt"

	"github.com/bank/clients"
)

func Saving(client clients.Client, agency int) *CurrentAccount {
	return &CurrentAccount{
		Client: client,
		Agency: agency,
	}
}

func (a *SavingAccount) Withdraw(amount float64) string {
	if amount > 0 && amount <= a.balance {
		a.balance -= amount
		return fmt.Sprintf("Withdrawal successful!New balance: %.2f", a.balance)
	} else {
		return "Something went wrong! Check your balance!"
	}
}

func (a *SavingAccount) Deposit(amount float64) string {
	if amount > 0 {
		a.balance += amount
		return fmt.Sprintf("Deposit successful! New balance: %.2f", a.balance)
	} else {
		return "Something went wrong! Try a valid amount!"
	}
}

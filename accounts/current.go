package accounts

import (
	"fmt"

	"github.com/bank/clients"
)

func Current(client clients.Client, agency int) *CurrentAccount {
	return &CurrentAccount{
		Client: client,
		Agency: agency,
	}
}

func (a *CurrentAccount) Withdraw(amount float64) string {
	if amount > 0 && amount <= a.balance {
		a.balance -= amount
		return fmt.Sprintf("Withdrawal successful!New balance: %.2f", a.balance)
	} else {
		return "Something went wrong! Check your balance!"
	}
}

func (a *CurrentAccount) Deposit(amount float64) string {
	if amount > 0 {
		a.balance += amount
		return fmt.Sprintf("Deposit successful! New balance: %.2f", a.balance)
	} else {
		return "Something went wrong! Try a valid amount!"
	}
}

func (a *CurrentAccount) Transfer(amount float64, destinyAccount *CurrentAccount) string {
	if amount > 0 && amount < a.balance {
		a.balance -= amount
		destinyAccount.Deposit(amount)
		return fmt.Sprintf("Transfer successful for: %s! New balance: %.2f", destinyAccount.Client.Name, destinyAccount.balance)
	} else {
		return "Something went wrong! Check your balance!"
	}
}

func (a *CurrentAccount) ShowBalance() float64 {
	return a.balance
}

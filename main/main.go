package main

import (
	"fmt"

	"github.com/bank/accounts"
	"github.com/bank/clients"
	"github.com/bank/payment"
)

func main() {
	camila := clients.CreateClient("Camila", 22, "123154542215", "Programmer")
	camilaAccount := accounts.Current(*camila, 01)
	camilaAccount.Deposit(600)

	maria := clients.CreateClient("Maria", 29, "01548962300", "Teacher")
	mariaAccount := accounts.Saving(*maria, 02)
	mariaAccount.Deposit(300)

	oldCamilaBalance := camilaAccount.ShowBalance()
	withdraw := camilaAccount.Withdraw(200)
	oldMariaBalance := mariaAccount.ShowBalance()
	transfer := camilaAccount.Transfer(100, mariaAccount)
	mariabalance := mariaAccount.ShowBalance()
	payment.PayBill(camilaAccount, 100)

	fmt.Println("Camila's balance before transfers: ", oldCamilaBalance)
	fmt.Println(withdraw)
	fmt.Println("Maria's balance before transfers:", oldMariaBalance)
	fmt.Println(transfer)
	fmt.Println("New balance in Maria's account: ", mariabalance)
	fmt.Println("After paying a bill, Camila's balance is: ", camilaAccount.ShowBalance())
}

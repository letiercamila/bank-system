package payment

type verifyAccoubt interface{
	Withdraw(amount float64) string
}

func PayBill(account verifyAccoubt, billValue float64){
	account.Withdraw(billValue)
}
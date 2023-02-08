package main

import (
	"bank/accounts"
	"fmt"
)

func InvoicePay(accounts verifyAccount, invoiceAmount float64) {
	accounts.Withdrawn(invoiceAmount)
}

type verifyAccount interface {
	Withdrawn(amount float64) string
}

func main() {

	denisAccount := accounts.SavingsAccounts{}
	patriciaAccount := accounts.CheckingAccounts{}

	denisAccount.Deposit(100)
	InvoicePay(&denisAccount, 60)

	patriciaAccount.Deposit(500)
	InvoicePay(&patriciaAccount, 300)

	fmt.Println(denisAccount.ShowAccountBalance())
	fmt.Println(patriciaAccount.ShowAccountBalance())

}

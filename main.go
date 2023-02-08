package main

import (
	"bank/accounts"
	"fmt"
)

func main() {

	denisAccount := accounts.SavingsAccounts{}
	//patriciaAccount := accounts.CheckingAccounts{}

	denisAccount.Deposit(5000)
	denisAccount.Withdrawn(300)

	fmt.Println(denisAccount.ShowAccountBalance())
	//fmt.Println(patriciaAccount)

}

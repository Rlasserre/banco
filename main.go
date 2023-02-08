package main

import (
	"bank/accounts"
	"fmt"
)

func main() {

	denisAccount := accounts.SavingsAccounts{}
	patriciaAccount := accounts.CheckingAccounts{}

	fmt.Println(denisAccount)
	fmt.Println(patriciaAccount)

}

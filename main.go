package main

import (
	"bank/accounts"
	"fmt"
)

func main() {

	silviaAccounts := accounts.CheckingAccounts{AccountOwner: "Silvia", AccountBalance: 500}
	ricardoAccounts := accounts.CheckingAccounts{AccountOwner: "Ricardo", AccountBalance: 1000}

	status := silviaAccounts.Transfer(200, &ricardoAccounts)

	fmt.Println(status)
	fmt.Println(silviaAccounts)
	fmt.Println(ricardoAccounts)
}

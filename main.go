package main

import (
	"fmt"
)

func main() {

	silviaAccounts := accounts.checkingAccounts{AccountOwner: "Silvia", AccountBalance: 500}
	ricardoAccounts := accounts.checkingAccounts{AccountOwner: "Ricardo", AccountBalance: 1000}

	status := silviaAccounts.Transfer(200, &ricardoAccounts)

	fmt.Println(status)
	fmt.Println(silviaAccounts)
	fmt.Println(ricardoAccounts)
}

package main

import (
	"fmt"
)

func main() {

	silviaAccounts := checkingAccounts{accountOwner: "Silvia", accountBalance: 500}
	ricardoAccounts := checkingAccounts{accountOwner: "Ricardo", accountBalance: 1000}

	status := silviaAccounts.Transfer(200, &ricardoAccounts)

	fmt.Println(status)
	fmt.Println(silviaAccounts)
	fmt.Println(ricardoAccounts)
}

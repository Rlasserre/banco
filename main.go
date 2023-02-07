package main

import (
	"bank/accounts"
	bankclients "bank/bankClients"
	"fmt"
)

func main() {

	clientSilvia := bankclients.AccountOwner{"Silvia", "123.456.789-00", "Professora"}
	silviaAccount := accounts.CheckingAccounts{clientSilvia, 123, 123456, 1000}

	ricardoAccounts := accounts.CheckingAccounts{
		AccountOwner: bankclients.AccountOwner{
			Name:      "Ricardo",
			CPF:       "321.654.987-11",
			Ocupation: "Motorista",
		},
		BankBranch:     123,
		AccountNumber:  789654,
		AccountBalance: 1000}

	status := silviaAccounts.Transfer(200, &ricardoAccounts)

	fmt.Println(status)
	fmt.Println(silviaAccount)
	fmt.Println(ricardoAccounts)
}

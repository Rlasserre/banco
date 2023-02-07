package main

import (
	"bank/accounts"
	bankclients "bank/bankClients"
	"fmt"
)

func main() {

	clientSilvia := bankclients.AccountOwner{Name: "Silvia", CPF: "123.456.789-00", Ocupation: "Professora"}
	silviaAccounts := accounts.CheckingAccounts{AccountOwner: clientSilvia, BankBranch: 123, AccountNumber: 123456}

	clientRicardo := bankclients.AccountOwner{Name: "Ricardo", CPF: "321.654.987.23", Ocupation: "Motorista"}
	ricardoAccounts := accounts.CheckingAccounts{AccountOwner: clientRicardo, BankBranch: 456, AccountNumber: 789561}

	silviaAccounts.Deposit(700)
	ricardoAccounts.Deposit(1500)
	fmt.Println(silviaAccounts, silviaAccounts.ShowAccountBalance())
	fmt.Println(ricardoAccounts, ricardoAccounts.ShowAccountBalance())

	status := silviaAccounts.Transfer(200, &ricardoAccounts)

	fmt.Println(status)
	fmt.Println(silviaAccounts, silviaAccounts.ShowAccountBalance())
	fmt.Println(ricardoAccounts, ricardoAccounts.ShowAccountBalance())
}

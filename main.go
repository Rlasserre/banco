package main

import "fmt"

type currentAccounts struct {
	accountOwner   string
	bankBranch     int
	accountNumber  int
	accountBalance float64
}

func (c *currentAccounts) withdrawn(withdrawnAmount float64) string {

	authorized := withdrawnAmount > 0 && withdrawnAmount <= c.accountBalance

	if authorized {
		c.accountBalance -= withdrawnAmount
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func main() {

	silviaAccounts := currentAccounts{}
	silviaAccounts.accountOwner = "Silvia"
	silviaAccounts.accountBalance = 500

	fmt.Println(silviaAccounts.accountBalance)

	fmt.Println(silviaAccounts.withdrawn(600))

	fmt.Println(silviaAccounts.accountBalance)
}

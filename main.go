package main

import "fmt"

type checkingAccounts struct {
	accountOwner   string
	bankBranch     int
	accountNumber  int
	accountBalance float64
}

func (c *checkingAccounts) Withdrawn(withdrawnAmount float64) string {

	authorized := withdrawnAmount > 0 && withdrawnAmount <= c.accountBalance

	if authorized {
		c.accountBalance -= withdrawnAmount
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *checkingAccounts) Deposit(depositAmount float64) (string, float64) {

	if depositAmount >= 0 {
		c.accountBalance += depositAmount
		return "Deposito realizado com sucesso", c.accountBalance
	} else {
		return "O valor do deposito invalido.", c.accountBalance
	}
}

func main() {

	silviaAccounts := checkingAccounts{}
	silviaAccounts.accountOwner = "Silvia"
	silviaAccounts.accountBalance = 500

	fmt.Println(silviaAccounts.accountBalance)
	status, valor := silviaAccounts.Deposit(300)
	fmt.Println(status, valor)

	//fmt.Println(silviaAccounts.accountBalance)
}

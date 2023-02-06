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

func (c *checkingAccounts) Transfer(transferAmount float64, targetAccount *checkingAccounts) bool {

	if transferAmount < c.accountBalance {
		c.accountBalance -= transferAmount
		targetAccount.Deposit(transferAmount)
		return true
	} else {
		return false
	}
}

func main() {

	silviaAccounts := checkingAccounts{accountOwner: "Silvia", accountBalance: 500}
	ricardoAccounts := checkingAccounts{accountOwner: "Ricardo", accountBalance: 1000}

	status := silviaAccounts.Transfer(200, &ricardoAccounts)

	fmt.Println(status)
	fmt.Println(silviaAccounts)
	fmt.Println(ricardoAccounts)
}

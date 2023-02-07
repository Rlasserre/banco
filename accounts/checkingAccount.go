package accounts

import bankclients "bank/bankClients"

type CheckingAccounts struct {
	AccountOwner   bankclients.AccountOwner
	BankBranch     int
	AccountNumber  int
	accountBalance float64
}

func (c *CheckingAccounts) Withdrawn(withdrawnAmount float64) string {

	authorized := withdrawnAmount > 0 && withdrawnAmount <= c.accountBalance

	if authorized {
		c.accountBalance -= withdrawnAmount
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *CheckingAccounts) Deposit(depositAmount float64) (string, float64) {

	if depositAmount >= 0 {
		c.accountBalance += depositAmount
		return "Deposito realizado com sucesso", c.accountBalance
	} else {
		return "O valor do deposito invalido.", c.accountBalance
	}
}

func (c *CheckingAccounts) Transfer(transferAmount float64,
	targetAccount *CheckingAccounts) bool {

	if transferAmount < c.accountBalance {
		c.accountBalance -= transferAmount
		targetAccount.Deposit(transferAmount)
		return true
	} else {
		return false
	}
}

func (c *CheckingAccounts) ShowAccountBalance() float64 {
	return c.accountBalance
}

package accounts

import bankclients "bank/bankClients"

type SavingsAccounts struct {
	AccountOwner                         bankclients.AccountOwner
	BankBranch, AccountNumber, Operation int
	accountBalance                       float64
}

func (c *SavingsAccounts) Withdrawn(withdrawnAmount float64) string {

	authorized := withdrawnAmount > 0 && withdrawnAmount <= c.accountBalance

	if authorized {
		c.accountBalance -= withdrawnAmount
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *SavingsAccounts) Deposit(depositAmount float64) (string, float64) {

	if depositAmount >= 0 {
		c.accountBalance += depositAmount
		return "Deposito realizado com sucesso", c.accountBalance
	} else {
		return "O valor do deposito invalido.", c.accountBalance
	}
}

func (c *SavingsAccounts) ShowAccountBalance() float64 {
	return c.accountBalance
}

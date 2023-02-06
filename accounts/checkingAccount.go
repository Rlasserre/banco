package accounts

type CheckingAccounts struct {
	AccountOwner   string
	BankBranch     int
	AccountNumber  int
	AccountBalance float64
}

func (c *CheckingAccounts) Withdrawn(withdrawnAmount float64) string {

	authorized := withdrawnAmount > 0 && withdrawnAmount <= c.AccountBalance

	if authorized {
		c.AccountBalance -= withdrawnAmount
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *CheckingAccounts) Deposit(depositAmount float64) (string, float64) {

	if depositAmount >= 0 {
		c.AccountBalance += depositAmount
		return "Deposito realizado com sucesso", c.AccountBalance
	} else {
		return "O valor do deposito invalido.", c.AccountBalance
	}
}

func (c *CheckingAccounts) Transfer(transferAmount float64,
	targetAccount *CheckingAccounts) bool {

	if transferAmount < c.AccountBalance {
		c.AccountBalance -= transferAmount
		targetAccount.Deposit(transferAmount)
		return true
	} else {
		return false
	}
}

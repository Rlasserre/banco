package main

import "fmt"

type currentAccounts struct {
	accountOwner   string
	bankBranch     int
	accountNumber  int
	accountBalance float64
}

func main() {

	rafaelAccount := currentAccounts{
		accountOwner:   "Rafael",
		bankBranch:     589,
		accountNumber:  123456,
		accountBalance: 125.58}

	rafaelAccount2 := currentAccounts{
		accountOwner:   "Rafael",
		bankBranch:     589,
		accountNumber:  123456,
		accountBalance: 125.58}

	fmt.Println(rafaelAccount == rafaelAccount2)
	// maraisaAccount := currentAccounts{"Maraisa", 756, 369852, 287.32}

	// var chrisAccount *currentAccounts
	// chrisAccount = new(currentAccounts)
	// chrisAccount.accountOwner = "Chris"

	// fmt.Println(rafaelAccount, maraisaAccount, chrisAccount)
}

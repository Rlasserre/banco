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

	maraisaAccount := currentAccounts{"Maraisa", 756, 369852, 287.32}

	fmt.Println(rafaelAccount, maraisaAccount)

}

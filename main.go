package main

import "fmt"

type currentAccounts struct {
	accountOwner   string
	bankBranch     int
	accountNumber  int
	accountBalance float64
}

func main() {

	fmt.Println(currentAccounts{})

}

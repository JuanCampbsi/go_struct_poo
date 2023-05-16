package main

import (
	"fmt"
	c "projects/goStudyng/POO/Bank/src/account"
)

func main() {
	var accountLegal *c.AccountCurrent
	accountLegal = new(c.AccountCurrent)
	accountLegal.Title = "Julia"
	accountLegal.Balance = 1000

	fmt.Println("Before account legal:", accountLegal.Balance)
	fmt.Println(accountLegal.WithDrawMoney(200))
	fmt.Println("After account legal:", accountLegal.Balance)

	status, value := accountLegal.DepositMoney(400)
	fmt.Println(status, value)

	var accountPhisical *c.AccountCurrent
	accountPhisical = new(c.AccountCurrent)
	accountPhisical.Title = "Juan"
	accountPhisical.Balance = 1000

	fmt.Println("\nBefore account Phisical:", accountPhisical.Balance)
	fmt.Println(accountPhisical.WithDrawMoney(300))
	fmt.Println("After account Phisical:", accountPhisical.Balance)

	statusPhisical, valuePhisical := accountPhisical.DepositMoney(-100)
	fmt.Println(statusPhisical, valuePhisical)

	statusTransfer := accountPhisical.TransferAccount(-200, accountLegal)
	fmt.Println(statusTransfer, *accountPhisical, *accountLegal)
}

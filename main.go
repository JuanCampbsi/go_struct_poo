package main

import (
	"fmt"
	account "goStudyng/POO/Bank/Account"
)

func main() {
	var accountLegal *account.AccountCurrent
	accountLegal = new(account.AccountCurrent)
	accountLegal.Title = "Julia"
	accountLegal.Balance = 1000

	fmt.Println("Before account legal:", accountLegal.Balance)
	fmt.Println(accountLegal.WithDrawMoney(200))
	fmt.Println("After account legal:", accountLegal.Balance)

	status, value := accountLegal.DepositMoney(400)
	fmt.Println(status, value)

	var accountPhisical *account.AccountCurrent
	accountPhisical = new(account.AccountCurrent)
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

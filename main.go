package main

import "fmt"

type AccountCurrent struct {
	title         string
	numberAgency  int
	numberAccount int
	balance       float64
}

func main() {
	var accountLegal *AccountCurrent
	accountLegal = new(AccountCurrent)
	accountLegal.title = "Julia"
	accountLegal.balance = 1000

	fmt.Println("Before account legal:", accountLegal.balance)
	fmt.Println(accountLegal.withDrawMoney(200))
	fmt.Println("After account legal:", accountLegal.balance)

	status, value := accountLegal.depositMoney(400)
	fmt.Println(status, value)

	var accountPhisical *AccountCurrent
	accountPhisical = new(AccountCurrent)
	accountPhisical.title = "Juan"
	accountPhisical.balance = 1000

	fmt.Println("\nBefore account Phisical:", accountPhisical.balance)
	fmt.Println(accountPhisical.withDrawMoney(300))
	fmt.Println("After account Phisical:", accountPhisical.balance)

	statusPhisical, valuePhisical := accountPhisical.depositMoney(-100)
	fmt.Println(statusPhisical, valuePhisical)

	statusTransfer := accountPhisical.transferAccount(-200, accountLegal)
	fmt.Println(statusTransfer, *accountPhisical, *accountLegal)
}

func (c *AccountCurrent) withDrawMoney(value float64) string {
	if value > 0 && value <= c.balance {
		c.balance -= value
		return "Successful withdrawal"
	} else {
		return "Insufficient balance"
	}
}

func (c *AccountCurrent) depositMoney(value float64) (string, float64) {
	if value > 0 {
		c.balance += value
		return "\nDeposit made successfully Mr. " + c.title + " your current balance is: ", c.balance
	} else {
		return "\nDeposit has not been made Mr. " + c.title + " your current balance is: ", c.balance
	}
}

func (c *AccountCurrent) transferAccount(value float64, targetAccount *AccountCurrent) string {
	if value > 0 && value < c.balance {
		c.balance -= value
		targetAccount.depositMoney(value)
		return "\nMr. " + c.title + " Transfer successfully for " + targetAccount.title
	} else {
		return "\nMr. " + c.title + " Transfer unsuccessfuly for " + targetAccount.title
	}
}

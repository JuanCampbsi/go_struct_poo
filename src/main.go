package main

import (
	"fmt"
	c "projects/goStudyng/POO/Bank/src/account"
	client "projects/goStudyng/POO/Bank/src/client"
)

type verifyAccount interface {
	WithDrawMoney(value float64) string
}

func PayBill(account verifyAccount, valueBill float64) {
	account.WithDrawMoney(valueBill)
}

func main() {
	contaTeste := *&c.AccountSavings{}
	contaTeste.DepositMoney(200)
	fmt.Println(contaTeste.GetBalanceUser())
	PayBill(&contaTeste, 60)
	fmt.Println(contaTeste.GetBalanceUser())

	clientLegal := *&client.Customers{
		Name:         "Julia",
		CPF:          "121.232.232-12",
		Profissional: "Fisioterapia"}

	accountLegal := *&c.AccountCurrent{
		Title:         clientLegal,
		NumberAgency:  122,
		NumberAccount: 20033,
		Balance:       800}

	fmt.Println("Before account legal:", accountLegal.Balance)
	fmt.Println(accountLegal.WithDrawMoney(200))
	fmt.Println("After account legal:", accountLegal.Balance)

	status, value := accountLegal.DepositMoney(400)
	fmt.Println(status, value)

	clientPhisical := *&client.Customers{
		Name:         "Juan",
		CPF:          "157.232.232-12",
		Profissional: "Dev go"}
	accountPhisical := *&c.AccountCurrent{
		Title:         clientPhisical,
		NumberAgency:  122,
		NumberAccount: 20033,
		Balance:       1000}

	fmt.Println("\nBefore account Phisical:", accountPhisical.GetBalanceUser())
	fmt.Println(accountPhisical.WithDrawMoney(300))
	fmt.Println("After account Phisical:", accountPhisical.GetBalanceUser())

	statusPhisical, valuePhisical := accountPhisical.DepositMoney(-100)
	fmt.Println(statusPhisical, valuePhisical)

	statusTransfer := accountPhisical.TransferAccount(200, &accountLegal)
	fmt.Println("\n", statusTransfer)
	fmt.Println("\n", *&accountPhisical, "\n", *&accountLegal)

	fmt.Println(accountPhisical.GetBalanceUser())
	PayBill(&accountPhisical, 60)
	fmt.Println(accountPhisical.GetBalanceUser())
}

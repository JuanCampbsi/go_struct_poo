package main

import "fmt"

type AccountCurrent struct {
	title         string
	numberAgency  int
	numberAccount int
	balance       float64
}

func main() {
	accountPhysical := AccountCurrent{title: "Juan", numberAgency: 122, numberAccount: 01545, balance: 125.5}
	accountPhysical2 := AccountCurrent{title: "Juan", numberAgency: 122, numberAccount: 01545, balance: 125.5}
	fmt.Println(accountPhysical == accountPhysical2)

	var accountLegal *AccountCurrent
	accountLegal = new(AccountCurrent)
	accountLegal.title = "Julia"
	accountLegal.numberAccount = 999

	accountSlices := []*AccountCurrent{
		{
			title:         "Juan",
			numberAgency:  999,
			numberAccount: 123456,
			balance:       1000.0,
		},
		{
			title:         "julia",
			numberAgency:  888,
			numberAccount: 654321,
			balance:       2000.0,
		},
	}
	fmt.Println("before append: ", len(accountSlices))
	accountSlices = append(accountSlices, &AccountCurrent{
		title:         "Campos",
		numberAgency:  999,
		numberAccount: 123456,
		balance:       1000.0,
	})

	fmt.Println("after append: ", len(accountSlices))
	fmt.Println(*accountSlices[len(accountSlices)-1])
	fmt.Println(&accountLegal)
}

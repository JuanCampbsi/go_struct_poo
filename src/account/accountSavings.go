package account

import client "projects/goStudyng/POO/Bank/src/client"

type AccountSavings struct {
	Title                                  client.Customers
	NumberAgency, NumberAccount, Operation int
	Balance                                float64
}

func (c *AccountSavings) WithDrawMoney(value float64) string {
	if value > 0 && value <= c.Balance {
		c.Balance -= value
		return "Successful withdrawal"
	} else {
		return "Insufficient Balance"
	}
}

func (c *AccountSavings) DepositMoney(value float64) (string, float64) {
	if value > 0 {
		c.Balance += value
		return "\nDeposit made successfully Mr. " + c.Title.Name + " your current Balance is: ", c.Balance
	} else {
		return "\nDeposit has not been made Mr. " + c.Title.Name + " your current Balance is: ", c.Balance
	}
}

func (c *AccountSavings) TransferAccount(value float64, targetAccount *AccountSavings) string {
	if value > 0 && value < c.Balance {
		c.Balance -= value
		targetAccount.DepositMoney(value)
		return "\nMr. " + c.Title.Name + " Transfer successfully for " + targetAccount.Title.Name
	} else {
		return "\nMr. " + c.Title.Name + " Transfer unsuccessfuly for " + targetAccount.Title.Name
	}
}

func (c *AccountSavings) GetBalanceUser() float64 {
	return c.Balance
}

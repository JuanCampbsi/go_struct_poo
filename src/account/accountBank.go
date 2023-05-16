package account

type AccountCurrent struct {
	Title         string
	NumberAgency  int
	NumberAccount int
	Balance       float64
}

func (c *AccountCurrent) WithDrawMoney(value float64) string {
	if value > 0 && value <= c.Balance {
		c.Balance -= value
		return "Successful withdrawal"
	} else {
		return "Insufficient Balance"
	}
}

func (c *AccountCurrent) DepositMoney(value float64) (string, float64) {
	if value > 0 {
		c.Balance += value
		return "\nDeposit made successfully Mr. " + c.Title + " your current Balance is: ", c.Balance
	} else {
		return "\nDeposit has not been made Mr. " + c.Title + " your current Balance is: ", c.Balance
	}
}

func (c *AccountCurrent) TransferAccount(value float64, targetAccount *AccountCurrent) string {
	if value > 0 && value < c.Balance {
		c.Balance -= value
		targetAccount.DepositMoney(value)
		return "\nMr. " + c.Title + " Transfer successfully for " + targetAccount.Title
	} else {
		return "\nMr. " + c.Title + " Transfer unsuccessfuly for " + targetAccount.Title
	}
}

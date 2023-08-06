package model

import "errors"

type Bank struct {
	Accounts []Account
}

func (b *Bank) FindAccount(accountNumber string) (Account, error) {
	for _, v := range b.Accounts {
		if v.AccountNumber == accountNumber {
			return v, nil
		}
	}
	return Account{}, errors.New("account not found")
}

// InsufficientFundsError is raised when the account doesn't have enough money.
type InsufficientFundsError struct{}

func (m *InsufficientFundsError) Error() string {
	return "Insufficient Funds"
}

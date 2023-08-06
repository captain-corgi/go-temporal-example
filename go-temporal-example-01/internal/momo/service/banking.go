package service

import (
	"errors"
	"math/rand"

	"github.com/captain-corgi/go-temporal-example-01/internal/momo/model"
)

// BankingService mocks interaction with a bank API. It supports withdrawals
// and deposits
type BankingService struct {
	// the hostname is to make it more realistic. This code does not
	// actually make any network calls.
	Hostname string
}

// Withdraw simulates a Withdrawwal from a bank.
// Acceptsthe account number (string), amount (int), and a reference ID (string)
// for idempotent transaction tracking.
// Returns a transaction id when successful
// Returns various errors based on amount and account number.
func (client BankingService) Withdraw(accountNumber string, amount int, referenceID string) (string, error) {
	acct, err := mockBank.FindAccount(accountNumber)

	if err != nil {
		return "", &model.InvalidAccountError{}
	}

	if amount > int(acct.Balance) {
		return "", &model.InsufficientFundsError{}
	}

	return generateTransactionID("W", 10), nil
}

// Deposit simulates a Withdrawwal from a bank.
// Accepts the account number (string), amount (int), and a reference ID (string)
// for idempotent transaction tracking.
// Returns a transaction id when successful
// Returns InvalidAccountError if the account is invalid
func (client BankingService) Deposit(accountNumber string, amount int, referenceID string) (string, error) {

	_, err := mockBank.FindAccount(accountNumber)
	if err != nil {
		return "", &model.InvalidAccountError{}
	}

	return generateTransactionID("D", 10), nil
}

// DepositThatFails simulates an unknonw error.
func (client BankingService) DepositThatFails(accountNumber string, amount int, referenceID string) (string, error) {
	return "", errors.New("this deposit has failed")
}

func generateTransactionID(prefix string, length int) string {
	randChars := make([]byte, length)
	for i := range randChars {
		allowedChars := "0123456789"
		randChars[i] = allowedChars[rand.Intn(len(allowedChars))]
	}
	return prefix + string(randChars)
}

package model

type Account struct {
	AccountNumber string
	Balance       int64
}

// InvalidAccountError is raised when the account number is invalid
type InvalidAccountError struct{}

func (m *InvalidAccountError) Error() string {
	return "Account number supplied is invalid"
}

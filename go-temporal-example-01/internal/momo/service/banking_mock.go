package service

import "github.com/captain-corgi/go-temporal-example-01/internal/momo/model"

// our mock bank
var mockBank = &model.Bank{
	Accounts: []model.Account{
		{AccountNumber: "85-150", Balance: 2000},
		{AccountNumber: "43-812", Balance: 0},
	},
}

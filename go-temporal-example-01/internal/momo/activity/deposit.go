package activity

import (
	"context"
	"fmt"
	"log"

	"github.com/captain-corgi/go-temporal-example-01/internal/momo/model"
	"github.com/captain-corgi/go-temporal-example-01/internal/momo/service"
)

func Deposit(ctx context.Context, data model.PaymentDetails) (string, error) {
	log.Printf("Depositing $%d into account %s.\n\n",
		data.Amount,
		data.TargetAccount,
	)

	referenceID := fmt.Sprintf("%s-deposit", data.ReferenceID)
	bank := service.BankingService{Hostname: "bank-api.example.com"}
	// Uncomment the next line and comment the one after that to simulate an unknown failure
	// confirmation, err := bank.DepositThatFails(data.TargetAccount, data.Amount, referenceID)
	confirmation, err := bank.Deposit(data.TargetAccount, data.Amount, referenceID)
	return confirmation, err
}

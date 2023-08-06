package activity

import (
	"context"
	"fmt"
	"log"

	"github.com/captain-corgi/go-temporal-example-01/internal/momo/model"
	"github.com/captain-corgi/go-temporal-example-01/internal/momo/service"
)

func Withdraw(ctx context.Context, data model.PaymentDetails) (string, error) {
	log.Printf("Withdrawing $%d from account %s.\n\n",
		data.Amount,
		data.SourceAccount,
	)

	referenceID := fmt.Sprintf("%s-withdrawal", data.ReferenceID)
	bank := service.BankingService{Hostname: "bank-api.example.com"}
	confirmation, err := bank.Withdraw(data.SourceAccount, data.Amount, referenceID)
	return confirmation, err
}

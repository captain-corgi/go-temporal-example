package activity

import (
	"context"
	"fmt"
	"log"

	"github.com/captain-corgi/go-temporal-example-01/internal/momo/model"
	"github.com/captain-corgi/go-temporal-example-01/internal/momo/service"
)

func Refund(ctx context.Context, data model.PaymentDetails) (string, error) {
	log.Printf("Refunding $%v back into account %v.\n\n",
		data.Amount,
		data.SourceAccount,
	)

	referenceID := fmt.Sprintf("%s-refund", data.ReferenceID)
	bank := service.BankingService{Hostname: "bank-api.example.com"}
	confirmation, err := bank.Deposit(data.SourceAccount, data.Amount, referenceID)
	return confirmation, err
}

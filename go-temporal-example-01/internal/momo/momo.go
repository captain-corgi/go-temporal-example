package momo

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"

	"github.com/captain-corgi/go-temporal-example-01/internal/momo/model"
	"github.com/captain-corgi/go-temporal-example-01/internal/momo/workflow"
)

type (
	Executer interface {
		TransferMoney() error
	}
	ExecuterImpl struct {
		TemporalClient client.Client
	}
)

func NewExecuter(temporalClient client.Client) *ExecuterImpl {
	return &ExecuterImpl{
		TemporalClient: temporalClient,
	}
}

func (r *ExecuterImpl) TransferMoney() error {
	// Prepare data
	var (
		input = model.PaymentDetails{
			SourceAccount: "85-150",
			TargetAccount: "43-812",
			Amount:        250,
			ReferenceID:   "12345",
		}
		options = client.StartWorkflowOptions{
			ID:        "pay-invoice-701",
			TaskQueue: model.MoneyTransferTaskQueueName,
		}
		result string
	)
	log.Printf("Starting transfer from account %s to account %s for %d", input.SourceAccount, input.TargetAccount, input.Amount)

	// Execute workflow
	we, err := r.TemporalClient.ExecuteWorkflow(context.Background(), options, workflow.MoneyTransfer, input)
	if err != nil {
		log.Printf("Unable to start the Workflow: %s", err.Error())
		return err
	}
	log.Printf("WorkflowID: %s RunID: %s\n", we.GetID(), we.GetRunID())

	// Get result
	if err = we.Get(context.Background(), &result); err != nil {
		log.Printf("Unable to get Workflow result: %s", err.Error())
		return err
	}
	log.Println(result)

	return nil
}

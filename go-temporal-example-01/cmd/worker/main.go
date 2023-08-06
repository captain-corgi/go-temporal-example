package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/captain-corgi/go-temporal-example-01/internal/momo/activity"
	"github.com/captain-corgi/go-temporal-example-01/internal/momo/model"
	"github.com/captain-corgi/go-temporal-example-01/internal/momo/workflow"
)

func main() {
	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client.", err)
	}
	defer c.Close()

	// Create worker
	w := worker.New(c, model.MoneyTransferTaskQueueName, worker.Options{})

	// This worker hosts both Workflow and Activity functions.
	w.RegisterWorkflow(workflow.MoneyTransfer)
	w.RegisterActivity(activity.Withdraw)
	w.RegisterActivity(activity.Deposit)
	w.RegisterActivity(activity.Refund)

	// Start listening to the Task Queue.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}

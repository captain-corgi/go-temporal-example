package model

const MoneyTransferTaskQueueName = "TRANSFER_MONEY_TASK_QUEUE"

type PaymentDetails struct {
	ReferenceID string

	SourceAccount string
	TargetAccount string
	Amount        int
}

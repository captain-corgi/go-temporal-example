package main

import (
	"log"

	"go.temporal.io/sdk/client"

	"github.com/captain-corgi/go-temporal-example-01/internal/momo"
)

// main function represent transport layer
func main() {
	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
	}
	defer c.Close()

	// Execute each usecases
	momoTasks := momo.NewExecuter(c)
	if err := momoTasks.TransferMoney(); err != nil {
		log.Fatalln(err)
	}
}

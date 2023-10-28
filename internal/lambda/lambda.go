package lambda

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Gabriel-Jeronimo/eth-keygen-api/internal/eth"
	"github.com/Gabriel-Jeronimo/eth-keygen-api/internal/keypair"
	"github.com/aws/aws-lambda-go/events"
)

type Transaction struct {
	To    string `json:"to"`
	Value string `json:"value"`
}

func Handler(ctx context.Context, event events.SQSEvent) error {
	var transaction Transaction

	for _, record := range event.Records {
		fmt.Printf("Processing message: %v", record.Body)
		err := json.Unmarshal([]byte(record.Body), &transaction)

		if err != nil {
			log.Printf("ERROR: Failed unmarshal record body into a struct: %v", err)
			return err
		}

		ethClient, _ := eth.Connect()

		privateKeyString, publicKeyString, err := keypair.GenerateKeypair()

		if err != nil {
			return err
		}

		address, err := keypair.GetAddress(publicKeyString)

		if err != nil {
			return err
		}

		eth.SignAndPushTransaction(ethClient, address, transaction.To, transaction.Value, privateKeyString)

	}
	return nil
}

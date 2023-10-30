package lambda

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Gabriel-Jeronimo/eth-keygen-api/internal/eth"
	"github.com/Gabriel-Jeronimo/eth-keygen-api/internal/keypair"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Transaction struct {
	To    string `json:"to"`
	Value string `json:"value"`
}

type Body struct {
	TransactionHash string `json:"transactionHash"`
}

type Response struct {
	StatusCode uint64 `json:"statusCode"`
}

func Handler(ctx context.Context, event events.SQSEvent) (Response, error) {
	var transaction Transaction

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	}))

	sqsClient := sqs.New(sess)

	queueURL := "https://sqs.us-west-2.amazonaws.com/029646376518/KeygenQueue"

	for _, record := range event.Records {
		fmt.Printf("Processing message: %v", record.Body)
		err := json.Unmarshal([]byte(record.Body), &transaction)

		if err != nil {
			log.Printf("ERROR: Failed unmarshal record body into a struct: %v", err)
			return Response{}, err

		}

		ethClient, _ := eth.Connect()

		privateKeyString, publicKeyString, err := keypair.GenerateKeypair()

		if err != nil {
			return Response{StatusCode: 500}, err
		}

		address, err := keypair.GetAddress(publicKeyString)

		if err != nil {
			return Response{StatusCode: 500}, err
		}

		faucetTx, err := eth.FaucetToAddress(ethClient, transaction.Value, address)

		if err != nil {
			return Response{StatusCode: 500}, err
		}

		eth.WaitTx(ethClient, faucetTx)

		tx, err := eth.SignAndPushTransaction(ethClient, address, transaction.To, transaction.Value, privateKeyString)

		if err != nil {
			return Response{StatusCode: 500}, err
		}

		fmt.Printf("Transaction confirmed: %s\n", tx)

		_, err = sqsClient.DeleteMessage(&sqs.DeleteMessageInput{
			QueueUrl:      &queueURL,
			ReceiptHandle: &record.ReceiptHandle,
		})

		if err != nil {
			return Response{StatusCode: 500}, err
		}
	}

	return Response{
		StatusCode: 200,
	}, nil
}

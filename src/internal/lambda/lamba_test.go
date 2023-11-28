package lambda_test

import (
	"context"
	"testing"

	"github.com/Gabriel-Jeronimo/eth-keygen-api/src/internal/lambda"
	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	request := events.SQSEvent{
		Records: []events.SQSMessage{},
	}

	_, err := lambda.Handler(context.Background(), request)

	if err != nil {
		t.Fatalf("Handler() returned an err: %s", err)
	}
}

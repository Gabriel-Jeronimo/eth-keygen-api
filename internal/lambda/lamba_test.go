package lambda_test

import (
	"context"
	"testing"

	"github.com/Gabriel-Jeronimo/eth-keygen-api/internal/lambda"
	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	request := events.SQSEvent{
		Records: []events.SQSMessage{
			{
				Body: `{
					"to": "0xAbC123aaa",
					"value": "0x1bc16d674ec80000"
					}`,
			},
		},
	}

	_, err := lambda.Handler(context.Background(), request)

	if err != nil {
		t.Fatalf("Handler failed")
	}
}

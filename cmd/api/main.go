package main

import (
	"log"
	"net/http"

	"github.com/Gabriel-Jeronimo/eth-keygen-api/internal/keypair"
	"github.com/Gabriel-Jeronimo/eth-keygen-api/internal/lambda"
	aws "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	keypair.InitRoutes()

	aws.Start(lambda.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package eth_test

import (
	"math/big"
	"os"
	"testing"

	"github.com/Gabriel-Jeronimo/eth-keygen-api/internal/eth"
)

func TestConnect(t *testing.T) {
	os.Setenv("INFURA_PROJECT_ID", "3123122312")

	_, err := eth.Connect()

	if err != nil {
		t.Fatalf("TestConnect() returned an error: %v", err)
	}
}

func TestCalculateTransaction(t *testing.T) {
	transactionCost, err := eth.CalculateTransactionCost(2100, big.NewInt(10), "0x7f080c0a4000")

	if err != nil {
		t.Fatalf("CalculateTransaction() returned an error: %v", err)
	}

	if transactionCost == "0x2d4c1870" {
		t.Logf("Output equal 0x2d4c1870, as expected")
	}
}

func TestCalculateTransactionWithInvalidParameters(t *testing.T) {
	_, err := eth.CalculateTransactionCost(2100, big.NewInt(10), "invalid")

	if err == nil {
		t.Errorf("CalculateTransactionCost() should have returned an error for an invalid hex value")
	}
}

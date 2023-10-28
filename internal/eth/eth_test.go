package eth_test

import (
	"os"
	"testing"

	"github.com/Gabriel-Jeronimo/eth-keygen-api/internal/eth"
)

func TestConnect(t *testing.T) {
	os.Setenv("INFURA_PROJECT_ID", "abc")
	os.Setenv("INFURA_SECRET_API", "another one")

	_, err := eth.Connect()

	if err != nil {
		t.Fatalf("TestConnect() returned an error: %v", err)
	}
}

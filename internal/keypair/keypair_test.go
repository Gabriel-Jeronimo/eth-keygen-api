package keypair_test

import (
	"fmt"
	"testing"

	"github.com/Gabriel-Jeronimo/eth-keygen-api/internal/keypair"
)

func TestGenerateKeypair(t *testing.T) {
	privateKey, publicKey, err := keypair.GenerateKeypair()

	if err != nil {
		t.Fatalf("GenerateKeypair() = [%q, %q], %v", privateKey, publicKey, err)
	}

	fmt.Printf("privateKey: %q, publicKey: %q\n", privateKey, publicKey)

}

func TestGetAddress(t *testing.T) {
	address, err := keypair.GetAddress()

	if err != nil {
		t.Fatalf("GetAddress() = [%q], %v", address, err)
	}

	fmt.Printf("Address: %q\n", address)
}

package keypair_test

import (
	"crypto/ecdsa"
	"encoding/hex"
	"testing"

	"github.com/Gabriel-Jeronimo/eth-keygen-api/internal/keypair"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestGenerateKeypair(t *testing.T) {
	privateKey, publicKey, err := keypair.GenerateKeypair()

	if err != nil {
		t.Fatalf("GenerateKeypair() returned an error: %v", err)
	}

	if privateKey == "" || publicKey == "" {
		t.Fatalf("GenerateKeypair() returned empty keys")
	}

	t.Logf("privateKey: %q, publicKey: %q", privateKey, publicKey)
}

func TestGetAddress(t *testing.T) {
	privateKey, _ := crypto.GenerateKey()

	publicKey := privateKey.Public()

	publicKeyBytes := crypto.FromECDSAPub(publicKey.(*ecdsa.PublicKey))

	address, err := keypair.GetAddress(hex.EncodeToString(publicKeyBytes))

	if err != nil {
		t.Fatalf("GetAddress() = [%q], %v", address, err)
	}

	t.Logf("Address: %q", address)
}

func TestGetAddressWithInvalidPublicKey(t *testing.T) {
	_, err := keypair.GetAddress("definitely not a invalid public key")

	if err == nil {
		t.Errorf("GetAddress() should have returned an error for an invalid public key")
	}
}

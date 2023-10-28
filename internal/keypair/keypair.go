package keypair

import (
	"crypto/ecdsa"
	"encoding/hex"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateKeypair() (string, string, error) {
	privateKey, err := crypto.GenerateKey()

	if err != nil {
		log.Printf("ERROR: Failed to generate keypair: %v", err)
		return "", "", err
	}

	publicKey := privateKey.Public()

	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKeyBytes := crypto.FromECDSAPub(publicKey.(*ecdsa.PublicKey))

	return hex.EncodeToString(privateKeyBytes), hex.EncodeToString(publicKeyBytes), nil
}

func GetAddress(publicKey string) (string, error) {
	publicKeyBytes, err := hex.DecodeString(publicKey)

	if err != nil {
		log.Printf("ERROR: Failed to decode public key string into hex: %v", err)
		return "", err
	}

	pubKey, err := crypto.UnmarshalPubkey(publicKeyBytes)

	if err != nil {
		log.Printf("ERROR: Failed to decode public key hex into secp256k1 public key: %v", err)
		return "", err
	}

	address := crypto.PubkeyToAddress(*pubKey)

	return address.String(), nil
}

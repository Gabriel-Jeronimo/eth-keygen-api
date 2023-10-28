package eth

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

func Connect() (*ethclient.Client, error) {
	// infuraProjectId := os.Getenv("INFURA_PROJECT_ID")
	// InfuraSecretApi := os.Getenv("INFURA_SECRET_API")

	infuraProjectID := "c510981aace04b1e80506a0ff746b7d2"
	infuraURL := "https://sepolia.infura.io/v3/" + infuraProjectID

	ethClient, err := ethclient.Dial(infuraURL)

	if err != nil {
		log.Printf("ERROR: Failed to connect with the eth client: %v", err)
		return nil, err
	}

	return ethClient, nil
}

func SignAndPushTransaction(ethClient *ethclient.Client, from string, to string, value string, privateKeyString string) (string, error) {
	toAddress := common.HexToAddress(to)
	amount := new(big.Int)
	amount.SetString(value[2:], 16)

	gasLimit := uint64(21000)
	gasPrice := big.NewInt(10 * params.GWei)

	nonce, err := getNounce(ethClient, from)

	if err != nil {
		return "", err
	}

	tx := types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, nil)

	privateKey, err := stringToPrivateKey(privateKeyString)

	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(11155111)), privateKey)

	if err != nil {
		log.Printf("ERROR: Failed to sign transaction: %v", err)
		return "", err
	}

	err = ethClient.SendTransaction(context.Background(), signedTx)

	if err != nil {
		log.Printf("ERROR: Failed to send transaction: %v", err)
		return "", err
	}

	hash := signedTx.Hash().Hex()

	return hash, nil
}

func getNounce(ethClient *ethclient.Client, from string) (uint64, error) {
	nonce, err := ethClient.PendingNonceAt(context.Background(), common.HexToAddress(from))

	if err != nil {
		log.Printf("ERROR: Failed to get nonce: %v", err)
		return 0, err
	}

	return nonce, nil
}

func stringToPrivateKey(privateKeyString string) (*ecdsa.PrivateKey, error) {
	privateKeyBytes, err := hex.DecodeString(privateKeyString)
	if err != nil {
		log.Fatalf("Failed to decode private key: %v", err)
		return nil, err
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		log.Fatalf("Failed to convert private key to ECDSA Private Key: %v", err)
		return nil, err
	}

	return privateKey, nil
}

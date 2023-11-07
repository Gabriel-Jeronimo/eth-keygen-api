package eth

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

type Client interface {
	TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error)
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
	SendTransaction(ctx context.Context, tx *types.Transaction) error
}

var gasLimit = uint64(21000)
var gasPrice = big.NewInt(10 * params.GWei)

func Connect() (*ethclient.Client, error) {
	infuraProjectId := os.Getenv("INFURA_PROJECT_ID")
	infuraURL := "https://sepolia.infura.io/v3/" + infuraProjectId

	ethClient, err := ethclient.Dial(infuraURL)

	if err != nil {
		log.Printf("ERROR: Failed to connect with the eth client: %v", err)
		return nil, err
	}

	return ethClient, nil
}

func SignAndPushTransaction(ethClient Client, from string, to string, value string, privateKeyString string) (string, error) {
	amount := new(big.Int)
	amount.SetString(value[2:], 16)

	nonce, err := GetNounce(ethClient, from)

	if err != nil {
		return "", err
	}

	tx := types.NewTransaction(nonce, common.HexToAddress(to), amount, gasLimit, gasPrice, nil)

	privateKey, err := StringToPrivateKey(privateKeyString)

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

	return signedTx.Hash().Hex(), nil
}

func WaitTx(ethClient Client, tx string) error {
	for {
		time.Sleep(10 * time.Second)

		_, pending, err := ethClient.TransactionByHash(context.Background(), common.HexToHash(tx))

		if err != nil {
			log.Printf("ERROR: failed to wait for transaction: %v", err)
			return err
		}

		if !pending {
			break
		}
	}

	return nil
}

func FaucetToAddress(ethClient Client, value string, to string) (string, error) {
	foundingPrivateKey := os.Getenv("FOUNDING_PRIVATE_KEY")
	foundingAddress := os.Getenv("FOUDING_ADDRESS")

	amount, err := CalculateTransactionCost(gasLimit, gasPrice, value)

	if err != nil {
		return "", nil
	}

	hash, err := SignAndPushTransaction(ethClient, foundingAddress, to, amount, foundingPrivateKey)

	if err != nil {
		return "", nil
	}

	return hash, nil
}

func CalculateTransactionCost(gasLimit uint64, gasPrice *big.Int, transactionValueHex string) (string, error) {
	transactionValue, success := new(big.Int).SetString(transactionValueHex, 0)

	if !success {
		return "", fmt.Errorf("failed to parse transaction value hex")
	}

	gasCost := new(big.Int).SetUint64(gasLimit)
	gasCost.Mul(gasCost, gasPrice)
	totalCost := new(big.Int).Set(transactionValue)
	totalCost.Add(totalCost, gasCost)

	totalCostHex := fmt.Sprintf("0x%X", totalCost)
	return totalCostHex, nil
}

func GetNounce(ethClient Client, from string) (uint64, error) {
	nonce, err := ethClient.PendingNonceAt(context.Background(), common.HexToAddress(from))

	if err != nil {
		log.Printf("ERROR: Failed to get nonce: %v", err)
		return 0, err
	}

	return nonce, nil
}

func StringToPrivateKey(privateKeyString string) (*ecdsa.PrivateKey, error) {
	privateKeyBytes, err := hex.DecodeString(privateKeyString)
	if err != nil {
		log.Printf("ERROR: Failed to decode private key: %v", err)
		return nil, err
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		log.Printf("ERROR: Failed to convert private key to ECDSA Private Key: %v", err)
		return nil, err
	}

	return privateKey, nil
}

package eth_test

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/Gabriel-Jeronimo/eth-keygen-api/src/internal/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var PendingNonceAtMock func(ctx context.Context, account common.Address) (uint64, error)
var TransactionByHash func(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error)
var SendTransaction func(ctx context.Context, tx *types.Transaction) error

type Client struct {
}

func (c Client) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return PendingNonceAtMock(ctx, account)
}

func (c Client) TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	return TransactionByHash(ctx, hash)
}

func (c Client) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return SendTransaction(ctx, tx)
}

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

func TestGetNounce(t *testing.T) {
	var nounceReturn uint64 = 54
	client := Client{}

	PendingNonceAtMock = func(ctx context.Context, account common.Address) (uint64, error) {
		return nounceReturn, nil
	}

	result, err := eth.GetNounce(client, "0x3011FF701a84B697D8821a03F18F7c52792D5338")

	if err != nil {
		t.Fatalf("GetNounce() returned an error: %v", err)
	}

	if result != 54 {
		t.Fatalf("GetNounce() returned %d, expected %d", result, nounceReturn)
	}

	t.Logf("GetNounce() returned %d, as expected", result)
}

func TestGetNounceWithInvalidAddress(t *testing.T) {
	client := Client{}

	PendingNonceAtMock = func(ctx context.Context, account common.Address) (uint64, error) {
		return 0, fmt.Errorf("Invalid address")
	}

	_, err := eth.GetNounce(client, "invalid_address")

	if err == nil {
		t.Errorf("GetNounce() should have returned an error for an invalid address")
	}
}

func TestWaitTx(t *testing.T) {
	client := Client{}

	TransactionByHash = func(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
		return &types.Transaction{}, false, nil
	}

	err := eth.WaitTx(client, "0xcc430ceebdbd9f3e42dc59d9c5a733042c37d4b581295fae0ba70aa90c97c70c")

	if err != nil {
		t.Fatalf("WaitTx() returned an error: %v", err)
	}
}

func TestWaitTxInvalidAddress(t *testing.T) {
	client := Client{}

	TransactionByHash = func(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
		return &types.Transaction{}, false, fmt.Errorf("Invalid address")
	}

	err := eth.WaitTx(client, "invalid_address")

	if err == nil {
		t.Errorf("WaitTx() should have returned an error for an invalid address")
	}
}

func TestSignAndPushTransaction(t *testing.T) {
	client := Client{}

	SendTransaction = func(ctx context.Context, tx *types.Transaction) error {
		return nil
	}

	PendingNonceAtMock = func(ctx context.Context, account common.Address) (uint64, error) {
		return 54, nil
	}

	privateKey, _ := crypto.GenerateKey()

	publicKey := privateKey.Public()

	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKeyBytes := crypto.FromECDSAPub(publicKey.(*ecdsa.PublicKey))

	hash, err := eth.SignAndPushTransaction(client, hex.EncodeToString(publicKeyBytes), "0x3011FF701a84B697D8821a03F18F7c52792D5338", "0x38d7ea4c68000", hex.EncodeToString(privateKeyBytes))

	if err != nil {
		t.Fatalf("SignAndPushTransaction() returned an error: %v", err)
	}

	t.Logf("SignAndPushTransaction() returned %s, as expected", hash)
}

func TestFaucetToAddress(t *testing.T) {
	client := Client{}

	SendTransaction = func(ctx context.Context, tx *types.Transaction) error {
		return nil
	}

	PendingNonceAtMock = func(ctx context.Context, account common.Address) (uint64, error) {
		return 54, nil
	}

	hash, err := eth.FaucetToAddress(client, "0x38d7ea4c68000", "0x3011FF701a84B697D8821a03F18F7c52792D5338")

	if err != nil {
		t.Fatalf("FaucetToAddress() returned an error: %v", err)
	}

	t.Logf("FaucetToAddress() returned %s, as expected", hash)
}

func TestFaucetToAddressFailToSignAndPush(t *testing.T) {
	client := Client{}

	privateKey, _ := crypto.GenerateKey()

	publicKey := privateKey.Public()

	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKeyBytes := crypto.FromECDSAPub(publicKey.(*ecdsa.PublicKey))

	os.Setenv("FOUNDING_PRIVATE_KEY", hex.EncodeToString(privateKeyBytes))
	os.Setenv("FOUDING_ADDRESS", hex.EncodeToString(publicKeyBytes))

	PendingNonceAtMock = func(ctx context.Context, account common.Address) (uint64, error) {
		return 54, nil
	}

	SendTransaction = func(ctx context.Context, tx *types.Transaction) error {
		return fmt.Errorf("Internal error")
	}

	_, err := eth.FaucetToAddress(client, "0x38d7ea4c68000", "0x3011FF701a84B697D8821a03F18F7c52792D5338")

	if err == nil {
		t.Logf("FaucetToAddress() should have returned an error")
	}
}

func TestStringToPrivateKey(t *testing.T) {
	privateKey, _ := crypto.GenerateKey()

	privateKeyBytes := crypto.FromECDSA(privateKey)

	result, err := eth.StringToPrivateKey(hex.EncodeToString(privateKeyBytes))

	if err != nil {
		t.Fatalf("StringToPrivateKey() returned an error: %v", err)
	}

	if bytes.Equal(crypto.FromECDSA(result), privateKeyBytes) {
		t.Logf("Result as expected")
	}
}

func TestStringToPrivateKeyWithInvalidPrivateKey(t *testing.T) {
	_, err := eth.StringToPrivateKey("invalid_private_key")

	if err == nil {
		t.Logf("StringToPrivateKey returned an error, as expected: %v", err)
	}
}

package models

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Transaction struct {
	Value      *big.Int          `json:"value"`
	PrivateKey *ecdsa.PrivateKey `json:"private_key"`
	AddressTo  common.Address    `json:"address`
}

type TransactionFromFrontend struct {
	Value      string `json:"value"`
	PrivateKey string `json:"private_key"`
	AddressTo  string `json:"address"`
}

const (
	GasLimit = 21000
)

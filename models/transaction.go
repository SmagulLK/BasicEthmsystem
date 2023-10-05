package models

import (
	"math/big"
)

type Transaction struct {
	Value      big.Int `json:"value"`
	PrivateKey string  `json:"private_key"`
	AddressTo  string  `json:"address"`
	Hex        string  `json:"hex,omitempty"`
	// this is no need on frontend
}

const (
	GasLimit = 210000
)

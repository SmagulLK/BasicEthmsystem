package models

import "math/big"

type Transaction struct {
	Value       string `json:"value"`
	PrivateKey  string `json:"private_key"`
	AddressTo   string `json:"address"`
	Hex         string `json:"hex,omitempty"`
	ValueBigInt *big.Int
	// this is no need on frontend
}

const (
	GasLimit = 2100000
)

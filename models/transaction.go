package models

import (
	"database/sql/driver"
	"fmt"
	"math/big"
)

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

type BigInt big.Int

func (b *BigInt) Value() (driver.Value, error) {
	if b != nil {
		return (*big.Int)(b).String(), nil
	}
	return nil, nil
}

func (b *BigInt) Scan(value interface{}) error {
	if value == nil {
		b = nil
	}

	switch t := value.(type) {
	case []uint8:
		_, ok := (*big.Int)(b).SetString(string(value.([]uint8)), 10)
		if !ok {
			return fmt.Errorf("failed to load value to []uint8: %v", value)
		}
	default:
		return fmt.Errorf("Could not scan type %T into BigInt", t)
	}

	return nil
}

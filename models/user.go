package models

import (
	"math/big"
)

type User struct {
	Balance    big.Int `json:"balance" binding:"required"`
	UserID     int     `json:"user_id,omitempty" `
	PublicKey  string  `json:"public_key" binding:"required"`
	PrivateKey string  `json:"private_key" binding:"required"`
	Address    string  `json:"address" binding:"required"`
}

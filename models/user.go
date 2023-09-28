package models

import (
	"crypto"
	"math/big"
)

type User struct {
	Balance    big.Int           `json:"balance" binding:"required"`
	UserID     int               `json:"user_id" `
	PublicKey  crypto.PublicKey  `json:"public_key" binding:"required"`
	PrivateKey crypto.PrivateKey `json:"private_key" binding:"required"`
}
